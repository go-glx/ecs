package ecs

type testMutateSystem struct {
}

func testCreateMutateSystem() *testMutateSystem {
	return &testMutateSystem{}
}

func (t *testMutateSystem) OnUpdate(w *World) {
	for _, cmp := range FindByComponent(w, &testMutableComponent{}) {
		cmp.counter++
	}
}
