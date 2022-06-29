package snapshot

import (
	"fmt"
	"sort"

	"github.com/fatih/structs"

	"github.com/fe3dback/glx-ecs/ecs"
)

func encodeWorld(w *ecs.World) StaticWorld {
	return StaticWorld{
		Systems:  encodeSystems(ecs.ExtractWorldSystems(w)),
		Entities: encodeEntities(ecs.ExtractWorldEntities(w)),
	}
}

func encodeSystems(systems []ecs.SystemTypeID) []StaticSystem {
	encoded := make([]StaticSystem, 0, len(systems))

	for _, system := range systems {
		encoded = append(encoded, StaticSystem{
			TypeID: string(system),
		})
	}

	sort.Slice(encoded, func(i, j int) bool {
		return encoded[i].TypeID <= encoded[j].TypeID
	})

	return encoded
}

func encodeEntities(entities []*ecs.Entity) []StaticEntity {
	encoded := make([]StaticEntity, 0, len(entities))

	for _, entity := range entities {
		encoded = append(encoded, encodeEntity(entity))
	}

	sort.Slice(encoded, func(i, j int) bool {
		return encoded[i].Name <= encoded[j].Name
	})

	return encoded
}

func encodeEntity(entity *ecs.Entity) StaticEntity {
	return StaticEntity{
		Name:       entity.Name(),
		Components: encodeComponents(ecs.ExtractEntityComponents(entity)),
	}
}

func encodeComponents(components []ecs.Component) []StaticComponent {
	encoded := make([]StaticComponent, 0, len(components))

	for _, component := range components {
		encoded = append(encoded, encodeComponent(component))
	}

	sort.Slice(encoded, func(i, j int) bool {
		return encoded[i].TypeID <= encoded[j].TypeID
	})

	return encoded
}

func encodeComponent(component ecs.Component) StaticComponent {
	return StaticComponent{
		TypeID: string(component.TypeID()),
		Props:  encodeComponentProps(component),
	}
}

func encodeComponentProps(c ecs.Component) []StaticComponentProperty {
	props := make([]StaticComponentProperty, 0)

	for _, field := range structs.Fields(c) {
		if !field.IsExported() {
			continue
		}

		props = append(props, StaticComponentProperty{
			Name:  field.Name(),
			Value: fmt.Sprintf("%v", field.Value()),
		})
	}

	sort.Slice(props, func(i, j int) bool {
		if props[i].Name != props[j].Name {
			return props[i].Name <= props[j].Name
		}

		return props[i].Value <= props[j].Value
	})

	return props
}
