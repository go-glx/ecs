package snapshot

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_marshalXML(t *testing.T) {
	actual, err := marshalXML(testCreateStaticWorld())
	assert.NoError(t, err)

	expected := testCreateStaticWorldXML()

	assert.Equal(t, expected, string(actual))
}

func testCreateStaticWorldXML() string {
	//language=XML
	return `<StaticWorld>
  <systems>
    <system id="GarbageCollector-adc3353dd900"></system>
  </systems>
  <entities>
    <entity name="ent1">
      <components>
        <component id="Deletable-a300548e4f48">
          <props>
            <prop name="Alive" value="true"></prop>
          </props>
        </component>
        <component id="Node2D-7c40b8e315a5">
          <props>
            <prop name="X" value="5"></prop>
            <prop name="Y" value="10"></prop>
          </props>
        </component>
      </components>
    </entity>
    <entity name="ent2">
      <components>
        <component id="Node2D-7c40b8e315a5">
          <props>
            <prop name="X" value="4"></prop>
            <prop name="Y" value="7"></prop>
          </props>
        </component>
      </components>
    </entity>
  </entities>
</StaticWorld>`
}
