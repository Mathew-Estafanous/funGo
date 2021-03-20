package stream

import . "github.com/Mathew-Estafanous/funGo/model"

// Predicate in a function that returns a boolean value either 'true'
// or 'false' depending on the passed in requirements.
type Predicate func(m Model) bool

// And combines two separate predicates together and into one predicate
// and requires that both predicates return true or else a return boolean
// of 'false' will be returned.
func (p Predicate) And(other Predicate) Predicate {
	return func(m Model) bool {
		if other == nil || p == nil {
			return false
		}
		return other(m) && p(m)
	}
}

// Or combines two predicates and requires that either one of them results
// in a return boolean of true or else it will return a final boolean
// result of 'false'
func (p Predicate) Or(other Predicate) Predicate {
	return func(m Model) bool {
		if other == nil || p == nil {
			return false
		}
		return other(m) || p(m)
	}
}

// Not returns a negated Predicate that returns the opposite boolean result.
func (p Predicate) Not() Predicate {
	return func(m Model) bool {
		return !p(m)
	}
}
