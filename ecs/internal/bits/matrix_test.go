package bits

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewMatrix(t *testing.T) {
	ref := func(bank, entry int64) Id {
		return Id{
			bankID:  bankID(bank),
			entryID: entryID(entry),
		}
	}

	tests := []struct {
		name     string
		assemble func(*Matrix)
		assert   func(*Matrix)
	}{
		{
			name:     "empty_bank",
			assemble: func(*Matrix) {},
			assert: func(mx *Matrix) {
				assert.Equal(t, make([]int64, 1), mx.bitmap)
			},
		},
		{
			name: "bank_add/remove",
			assemble: func(mx *Matrix) {
				mx.AddComponent(ref(0, 1))
				mx.AddComponent(ref(0, 3))
				mx.AddComponent(ref(0, 7))
				mx.AddComponent(ref(0, 9))
				mx.RemoveComponent(ref(0, 7))
				mx.RemoveComponent(ref(1, 3)) // #3 stay (another bank id)
				mx.AddComponent(ref(1, 5))
			},
			assert: func(mx *Matrix) {
				assert.True(t, mx.Contain(ref(0, 1)))
				assert.True(t, mx.Contain(ref(0, 3)))  // stayed
				assert.False(t, mx.Contain(ref(0, 7))) // removed
				assert.True(t, mx.Contain(ref(0, 9)))
				assert.True(t, mx.Contain(ref(1, 5)))

				firstBank := ref(0, 1).entryMask() |
					ref(0, 3).entryMask() |
					ref(0, 9).entryMask()

				secondBank := ref(1, 5).entryMask()

				assert.Equal(t, []int64{firstBank, secondBank}, mx.bitmap)
			},
		},
		{
			name: "has_all",
			assemble: func(mx *Matrix) {
				mx.AddComponent(ref(0, 1))
				mx.AddComponent(ref(0, 3))
				mx.AddComponent(ref(1, 5))
				mx.AddComponent(ref(1, 8))
				mx.AddComponent(ref(1, 15))
			},
			assert: func(mx *Matrix) {
				// true
				assert.True(t, mx.ContainAll(
					ref(0, 1),
					ref(1, 8),
					ref(0, 3),
					ref(1, 15),
				))
				assert.True(t, mx.ContainAll(
					ref(1, 15),
					ref(1, 8),
				))
				assert.True(t, mx.ContainAll(
					ref(1, 15),
				))

				// false
				assert.False(t, mx.ContainAll(
					ref(0, 1),
					ref(0, 2),
				))
				assert.False(t, mx.ContainAll(
					ref(0, 1),
					ref(1, 1),
				))
				assert.False(t, mx.ContainAll(
					ref(1, 1),
				))
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mat := NewMatrix()
			tt.assemble(mat)
			tt.assert(mat)
		})
	}
}
