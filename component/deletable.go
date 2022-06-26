package component

type Deletable struct {
	Alive bool
}

func NewDeletable() *Deletable {
	return &Deletable{
		Alive: true,
	}
}
