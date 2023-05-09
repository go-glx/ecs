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

func (s *UniqueCollection[K, V]) IterateInOrder(order []K, itt func(K, V)) {
	for _, key := range order {
		if val, exist := s.collection[key]; exist {
			itt(key, val)
		}
	}
}

func (s *UniqueCollection[K, V]) Get(id K) (V, bool) {
	data, exist := s.collection[id]
	return data, exist
}

func (s *UniqueCollection[K, V]) Has(id K) bool {
	_, exist := s.collection[id]
	return exist
}

func (s *UniqueCollection[K, V]) Values() []V {
	values := make([]V, 0, len(s.collection))

	for _, value := range s.collection {
		values = append(values, value)
	}

	return values
}

func (s *UniqueCollection[K, V]) Keys() []K {
	keys := make([]K, 0, len(s.collection))

	for key := range s.collection {
		keys = append(keys, key)
	}

	return keys
}
