package component

import "github.com/go-glx/ecs/ecs"

const Node3DTypeID = "Node3D-49f9d0aa8831"

type Node3D struct {
	X float64
	Y float64
	Z float64
}

func NewNode3D(x, y, z float64) *Node3D {
	return &Node3D{
		X: x,
		Y: y,
		Z: z,
	}
}

func (c Node3D) TypeID() ecs.ComponentTypeID {
	return Node3DTypeID
}
