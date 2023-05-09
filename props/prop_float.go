package props

import (
	"fmt"
	"strconv"
)

type Float struct {
	name  string
	value float32
}

func NewFloat(name string, defaultValue float32) *Float {
	return &Float{
		name:  name,
		value: defaultValue,
	}
}

func (f *Float) Set(newValue float32) {
	f.value = newValue
}

func (f *Float) Get() float32 {
	return f.value
}

func (f *Float) Name() string {
	return f.name
}

func (f *Float) Encode() string {
	return strconv.FormatFloat(float64(f.value), 'f', 5, 32)
}

func (f *Float) Decode(raw string) {
	v, err := strconv.ParseFloat(raw, 32)
	if err != nil {
		panic(fmt.Errorf("failed decode prop '%s': value '%s' is invalid float32", f.name, raw))
	}
	f.value = float32(v)
}
