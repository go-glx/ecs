package ecs

import (
	"fmt"
	"reflect"

	"github.com/fe3dback/glx-ecs/ecs/internal/collection"
	"github.com/fe3dback/glx-ecs/ecs/internal/ids"
)

type EngineState interface {
	DeltaTime() float64
}

type World struct {
	state    EngineState
	entities *collection.UniqueCollection[entityID, *Entity]
	systems  *collection.UniqueCollection[ids.ObjectID, System]

	createEntitiesQueue []*Entity
	createSystemsQueue  []System
	dropEntitiesQueue   []entityID
	dropSystemsQueue    []ids.ObjectID
}

func NewWorld() *World {
	return &World{
		entities: collection.NewUniqueCollection[entityID, *Entity](),
		systems:  collection.NewUniqueCollection[ids.ObjectID, System](),

		createEntitiesQueue: make([]*Entity, 0),
		createSystemsQueue:  make([]System, 0),
		dropEntitiesQueue:   make([]entityID, 0),
		dropSystemsQueue:    make([]ids.ObjectID, 0),
	}
}

// AddSystem will not add System immediately after call,
// instead it will add System to queue,
// all systems will be created right before world Update
func (w *World) AddSystem(system System) {
	w.assertSystemValid(system, "AddSystem")
	w.createSystemsQueue = append(w.createSystemsQueue, system)
}

// RemoveSystem will not remove system immediately after call
// this just mark system as deleted
// all systems will be deleted after world Update
func (w *World) RemoveSystem(system System) {
	w.dropSystemsQueue = append(w.dropSystemsQueue, ids.Of(system))
}

// AddEntity will not add Entity immediately after call,
// instead it will add Entity to queue,
// all entities will be created right before world Update
func (w *World) AddEntity(entity *Entity) {
	w.createEntitiesQueue = append(w.createEntitiesQueue, entity)
}

// RemoveEntity will not remove entity immediately after call
// this just mark entity as deleted
// all entities will be deleted after world Update
func (w *World) RemoveEntity(entity *Entity) {
	w.dropEntitiesQueue = append(w.dropEntitiesQueue, entity.id)
}

// Entities should be used only for inspection
// and read-only engine wide operations
// all normal work with Entities, should be done
// wia ECS systems
func (w *World) Entities() []Entity {
	list := make([]Entity, 0, w.entities.Len())

	for _, e := range w.entities.Iterate() {
		list = append(list, *e)
	}

	return list
}

func (w *World) assertSystemValid(system System, inAction string) {
	if system == nil {
		panic(fmt.Errorf("failed %s: trying to add nil system to world",
			inAction,
		))
	}

	if reflect.ValueOf(system).Kind() != reflect.Ptr {
		panic(fmt.Errorf("failed %s: system '%s': should by passed as mutable pointer",
			inAction,
			ids.Of(system),
		))
	}
}
