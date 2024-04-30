package builder

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"strings"
)

type Record struct {
	Name     []byte
	Type     uint16
	Class    uint16
	TTL      uint32
	RdLength uint16
	Rdata    string
}

const RESPONSE uint16 = 1<<15 | 1<<7

func ParseHeader(responseReader *bytes.Reader) *Header {
	var header Header

	binary.Read(responseReader, binary.BigEndian, &header.Id)
	binary.Read(responseReader, binary.BigEndian, &header.Flags)
	binary.Read(responseReader, binary.BigEndian, &header.QdCount)
	binary.Read(responseReader, binary.BigEndian, &header.AnCount)
	binary.Read(responseReader, binary.BigEndian, &header.NsCount)
	binary.Read(responseReader, binary.BigEndian, &header.ArCount)
	return &header
}

func (h *Header) Print() {
	fmt.Println("--- HEADER ---")
	fmt.Printf("ID: %v\n", h.Id)
	fmt.Printf("AnCount: %v\n", h.AnCount)
	fmt.Printf("ArCount: %v\n", h.ArCount)
	fmt.Printf("NsCount: %v\n", h.NsCount)
	fmt.Printf("QdCount: %v\n", h.QdCount)
	fmt.Printf("Flags: 0x%X\n", h.Flags)
}

func DecodeDnsName(reader *bytes.Reader) string {
	var parts []string

	for {
		length, _ := reader.ReadByte()
		if length == 0 {
			break
		}

		part := make([]byte, length)
		reader.Read(part)
		parts = append(parts, string(part))
	}

	return strings.Join(parts, ".")
}

func ParseQuestion(responseReader *bytes.Reader) *Question {
	var question Question

	question.QName = []byte(DecodeDnsName(responseReader))
	binary.Read(responseReader, binary.BigEndian, &question.QType)
	binary.Read(responseReader, binary.BigEndian, &question.QClass)

	return &question
}

func (q *Question) Print(n uint16) {
	fmt.Printf("--- QUESTION %d ---\n", n)
	fmt.Printf("Name: %s\n", q.QName)
	fmt.Printf("Type: 0x%X\n", q.QType)
	fmt.Printf("Class: 0x%X\n", q.QClass)
}

func ParseRecord(reader *bytes.Reader) *Record {
	var record Record

	reader.Read(make([]byte, 2))
	binary.Read(reader, binary.BigEndian, &record.Type)
	binary.Read(reader, binary.BigEndian, &record.Class)
	binary.Read(reader, binary.BigEndian, &record.TTL)
	binary.Read(reader, binary.BigEndian, &record.RdLength)
	dataBytes := make([]byte, record.RdLength)
	binary.Read(reader, binary.BigEndian, &dataBytes)
	record.Rdata = fmt.Sprintf("%d.%d.%d.%d", dataBytes[0], dataBytes[1], dataBytes[2], dataBytes[3])

	return &record
}

const (
	QUESTION = iota
	ANSWER
	AUTHORITIES
	ADDITIONALS
)

func (r *Record) Print(t int, n uint16) {
	switch t {
	case ANSWER:
		fmt.Printf("--- ANSWER %d ---\n", n)
	case AUTHORITIES:
		fmt.Printf("--- AUTHORITIES %d ---\n", n)
	case ADDITIONALS:
		fmt.Printf("---  ADDITIONALS %d ---\n", n)
	}
	fmt.Printf("Name: %s\n", r.Name)
	fmt.Printf("Type: 0x%X\n", r.Type)
	fmt.Printf("Class: 0x%X\n", r.Class)
	fmt.Printf("TTL: %d\n", r.TTL)
	fmt.Printf("Length: %d\n", r.RdLength)
	fmt.Printf("Data: %s\n", r.Rdata)
}
