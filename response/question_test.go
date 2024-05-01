package response

import (
	"bytes"
	"encoding/hex"
	"testing"

	"github.com/dlion/runne/packet"
	"github.com/stretchr/testify/assert"
)

func TestQuestion(t *testing.T) {

	t.Run("Should read domain names from a question", func(t *testing.T) {
		response, _ := hex.DecodeString("00168080000100020000000003646e7306676f6f676c6503636f6d00") //truncated for readability
		reader := bytes.NewReader(response)
		const QUESTION_STARTING_POINT = 12
		skipResponseTill(t, reader, response, QUESTION_STARTING_POINT)

		dnsName := decodeName(reader)

		assert.NotEmpty(t, dnsName)
		assert.Equal(t, "dns.google.com", dnsName)
	})

	t.Run("Should create a question from a response", func(t *testing.T) {
		response, _ := hex.DecodeString("00168080000100020000000003646e7306676f6f676c6503636f6d0000010001c00c0001000100000214000408080808c00c0001000100000214000408080404")
		reader := bytes.NewReader(response)
		const QUESTION_STARTING_POINT = 12
		skipResponseTill(t, reader, response, QUESTION_STARTING_POINT)

		question := ParseQuestion(bytes.NewReader(response))

		assert.NotEmpty(t, question)
		assert.Equal(t, &packet.Question{
			QName:  []byte("dns.google.com"),
			QType:  packet.TYPE_A,
			QClass: packet.CLASS_IN,
		}, question)
	})
}
