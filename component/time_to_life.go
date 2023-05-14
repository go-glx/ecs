package component

import "github.com/go-glx/ecs/ecs"

const TimeToLifeTypeID = "internal/TimeToLife"

// TimeToLife is example component. See example system processing in system.GarbageCollector
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
