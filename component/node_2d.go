package component

import "github.com/go-glx/ecs/ecs"

const Node2DTypeID = "internal/Node2D"

type Node2D struct {
	X float64
	Y float64
}

func NewNode2D(x, y float64) *Node2D {
	return &Node2D{
		X: x,
		Y: y,
	}
}

func (c Node2D) TypeID() ecs.ComponentTypeID {
	return Node2DTypeID
}
