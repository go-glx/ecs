package props

import (
	"fmt"
	"strconv"
)

type Int struct {
	name  string
	value int
}

func NewInt(name string, defaultValue int) *Int {
	return &Int{
		name:  name,
		value: defaultValue,
	}
}

func (i *Int) Set(newValue int) {
	i.value = newValue
}

func (i *Int) Get() int {
	return i.value
}

func (i *Int) Name() string {
	return i.name
}

func (i *Int) Encode() string {
	return strconv.FormatInt(int64(i.value), 10)
}

func (i *Int) Decode(raw string) {
	v, err := strconv.ParseInt(raw, 10, 64)
	if err != nil {
		panic(fmt.Errorf("failed decode prop '%s': value '%s' is invalid int", i.name, raw))
	}
	i.value = int(v)
}
