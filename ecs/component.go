package ecs

type ComponentTypeID string

type Component interface {
	// TypeID is Unique component type identifier
	// some static UUIDv4 is best option
	// ID used in snapshots for encode/decode
	// world data to XML, JSON, or other format
	TypeID() ComponentTypeID
}

type ComponentWithRequirements interface {
	RequireComponents() []ComponentTypeID
}
