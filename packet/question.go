package packet

import (
	"bytes"
	"encoding/binary"
)

const TYPE_A uint16 = 1
const TYPE_NS uint16 = 2
const CLASS_IN uint16 = 1

type Question struct {
	QName  []byte
	QType  uint16
	QClass uint16
}

func NewQuestion(qname string, qtype, qclass uint16) *Question {
	return &Question{
		QName:  encodeDnsName([]byte(qname)),
		QType:  qtype,
		QClass: qclass,
	}
}

func (q *Question) ToBytes() []byte {
	encodedQuestion := new(bytes.Buffer)
	binary.Write(encodedQuestion, binary.BigEndian, q.QName)
	binary.Write(encodedQuestion, binary.BigEndian, q.QType)
	binary.Write(encodedQuestion, binary.BigEndian, q.QClass)

	return encodedQuestion.Bytes()
}

func encodeDnsName(qname []byte) []byte {
	var encoded []byte
	parts := bytes.Split([]byte(qname), []byte{'.'})
	for _, part := range parts {
		encoded = append(encoded, byte(len(part)))
		encoded = append(encoded, part...)
	}
	return append(encoded, 0x00)
}

func ParseQuestion(reader *bytes.Reader) *Question {
	var question Question

	question.QName = []byte(DecodeName(reader))
	binary.Read(reader, binary.BigEndian, &question.QType)
	binary.Read(reader, binary.BigEndian, &question.QClass)

	return &question
}
