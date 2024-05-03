package packet

import (
	"bytes"
	"encoding/binary"
	"errors"
)

const RECURSION_FLAG uint16 = 1 << 8

type Header struct {
	Id      uint16
	Flags   uint16
	QdCount uint16
	AnCount uint16
	NsCount uint16
	ArCount uint16
}

func NewHeader(id, flags, qdcount, ancount, nscount, arcount uint16) *Header {
	return &Header{
		Id:      id,
		Flags:   flags,
		QdCount: qdcount,
		AnCount: ancount,
		NsCount: nscount,
		ArCount: arcount,
	}
}

func (h *Header) ToBytes() []byte {
	encodedHeader := new(bytes.Buffer)
	binary.Write(encodedHeader, binary.BigEndian, h.Id)
	binary.Write(encodedHeader, binary.BigEndian, h.Flags)
	binary.Write(encodedHeader, binary.BigEndian, h.QdCount)
	binary.Write(encodedHeader, binary.BigEndian, h.AnCount)
	binary.Write(encodedHeader, binary.BigEndian, h.NsCount)
	binary.Write(encodedHeader, binary.BigEndian, h.ArCount)

	return encodedHeader.Bytes()
}

func ParseHeader(reader *bytes.Reader) (*Header, error) {
	var header Header

	binary.Read(reader, binary.BigEndian, &header.Id)
	binary.Read(reader, binary.BigEndian, &header.Flags)
	switch header.Flags & 0b1111 {
	case 1:
		return nil, errors.New("error with the query")
	case 2:
		return nil, errors.New("error with the server")
	case 3:
		return nil, errors.New("the domain doesn't exist")
	}
	binary.Read(reader, binary.BigEndian, &header.QdCount)
	binary.Read(reader, binary.BigEndian, &header.AnCount)
	binary.Read(reader, binary.BigEndian, &header.NsCount)
	binary.Read(reader, binary.BigEndian, &header.ArCount)

	return &header, nil
}
