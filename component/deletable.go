package component

import "github.com/fe3dback/glx-ecs/ecs"

const DeletableTypeID = "Deletable-a300548e4f48"

type Deletable struct {
	Alive bool
}

func NewDeletable() *Deletable {
	return &Deletable{
		Alive: true,
	}
}

func (c Deletable) TypeID() ecs.ComponentTypeID {
	return DeletableTypeID
}
