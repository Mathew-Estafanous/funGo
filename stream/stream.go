package stream

import (
	. "github.com/Mathew-Estafanous/funGo/model"
)

type Stream struct {
	ch chan Model
}

type Consumer func(m Model)

type Predicate func(m Model) bool

type Operator func(m Model) Model

func NewStream(c chan Model) Stream {
	return Stream{
		ch: c,
	}
}

func NewStreamFromSlice(slice ModelSlice) Stream {
	openChan := make(chan Model)

	go func() {
		defer close(openChan)
		for _, model := range slice {
			openChan <- model
		}
	}()

	return NewStream(openChan)
}

func (s Stream) Filter(pred Predicate) Stream {
	openChan := make(chan Model)

	go func() {
		defer close(openChan)
		for model := range s.ch {
			if pred(model) {
				openChan <- model
			}
		}
	}()

	return NewStream(openChan)
}

func (s Stream) Map(op Operator) Stream {
	openChan := make(chan Model)

	go func() {
		defer close(openChan)
		for model := range s.ch {
			openChan <- op(model)
		}
	}()

	return NewStream(openChan)
}

func (s Stream) ForEach(con Consumer)  {
	if s.ch == nil {
		return
	}

	for m := range s.ch {
		con(m)
	}
}