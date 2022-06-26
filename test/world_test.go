package test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/fe3dback/glx-ecs/component"
	"github.com/fe3dback/glx-ecs/ecs"
	"github.com/fe3dback/glx-ecs/system"
)

type stabEngine struct {
}

func (s *stabEngine) DeltaTime() float64 {
	return 0.00001
}

func TestWorld(t *testing.T) {
	// -- create world
	ent1 := ecs.NewEntity("ent1")
	assert.Equal(t, "ent1 (1)", ent1.String())
	assert.Equal(t, uint64(1), ent1.ID())
	assert.Panics(t, func() {
		// panic, because nil component
		ent1.AddComponent(nil)
	})
	assert.Panics(t, func() {
		// panic, because not pointer
		ent1.AddComponent(component.Deletable{})
	})
	ent1.AddComponent(component.NewDeletable())

	ent2 := ecs.NewEntity("ent2")
	assert.Equal(t, "ent2 (2)", ent2.String())
	assert.Equal(t, uint64(2), ent2.ID())
	assert.Panics(t, func() {
		// panic, because TimeToLife require Deletable
		ent2.AddComponent(component.NewTimeToLife(3))
	})
	assert.NotPanics(t, func() {
		// now, this should not panic, because we meet requirements
		ent2.AddComponent(component.NewDeletable())
		ent2.AddComponent(component.NewTimeToLife(3))
	})

	// -- create world
	world := ecs.NewWorld()

	// -- - with entities
	world.AddEntity(ent1)
	world.AddEntity(ent2)

	// -- - and systems
	world.AddSystem(system.NewGarbageCollector())

	// -- emulate step

	// -- - step 0
	assert.Equal(t, 0, len(world.Entities()), "queued, but not created yet")

	// -- - step 1
	world.Update()
	assert.Equal(t, 2, len(world.Entities()), "entities created, both alive")

	world.RemoveEntity(ent1)
	assert.Equal(t, 2, len(world.Entities()), "ent1 delete is queued, but not executed")

	// -- - step 2
	world.Update()
	assert.Equal(t, 1, len(world.Entities()), "ent1 deletion executed")

	// -- - step 3
	world.Update()
	assert.Equal(t, 0, len(world.Entities()), "ent2 deleted by TTL component")
}
