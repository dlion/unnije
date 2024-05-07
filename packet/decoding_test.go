package packet

import (
	"bytes"
	"encoding/hex"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDecoding(t *testing.T) {
	t.Run("Should read domain names from a question", func(t *testing.T) {
		response, _ := hex.DecodeString("00168080000100020000000003646e7306676f6f676c6503636f6d00") //truncated for readability
		reader := bytes.NewReader(response)
		const QUESTION_STARTING_POINT = 12
		skipResponseTill(t, reader, response, QUESTION_STARTING_POINT)

		dnsName := DecodeName(reader)

		assert.NotEmpty(t, dnsName)
		assert.Equal(t, "dns.google.com", dnsName)
	})
}
