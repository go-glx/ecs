package ecs

type System interface{}

type SystemUpdatable interface {
	System
	OnUpdate(w *World)
}

type SystemSyncable interface {
	System
	OnSync(w *World)
}
