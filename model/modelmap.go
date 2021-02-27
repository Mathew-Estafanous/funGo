package model

// ModelMap is a Model for the type map[Model]Model
type ModelMap map[Model]Model

// Equals checks and returns 'true' if m is equal to mm
func (mm ModelMap) Equals(m Model) bool {
	mappedModel, ok := m.(ModelMap)
	if ok == false {
		return false
	}

	for key, val := range mappedModel {
		if val2, ok := mm[key]; ok {
			if result := val.Equals(val2); !result {
				return false
			}
		} else {
			return false
		}
	}
	return true
}