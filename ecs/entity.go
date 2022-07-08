package ecs

import (
	"fmt"
	"reflect"

	"github.com/go-glx/ecs/ecs/internal/collection"
	"github.com/go-glx/ecs/ecs/internal/ids"
)

var entityIDs = ids.NewGlobalID()

type (
	entityID = uint64

	Entity struct {
		name       string
		id         entityID
		components *collection.UniqueCollection[ComponentTypeID, Component]
	}
)

func NewEntity(name string) *Entity {
	return &Entity{
		id:         entityIDs.Next(),
		name:       name,
		components: collection.NewUniqueCollection[ComponentTypeID, Component](),
	}
}

func (e *Entity) ID() uint64 {
	return e.id
}

func (e *Entity) String() string {
	return fmt.Sprintf("%s (%d)", e.name, e.id)
}

func (e *Entity) Name() string {
	return e.name
}

func (e *Entity) AddComponent(cmp Component) {
	e.assertComponentValid(cmp)
	e.components.Set(cmp.TypeID(), cmp)
}

func (e *Entity) RemoveComponent(typeID ComponentTypeID) {
	e.components.Remove(typeID)
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
