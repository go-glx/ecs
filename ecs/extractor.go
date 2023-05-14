package ecs

// ExtractWorldSystems is utility function for read-only dump
// do not use it directly, if you not 100% sure.
// this function primary created for encode world into snapshot
func ExtractWorldSystems(w *World) []System {
	systems := make([]System, 0, len(w.systemsOrder))

	w.IterateOverSystems(func(systemID string, system System) {
		systems = append(systems, system)
	})

	return systems
}

// ExtractWorldEntities is utility function for read-only dump
// do not use it directly, if you not 100% sure.
// this function primary created for encode world into snapshot
func ExtractWorldEntities(w *World) []*Entity {
	return w.entities.Values()
}

// ExtractEntityComponents is utility function for read-only dump
// do not use it directly, if you not 100% sure.
// this function primary created for encode world into snapshot
func ExtractEntityComponents(ent *Entity) []Component {
	return ent.components.Values()
}
