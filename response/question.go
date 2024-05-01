package response

import (
	"bytes"
	"encoding/binary"

	"github.com/dlion/runne/packet"
)

func ParseQuestion(reader *bytes.Reader) *packet.Question {
	var question packet.Question

	question.QName = []byte(decodeName(reader))
	binary.Read(reader, binary.BigEndian, &question.QType)
	binary.Read(reader, binary.BigEndian, &question.QClass)

	return &question
}
