package ecs

import (
	"fmt"

	"github.com/fe3dback/glx-ecs/ecs/internal/ids"
)

// todo: bitmaps filters map/sets
// todo: fast query and filters
// todo: component caches

func FindByComponent[T any](w *World, sample T) map[*Entity]T {
	sampleID := ids.Of(sample)
	result := make(map[*Entity]T)

	for _, ent := range w.entities.Iterate() {
		if cmp, exist := ent.components.Get(sampleID); exist {
			result[ent] = cmp.(T)
		}
	}

	return result
}

func FindByComponentWhere[T any](w *World, sample T, filter func(T) bool) map[*Entity]T {
	result := make(map[*Entity]T)

	for ent, cmp := range FindByComponent(w, sample) {
		if !filter(cmp) {
			continue
		}

		result[ent] = cmp
	}

	return result
}

func MustFindComponentOf[T any](ent *Entity, sample T) T {
	cmp, exist := ent.components.Get(ids.Of(sample))
	if !exist {
		panic(fmt.Errorf("failed get component '%s' from entity '%s': not exist",
			ids.Of(sample),
			ent.String(),
		))
	}

	return cmp.(T)
}
