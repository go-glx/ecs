package props

type String struct {
	name  string
	value string
}

func NewString(name string, defaultValue string) *String {
	return &String{
		name:  name,
		value: defaultValue,
	}
}

func (s *String) Set(newValue string) {
	s.value = newValue
}

func (s *String) Get() string {
	return s.value
}

func (s *String) Name() string {
	return s.name
}

func (s *String) Encode() string {
	return s.value
}

func (s *String) Decode(raw string) {
	s.value = raw
}
