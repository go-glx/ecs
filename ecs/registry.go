package ecs

import (
	"fmt"
	"reflect"
)

type Registry struct {
	components map[ComponentTypeID]Component
	systems    map[SystemTypeID]System
}

func NewRegistry() *Registry {
	return &Registry{
		components: make(map[ComponentTypeID]Component),
		systems:    make(map[SystemTypeID]System),
	}
}

// RegisterComponent will store example Component values
// as default component state.
//
// When World restored from snapshot, all unknown fields will
// be written as default values from default state
func (r *Registry) RegisterComponent(component Component) {
	r.assertTypeIDNotRegistered(string(component.TypeID()), component, "RegisterComponent")
	r.components[component.TypeID()] = component
}

// CreateComponentWithDefaultValues will return gold copy of
// component with provided typeID (default values)
//
// It is useful in snapshot tools for restoring World from
// file storage
func (r *Registry) CreateComponentWithDefaultValues(typeID ComponentTypeID) (Component, bool) {
	cmp, exist := r.components[typeID]
	if !exist {
		return nil, false
	}

	var createdCopy Component
	createdCopy = reflect.New(reflect.ValueOf(cmp).Elem().Type()).Interface().(Component)

	return createdCopy, exist
}

// RegisterSystem will store System by their typeID
//
// This allows ECS marshal/unmarshal world from
// snapshot automatically
func (r *Registry) RegisterSystem(system System) {
	r.assertTypeIDNotRegistered(string(system.TypeID()), system, "RegisterSystem")
	r.assertSystemValid(system)
	r.systems[system.TypeID()] = system
}

func (r *Registry) assertTypeIDNotRegistered(id string, obj any, action string) {
	for systemID, system := range r.systems {
		if systemID == SystemTypeID(id) {
			panic(fmt.Errorf(
				"%s: failed register '%s': already has `system` '%s' with this ID '%s'",
				action,
				reflect.TypeOf(obj).String(),
				reflect.TypeOf(system).String(),
				systemID,
			))
		}
	}

	for componentID, component := range r.components {
		if componentID == ComponentTypeID(id) {
			panic(fmt.Errorf(
				"%s: failed register '%s': already has `component` '%s' with this ID '%s'",
				action,
				reflect.TypeOf(obj).String(),
				reflect.TypeOf(component).String(),
				componentID,
			))
		}
	}
}

func (r *Registry) getSystemOfType(typeID SystemTypeID) System {
	if system, exist := r.systems[typeID]; exist {
		return system
	}

	panic(fmt.Errorf("system '%s' not exist in registry", typeID))
}

func (r *Registry) getDefaultComponentOfType(typeID ComponentTypeID) Component {
	if component, exist := r.components[typeID]; exist {
		return component
	}

	panic(fmt.Errorf("component '%s' not exist in registry", typeID))
}

func (r *Registry) assertSystemValid(system System) {
	if system == nil {
		panic(fmt.Errorf("failed registry system: trying to add nil system to world"))
	}

	if reflect.ValueOf(system).Kind() != reflect.Ptr {
		panic(fmt.Errorf("failed registry system: '%s': should by passed as mutable pointer",
			system.TypeID(),
		))
	}
}
