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
//
// This NOT suitable for making save/load systems
// by snapshotting ecs.World state, snapshot will
// not have any private component state, evaluated
// during world updates
//
// It should be used only for first time initialize
// levels from file, and store back world after editing
// in level-editor (when game support live edit mode)
//
// But anyway in very simple games, with simple state
// and if all your component update state only in public
// properties, it can work as save/load system.
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
	return decodeWorld(r, s)
}

// MarshalToXML will transform StaticWorld into xml
// string (in []byte array)
// you can write result bytes to any storage (like files)
//
// This function return human-readable XML
// that can be edited externally (level editors, notepad, etc..)
func MarshalToXML(sw StaticWorld) ([]byte, error) {
	return marshalXML(sw)
}

// UnmarshalFromXML doing reverse of MarshalToXML
// it converts XML string (in bytes) back to StaticWorld
//
// After getting StaticWorld, you can decode it into
// normal *ecs.World and process any game updates on it
func UnmarshalFromXML(src []byte) (StaticWorld, error) {
	return unmarshalXML(src)
}
