package component

import "github.com/fe3dback/glx-ecs/ecs"

const TimeToLifeTypeID = "TimeToLife-b935a9df3cdd"

type TimeToLife struct {
	TicksLeft uint
}

func NewTimeToLife(ticksLeft uint) *TimeToLife {
	if ticksLeft < 0 {
		ticksLeft = 0
	}

	return &TimeToLife{
		TicksLeft: ticksLeft,
	}
}

func (c TimeToLife) TypeID() ecs.ComponentTypeID {
	return TimeToLifeTypeID
}

func (c TimeToLife) RequireComponents() []ecs.ComponentTypeID {
	return []ecs.ComponentTypeID{
		DeletableTypeID,
	}
}
