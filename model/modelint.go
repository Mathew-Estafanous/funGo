package model

// ModelInt is a Model for the type int
type ModelInt int

// Equals checks and returns 'true' if m is equal to mi
func (mi ModelInt) Equals(m Model) bool {
	return mi == m
}
