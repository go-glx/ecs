package ecs

import (
	"fmt"
	"reflect"

	"github.com/fe3dback/glx-ecs/ecs/internal/collection"
	"github.com/fe3dback/glx-ecs/ecs/internal/ids"
)

var entityIDs = ids.NewGlobalID()

type (
	entityID = uint64

	Entity struct {
		name       string
		id         entityID
		components *collection.UniqueCollection[ids.ObjectID, Component]
	}
)

func NewEntity(name string) *Entity {
	return &Entity{
		name:       name,
		id:         entityIDs.Next(),
		components: collection.NewUniqueCollection[ids.ObjectID, Component](),
	}
}

func (e *Entity) ID() uint64 {
	return e.id
}

func (e *Entity) String() string {
	return fmt.Sprintf("%s (%d)", e.name, e.id)
}

func (e *Entity) AddComponent(cmp Component) {
	e.assertComponentValid(cmp, "AddComponent")
	e.components.Set(ids.Of(cmp), cmp)
}

func (e *Entity) RemoveComponent(sample Component) {
	e.components.Remove(ids.Of(sample))
}

func (e *Entity) assertComponentValid(cmp Component, inAction string) {
	if cmp == nil {
		panic(fmt.Errorf("failed %s: trying to add nil component to entity '%s'",
			inAction,
			e.String(),
		))
	}

	if reflect.ValueOf(cmp).Kind() != reflect.Ptr {
		panic(fmt.Errorf("failed %s: '%s' component '%s': should by passed as mutable pointer",
			inAction,
			e.String(),
			ids.Of(cmp),
		))
	}

	if reqCmp, ok := cmp.(ComponentWithRequirements); ok {
		for _, requirement := range reqCmp.RequireComponents() {
			requirementComponentID := ids.Of(requirement)
			if e.components.Has(requirementComponentID) {
				continue
			}

			panic(fmt.Errorf(
				"failed %s: component '%s' of entity '%s': requirement of '%s' not satisfied",
				inAction,
				ids.Of(cmp),
				e.String(),
				requirementComponentID,
			))
		}
	}
}
