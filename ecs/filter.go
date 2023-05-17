package ecs

import "github.com/go-glx/ecs/ecs/internal/bits"

// todo: cache results in world (until entities created/destroyed)
// todo: T1..T8 variants (code generation)

type (
	filter1[T1 Component] struct {
		world   *World
		include []bits.Id
		exclude []bits.Id
	}
	filter2[T1, T2 Component] struct {
		world   *World
		include []bits.Id
		exclude []bits.Id
	}
	filter3[T1, T2, T3 Component] struct {
		world   *World
		include []bits.Id
		exclude []bits.Id
	}

	Filter1[T1 Component] interface {
		Find() Result1[T1]
	}
	Filter2[T1, T2 Component] interface {
		Find() Result2[T1, T2]
	}
	Filter3[T1, T2, T3 Component] interface {
		Find() Result3[T1, T2, T3]
	}
)

func NewFilter1[T1 Component](w RuntimeWorld, exclude ...Component) Filter1[T1] {
	var TType1 T1

	flt := &filter1[T1]{
		world: w.(*World),
		include: []bits.Id{
			w.(*World).registry.bitsRegistry.ComponentBits(string(TType1.TypeID())),
		},
		exclude: make([]bits.Id, len(exclude)),
	}

	for ind, excludeCmp := range exclude {
		flt.exclude[ind] = w.(*World).registry.bitsRegistry.ComponentBits(string(excludeCmp.TypeID()))
	}

	return flt
}
func NewFilter2[T1, T2 Component](w RuntimeWorld, exclude ...Component) Filter2[T1, T2] {
	var TType1 T1
	var TType2 T2

	flt := &filter2[T1, T2]{
		world: w.(*World),
		include: []bits.Id{
			w.(*World).registry.bitsRegistry.ComponentBits(string(TType1.TypeID())),
			w.(*World).registry.bitsRegistry.ComponentBits(string(TType2.TypeID())),
		},
		exclude: make([]bits.Id, len(exclude)),
	}

	for ind, excludeCmp := range exclude {
		flt.exclude[ind] = w.(*World).registry.bitsRegistry.ComponentBits(string(excludeCmp.TypeID()))
	}

	return flt
}
func NewFilter3[T1, T2, T3 Component](w RuntimeWorld, exclude ...Component) Filter3[T1, T2, T3] {
	var TType1 T1
	var TType2 T2
	var TType3 T3

	flt := &filter3[T1, T2, T3]{
		world: w.(*World),
		include: []bits.Id{
			w.(*World).registry.bitsRegistry.ComponentBits(string(TType1.TypeID())),
			w.(*World).registry.bitsRegistry.ComponentBits(string(TType2.TypeID())),
			w.(*World).registry.bitsRegistry.ComponentBits(string(TType3.TypeID())),
		},
		exclude: make([]bits.Id, len(exclude)),
	}

	for ind, excludeCmp := range exclude {
		flt.exclude[ind] = w.(*World).registry.bitsRegistry.ComponentBits(string(excludeCmp.TypeID()))
	}

	return flt
}

func (flt *filter1[T1]) Find() Result1[T1] {
	var TType1 T1
	result := &filterResult1[T1]{}

	for _, ent := range flt.world.entities.Iterate() {
		if !ent.componentMask.ContainAll(flt.include...) {
			continue
		}

		if ent.componentMask.ContainAny(flt.exclude...) {
			continue
		}

		cmp, exist := ent.components.Get(TType1.TypeID())
		if !exist {
			continue
		}

		result.ent = append(result.ent, ent)
		result.cmp1 = append(result.cmp1, cmp.(any).(*T1))
	}

	return result
}
func (flt *filter2[T1, T2]) Find() Result2[T1, T2] {
	var TType1 T1
	var TType2 T2
	result := &filterResult2[T1, T2]{}

	for _, ent := range flt.world.entities.Iterate() {
		if !ent.componentMask.ContainAll(flt.include...) {
			continue
		}

		if ent.componentMask.ContainAny(flt.exclude...) {
			continue
		}

		cmp1, exist := ent.components.Get(TType1.TypeID())
		if !exist {
			continue
		}
		cmp2, exist := ent.components.Get(TType2.TypeID())
		if !exist {
			continue
		}

		result.ent = append(result.ent, ent)
		result.cmp1 = append(result.cmp1, cmp1.(any).(*T1))
		result.cmp2 = append(result.cmp2, cmp2.(any).(*T2))
	}

	return result
}
func (flt *filter3[T1, T2, T3]) Find() Result3[T1, T2, T3] {
	var TType1 T1
	var TType2 T2
	var TType3 T3
	result := &filterResult3[T1, T2, T3]{}

	for _, ent := range flt.world.entities.Iterate() {
		if !ent.componentMask.ContainAll(flt.include...) {
			continue
		}

		if ent.componentMask.ContainAny(flt.exclude...) {
			continue
		}

		cmp1, exist := ent.components.Get(TType1.TypeID())
		if !exist {
			continue
		}
		cmp2, exist := ent.components.Get(TType2.TypeID())
		if !exist {
			continue
		}
		cmp3, exist := ent.components.Get(TType3.TypeID())
		if !exist {
			continue
		}

		result.ent = append(result.ent, ent)
		result.cmp1 = append(result.cmp1, cmp1.(any).(*T1))
		result.cmp2 = append(result.cmp2, cmp2.(any).(*T2))
		result.cmp3 = append(result.cmp3, cmp3.(any).(*T3))
	}

	return result
}
