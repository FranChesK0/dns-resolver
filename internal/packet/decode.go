package packet

import (
	"bytes"
	"io"
	"strings"
)

func DecodeName(reader *bytes.Reader) string {
	var name bytes.Buffer

	for {
		length, _ := reader.ReadByte()
		if (length & 0xC0) == 0xC0 {
			name.WriteString(getBackDomainFromHeader(reader, length))
			break
		}
		if length == 0 {
			break
		}

		label := make([]byte, length)
		io.ReadFull(reader, label)
		name.Write(label)
		name.WriteByte('.')
	}

	res, _ := strings.CutSuffix(name.String(), ".")
	return res
}

func getBackDomainFromHeader(reader *bytes.Reader, length byte) string {
	nextByte, _ := reader.ReadByte()
	pointer := uint16((uint16(length) & 0x3F) | uint16(nextByte))
	currPos, _ := reader.Seek(0, io.SeekCurrent)
	reader.Seek(int64(pointer), io.SeekStart)
	decodedName := DecodeName(reader)
	reader.Seek(currPos, io.SeekStart)
	return decodedName
}
