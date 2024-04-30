package packet

import (
	"encoding/hex"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
	t.Run("Should create a query", func(t *testing.T) {
		header := NewHeader(22, RECURSION_FLAG, 1, 0, 0, 0)
		question := NewQuestion("dns.google.com", TYPE_A, CLASS_IN)

		query := NewQuery(header, question)

		expected, err := hex.DecodeString("00160100000100000000000003646e7306676f6f676c6503636f6d0000010001")
		assert.Nil(t, err)
		assert.Equal(t, expected, query)
	})
}
