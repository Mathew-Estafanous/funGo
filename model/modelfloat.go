package model

// ModelFloat is a model of the 'float32' type
type ModelFloat float32

// Equals checks and returns 'true' if m is equal to mf
func (mf ModelFloat) Equals(m Model) bool {
	return mf == m
}
