package ecs

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
