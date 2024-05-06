package packet

import (
	"bytes"
	"io"
	"strings"
)

func DecodeName(reader *bytes.Reader) string {
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
		name.Write(label)
		name.WriteByte('.')

	}

	result, _ := strings.CutSuffix(name.String(), ".")
	return result
}

func getBackTheDomainFromTheHeader(reader *bytes.Reader, lengthByte byte) string {
	nextByte, _ := reader.ReadByte()
	pointer := uint16((uint16(lengthByte) & 0x3F) | uint16(nextByte))

	currentPos, _ := reader.Seek(0, io.SeekCurrent)

	reader.Seek(int64(pointer), io.SeekStart)

	decodedName := DecodeName(reader)

	reader.Seek(currentPos, io.SeekStart)

	return decodedName
}
