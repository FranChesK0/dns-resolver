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

func NewHeader(id, flags, qdCount, anCount, nsCount, arCount uint16) *Header {
	return &Header{
		Id:      id,
		Flags:   flags,
		QdCount: qdCount,
		AnCount: anCount,
		NsCount: nsCount,
		ArCount: arCount,
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
		return nil, errors.New("query error")
	case 2:
		return nil, errors.New("server error")
	case 3:
		return nil, errors.New("the domain does not exist")
	}
	binary.Read(reader, binary.BigEndian, &header.QdCount)
	binary.Read(reader, binary.BigEndian, &header.AnCount)
	binary.Read(reader, binary.BigEndian, &header.NsCount)
	binary.Read(reader, binary.BigEndian, &header.ArCount)
	return &header, nil
}
