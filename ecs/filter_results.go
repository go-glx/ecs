package ecs

type (
	filterResult1[T1 any] struct {
		ent  []*Entity
		cmp1 []*T1
		iter int
	}
	filterResult2[T1, T2 any] struct {
		ent  []*Entity
		cmp1 []*T1
		cmp2 []*T2
		iter int
	}
	filterResult3[T1, T2, T3 any] struct {
		ent  []*Entity
		cmp1 []*T1
		cmp2 []*T2
		cmp3 []*T3
		iter int
	}

	Result1[T1 any] interface {
		Next() bool
		Get() (*Entity, *T1)
	}
	Result2[T1, T2 any] interface {
		Next() bool
		Get() (*Entity, *T1, *T2)
	}
	Result3[T1, T2, T3 any] interface {
		Next() bool
		Get() (*Entity, *T1, *T2, *T3)
	}
	Result4[T1, T2, T3, T4 any] interface {
		Next() bool
		Get() (*Entity, *T1, *T2, *T3, *T4)
	}
)

func (q *filterResult1[T1]) Next() bool {
	if q.iter >= len(q.ent) {
		return false
	}

	q.iter++
	return true
}
func (q *filterResult2[T1, T2]) Next() bool {
	if q.iter >= len(q.ent) {
		return false
	}

	q.iter++
	return true
}
func (q *filterResult3[T1, T2, T3]) Next() bool {
	if q.iter >= len(q.ent) {
		return false
	}

	q.iter++
	return true
}

func (q *filterResult1[T1]) Get() (*Entity, *T1) {
	return q.ent[q.iter-1], q.cmp1[q.iter-1]
}
func (q *filterResult2[T1, T2]) Get() (*Entity, *T1, *T2) {
	return q.ent[q.iter-1], q.cmp1[q.iter-1], q.cmp2[q.iter-1]
}
func (q *filterResult3[T1, T2, T3]) Get() (*Entity, *T1, *T2, *T3) {
	return q.ent[q.iter-1], q.cmp1[q.iter-1], q.cmp2[q.iter-1], q.cmp3[q.iter-1]
}
