# go-glx / ECS

Simple and powerful __Entity Components Systems__ (ECS) pattern library
for using in any game engine

Library use generics for typing and require go1.18

- Not Safe for concurrent use
- [WIP] Not stable API until version 1.0

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
  world.Update()
  world.Draw()
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
// recommended value: namespace <author>/<componentName>
const Node2DTypeID = "fe3dback/Node2D" 

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

- `OnInit(w RuntimeWorld)`
- `OnUpdate(w RuntimeWorld)`
- `OnSync(w RuntimeWorld)`

And system method `TypeID() ecs.SystemTypeID`

Any struct can implement one or more of this methods
ECS will help with entities/components fast filtering:

```go
import "github.com/fe3dback/glx-ecs/ecs"

type Gravity struct {
  engine engine
  filter Filter1[Node2D]
}

func NewGravity(engine engine) *Gravity {
  return &Gravity{engine: engine}
}

func (g *Gravity) TypeID() ecs.SystemTypeID {
  return "fe3dback/Gravity"
}

func (s *Gravity) OnInit(w RuntimeWorld) {
  t.filter = NewFilter1[Node2D](w)
}

func (s *Gravity) OnUpdate(w ecs.RuntimeWorld) {
  // full typing support, because of go1.18 generics
  found := t.filter.Find()
  for found.Next() {
    // entity = *Entity instance
    // cmp    = *Node2D instance
    ent, cmp := found.Get()

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

// ..

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

### System Properties

System can have unique properties, that can be same or different in different worlds.
All system properties will be encoded/decode to snapshot. This very helpful for
making map editors and levels loading.

By default, 3 most common property types is supported out of the box:
- props.Int
- props.Float (32bit)
- props.String

but you can make any custom property, just implement `props.Property` interface

```go
type Gravity struct {
  propForce *props.Float
}

func NewGravity() *Gravity {
  return &Gravity{
    propForce: props.NewFloat("force", 9.8),
  }
}

func (s *Gravity) Props() []props.Property {
  return []props.Property{
    s.propForce,
  }
}

func (s *Gravity) OnUpdate(w ecs.RuntimeWorld) {
  ...
  cmp.y += 100 * s.propForce.Get() * s.engine.DeltaTime()
  //             ^
  //             current property value (9.8)
  ...
}
```

This will be encoded as:

```xml
<StaticWorld>
  <systems>
    <system id="fe3dback/Gravity">
      <props>
        <prop name="force" value="9.8"></prop>
      </props>
    </system>
  </systems>
</StaticWorld>
```

Do not make all system consts as properties, because it will make your snapshots bigger.
Also, keep in mind that any snapshot value will ALWAYS override default property value.

Go code value and snapshot value will be mapped by specified name

**Tips**: 
- properties is absolutely useless if you not use snapshot save/load feature. (example: levels loading)
- for any runtime properties (that not needed to store in files) - just use normal go struct values
- use shared mutable variables (just `*string` ptr for example in constructor) for linking different systems together   

### Registry

All __Systems__ and __Components__ MUST be registered
before adding them to __World__

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

Notes:
- Registry allows to marshal/unmarshal world automatically to XML or other formats, it`s very useful for storing levels in files.
- Also, registry used for compute components bitmaps hash for fast filtering.

## Drawing world

You can define some system that can draw
all game objects, or some specified components

```go
type renderer interface {          // it`s your engine stuff
  DrawTexture2D(x, y, assetID int) // for example
}

type Drawer struct {
  renderer renderer
}

func (s *Drawer) OnDraw(w ecs.RuntimeWorld) {
  // OnDraw called right after world.OnUpdate
  // its best place to draw world

  found := NewFilter2[Texture2D, Transform2D](w).Find()
  for found.Next() {
    _, texture, transform := found.Get()
    
    s.renderer.DrawTexture2D(
      transform.x, 
      transform.y, 
      texture.assetID
    )
  }
}
```

## Snapshot (Save/Load world to file)

Lib can create snapshots of the World and marshal it into
XML/json/etc.. format

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
    <system id="fe3dback/Gravity">
      <props>
        <prop name="force" value="9.8"></prop>
      </props>
    </system>
  </systems>
  <entities>
    <entity name="my entity">
      <components>
        <component id="fe3dback/Node2D">
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
Tips (what you can do with snapshots):
- save/load world from files. (useful for map editors)
- in-mem save/load state inside custom map editor (immediate-mode testing, like unity/unreal "play" button)
- don`t use snapshots as game save/load system. Snapshot will have only public fields from all components, but not have any private fields evaluated during World.Update(). 

