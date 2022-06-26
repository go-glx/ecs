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

func isDeadComponent(cmp *component.Deletable) bool {
	return !cmp.Alive
}

func (g *GarbageCollector) OnUpdate(w *ecs.World) {
	g.updateTTLComponents(w)
	g.deleteDeadEntities(w)
}

func (g *GarbageCollector) updateTTLComponents(w *ecs.World) {
	ttlComponents := ecs.FindByComponent(w, &component.TimeToLife{})
	for ent, cmp := range ttlComponents {
		if cmp.TicksLeft > 0 {
			cmp.TicksLeft--
		}

		if cmp.TicksLeft <= 0 {
			ecs.MustFindComponentOf(ent, &component.Deletable{}).Alive = false
		}
	}
}

func (g *GarbageCollector) deleteDeadEntities(w *ecs.World) {
	deadComponents := ecs.FindByComponentWhere(w, &component.Deletable{}, isDeadComponent)
	if len(deadComponents) == 0 {
		return
	}

	for ent := range deadComponents {
		w.RemoveEntity(ent)
	}
}
