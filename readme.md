# Galaxy ECS

Version: 0.5

--

Simple and powerful __Entity Components Systems__ (ECS) pattern library
for using in any game engine

Library use generics for typing and require go1.18

Not Safe for concurrent use

--

Features:
- Generics (typed) API for components
- Simple API
- Support marshal/unmarshal world to XML and other formats
- written with SOLID in mind, easy integrate with any engine, that respects dependency injection pattern
- Fast component search/query (bitmaps, filters, etc..) (__TODO__) 

## Usage

### Game loop (empty world)

```go
import "github.com/fe3dback/glx-ecs/ecs"

registry := ecs.NewRegistry()
world := ecs.NewWorld(registry)

for {
  world.Update() // update
  world.Sync()   // draw
}
```

### Entities

Entity is ecs struct that has id and some other system fields

But for simplicity is just **Collection** of **Components** 

```go
import "github.com/fe3dback/glx-ecs/component"
import "github.com/my/engine/owncmp"

ent := ecs.NewEntity("my entity")

// add some std components
ent.AddComponent(component.NewDeletable())
ent.AddComponent(component.NewTimeToLife(100))

// add custom components
ent.AddComponent(owncmp.NewNode2D(10, 5))

// will be queued to next frame
world.AddEntity(ent)
```

### Components

Component is just mutable __data struct__, their not 
have any behavior

```go
import "github.com/fe3dback/glx-ecs/ecs"

// some unique type ID
// this used for marshal/unmarshal world to XML,JSON,etc..
// this should NOT change during any code refactoring
// recommended value: HumanReadableName + UUIDv4 (last part)
const Node2DTypeID = "Node2D-a300548e4f48" 

// Component data
type Node2D struct {
  X float64
  Y float64
}

// Constructor
func NewNode2D(x, y float64) *Node2D {
  return &Node2D{
    X: x,
    Y: y,
  }
}

// Single required ECS method for every component
func (c Node2D) TypeID() ecs.ComponentTypeID {
  return Node2DTypeID
}

// Optionally you can specify other components
// that MUST be added to entity, before adding this
// ECS will assert, that all required components exist
// in Entity
func (c Node2D) RequireComponents() []ecs.ComponentTypeID {
  return []ecs.ComponentTypeID{
    DeletableTypeID,
  }
}

```

After component definition, you can add it to entity

```go
ent.AddComponent(NewNode2D(10, 5))
```

### Systems - Apply live to world

Systems is just composition of __interface__`s:

- `OnUpdate(w *World)`
- `OnSync(w *World)`

And system method `TypeID() ecs.SystemTypeID`

Any struct can implement one or both of this methods:

```go
import "github.com/fe3dback/glx-ecs/ecs"

const GravityTypeID = "Gravity-adc3353dd900"

type Gravity struct {}

func NewGravity() *Gravity { 
  return &Gravity{}
}

func (g Gravity) TypeID() ecs.SystemTypeID {
  return GravityTypeID
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
  found := ecs.FindComponent[Node2D](w)
  
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

### Registry

All __Systems__ and __Components__ MUST be registered
before adding them to __World__

Registry allows to marshal/unmarshal world automatically
to XML or other formats, it`s very useful for storing levels in files

```go
r := ecs.NewRegistry()
r.RegisterSystem(system.NewGarbageCollector())
r.RegisterComponent(component.NewDeletable())

// [5, 10] - is default values for this component
// when world unmarshalled for example from XML
// it can be previous version of this component
// in this case ECS will use default values for
// all new fields, not exist in XML snapshot
r.RegisterComponent(owncmp.NewNode2D(5, 10))

// create world
world := ecs.NewWorld(r)
```

## Drawing world

You can define some system that can draw
all game objects, or some specified components

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

  textures := ecs.FindComponent[Texture](w)
  for _, tex := range textures {
    s.renderer.DrawBox(tex.x, tex.y, tex.w, tex.h)
  }
}
```

## Snapshot (Save/Load world to file)

You can create snapshot from World and marshal it into
XML or other format

This allows for example to load World from external editor.
Or maintain world save/load in custom game editor.

This function not suitable for save/load game systems.
Snapshot will have only public fields from all components,
but not have any private fields evaluated during World.Update()

Anyway in can be used as save/load game system in super simple games
and if you maintain all components state only in Public properties

```go
import "github.com/fe3dback/glx-ecs/snapshot"

w := ecs.NewWorld( .. )

// encode
snap := snapshot.Create(w)
xml := snapshot.MarshalToXML(snap)

// decode
newSnap := snapshot.UnmarshalFromXML(xml)
newWorld := snapshot.Restore(newSnap)
```

Marshalled XML:
```xml
<StaticWorld>
  <systems>
    <system id="Gravity-adc3353dd900"></system>
  </systems>
  <entities>
    <entity name="my entity">
      <components>
        <component id="Node2D-a300548e4f48">
          <props>
            <prop name="X" value="5"></prop>
            <prop name="Y" value="10"></prop>
          </props>
        </component>
      </components>
    </entity>
  </entities>
</StaticWorld>
```
