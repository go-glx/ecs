package ecs

import (
	"fmt"
	"reflect"

	"github.com/fe3dback/glx-ecs/ecs/internal/collection"
	"github.com/fe3dback/glx-ecs/ecs/internal/ids"
)

type World struct {
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

// Update should be called every game tick (frame)
// this function will process all world queues
// and call SystemUpdatable.OnUpdate in all registered systems
func (w *World) Update() {
	w.createQueued()
	w.update()
	w.dropQueued()
}

// Sync should be called after every game tick (frame)
// this function will call SystemSyncable.OnSync in all registered systems
// it`s useful for drawing World
func (w *World) Sync() {
	for _, system := range w.systems.Iterate() {
		if systemSyncer, ok := system.(SystemSyncable); ok {
			systemSyncer.OnSync(w)
		}
	}
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

func (w *World) update() {
	for _, system := range w.systems.Iterate() {
		if systemUpdatable, ok := system.(SystemUpdatable); ok {
			systemUpdatable.OnUpdate(w)
		}
	}
}

func (w *World) createQueued() {
	for _, entity := range w.createEntitiesQueue {
		w.entities.Set(entity.id, entity)
	}

	for _, system := range w.createSystemsQueue {
		w.systems.Set(ids.Of(system), system)
	}

	w.createEntitiesQueue = nil
	w.createSystemsQueue = nil
}

func (w *World) dropQueued() {
	for _, systemID := range w.dropSystemsQueue {
		w.systems.Remove(systemID)
	}

	for _, entityID := range w.dropEntitiesQueue {
		w.entities.Remove(entityID)
	}

	w.dropSystemsQueue = nil
	w.dropEntitiesQueue = nil
}
