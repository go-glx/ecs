package snapshot

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/go-glx/ecs/component"
	"github.com/go-glx/ecs/ecs"
	"github.com/go-glx/ecs/system"
)

func TestSnapshot(t *testing.T) {
	// world -> xml
	worldOrigin := testCreateWorld()
	worldOriginStatic := Create(worldOrigin)
	worldOriginXML, err := MarshalToXML(worldOriginStatic)
	assert.NoError(t, err)

	// xml -> world
	worldBackStatic, err := UnmarshalFromXML(worldOriginXML)
	assert.NoError(t, err)
	worldBack := Restore(testCreateRegistry(), worldBackStatic)

	// smoke assert
	// all actual asserts written in encode/decode marshal/unmarshal files
	assert.Equal(t, len(worldOrigin.Entities()), len(worldBack.Entities()))
}

func testCreateRegistry() *ecs.Registry {
	r := ecs.NewRegistry()
	r.RegisterSystem(system.NewGarbageCollector())
	r.RegisterComponent(component.NewNode2D(1, 2))
	r.RegisterComponent(component.NewDeletable())

	return r
}

func testCreateWorld() *ecs.World {
	ent1 := ecs.NewEntity("ent1")
	ent1.AddComponent(component.NewNode2D(5, 10))
	ent1.AddComponent(component.NewDeletable())

	ent2 := ecs.NewEntity("ent2")
	ent2.AddComponent(component.NewNode2D(4, 7))

	return ecs.NewWorld(
		testCreateRegistry(),
		ecs.WithInitialEntities(ent1, ent2),
		ecs.WithInitialSystems(system.GarbageCollectorTypeID),
	)
}

// should match testCreateWorld data
func testCreateStaticWorld() StaticWorld {
	return StaticWorld{
		Systems: []StaticSystem{
			{
				TypeID: system.GarbageCollectorTypeID,
			},
		},
		Entities: []StaticEntity{
			{
				Name: "ent1",
				Components: []StaticComponent{
					{
						TypeID: "Deletable-a300548e4f48",
						Order:  0,
						Props: []StaticComponentProperty{
							{Name: "Alive", Value: "true"},
						},
					},
					{
						TypeID: "Node2D-7c40b8e315a5",
						Order:  1,
						Props: []StaticComponentProperty{
							{Name: "X", Value: "5"},
							{Name: "Y", Value: "10"},
						},
					},
				},
			},
			{
				Name: "ent2",
				Components: []StaticComponent{
					{
						TypeID: "Node2D-7c40b8e315a5",
						Order:  0,
						Props: []StaticComponentProperty{
							{Name: "X", Value: "4"},
							{Name: "Y", Value: "7"},
						},
					},
				},
			},
		},
	}
}
