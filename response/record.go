package response

import (
	"bytes"
	"encoding/binary"
	"fmt"

	"github.com/dlion/runne/packet"
)

func ParseRecord(reader *bytes.Reader) *packet.Record {
	var record packet.Record
	record.Name = []byte(decodeName(reader))
	binary.Read(reader, binary.BigEndian, &record.Type)
	binary.Read(reader, binary.BigEndian, &record.Class)
	binary.Read(reader, binary.BigEndian, &record.TTL)
	binary.Read(reader, binary.BigEndian, &record.RdLength)
	switch record.Type {
	case packet.TYPE_A:
		record.Rdata = readIP(reader, record.RdLength)
	case packet.TYPE_NS:
		record.Rdata = decodeName(reader)
	default:
		record.Rdata = string(readData(reader, record.RdLength))
	}
	return &record
}

func readIP(reader *bytes.Reader, length uint16) string {
	dataBytes := readData(reader, length)
	return fmt.Sprintf("%d.%d.%d.%d", dataBytes[0], dataBytes[1], dataBytes[2], dataBytes[3])
}

func readData(reader *bytes.Reader, length uint16) []byte {
	dataBytes := make([]byte, length)
	binary.Read(reader, binary.BigEndian, &dataBytes)
	return dataBytes
}
