package ecs

import "github.com/go-glx/ecs/props"

type SystemTypeID string

type System interface {
	// TypeID is Unique component type identifier
	// some static UUIDv4 is best option
	// ID used in snapshots for encode/decode
	// world data to XML, JSON, or other format
	TypeID() SystemTypeID
}

type SystemUpdatable interface {
	System
	OnUpdate(w *World)
}

type SystemSyncable interface {
	System
	OnSync(w *World)
}

type SystemConfigurable interface {
	System
	Props() []props.Property
}
