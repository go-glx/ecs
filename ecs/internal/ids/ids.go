package ids

import "sync"

type (
	id = uint64

	globalID struct {
		value id
		sync.Mutex
	}
)

func NewGlobalID() *globalID {
	return &globalID{
		value: 0,
	}
}

func (gid *globalID) Next() uint64 {
	gid.Lock()
	gid.value++
	next := gid.value
	gid.Unlock()

	return next
}
