package packet

import (
	"bytes"
	"testing"
)

func skipResponseTill(t *testing.T, reader *bytes.Reader, response []byte, startingPoint int64) {
	t.Helper()
	reader.ReadAt(response, startingPoint)
}
