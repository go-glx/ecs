package ecs

const testMutableComponentTypeID = "MutableComponent-d272cd8609c8"
const testComplexComponentTypeID = "ComplexComponent-ad0cd3e9a49c"

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
