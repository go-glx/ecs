package ecs

func ExtractWorldSystems(w *World) []System {
	systems := make([]System, 0, len(w.systemsOrder))

	w.IterateOverSystems(func(systemID string, system System) {
		systems = append(systems, system)
	})

	return systems
}

func ExtractWorldEntities(w *World) []*Entity {
	return w.entities.Values()
}

func ExtractEntityComponents(ent *Entity) []Component {
	return ent.components.Values()
}
