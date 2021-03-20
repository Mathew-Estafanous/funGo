package model

// ModelByte is a Model for the type byte
type ModelByte byte

// Equals checks and returns 'true' if m is equal to mb
func (mb ModelByte) Equals(m Model) bool {
	return mb == m
}
