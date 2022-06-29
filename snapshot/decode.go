package snapshot

import (
	"fmt"

	"github.com/mitchellh/mapstructure"

	"github.com/fe3dback/glx-ecs/ecs"
)

func decodeWorld(r *ecs.Registry, sw StaticWorld) *ecs.World {
	return ecs.NewWorld(
		r,
		ecs.WithInitialSystems(decodeSystems(sw)...),
		ecs.WithInitialEntities(decodeEntities(r, sw)...),
	)
}

func decodeSystems(sw StaticWorld) []ecs.SystemTypeID {
	result := make([]ecs.SystemTypeID, 0, len(sw.Systems))

	for _, system := range sw.Systems {
		result = append(result, ecs.SystemTypeID(system.TypeID))
	}

	return result
}

func decodeEntities(r *ecs.Registry, sw StaticWorld) []*ecs.Entity {
	result := make([]*ecs.Entity, 0, len(sw.Entities))

	for _, entity := range sw.Entities {
		result = append(result, decodeEntity(r, entity))
	}

	return result
}

func decodeEntity(r *ecs.Registry, sw StaticEntity) *ecs.Entity {
	ent := ecs.NewEntity(sw.Name)

	for _, component := range sw.Components {
		ent.AddComponent(decodeComponent(r, component))
	}

	return ent
}

func decodeComponent(r *ecs.Registry, sw StaticComponent) ecs.Component {
	defComponent, exist := r.CreateComponentWithDefaultValues(ecs.ComponentTypeID(sw.TypeID))
	if !exist {
		panic(fmt.Errorf("failed decode world, because component '%s' not registered",
			sw.TypeID,
		))
	}

	props := make(map[string]string, len(sw.Props))
	for _, resProp := range sw.Props {
		props[resProp.Name] = resProp.Value
	}

	decoder, err := mapstructure.NewDecoder(&mapstructure.DecoderConfig{
		ErrorUnused:      false,
		ZeroFields:       false,
		WeaklyTypedInput: true,
		Squash:           false,
		Result:           &defComponent,
	})
	if err != nil {
		panic(fmt.Errorf("failed create decoder for decoding '%s' component: %w", sw.TypeID, err))
	}

	err = decoder.Decode(props)
	if err != nil {
		panic(fmt.Errorf("failed decode component '%s' props: %w (props: %#v)", sw.TypeID, err, props))
	}

	return defComponent
}
