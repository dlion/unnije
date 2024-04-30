package packet

import (
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
}
