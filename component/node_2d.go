package component

type Node2D struct {
	x float64
	y float64
}

func NewNode2D(x, y float64) *Node2D {
	return &Node2D{
		x: x,
		y: y,
	}
}
