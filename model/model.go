package model

// Simple interface that every functional type that uses this
// library must use.
type Model interface {
	Equals(Model) bool
}

// ModelsEqual checks if two given models are equal.
func ModelsEqual(m1, m2 Model) bool {
	return (m1 == nil && m2 == nil) || (m1 != nil && m1.Equals(m2))
}