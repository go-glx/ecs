package component

import "github.com/fe3dback/glx-ecs/ecs"

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

func (ttl *TimeToLife) RequireComponents() []ecs.Component {
	return []ecs.Component{
		Deletable{},
	}
}
