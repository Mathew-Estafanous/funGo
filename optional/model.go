package optional

type Model interface {
	Equals(Model) bool
}

type ModelInt int

func (m ModelInt) Equals(n Model) bool {
	return m == n
}