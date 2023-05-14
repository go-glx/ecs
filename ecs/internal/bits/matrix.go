package bits

type Matrix struct {
	bitmap []int64
}

func NewMatrix() *Matrix {
	return &Matrix{
		bitmap: make([]int64, 1),
	}
}

func (m *Matrix) AddComponent(bits Id) {
	if int(bits.bankID) >= len(m.bitmap) {
		m.bitmap = append(m.bitmap, 0)
	}

	m.bitmap[int(bits.bankID)] |= bits.entryMask()
}

func (m *Matrix) RemoveComponent(bits Id) {
	if int(bits.bankID) >= len(m.bitmap) {
		return
	}

	m.bitmap[int(bits.bankID)] ^= bits.entryMask()
}

func (m *Matrix) Clear() {
	for ind := range m.bitmap {
		m.bitmap[ind] = 0
	}
}

func (m *Matrix) ContainAll(ids ...Id) bool {
	for _, id := range ids {
		if !m.Contain(id) {
			return false
		}
	}

	return true
}

func (m *Matrix) ContainAny(ids ...Id) bool {
	for _, id := range ids {
		if m.Contain(id) {
			return true
		}
	}

	return false
}

func (m *Matrix) Contain(bits Id) bool {
	if int(bits.bankID) >= len(m.bitmap) {
		return false
	}

	return m.bitmap[int(bits.bankID)]&bits.entryMask() != 0
}
