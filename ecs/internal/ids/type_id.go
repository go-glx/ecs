package ids

import (
	"reflect"
)

type ObjectID = string

// Of return object classname:
//  - `pkg.Obj{}`  => `pkg.Obj`
//  - `&pkg.Obj{}` => `pkg.Obj`
//  - `Obj{}`      => `pkg.Obj`
//  - `&Obj{}`     => `pkg.Obj`
func Of(obj any) ObjectID {
	s := reflect.TypeOf(obj).String()

	if len(s) > 1 {
		if s[0] == '*' {
			return s[1:]
		}
	}

	return s
}
