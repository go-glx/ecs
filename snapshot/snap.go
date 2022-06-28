package snapshot

import "github.com/fe3dback/glx-ecs/ecs"

// Static snapshot of world
// can be marshaled/unmarshalled to another data view (json, xml, etc..)
type Static struct {
	// todo
}

func Create(w *ecs.World) Static {
	return Static{}
}

func Restore(s Static) *ecs.World {
	return nil
	// return ecs.NewWorld()
}
