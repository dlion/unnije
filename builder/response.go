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
