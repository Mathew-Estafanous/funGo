package stream

import (
	. "github.com/Mathew-Estafanous/funGo/model"
)

// Stream is a struct that acts as a wrapper around channels and uses
// goroutines to pass down relevant models down the stream pipeline until
// a terminating process is reached. When building an entire stream pipeline,
// there are three main steps that are involved. Creation, Non-Terminal
// and Termination steps.
//
// Creation
//
// First is the Creation, which involves generating a Stream usually
// using a given ModelSlice or by providing a channel. If you provide a channel,
// you are responsible for closing it when finished.
type Stream struct {
	ch chan Model
}

type Consumer func(m Model)

// NewStream creates and returns a new stream struct that contains the
// passed in channel.
//
// The responsibility of closing the channel is left to the caller
// of the method and not the method itself.
func NewStream(c chan Model) Stream {
	return Stream{
		ch: c,
	}
}

// NewStreamFromSlice takes a model slice and generates a stream containing
// all the Models that were within that slice.
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

// Filter takes in a Predicate and uses it to filter out all models that do not
// match the given requirements. If the predicate returns 'true' then that model
// will be passed on to the next stream. If it is false, then it will not be sent
// to the next stream.
func (s Stream) Filter(pred Predicate) Stream {
	nextChan := make(chan Model)

	go func() {
		defer close(nextChan)
		for model := range s.ch {
			if pred(model) {
				nextChan <- model
			}
		}
	}()

	return NewStream(nextChan)
}

// Map takes in an Operator and returns a Stream that contains the list of
// models that the operator was used on.
func (s Stream) Map(op Operator) Stream {
	nextChan := make(chan Model)

	go func() {
		defer close(nextChan)
		for model := range s.ch {
			nextChan <- op(model)
		}
	}()

	return NewStream(nextChan)
}

// FlatMap applies and returns a Stream of models that have applied the
// MultiOperator to each given model. This acts as a one to many
// relationship operation that converts one Model into several models.
func (s Stream) FlatMap(multiOp MultiOperator) Stream {
	nextChan := make(chan Model)

	go func() {
		defer close(nextChan)
		for model := range s.ch {
			for _, m := range multiOp(model) {
				nextChan <- m
			}
		}
	}()

	return NewStream(nextChan)
}

// Limit takes in a given maximum and limits the number of models that
// are present within the stream and returns a stream that has a total
// number of elements that does not exceed the maximum limit.
//
// If the limit is already greater than the initial stream, then that
// stream will remain unchanged.
func (s Stream) Limit(max int) Stream {
	nextChan := make(chan Model)

	go func() {
		defer close(nextChan)
		count := 0
		for m := range s.ch {
			if count >= max {
				return
			}
			nextChan <- m
			count++
		}
	}()

	return NewStream(nextChan)
}

func (s Stream) Distinct() Stream {
	var modelList ModelSlice
	for m := range s.ch {
		if contains(modelList, m) {
			continue
		}
		modelList = append(modelList, m)
	}

	nextChan := make(chan Model)

	go func() {
		defer close(nextChan)
		for _, v := range modelList {
			nextChan <- v
		}
	}()

	return NewStream(nextChan)
}

// ForEach is a terminating process that does return anything. For each
// Model in the stream, the Consumer will be called on that model.
func (s Stream) ForEach(con Consumer)  {
	if s.ch == nil {
		return
	}

	for m := range s.ch {
		con(m)
	}
}

func contains(s []Model, m Model) bool {
	for _, v := range s {
		if v.Equals(m) {
			return true
		}
	}
	return false
}