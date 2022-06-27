package system

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/fe3dback/glx-ecs/component"
	"github.com/fe3dback/glx-ecs/ecs"
)

func TestGarbageCollector(t *testing.T) {
	// -- create world
	ent1 := ecs.NewEntity("ent1")
	ent1.AddComponent(component.NewDeletable())

	ent2 := ecs.NewEntity("ent2")
	ent2.AddComponent(component.NewDeletable())
	ent2.AddComponent(component.NewTimeToLife(3))

	// -- create world
	world := ecs.NewWorld()

	// -- - with entities
	world.AddEntity(ent1)
	world.AddEntity(ent2)

	// -- - and systems
	world.AddSystem(NewGarbageCollector())

	// -- emulate step

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
