package system

import (
	"github.com/go-glx/ecs/component"
	"github.com/go-glx/ecs/ecs"
)

const GarbageCollectorTypeID = "internal/GarbageCollector"

// GarbageCollector is example system that works with example component.TimeToLife
type GarbageCollector struct {
	filter ecs.Filter1[component.TimeToLife]
}

func NewGarbageCollector() *GarbageCollector {
	return &GarbageCollector{}
}

func (g *GarbageCollector) TypeID() ecs.SystemTypeID {
	return GarbageCollectorTypeID
}

func (g *GarbageCollector) OnInit(w ecs.RuntimeWorld) {
	g.filter = ecs.NewFilter1[component.TimeToLife](w)
}

func (g *GarbageCollector) OnUpdate(w ecs.RuntimeWorld) {
	found := g.filter.Find()

	for found.Next() {
		ent, cmp := found.Get()

		if cmp.TicksLeft > 0 {
			cmp.TicksLeft--
		}

		if cmp.TicksLeft <= 0 {
			w.RemoveEntity(ent)
		}
	}
}
