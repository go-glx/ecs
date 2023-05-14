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
    <system id="internal/GarbageCollector">
      <props></props>
    </system>
  </systems>
  <entities>
    <entity name="ent1" prefab="">
      <components>
        <component id="internal/Node2D" order="0">
          <props>
            <prop name="X" value="5"/>
            <prop name="Y" value="10"/>
          </props>
        </component>
      </components>
    </entity>
    <entity name="ent2" prefab="">
      <components>
        <component id="internal/Node2D" order="0">
          <props>
            <prop name="X" value="4"/>
            <prop name="Y" value="7"/>
          </props>
        </component>
      </components>
    </entity>
  </entities>
</StaticWorld>`
}
