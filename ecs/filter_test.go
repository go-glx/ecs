package ecs

import (
	"testing"
)

func BenchmarkFindByComponent(b *testing.B) {
	r := NewRegistry()
	r.RegisterComponent(testMutableComponent{})
	w := NewWorld(r)

	ent := NewEntity("entWithNode")
	ent.AddComponent(testCreateMutableComponent())

	w.AddEntity(NewEntity("ent1"))
	w.AddEntity(ent)
	w.AddEntity(NewEntity("ent2"))

	w.Update() // simulate tick, add queued entities

	flt := NewFilter1[testMutableComponent](w)

	for i := 0; i < b.N; i++ {
		found := flt.Find()
		count := 0

		for found.Next() {
			count++
		}

		if count != 1 {
			b.Fatalf("expected only one entity with Node2D, found: %d", count)
		}
	}
}
