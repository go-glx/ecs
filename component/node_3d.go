package component

type Node3D struct {
	x float64
	y float64
	z float64
}

func NewNode3D(x, y, z float64) *Node3D {
	return &Node3D{
		x: x,
		y: y,
		z: z,
	}
}
