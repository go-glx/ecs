package ecs

type testMutableComponent struct {
	counter int
}

type testComplexComponent struct {
	// will require testMutableComponent
}

func testCreateMutableComponent() *testMutableComponent {
	return &testMutableComponent{
		counter: 0,
	}
}

func testCreateComplexComponent() *testComplexComponent {
	return &testComplexComponent{}
}

func (c *testComplexComponent) RequireComponents() []Component {
	return []Component{
		testMutableComponent{},
	}
}
