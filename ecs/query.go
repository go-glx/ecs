package ecs

import (
	"fmt"
)

// todo: bitmaps filters map/sets
// todo: fast query and filters
// todo: component caches

func FindComponent[T Component](w *World) map[*Entity]*T {
	var TType T
	findResult := make(map[*Entity]*T)

	for _, ent := range w.entities.Iterate() {
		if cmp, exist := ent.components.Get(TType.TypeID()); exist {
			findResult[ent] = cmp.(any).(*T)
		}
	}

	return findResult
}

func FindComponentWhere[T Component](w *World, filter func(*T) bool) map[*Entity]*T {
	result := make(map[*Entity]*T)

	for ent, cmp := range FindComponent[T](w) {
		if !filter(cmp) {
			continue
		}

		result[ent] = cmp
	}

	return result
}

func MustFindComponentOf[T Component](ent *Entity) *T {
	var TType T

	cmp, exist := ent.components.Get(TType.TypeID())
	if !exist {
		panic(fmt.Errorf("failed get component '%s' from entity '%s': not exist",
			TType.TypeID(),
			ent.String(),
		))
	}

	return cmp.(any).(*T)
}
