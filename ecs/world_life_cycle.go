package ecs

import "github.com/fe3dback/glx-ecs/ecs/internal/ids"

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
