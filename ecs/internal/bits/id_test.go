package bits

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_id_entryMask(t *testing.T) {
	tests := []struct {
		name    string
		entryID entryID
		want    int64
	}{
		{
			name:    "0",
			entryID: 0,
			want:    1,
		},
		{
			name:    "1",
			entryID: 1,
			want:    2,
		},
		{
			name:    "2",
			entryID: 2,
			want:    4,
		},
		{
			name:    "63",
			entryID: 63,
			want:    math.MinInt64,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			x := Id{entryID: tt.entryID}
			assert.Equalf(t, tt.want, x.entryMask(), "entryMask()")
		})
	}
}
