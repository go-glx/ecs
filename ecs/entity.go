package ecs

import (
	"fmt"
	"reflect"

	"github.com/go-glx/ecs/ecs/internal/bits"
	"github.com/go-glx/ecs/ecs/internal/collection"
	"github.com/go-glx/ecs/ecs/internal/ids"
)

var entityIDs = ids.NewGlobalID()

type (
	entityID = uint64

	Entity struct {
		name          string   // not unique name (user space)
		prefab        string   // optional prefab name (if entity created from prefab) or empty string
		id            entityID // unique auto increment entity ID
		components    *collection.UniqueCollection[ComponentTypeID, Component]
		componentMask *bits.Matrix
		world         *World // back ptr to world
	}
)

func NewEntity(name string) *Entity {
	return &Entity{
		id:            entityIDs.Next(),
		name:          name,
		prefab:        "",
		components:    collection.NewUniqueCollection[ComponentTypeID, Component](),
		componentMask: bits.NewMatrix(),
	}
}

// -------------------------------------
// PUBLIC API
// -------------------------------------

func (e *Entity) ID() uint64 {
	return e.id
}

func (e *Entity) String() string {
	return fmt.Sprintf("%s (%d)", e.name, e.id)
}

func (e *Entity) Name() string {
	return e.name
}

func (e *Entity) IsPrefab() bool {
	return e.prefab != ""
}

func (e *Entity) PrefabID() string {
	return e.prefab
}

func (e *Entity) AddComponent(cmp Component) {
	e.assertIsNotPrefab()
	e.assertComponentValid(cmp)
	e.components.Set(cmp.TypeID(), cmp)
	e.recalculateHashes()
}

func (e *Entity) RemoveComponent(typeID ComponentTypeID) {
	e.assertIsNotPrefab()
	e.assertComponentCanBeDeleted(typeID)
	e.components.Remove(typeID)
	e.recalculateHashes()
}

// -------------------------------------
// INTERNAL API
// -------------------------------------

// recalculateHashes should be called after entity added to world
// and when components is changed
func (e *Entity) recalculateHashes() {
	if e.world == nil {
		// not attached to world yet, so we will recalculate after attach
		return
	}

	// calculate components bit hash
	e.componentMask.Clear()
	for cmpID := range e.components.Iterate() {
		cmpBits := e.world.registry.bitsRegistry.ComponentBits(string(cmpID))
		e.componentMask.AddComponent(cmpBits)
	}
}

func (e *Entity) assertIsNotPrefab() {
	if !e.IsPrefab() {
		return
	}

	panic(fmt.Errorf("failed modify prefab '%s' entity '%s', id=%d", e.prefab, e.name, e.id))
}

func (e *Entity) assertComponentValid(cmp Component) {
	if cmp == nil {
		panic(fmt.Errorf("invalid component: trying to add nil component to entity '%s'",
			e.String(),
		))
	}

	if reflect.ValueOf(cmp).Kind() != reflect.Ptr {
		panic(fmt.Errorf("invalid component: component '%s' of entity '%s': should by passed as mutable pointer",
			cmp.TypeID(),
			e.String(),
		))
	}

	if reqCmp, ok := cmp.(ComponentWithRequirements); ok {
		for _, requirement := range reqCmp.RequireComponents() {
			if e.components.Has(requirement) {
				continue
			}

			panic(fmt.Errorf(
				"invalid component: component '%s' of entity '%s': requirement of '%s' not satisfied",
				cmp.TypeID(),
				e.String(),
				requirement,
			))
		}
	}
}

func (e *Entity) assertComponentCanBeDeleted(typeID ComponentTypeID) {
	for _, component := range e.components.Iterate() {
		complexComponent, ok := component.(ComponentWithRequirements)
		if !ok {
			continue
		}

		for _, requirement := range complexComponent.RequireComponents() {
			if typeID != requirement {
				continue
			}

			panic(fmt.Errorf(
				"component '%s' of entity '%s': can be deleted, because it required be '%s'",
				typeID,
				e.String(),
				component.TypeID(),
			))
		}
	}
}
