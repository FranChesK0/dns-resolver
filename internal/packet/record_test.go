package packet

import (
	"bytes"
	"encoding/hex"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateRecordFromResponse(t *testing.T) {
	response, _ := hex.DecodeString("00168080000100020000000003646e7306676f6f676c6503636f6d0000010001c00c0001000100000214000408080808c00c0001000100000214000408080404")
	reader := bytes.NewReader(response)
	const RECORD_STARTING_POINT = 32
	skipResponseTill(t, reader, response, RECORD_STARTING_POINT)

	record := ParseRecord(reader)

	assert.NotEmpty(t, record)
	assert.Equal(t, TYPE_A, record.Type)
	assert.Equal(t, CLASS_IN, record.Class)
	assert.Greater(t, record.TTL, uint32(0))
	assert.Greater(t, record.RdLength, uint16(0))
	assert.Equal(t, "8.8.8.8", record.RData)

	record = ParseRecord(reader)

	assert.NotEmpty(t, record)
	assert.Equal(t, TYPE_A, record.Type)
	assert.Equal(t, CLASS_IN, record.Class)
	assert.Greater(t, record.TTL, uint32(0))
	assert.Greater(t, record.RdLength, uint16(0))
	assert.Equal(t, "8.8.4.4", record.RData)
}
