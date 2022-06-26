package collection

import (
	"sync"
)

type UniqueCollection[K comparable, V any] struct {
	collection map[K]V

	mux sync.RWMutex
}

func NewUniqueCollection[K comparable, V any]() *UniqueCollection[K, V] {
	return &UniqueCollection[K, V]{
		collection: make(map[K]V),
	}
}

func (s *UniqueCollection[K, V]) Len() int {
	return len(s.collection)
}

func (s *UniqueCollection[K, V]) Set(id K, item V) {
	s.mux.RLock()
	_, exist := s.collection[id]
	s.mux.RUnlock()

	if exist {
		return
	}

	s.mux.Lock()
	s.collection[id] = item
	s.mux.Unlock()
}

func (s *UniqueCollection[K, V]) Remove(id K) {
	s.mux.RLock()
	_, exist := s.collection[id]
	s.mux.RUnlock()

	if !exist {
		return
	}

	s.mux.Lock()
	delete(s.collection, id)
	s.mux.Unlock()
}

func (s *UniqueCollection[K, V]) Iterate() map[K]V {
	s.mux.RLock()
	collectionCopy := s.collection
	s.mux.RUnlock()

	return collectionCopy
}

func (s *UniqueCollection[K, V]) Get(id K) (V, bool) {
	s.mux.RLock()
	data, exist := s.collection[id]
	s.mux.RUnlock()

	return data, exist
}

func (s *UniqueCollection[K, V]) Has(id K) bool {
	s.mux.RLock()
	_, exist := s.collection[id]
	s.mux.RUnlock()

	return exist
}
