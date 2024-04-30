package response

import (
	"bytes"
	"encoding/binary"
	"fmt"

	"github.com/dlion/cue/packet"
)

func ParseRecord(reader *bytes.Reader) *packet.Record {
	var record packet.Record

	record.Name = []byte(decodeName(reader))

	var recordType uint16
	if err := binary.Read(reader, binary.BigEndian, &recordType); err != nil {
		fmt.Printf("ERRRR: %v", err)
		return nil
	}
	record.Type = recordType

	var recordClass uint16
	if err := binary.Read(reader, binary.BigEndian, &recordClass); err != nil {
		fmt.Printf("ERRRR: %v", err)

		return nil
	}
	record.Class = recordClass

	if err := binary.Read(reader, binary.BigEndian, &record.TTL); err != nil {
		fmt.Printf("ERRRR: %v", err)

		return nil
	}

	var rdLength uint16
	if err := binary.Read(reader, binary.BigEndian, &rdLength); err != nil {
		fmt.Printf("ERRRR: %v", err)

		return nil
	}
	record.RdLength = rdLength

	dataBytes := make([]byte, record.RdLength)
	if err := binary.Read(reader, binary.BigEndian, &dataBytes); err != nil {
		fmt.Printf("ERRRR: %v", err)

		return nil
	}

	record.Rdata = fmt.Sprintf("%d.%d.%d.%d", dataBytes[0], dataBytes[1], dataBytes[2], dataBytes[3])

	return &record
}
