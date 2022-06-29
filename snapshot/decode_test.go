package snapshot

import (
	"fmt"
	"sort"
	"testing"

	"github.com/fatih/structs"
	"github.com/stretchr/testify/assert"

	"github.com/fe3dback/glx-ecs/ecs"
)

func Test_decodeWorld(t *testing.T) {
	worldOrigin := testCreateWorld()
	worldBack := decodeWorld(testCreateRegistry(), testCreateStaticWorld())

	// assert same

	// --- systems
	origSystems := ecs.ExtractWorldSystems(worldOrigin)
	backSystems := ecs.ExtractWorldSystems(worldBack)

	assert.Equal(t, len(origSystems), len(backSystems))
	assert.EqualValues(t, origSystems, backSystems)

	// --- entities
	origEntities := ecs.ExtractWorldEntities(worldOrigin)
	backEntities := ecs.ExtractWorldEntities(worldBack)

	// lite check for count
	assert.Equal(t, len(origEntities), len(backEntities))
	if len(origEntities) != len(backEntities) {
		return
	}

	// if we have equal cnt of entities
	// sort them all, and check each pair
	testSortEntities(origEntities)
	testSortEntities(backEntities)

	for h := 0; h < len(origEntities); h++ {
		entA := origEntities[h]
		entB := backEntities[h]

		assert.Equal(t, entA.Name(), entB.Name())

		entACmp := ecs.ExtractEntityComponents(entA)
		entBCmp := ecs.ExtractEntityComponents(entB)

		assert.Equal(t, len(entACmp), len(entBCmp))
		if len(entACmp) != len(entBCmp) {
			return
		}

		testSortComponents(entACmp)
		testSortComponents(entBCmp)

		// compare component pairs
		for j := 0; j < len(entACmp); j++ {
			cmpA := entACmp[j]
			cmpB := entBCmp[j]

			assert.Equal(t, cmpA.TypeID(), cmpB.TypeID())
			assert.Equal(
				t,
				fmt.Sprintf("%v", testExtractComponentValues(cmpA)),
				fmt.Sprintf("%v", testExtractComponentValues(cmpB)),
			)
		}
	}
}

func testExtractComponentValues(c ecs.Component) map[string]string {
	res := make(map[string]string)

	for _, field := range structs.Fields(c) {
		if !field.IsExported() {
			continue
		}

		res[field.Name()] = fmt.Sprintf("%v", field.Value())
	}

	return res
}

func testSortEntities(list []*ecs.Entity) {
	sort.Slice(list, func(i, j int) bool {
		a, b := list[i], list[j]

		if a.Name() != b.Name() {
			return a.Name() <= b.Name()
		}

		aCmps := ecs.ExtractEntityComponents(a)
		bCmps := ecs.ExtractEntityComponents(b)

		if len(aCmps) != len(bCmps) {
			return len(aCmps) <= len(bCmps)
		}

		testSortComponents(aCmps)
		testSortComponents(bCmps)

		for h := 0; h < len(aCmps); h++ {
			cmpA := aCmps[h]
			cmpB := bCmps[h]

			if cmpA.TypeID() != cmpB.TypeID() {
				return cmpA.TypeID() <= cmpB.TypeID()
			}
		}

		return false
	})
}

func testSortComponents(list []ecs.Component) {
	sort.Slice(list, func(i, j int) bool {
		a, b := list[i], list[j]

		if a.TypeID() != b.TypeID() {
			return a.TypeID() <= b.TypeID()
		}

		return false
	})
}
