package system

import (
	"github.com/fe3dback/glx-ecs/component"
	"github.com/fe3dback/glx-ecs/ecs"
)

type GarbageCollector struct {
}

func NewGarbageCollector() *GarbageCollector {
	return &GarbageCollector{}
}

func (g *GarbageCollector) OnUpdate(w *ecs.World) {
	g.updateTTLComponents(w)
	g.deleteDeadEntities(w)
}

func (g *GarbageCollector) updateTTLComponents(w *ecs.World) {
	ttlComponents := ecs.FindComponent[component.TimeToLife](w)
	for ent, cmp := range ttlComponents {
		if cmp.TicksLeft > 0 {
			cmp.TicksLeft--
		}

		if cmp.TicksLeft <= 0 {
			ecs.MustFindComponentOf[component.Deletable](ent).Alive = false
		}
	}
}

func (g *GarbageCollector) deleteDeadEntities(w *ecs.World) {
	deadComponents := ecs.FindComponentWhere[component.Deletable](w, isDeadComponent)
	if len(deadComponents) == 0 {
		return
	}

	for ent := range deadComponents {
		w.RemoveEntity(ent)
	}
}

func isDeadComponent(cmp *component.Deletable) bool {
	return !cmp.Alive
}
