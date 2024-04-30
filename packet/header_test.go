package packet

import (
	"bytes"
	"encoding/hex"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHeader(t *testing.T) {
	t.Run("Should encode an header into bytes", func(t *testing.T) {
		header := NewHeader(22, RECURSION_FLAG, 1, 0, 0, 0)

		encodedHeader := header.ToBytes()

		expected, err := hex.DecodeString("0016010000010000000000000")
		assert.NotNil(t, err)
		assert.Equal(t, expected, encodedHeader)
	})

	t.Run("Should create an header from a response", func(t *testing.T) {
		response, _ := hex.DecodeString("001680800001000200000000")
		header, _ := ParseHeader(bytes.NewReader(response))

		assert.Equal(t, &Header{
			Id:      0x16,
			Flags:   1<<15 | 1<<7, // QR (Response) bit = 1, OPCODE = 0 (standard query), AA = 1, TC = 0, RD (Recursion Desired) bit = 1, RA = 1, Z = 0, RCODE = 0
			QdCount: 0x1,
			AnCount: 0x2,
			NsCount: 0x0,
			ArCount: 0x0,
		}, header)
	})

	t.Run("Should return an error if the header flags contains a query error", func(t *testing.T) {
		response, _ := hex.DecodeString("001680810001000200000000")

		header, err := ParseHeader(bytes.NewReader(response))

		assert.Nil(t, header)
		assert.NotNil(t, err)
		assert.EqualError(t, err, "error with the query")
	})

	t.Run("Should return an error if the header flags contains a server error", func(t *testing.T) {
		response, _ := hex.DecodeString("001680820001000200000000")

		header, err := ParseHeader(bytes.NewReader(response))

		assert.Nil(t, header)
		assert.NotNil(t, err)
		assert.EqualError(t, err, "error with the server")
	})

	t.Run("Should return an error if the header flags contains a domain not exist error", func(t *testing.T) {
		response, _ := hex.DecodeString("001680830001000200000000")

		header, err := ParseHeader(bytes.NewReader(response))

		assert.Nil(t, header)
		assert.NotNil(t, err)
		assert.EqualError(t, err, "the domain doesn't exist")
	})
}
