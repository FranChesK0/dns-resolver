package packet

import (
	"bytes"
	"encoding/binary"

	"github.com/FranChesK0/dns-resolver/internal/decode"
)

const (
	TYPE_A   uint16 = 1
	TYPE_NS  uint16 = 2
	CLASS_IN uint16 = 1
)

type Question struct {
	QName  []byte
	QType  uint16
	QClass uint16
}

func NewQuestion(qName string, qType uint16, qClass uint16) *Question {
	return &Question{
		QName:  encodeDnsName([]byte(qName)),
		QType:  qType,
		QClass: qClass,
	}
}

func (q *Question) ToBytes() []byte {
	encodedQuestion := new(bytes.Buffer)
	binary.Write(encodedQuestion, binary.BigEndian, q.QName)
	binary.Write(encodedQuestion, binary.BigEndian, q.QType)
	binary.Write(encodedQuestion, binary.BigEndian, q.QClass)
	return encodedQuestion.Bytes()
}

func ParseQuestion(reader *bytes.Reader) *Question {
	var question Question
	question.QName = []byte(decode.DecodeName(reader))
	binary.Read(reader, binary.BigEndian, &question.QType)
	binary.Read(reader, binary.BigEndian, &question.QClass)
	return &question
}

func encodeDnsName(qName []byte) []byte {
	encoded := make([]byte, 0)
	parts := bytes.Split([]byte(qName), []byte{'.'})
	for _, part := range parts {
		encoded = append(encoded, byte(len((part))))
		encoded = append(encoded, part...)
	}
	return append(encoded, 0x00)
}
