package ecs

type Component interface {
}

type ComponentWithRequirements interface {
	RequireComponents() []Component
}
