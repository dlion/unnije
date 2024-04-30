package network

import (
	"encoding/hex"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
	t.Run("Should check if the response starts with the same ID as the query", func(t *testing.T) {
		query, _ := hex.DecodeString("00160100000100000000000003646e7306676f6f676c6503636f6d0000010001")
		response, _ := hex.DecodeString("00168080000100020000000003646e7306676f6f676c6503636f6d0000010001c00c0001000100000214000408080808c00c0001000100000214000408080404")

		assert.True(t, hasTheSameID(query, response))
	})
}
