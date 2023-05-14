package ecs

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func testPrepareExtractableWorld() *World {
	r := NewRegistry()
	r.RegisterSystem(testCreateMutateSystem())
	r.RegisterComponent(testCreateMutableComponent())
	r.RegisterComponent(testCreateComplexComponent())

	ent1 := NewEntity("ent1")
	ent1.AddComponent(testCreateMutableComponent())

	ent2 := NewEntity("ent2")
	ent2.AddComponent(testCreateMutableComponent())
	ent2.AddComponent(testCreateComplexComponent())

	return NewWorld(
		r,
		WithInitialSystems(testMutateSystemTypeID),
		WithInitialEntities(ent1, ent2),
	)
}

func TestExtractWorldSystems(t *testing.T) {
	actual := ExtractWorldSystems(testPrepareExtractableWorld())

	assert.Equal(t, testMutateSystemTypeID, string(actual[0].TypeID()))
}

func TestExtractWorldEntities(t *testing.T) {
	actual := ExtractWorldEntities(testPrepareExtractableWorld())

	assert.Len(t, actual, 2)
}

func TestExtractEntityComponents(t *testing.T) {
	ent2 := NewEntity("ent2")
	ent2.AddComponent(testCreateMutableComponent())
	ent2.AddComponent(testCreateComplexComponent())

	actual := ExtractEntityComponents(ent2)
	expected := []Component{testCreateMutableComponent(), testCreateComplexComponent()}

	assert.EqualValues(t, len(expected), len(actual))
}
