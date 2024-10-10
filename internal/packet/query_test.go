package packet

import (
	"encoding/hex"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateQuery(t *testing.T) {
	header := NewHeader(22, RECURSION_FLAG, 1, 0, 0, 0)
	question := NewQuestion("test.domain.name", TYPE_A, CLASS_IN)

	query := NewQuery(header, question)

	expected, err := hex.DecodeString("001601000001000000000000047465737406646f6d61696e046e616d650000010001")
	assert.Nil(t, err)
	assert.Equal(t, expected, query)
}
