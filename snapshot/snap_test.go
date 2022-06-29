package snapshot

import (
	"github.com/fe3dback/glx-ecs/component"
	"github.com/fe3dback/glx-ecs/ecs"
	"github.com/fe3dback/glx-ecs/system"
)

func testCreateWorld() *ecs.World {
	r := ecs.NewRegistry()
	r.RegisterSystem(system.NewGarbageCollector())
	r.RegisterComponent(component.NewNode2D(1, 2))
	r.RegisterComponent(component.NewDeletable())

	ent1 := ecs.NewEntity("ent1")
	ent1.AddComponent(component.NewNode2D(5, 10))
	ent1.AddComponent(component.NewDeletable())

	ent2 := ecs.NewEntity("ent2")
	ent2.AddComponent(component.NewNode2D(4, 7))

	return ecs.NewWorld(
		r,
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
						Props: []StaticComponentProperty{
							{Name: "Alive", Value: "true"},
						},
					},
					{
						TypeID: "Node2D-7c40b8e315a5",
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
