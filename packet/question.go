package packet

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

const TYPE_A uint16 = 1
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

func (q *Question) Print(n uint16) {
	fmt.Printf("--- QUESTION %d ---\n", n)
	fmt.Printf("Name: %s\n", q.QName)
	fmt.Printf("Type: 0x%X\n", q.QType)
	fmt.Printf("Class: 0x%X\n", q.QClass)
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