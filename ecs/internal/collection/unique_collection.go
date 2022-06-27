package collection

type UniqueCollection[K comparable, V any] struct {
	collection map[K]V
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
	if _, exist := s.collection[id]; exist {
		return
	}

	s.collection[id] = item
}

func (s *UniqueCollection[K, V]) Remove(id K) {
	if _, exist := s.collection[id]; !exist {
		return
	}

	delete(s.collection, id)
}

func (s *UniqueCollection[K, V]) Iterate() map[K]V {
	return s.collection
}

func (s *UniqueCollection[K, V]) Get(id K) (V, bool) {
	data, exist := s.collection[id]
	return data, exist
}

func (s *UniqueCollection[K, V]) Has(id K) bool {
	_, exist := s.collection[id]
	return exist
}
