package snapshot

import "github.com/fe3dback/glx-ecs/ecs"

// StaticWorld snapshot of world
// can be marshaled/unmarshalled to another data view (json, xml, etc..)
type (
	StaticWorld struct {
		Systems  []StaticSystem `xml:"systems>system" json:"systems"`
		Entities []StaticEntity `xml:"entities>entity" json:"entities"`
	}

	StaticSystem struct {
		TypeID string `xml:"id,attr" json:"id"`
	}

	StaticEntity struct {
		Name       string            `xml:"name,attr" json:"name"`
		Components []StaticComponent `xml:"components>component,omitempty" json:"components"`
	}

	StaticComponent struct {
		TypeID string                    `xml:"id,attr" json:"id"`
		Props  []StaticComponentProperty `xml:"props>prop,omitempty" json:"props"`
	}

	StaticComponentProperty struct {
		Name  string `xml:"name,attr" json:"name"`
		Value string `xml:"value,attr" json:"value"`
	}
)

// Create will transfer ecs.World to StaticWorld
// (world snapshot). Next you can marshal StaticWorld
// to other []byte format for storing in files
// for example see MarshalToXML
func Create(w *ecs.World) StaticWorld {
	return encodeWorld(w)
}

// Restore will create new World instance from
// snapshot (StaticWorld) object.
// previously you can unmarshal some []byte format
// to StaticWorld, for example from XML, see UnmarshalFromXML
//
// You should provide ecs.Registry that known all systems
// and components, used previously for creating snapshot
//
// Registry contain default values for all components,
// when you add new component fields, that not exist in snapshot
// default value will be used from registry
func Restore(r *ecs.Registry, s StaticWorld) *ecs.World {
	return nil
	// return ecs.NewWorld()
}

func MarshalToXML(sw StaticWorld) ([]byte, error) {
	return marshalXML(sw)
}

func UnmarshalFromXML(src []byte) (StaticWorld, error) {
	return StaticWorld{}, nil // todo
}
