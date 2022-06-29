package snapshot

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_unmarshalXML(t *testing.T) {
	xml := testCreateStaticWorldXML()
	expected := testCreateStaticWorld()

	actual, err := unmarshalXML([]byte(xml))
	assert.NoError(t, err)

	assert.Equal(t, expected, actual)
}
