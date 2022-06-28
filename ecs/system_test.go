package ecs

const testMutateSystemTypeID = "MutateSystem-4aa50a28aa2b"

type testMutateSystem struct {
}

func testCreateMutateSystem() *testMutateSystem {
	return &testMutateSystem{}
}

func (t testMutateSystem) TypeID() SystemTypeID {
	return testMutateSystemTypeID
}

func (t *testMutateSystem) OnUpdate(w *World) {
	for _, cmp := range FindComponent[testMutableComponent](w) {
		cmp.counter++
	}
}
