package builder

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
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

func ParseQuestion(responseReader *bytes.Reader) *Question {
	var question Question

	question.QName = []byte(decodeName(responseReader))
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

func decodeName(reader *bytes.Reader) string {
	var name bytes.Buffer

	for {
		lengthByte, _ := reader.ReadByte()

		if (lengthByte & 0xC0) == 0xC0 {

			nextByte, err := reader.ReadByte()
			if err != nil {
				fmt.Printf("ERRRR: %v", err)
				break
			}

			pointer := uint16((uint16(lengthByte) & 0x3F) | uint16(nextByte))

			currentPos, err := reader.Seek(0, io.SeekCurrent)
			if err != nil {
				fmt.Printf("ERRRR: %v", err)
				break
			}

			_, err = reader.Seek(int64(pointer), io.SeekStart)
			if err != nil {
				fmt.Printf("ERRRR: %v", err)
				break
			}

			decodedName := decodeName(reader)
			_, err = name.WriteString(decodedName)
			if err != nil {
				fmt.Printf("ERRRR: %v", err)
				break
			}

			_, err = reader.Seek(currentPos, io.SeekStart)
			if err != nil {
				fmt.Printf("ERRRR: %v", err)
				break
			}

			break
		} else {
			label := make([]byte, lengthByte)
			_, err := reader.Read(label)
			if err != nil {
				fmt.Printf("ERRRR: %v", err)
				break
			}
			_, err = name.Write(label)
			if err != nil {
				fmt.Printf("ERRRR: %v", err)
				break
			}

			if lengthByte == 0 {
				break
			}

			_, err = name.WriteString(".")
			if err != nil {
				fmt.Printf("ERRRR: %v", err)
				break
			}
		}
	}

	result := name.String()
	if len(result) > 0 && result[len(result)-1] == '.' {
		result = result[:len(result)-1]
	}

	return result
}

func ParseRecord(reader *bytes.Reader) *Record {
	var record Record

	record.Name = []byte(decodeName(reader))

	// Read the record.Type (16-bit unsigned integer)
	var recordType uint16
	if err := binary.Read(reader, binary.BigEndian, &recordType); err != nil {
		fmt.Printf("ERRRR: %v", err)
		return nil
	}
	record.Type = recordType

	// Read the record.Class (16-bit unsigned integer)
	var recordClass uint16
	if err := binary.Read(reader, binary.BigEndian, &recordClass); err != nil {
		fmt.Printf("ERRRR: %v", err)

		return nil
	}
	record.Class = recordClass

	// Read the record.TTL (32-bit unsigned integer)
	if err := binary.Read(reader, binary.BigEndian, &record.TTL); err != nil {
		fmt.Printf("ERRRR: %v", err)

		return nil
	}

	// Read the record.RdLength (16-bit unsigned integer)
	var rdLength uint16
	if err := binary.Read(reader, binary.BigEndian, &rdLength); err != nil {
		fmt.Printf("ERRRR: %v", err)

		return nil
	}
	record.RdLength = rdLength

	// Read the record.Rdata (variable length, based on RdLength)
	dataBytes := make([]byte, record.RdLength)
	if err := binary.Read(reader, binary.BigEndian, &dataBytes); err != nil {
		fmt.Printf("ERRRR: %v", err)

		return nil
	}

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
