package ecs

import "fmt"

type Prefab struct {
	id      string
	factory func() *Entity
}

func NewPrefab(name string, factory func() *Entity) *Prefab {
	if name == "" {
		panic(fmt.Errorf("can`t use empty string as prefab name"))
	}

	return &Prefab{
		id: name,
		factory: func() *Entity {
			ent := factory()
			ent.prefab = name

			return ent
		},
	}
}
