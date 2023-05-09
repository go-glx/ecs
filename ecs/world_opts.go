package ecs

import "fmt"

type WorldInitializer = func(*World)

func WithInitialSystems(systems ...SystemTypeID) WorldInitializer {
	return func(world *World) {
		for _, system := range systems {
			world.addSystemInternal(system)
		}
	}
}

func WithInitialEntities(entities ...*Entity) WorldInitializer {
	return func(world *World) {
		for _, entity := range entities {
			world.addEntityInternal(entity)
		}
	}
}

func WithInitialSystemProperties(system SystemTypeID, initialProps map[string]string) WorldInitializer {
	return func(world *World) {
		sys, exist := world.systems.Get(system)
		if !exist {
			panic(fmt.Errorf("failed apply default system properties: system '%s' not exist in world (call WithInitialSystems first)", system))
		}

		configurableSystem, ok := sys.(SystemConfigurable)
		if !ok {
			return
		}

		for _, property := range configurableSystem.Props() {
			for propName, propRawValue := range initialProps {
				if propName == property.Name() {
					property.Decode(propRawValue)
					break
				}
			}
		}
	}
}
