package ecs

import (
	"fmt"

	"github.com/fe3dback/glx-ecs/ecs/internal/ids"
)

// todo: bitmaps filters map/sets
// todo: fast query and filters
// todo: component caches

func FindComponent[T any](w *World) map[*Entity]*T {
	var TType T
	sampleID := ids.Of(TType)
	findResult := make(map[*Entity]*T)

	for _, ent := range w.entities.Iterate() {
		if cmp, exist := ent.components.Get(sampleID); exist {
			findResult[ent] = cmp.(*T)
		}
	}

	return findResult
}

func FindComponentWhere[T any](w *World, filter func(*T) bool) map[*Entity]*T {
	result := make(map[*Entity]*T)

	for ent, cmp := range FindComponent[T](w) {
		if !filter(cmp) {
			continue
		}

		result[ent] = cmp
	}

	return result
}

func MustFindComponentOf[T any](ent *Entity) *T {
	var TType T
	cmp, exist := ent.components.Get(ids.Of(TType))
	if !exist {
		panic(fmt.Errorf("failed get component '%s' from entity '%s': not exist",
			ids.Of(TType),
			ent.String(),
		))
	}

	return cmp.(*T)
}
