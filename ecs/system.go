package ecs

import "github.com/go-glx/ecs/props"

type SystemTypeID string

type (
	// RuntimeWorld is minimal world interface for interacting from systems update/sync/etc..
	RuntimeWorld interface {
		AddEntity(entity *Entity)
		AddPrefabEntity(prefabID string)
		RemoveEntity(entity *Entity)
	}
)

type System interface {
	// TypeID is Unique component type identifier
	// some static UUIDv4 is best option
	// ID used in snapshots for encode/decode
	// world data to XML, JSON, or other format
	TypeID() SystemTypeID
}

type SystemInitializable interface {
	System
	OnInit(w RuntimeWorld)
}

type SystemUpdatable interface {
	System
	OnUpdate(w RuntimeWorld)
}

type SystemDrawable interface {
	System
	OnDraw(w RuntimeWorld)
}

type SystemConfigurable interface {
	System
	Props() []props.Property
}
