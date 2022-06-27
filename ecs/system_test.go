package ecs

type testMutateSystem struct {
}

func testCreateMutateSystem() *testMutateSystem {
	return &testMutateSystem{}
}

func (t *testMutateSystem) OnUpdate(w *World) {
	for _, cmp := range FindComponent[testMutableComponent](w) {
		cmp.counter++
	}
}
