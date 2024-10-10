package query

import (
	"bytes"
	"encoding/binary"
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
