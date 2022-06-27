# Galaxy ECS

__NOT READY FOR REAL USE__

Version: 0.2

--

Simple Entity Component Systems (ECS) pattern library
for galaxy engine

Library use generics for typing and require go1.18

Safe for concurent use

## Usage

#### Basic world (game loop)

```go
world := e.createWorld()

for {
  world.Update() // update
  world.Sync()   // draw
}
```

#### Adding entities to world

```go
ent := ecs.NewEntity("my entity")
ent.AddComponent(component.NewDeletable())
ent.AddComponent(component.NewTimeToLife(100))

world.AddEntity(ent) // will be queued to next frame
```

#### Define custom components

Component in glx-ecs is just mutable __data__ struct, it not
has any behavior

```go
type Node2D struct {
  x float64
  y float64
}

func NewNode2D(x, y float64) *Node2D {
  return &Node2D{
    x: x,
    y: y,
  }
}
```

After component definition, you can add it to entity

```go
ent.AddComponent(NewNode2D(10, 5))
```

#### Add live to world with systems

Systems is just composition of __interface__`s:

- `OnUpdate(w *World)`
- `OnSync(w *World)`

Any struct can implement one or both of this methods:

```go
type Gravity struct {
 // ..
}

func (s *Gravity) OnUpdate(w *ecs.World) {
  // some update code
  // will be executed every world.Update()
}
```

ECS will help with entities/components fast filter:

```go
func (s *Gravity) OnUpdate(w *ecs.World) {
  // full typing support, because of go1.18 generics
  found := ecs.FindByComponent(w, &Node2D{})
  
  for entity, cmp := range found {
    // entity = *Entity instance
    // cmp    = *Node2D instance
	
    // for example, update y value of Node2D
    // in all world entities to 100px per second
    cmp.y += 100 * s.engine.DeltaTime()
  }
}
```

Ok, but what exactly s.engine and where ECS get DeltaTime()?

ECS - __do nothing with engine stuff__

You should provide all System deps, like _engine_ in SOLID
manner with dependency injection pattern:

```go
type engine interface {
  DeltaTime() float64
}

type Gravity struct {
  engine engine
}

func NewGravity(engine engine) *Gravity {
  return &Gravity{engine: engine}
}

// ..

func (e *MyEngine) createWorld() *ecs.World {
  world := ecs.NewWorld()

  world.AddSystem(NewGravity(e)) // will be queued like in entities
}
```

#### Drawing world

You can define some system that can draw
all game object, or some specified components

```go
type renderer interface {
  DrawBox(x, y, w, h int) // for example
}

type Drawer struct {
  renderer renderer
}

func (s *Drawer) OnSync(w *ecs.World) {
  // OnSync called right after world.OnUpdate
  // its best place to draw world

  textures := ecs.FindByComponent(w, &Texture{})
  for _, tex := range textures {
    s.renderer.DrawBox(tex.x, tex.y, tex.w, tex.h)
  }
}
```
