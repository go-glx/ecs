package bits

import (
	"fmt"
)

type Registry struct {
	banksUsed   uint8
	entriesUsed uint8
	hash        map[string]Id
}

func NewRegistry() *Registry {
	return &Registry{
		hash: make(map[string]Id),
	}
}

func (r *Registry) RegisterComponent(cmpID string) {
	if r.entriesUsed >= maxEntries {
		r.banksUsed++
		r.entriesUsed = 0
	}
	if r.entriesUsed >= maxBanks {
		panic(fmt.Errorf("failed register component: max components limit is reached (%d banks, %d total components)", maxBanks, maxBanks*maxEntries))
	}

	if _, exist := r.hash[cmpID]; exist {
		panic(fmt.Errorf("component '%s' already registered", cmpID))
	}

	// combine id and return
	newID := Id{
		bankID:  bankID(r.banksUsed),
		entryID: entryID(r.entriesUsed),
	}
	r.hash[cmpID] = newID

	// calculate next val
	r.entriesUsed++
}

func (r *Registry) ComponentBits(cmpID string) Id {
	if id, exist := r.hash[cmpID]; exist {
		return id
	}

	panic(fmt.Errorf("failed get componentBits, component '%s' not registered yet", cmpID))
}
