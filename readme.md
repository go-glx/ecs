# Galaxy ECS

__NOT READY FOR REAL USE__

Version: 0.1

--

Simple Entity Component Systems (ECS) pattern library
for galaxy engine

Library use generics for typing and require go1.18

## Usage

```go
type MySystem struct {
  // custom deps
  myEngine interface {
    DeltaTime() float64	  
  }
}

func NewMySystem(myEngine) *MySystem { .. }

func (s *MySystem) OnUpdate(w *ecs.World) {
  ttlComponents := ecs.FindByComponent(w, &component.TimeToLife{})

  for entity, cmp := range ttlComponents {
    // cmp is typed as &component.TimeToLife{}
    // because of go1.18 generics
    if cmp.TicksLeft > 0 {
      // ..
    }
  }
  
  if s.myEngine.DeltaTime() > 0.1 {
    // ECS do nothing with your engine state
    // you should provide it directly to system
    // constructor in your DI
  }
}

func (s *MySystem) OnSync(w *ecs.World) {
  // draw somehow system state
}

// ---------------------------------

func (e *myEngine) createWorld() *ecs.World {
  world := ecs.NewWorld()
  
  ent := ecs.NewEntity("my entity")
  ent.AddComponent(component.NewDeletable())
  ent.AddComponent(component.NewTimeToLife(100))
  
  world.AddEntity(ent)
  
  world.AddSystem(system.NewGarbageCollector()) // std systems
  world.AddSystem(NewMySystem(e))               // custom systems
}

func (e *myEngine) gameLoop() {
  world := e.createWorld()

  for {
    world.Update() // update
    world.Sync()   // draw
  }
}

```
