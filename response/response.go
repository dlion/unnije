package response

import (
	"bytes"
	"fmt"
	"io"
)

func decodeName(reader *bytes.Reader) string {
	var name bytes.Buffer

	for {
		lengthByte, _ := reader.ReadByte()

		if (lengthByte & 0xC0) == 0xC0 {

			nextByte, err := reader.ReadByte()
			if err != nil {
				fmt.Printf("ERRRR: %v", err)
				break
			}

			pointer := uint16((uint16(lengthByte) & 0x3F) | uint16(nextByte))

			currentPos, err := reader.Seek(0, io.SeekCurrent)
			if err != nil {
				fmt.Printf("ERRRR: %v", err)
				break
			}

			_, err = reader.Seek(int64(pointer), io.SeekStart)
			if err != nil {
				fmt.Printf("ERRRR: %v", err)
				break
			}

			decodedName := decodeName(reader)
			_, err = name.WriteString(decodedName)
			if err != nil {
				fmt.Printf("ERRRR: %v", err)
				break
			}

			_, err = reader.Seek(currentPos, io.SeekStart)
			if err != nil {
				fmt.Printf("ERRRR: %v", err)
				break
			}

			break
		} else {
			label := make([]byte, lengthByte)
			_, err := reader.Read(label)
			if err != nil {
				fmt.Printf("ERRRR: %v", err)
				break
			}
			_, err = name.Write(label)
			if err != nil {
				fmt.Printf("ERRRR: %v", err)
				break
			}

			if lengthByte == 0 {
				break
			}

			_, err = name.WriteString(".")
			if err != nil {
				fmt.Printf("ERRRR: %v", err)
				break
			}
		}
	}

	result := name.String()
	if len(result) > 0 && result[len(result)-1] == '.' {
		result = result[:len(result)-1]
	}

	return result
}
