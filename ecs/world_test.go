package ecs

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewWorld(t *testing.T) {
	r := NewRegistry()
	r.RegisterSystem(testCreateMutateSystem())
	r.RegisterComponent(testCreateMutableComponent())
	r.RegisterComponent(testCreateComplexComponent())

	w := NewWorld(r)

	// -- create entities

	assert.Equal(t, 0, w.entities.Len(), "by default world is empty")

	mutEntity := testCreateEntity("ent1")
	mutEntity.AddComponent(testCreateMutableComponent()) // will be updated with MutateSystem

	secondsEntity := testCreateEntity("ent2")

	w.AddEntity(mutEntity)
	w.AddEntity(secondsEntity)
	assert.Equal(t, 0, w.entities.Len(), "entities queued, but not created yet")
	assert.Equal(t, 0, len(w.Entities()), "entities queued, but not created yet")

	w.AddSystem(testMutateSystemTypeID)
	assert.Equal(t, 0, w.systems.Len(), "system queued, but not created yet")

	// -- check component state
	tmpCmp, _ := mutEntity.components.Get(testMutableComponentTypeID)
	cmpState := tmpCmp.(*testMutableComponent)

	assert.Equal(t, 0, cmpState.counter)

	// -- emulate step: 1
	w.Update()
	w.Draw()
	// -- -- result step: 1

	assert.Equal(t, 1, w.systems.Len(), "system created")
	assert.Equal(t, 2, w.entities.Len(), "entities created")
	assert.Equal(t, 1, cmpState.counter, "system inc counter to 1")

	// -- -- modify world: 1
	w.RemoveEntity(secondsEntity)
	assert.Equal(t, 2, w.entities.Len(), "remove queued, but not deleted yet")

	// -- emulate step: 2
	w.Update()
	w.Draw()
	// -- -- result step: 2

	assert.Equal(t, 1, w.entities.Len(), "second entity deleted")
	assert.Equal(t, 2, cmpState.counter, "system inc counter to 2")

	// -- emulate step: 3
	w.Update()
	w.Draw()
	// -- -- result step: 3

	assert.Equal(t, 3, cmpState.counter, "system inc counter to 3")
}
