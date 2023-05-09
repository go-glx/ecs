package props

type Property interface {
	Name() string
	Encode() string
	Decode(raw string)
}
