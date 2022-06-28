package system

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/fe3dback/glx-ecs/component"
	"github.com/fe3dback/glx-ecs/ecs"
)

func TestGarbageCollector(t *testing.T) {
	r := ecs.NewRegistry()
	r.RegisterSystem(NewGarbageCollector())
	r.RegisterComponent(component.NewDeletable())
	r.RegisterComponent(component.NewTimeToLife(0))

	// -- create world
	ent1 := ecs.NewEntity("ent1")
	ent1.AddComponent(component.NewDeletable())

	ent2 := ecs.NewEntity("ent2")
	ent2.AddComponent(component.NewDeletable())
	ent2.AddComponent(component.NewTimeToLife(3))

	// -- create world
	world := ecs.NewWorld(
		r,
		ecs.WithInitialSystems(GarbageCollectorTypeID),
		ecs.WithInitialEntities(ent1, ent2),
	)

	assert.Equal(t, 2, len(world.Entities()), "entities created, both alive")

	// -- emulate step
	// -- - step 1
	world.Update()
	assert.Equal(t, 2, len(world.Entities()), "both alive for now")

	world.RemoveEntity(ent1)
	assert.Equal(t, 2, len(world.Entities()), "ent1 delete is queued, but not executed")

	// -- - step 2
	world.Update()
	assert.Equal(t, 1, len(world.Entities()), "ent1 deletion executed")

	// -- - step 3
	world.Update()
	assert.Equal(t, 0, len(world.Entities()), "ent2 deleted by TTL component")
}
