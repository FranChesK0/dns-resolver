package packet

import (
	"bytes"
	"encoding/hex"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEncodeQuestionIntoBytes(t *testing.T) {
	question := NewQuestion("test.domain.name", TYPE_A, CLASS_IN)
	encodedQuestion := question.ToBytes()

	expected, _ := hex.DecodeString("047465737406646f6d61696e046e616d650000010001")
	assert.NotNil(t, expected)
	assert.Equal(t, expected, encodedQuestion)
}

func TestEncodeDNSName(t *testing.T) {
	encodedDNSName := encodeDnsName([]byte("test.domain.name"))
	assert.Equal(t, []byte("\x04test\x06domain\x04name\x00"), encodedDNSName)
}

func TestCreateQuestionFromResponse(t *testing.T) {
	response, _ := hex.DecodeString("001680800001000200000000047465737406646f6d61696e046e616d650000010001c00c0001000100000214000408080808c00c0001000100000214000408080404")
	reader := bytes.NewReader(response)
	const QUESTION_STARTING_POINT = 12
	skipResponseTill(t, reader, response, QUESTION_STARTING_POINT)

	question := ParseQuestion(bytes.NewReader(response))

	assert.NotEmpty(t, question)
	assert.Equal(t, &Question{
		QName:  []byte("test.domain.name"),
		QType:  TYPE_A,
		QClass: CLASS_IN,
	}, question)
}

func skipResponseTill(t *testing.T, reader *bytes.Reader, response []byte, startingPoint int64) {
	t.Helper()
	reader.ReadAt(response, startingPoint)
}
