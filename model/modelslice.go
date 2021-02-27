package model

// ModelSlice is a Model for the type []Model
type ModelSlice []Model

// Equals checks and returns 'true' if m is equal to ms
func (ms ModelSlice) Equals(m Model) bool {
	modelSlice, ok := m.(ModelSlice)
	if ok == false || len(modelSlice) != len(ms) {
		return false
	}

	for i, _ := range ms {
		m1, m2 := modelSlice[i], ms[i]
		if !m1.Equals(m2) {
			return false
		}
	}
	return true
}

