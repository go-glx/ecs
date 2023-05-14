package ecs

const testMutableComponentTypeID = "test/MutableComponent"
const testComplexComponentTypeID = "test/ComplexComponent"

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

func (c testMutableComponent) TypeID() ComponentTypeID {
	return testMutableComponentTypeID
}

func (c testComplexComponent) TypeID() ComponentTypeID {
	return testComplexComponentTypeID
}

func (c *testComplexComponent) RequireComponents() []ComponentTypeID {
	return []ComponentTypeID{
		testMutableComponentTypeID,
	}
}
