package ecs

const testMutateSystemTypeID = "internal/MutateSystem"

type testMutateSystem struct {
	filter Filter1[testMutableComponent]
}

func testCreateMutateSystem() *testMutateSystem {
	return &testMutateSystem{}
}

func (t *testMutateSystem) TypeID() SystemTypeID {
	return testMutateSystemTypeID
}

func (t *testMutateSystem) OnInit(w RuntimeWorld) {
	t.filter = NewFilter1[testMutableComponent](w)
}

func (t *testMutateSystem) OnUpdate(w RuntimeWorld) {
	found := t.filter.Find()

	for found.Next() {
		_, cmp := found.Get()
		cmp.counter++
	}
}
