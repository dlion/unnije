package packet

import (
	"bytes"
	"encoding/hex"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQuestion(t *testing.T) {
	t.Run("Should encode a question into bytes", func(t *testing.T) {
		question := NewQuestion("dns.google.com", TYPE_A, CLASS_IN)

		encodedQuestion := question.ToBytes()

		expected, _ := hex.DecodeString("03646e7306676f6f676c6503636f6d0000010001")
		assert.NotNil(t, expected)
		assert.Equal(t, expected, encodedQuestion)
	})

	t.Run("Should encode the dns name", func(t *testing.T) {
		encodedDnsName := encodeDnsName([]byte("dns.google.com"))
		assert.Equal(t, []byte("\x03dns\x06google\x03com\x00"), encodedDnsName)
	})

	t.Run("Should create a question from a response", func(t *testing.T) {
		response, _ := hex.DecodeString("00168080000100020000000003646e7306676f6f676c6503636f6d0000010001c00c0001000100000214000408080808c00c0001000100000214000408080404")
		reader := bytes.NewReader(response)
		const QUESTION_STARTING_POINT = 12
		skipResponseTill(t, reader, response, QUESTION_STARTING_POINT)

		question := ParseQuestion(bytes.NewReader(response))

		assert.NotEmpty(t, question)
		assert.Equal(t, &Question{
			QName:  []byte("dns.google.com"),
			QType:  TYPE_A,
			QClass: CLASS_IN,
		}, question)
	})
}
