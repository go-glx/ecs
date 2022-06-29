package ecs

func ExtractWorldSystems(w *World) []SystemTypeID {
	return w.systems.Keys()
}

func ExtractWorldEntities(w *World) []*Entity {
	return w.entities.Values()
}

func ExtractEntityComponents(ent *Entity) []Component {
	return ent.components.Values()
}
