package response

import (
	"bytes"
	"io"
)

func decodeName(reader *bytes.Reader) string {
	var name bytes.Buffer

	for {
		lengthByte, _ := reader.ReadByte()

		if (lengthByte & 0xC0) == 0xC0 {
			name.WriteString(getBackTheDomainFromTheHeader(reader, lengthByte))
			break
		}

		if lengthByte == 0 {
			break
		}

		label := make([]byte, lengthByte)
		io.ReadFull(reader, label)
		if name.Len() > 0 {
			name.WriteByte('.')
		}
		name.Write(label)
	}

	return name.String()
}

func getBackTheDomainFromTheHeader(reader *bytes.Reader, lengthByte byte) string {
	nextByte, _ := reader.ReadByte()
	pointer := uint16((uint16(lengthByte) & 0x3F) | uint16(nextByte))

	currentPos, _ := reader.Seek(0, io.SeekCurrent)

	reader.Seek(int64(pointer), io.SeekStart)

	decodedName := decodeName(reader)

	reader.Seek(currentPos, io.SeekStart)

	return decodedName
}
