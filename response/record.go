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
	record.Rdata = readData(reader, record.RdLength)
	return &record
}

func readData(reader *bytes.Reader, length uint16) string {
	dataBytes := make([]byte, length)
	binary.Read(reader, binary.BigEndian, &dataBytes)
	return fmt.Sprintf("%d.%d.%d.%d", dataBytes[0], dataBytes[1], dataBytes[2], dataBytes[3])
}
