package ecs

import (
	"testing"
)

func BenchmarkFindByComponent(b *testing.B) {
	w := NewWorld()

	ent := NewEntity("entWithNode")
	ent.AddComponent(testCreateMutableComponent())

	w.AddEntity(NewEntity("ent1"))
	w.AddEntity(ent)
	w.AddEntity(NewEntity("ent2"))

	w.Update() // simulate tick, add queued entities

	for i := 0; i < b.N; i++ {
		found := FindComponent[testMutableComponent](w)
		if len(found) != 1 {
			b.Fatalf("expected only one entity with Node2D, found: %d", len(found))
		}
	}
}
