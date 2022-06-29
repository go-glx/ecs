package snapshot

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEncodeWorld(t *testing.T) {
	snapshot := encodeWorld(testCreateWorld())
	expected := testCreateStaticWorld()

	assert.Equal(t, expected, snapshot)
}
