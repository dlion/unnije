package builder

import (
	"bytes"
	"encoding/hex"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestResponse(t *testing.T) {
	t.Run("Should create an header from a response", func(t *testing.T) {
		response, _ := hex.DecodeString("001680800001000200000000")
		header := ParseHeader(bytes.NewReader(response))

		assert.Equal(t, &Header{
			Id:      0x16,
			Flags:   RESPONSE, // QR (Response) bit = 1, OPCODE = 0 (standard query), AA = 1, TC = 0, RD (Recursion Desired) bit = 1, RA = 1, Z = 0, RCODE = 0
			QdCount: 0x1,      // QDCOUNT: 1
			AnCount: 0x2,      // ANCOUNT: 2
			NsCount: 0x0,      // NSCOUNT: 0
			ArCount: 0x0,      // ARCOUNT: 0
		}, header)
	})

	t.Run("Should read domain names from a response", func(t *testing.T) {
		response, _ := hex.DecodeString("03646e7306676f6f676c6503636f6d0000010001c00c0001000100000214000408080808c00c0001000100000214000408080404")

		dnsName := DecodeDnsName(bytes.NewReader(response))

		assert.NotEmpty(t, dnsName)
		assert.Equal(t, "dns.google.com", dnsName)
	})

	t.Run("Should create a question from a response", func(t *testing.T) {
		response, _ := hex.DecodeString("03646e7306676f6f676c6503636f6d0000010001c00c0001000100000214000408080808c00c0001000100000214000408080404")

		question := ParseQuestion(bytes.NewReader(response))

		assert.NotEmpty(t, question)
		assert.Equal(t, &Question{
			QName:  []byte("dns.google.com"),
			QType:  TYPE_A,
			QClass: CLASS_IN,
		}, question)
	})

	t.Run("Should create a record from a response", func(t *testing.T) {
		response, _ := hex.DecodeString("c00c0001000100000214000408080808c00c0001000100000214000408080404")
		reader := bytes.NewReader(response)

		record := ParseRecord(reader)

		assert.NotEmpty(t, record)
		assert.Equal(t, TYPE_A, record.Type)
		assert.Equal(t, CLASS_IN, record.Class)
		assert.Greater(t, record.TTL, uint32(0))
		assert.Greater(t, record.RdLength, uint16(0))
		assert.Equal(t, "8.8.8.8", record.Rdata)

		record = ParseRecord(reader)

		assert.NotEmpty(t, record)
		assert.Equal(t, TYPE_A, record.Type)
		assert.Equal(t, CLASS_IN, record.Class)
		assert.Greater(t, record.TTL, uint32(0))
		assert.Greater(t, record.RdLength, uint16(0))
		assert.Equal(t, "8.8.4.4", record.Rdata)
	})
}
