package builder

import (
	"bytes"
	"encoding/binary"
)

const TYPE_A uint16 = 1
const CLASS_IN uint16 = 1

type Question struct {
	QName  []byte
	QType  uint16
	QClass uint16
}

func NewQuestion(QName []byte, QType, QClass uint16) *Question {
	return &Question{
		QName:  EncodeDnsName(QName),
		QType:  QType,
		QClass: QClass,
	}
}

func (q *Question) ToBytes() []byte {
	encodedQuestion := new(bytes.Buffer)
	binary.Write(encodedQuestion, binary.BigEndian, q.QName)
	binary.Write(encodedQuestion, binary.BigEndian, q.QType)
	binary.Write(encodedQuestion, binary.BigEndian, q.QClass)

	return encodedQuestion.Bytes()
}

func EncodeDnsName(QName []byte) []byte {
	var encoded []byte
	parts := bytes.Split([]byte(QName), []byte{'.'})
	for _, part := range parts {
		encoded = append(encoded, byte(len(part)))
		encoded = append(encoded, part...)
	}

	return append(encoded, 0x00)
}
