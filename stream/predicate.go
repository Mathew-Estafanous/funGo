package stream

import . "github.com/Mathew-Estafanous/funGo/model"

type Predicate func(m Model) bool

func (p Predicate) And(pred Predicate) Predicate {
	return func(m Model) bool {
		if pred(m) && p(m) {
			return true
		}
		return false
	}
}

func (p Predicate) Or(pred Predicate) Predicate {
	return func(m Model) bool {
		if pred(m) || p(m) {
			return true
		}
		return false
	}
}

func (p Predicate) Not() Predicate {
	return func(m Model) bool {
		return !p(m)
	}
}