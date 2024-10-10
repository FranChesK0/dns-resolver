package packet

import (
	"bytes"
	"encoding/hex"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReadDomainNamesFromQuestion(t *testing.T) {
	response, _ := hex.DecodeString("001680800001000200000000047465737406646f6d61696e046e616d650000010001")
	reader := bytes.NewReader(response)
	const QUESTION_STARTING_POINT = 12
	skipResponseTill(t, reader, response, QUESTION_STARTING_POINT)

	dnsName := DecodeName(reader)

	assert.NotEmpty(t, dnsName)
	assert.Equal(t, "test.domain.name", dnsName)
}
