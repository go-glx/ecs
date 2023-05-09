package ecs

import (
	"fmt"

	"github.com/go-glx/ecs/ecs/internal/collection"
)

type World struct {
	registry *Registry

	entities     *collection.UniqueCollection[entityID, *Entity]
	systems      *collection.UniqueCollection[SystemTypeID, System]
	systemsOrder []SystemTypeID

	createEntitiesQueue []*Entity
	dropEntitiesQueue   []entityID
	createSystemsQueue  []SystemTypeID
	dropSystemsQueue    []SystemTypeID
}

func NewWorld(registry *Registry, initializers ...WorldInitializer) *World {
	world := &World{
		registry: registry,

		entities:     collection.NewUniqueCollection[entityID, *Entity](),
		systems:      collection.NewUniqueCollection[SystemTypeID, System](),
		systemsOrder: make([]SystemTypeID, 0, 128),

		createEntitiesQueue: make([]*Entity, 0),
		dropEntitiesQueue:   make([]entityID, 0),
		createSystemsQueue:  make([]SystemTypeID, 0),
		dropSystemsQueue:    make([]SystemTypeID, 0),
	}

	for _, init := range initializers {
		init(world)
	}

	return world
}

// InitializeWith allows to add some more initialization logic
// into world creation process, this MUST be called
// ONLY before first Update. (right after NewWorld)
func (w *World) InitializeWith(initializers ...WorldInitializer) {
	for _, init := range initializers {
		init(w)
	}
}

// AddSystem will not add System immediately after call,
// instead it will add System to queue,
// all systems will be created right before world Update
func (w *World) AddSystem(systemTypeID SystemTypeID) {
	if w.systems.Has(systemTypeID) {
		panic(fmt.Errorf("system '%s' already added to ECS world", systemTypeID))
	}

	w.createSystemsQueue = append(w.createSystemsQueue, systemTypeID)
}

// RemoveSystem will not remove system immediately after call
// this just mark system as deleted
// all systems will be deleted after world Update
func (w *World) RemoveSystem(systemTypeID SystemTypeID) {
	w.dropSystemsQueue = append(w.dropSystemsQueue, systemTypeID)
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

func (w *World) addEntityInternal(entity *Entity) {
	w.assertEntityHasKnownComponents(entity)
	w.entities.Set(entity.id, entity)
}

func (w *World) addSystemInternal(systemTypeID SystemTypeID) {
	w.systems.Set(systemTypeID, w.registry.getSystemOfType(systemTypeID))
	w.systemsOrder = append(w.systemsOrder, systemTypeID)
}

func (w *World) removeSystemInternal(systemTypeID SystemTypeID) {
	w.systems.Remove(systemTypeID)

	newSystemsOrder := make([]SystemTypeID, 0, len(w.systemsOrder))
	for _, id := range w.systemsOrder {
		if id == systemTypeID {
			continue
		}

		newSystemsOrder = append(newSystemsOrder, id)
	}

	w.systemsOrder = newSystemsOrder
}

func (w *World) assertEntityHasKnownComponents(entity *Entity) {
	for typeID := range entity.components.Iterate() {
		// function will panic, if component not known
		_ = w.registry.getDefaultComponentOfType(typeID)
	}
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
	w.systems.IterateInOrder(w.systemsOrder, func(id SystemTypeID, system System) {
		if systemSyncer, ok := system.(SystemSyncable); ok {
			systemSyncer.OnSync(w)
		}
	})
}

// IterateOverSystems can be used to iterate over all world systems
// and execute some own custom logic on it
func (w *World) IterateOverSystems(itt func(systemID string, system System)) {
	w.systems.IterateInOrder(w.systemsOrder, func(id SystemTypeID, system System) {
		itt(string(id), system)
	})
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

func (w *World) update() {
	w.systems.IterateInOrder(w.systemsOrder, func(id SystemTypeID, system System) {
		if systemUpdatable, ok := system.(SystemUpdatable); ok {
			systemUpdatable.OnUpdate(w)
		}
	})
}

func (w *World) createQueued() {
	for _, entity := range w.createEntitiesQueue {
		w.addEntityInternal(entity)
	}

	for _, system := range w.createSystemsQueue {
		w.addSystemInternal(system)
	}

	w.createEntitiesQueue = nil
	w.createSystemsQueue = nil
}

func (w *World) dropQueued() {
	for _, systemID := range w.dropSystemsQueue {
		w.removeSystemInternal(systemID)
	}

	for _, entityID := range w.dropEntitiesQueue {
		w.entities.Remove(entityID)
	}

	w.dropSystemsQueue = nil
	w.dropEntitiesQueue = nil
}
