package packet

import (
	"bytes"
	"encoding/hex"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEncodeHeaderIntoBytes(t *testing.T) {
	header := NewHeader(22, RECURSION_FLAG, 1, 0, 0, 0)
	encodedHeader := header.ToBytes()

	expected, err := hex.DecodeString("0016010000010000000000000")
	assert.NotNil(t, err)
	assert.Equal(t, expected, encodedHeader)
}

func TestCreateHeaderFromResponse(t *testing.T) {
	response, _ := hex.DecodeString("001680800001000200000000")
	header, _ := ParseHeader(bytes.NewReader(response))

	assert.Equal(t, &Header{
		Id:      0x16,
		Flags:   1<<15 | 1<<7,
		QdCount: 0x1,
		AnCount: 0x2,
		NsCount: 0x0,
		ArCount: 0x0,
	}, header)
}

func TestParseHeaderReturnQueryError(t *testing.T) {
	response, _ := hex.DecodeString("001680810001000200000000")
	header, err := ParseHeader(bytes.NewReader(response))

	assert.Nil(t, header)
	assert.NotNil(t, err)
	assert.EqualError(t, err, "query error")
}

func TestParseHeaderReturnServerError(t *testing.T) {
	response, _ := hex.DecodeString("001680820001000200000000")
	header, err := ParseHeader(bytes.NewReader(response))

	assert.Nil(t, header)
	assert.NotNil(t, err)
	assert.EqualError(t, err, "server error")
}

func TestParseHeaderReturnDomainNotExistError(t *testing.T) {
	response, _ := hex.DecodeString("001680830001000200000000")
	header, err := ParseHeader(bytes.NewReader(response))

	assert.Nil(t, header)
	assert.NotNil(t, err)
	assert.EqualError(t, err, "the domain does not exist")
}
