package ids

import (
	"fmt"
	"strings"
)

type ObjectID = string

// Of return object classname:
//  - `pkg.Obj{}`  => `pkg.Obj`
//  - `&pkg.Obj{}` => `pkg.Obj`
//  - `Obj{}`      => `pkg.Obj`
//  - `&Obj{}`     => `pkg.Obj`
func Of(obj any) ObjectID {
	return strings.Replace(
		fmt.Sprintf("%T", obj),
		"*",
		"",
		1,
	)
}
