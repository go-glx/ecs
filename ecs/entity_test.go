package ecs

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/fe3dback/glx-ecs/ecs/internal/ids"
)

func TestNewEntity(t *testing.T) {
	ent := testCreateEntity("basic")

	assert.Equal(t, uint64(1), ent.id)
	assert.Equal(t, uint64(1), ent.ID())
	assert.Equal(t, "basic", ent.name)
	assert.Equal(t, "basic (1)", ent.String())
	assert.Equal(t, 0, ent.components.Len(), "default entity not has components")

	// --

	ent2 := testCreateEntity("ent2")
	assert.Equal(t, uint64(2), ent2.id, "id should be global auto increment")

	// --
	assert.Panics(t, func() { ent.AddComponent(nil) }, "nil component not valid")
	assert.Panics(t, func() { ent.AddComponent(testMutableComponent{}) }, "component should be passed by reference")
	assert.Panics(t, func() { ent.AddComponent(testCreateComplexComponent()) }, "should require another component")

	assert.NotPanics(t, func() {
		ent.AddComponent(&testMutableComponent{counter: 42}) // first time creating, value = 42
		ent.AddComponent(testCreateMutableComponent())       // ignore, already exist
		ent.AddComponent(&testMutableComponent{counter: 5})  // ignore, already exist (value not changed from default 42)

		ent.AddComponent(testCreateComplexComponent()) // valid, because we meet requirement
	})

	assert.Equal(t, 2, ent.components.Len(), "expect two unique components")

	// -- inner black magic testing
	mutableCmp, mutableCmpExist := ent.components.Get(ids.Of(testMutableComponent{}))
	assert.True(t, mutableCmpExist)
	assert.Equal(t, 42, mutableCmp.(*testMutableComponent).counter, "should not be override to 5 from AddComponent duplicate call")
}

func testCreateEntity(name string) *Entity {
	return NewEntity(name)
}
