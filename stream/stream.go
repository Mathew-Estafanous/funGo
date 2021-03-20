package stream

import (
	. "github.com/Mathew-Estafanous/funGo/model"
	. "github.com/Mathew-Estafanous/funGo/optional"
)

// Stream is a struct that acts as a wrapper around channels and uses
// goroutines to pass down relevant models down the stream pipeline until
// a terminating process is reached. When building an entire stream pipeline,
// there are three main steps that are involved. Creation, Non-Terminal
// and Termination steps.
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

// Distinct alters the given stream by removing all duplicate elements
// and ensuring that the stream does not contain any equal values.
// If there are no duplicates, then the stream should remain unaltered.
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
		for _, k := range modelList {
			nextChan <- k
		}
	}()

	return NewStream(nextChan)
}

// contains is an unexported method that Distinct() when checking
// that there are no duplicates in the given ModelSlice.
func contains(slice ModelSlice, m Model) bool {
	for _, v := range slice {
		if v.Equals(m) {
			return true
		}
	}
	return false
}

// Peek is an operation that uses a consumer to peek into the given
// stream and observe the Models within. It is not meant to alter
// any of the elements or act as a terminal operation.
//
// The ForEach function is similar, but is meant as a terminal operation,
// unlike this.
func (s Stream) Peek(consumer Consumer) Stream {
	var modelList ModelSlice
	for m := range s.ch {
		consumer(m)
		modelList = append(modelList, m)
	}

	nexChan := make(chan Model)

	go func() {
		defer close(nexChan)
		for _, v := range modelList {
			nexChan <- v
		}
	}()

	return NewStream(nexChan)
}

// AnyMatch is a terminating process that uses a given predicate to
// check if the predicate is true on any of the models. If it matches
// with any of the models, then the entire process will return true.
func (s Stream) AnyMatch(predicate Predicate) bool {
	for m := range s.ch {
		if predicate(m) {
			return true
		}
	}
	return false
}

// AllMatch is a termination process and requires that all the models
// within the stream match the predicate or else the function will
// end up returning false. If all models match the predicate then the
// return bool will be true.
func (s Stream) AllMatch(predicate Predicate) bool {
	for m := range s.ch {
		if !predicate(m) {
			return false
		}
	}
	return true
}

// NoneMatch is a terminating process that is the reciprocal
// result to AllMatch. Returning true if all the models do not match
// the predicate and false if any of the models match the predicate.
func (s Stream) NoneMatch(predicate Predicate) bool {
	for m := range s.ch {
		if predicate(m) {
			return false
		}
	}
	return true
}

func (s Stream) FindFirst(predicate Predicate) Optional {
	for m := range s.ch {
		if predicate(m) {
			return OptionalOf(m)
		}
	}
	return OptionalEmpty()
}

// Count takes in the Stream and gets the total number of models that
// are remaining in the given Stream. This is a terminal operation and
// will return the count as an int.
func (s Stream) Count() int {
	count := 0
	for range s.ch {
		count++
	}
	return count
}

// Collect is an important terminal operator that allows flexibility in
// in outlining how the stream should be grouped and collected. Note that
// the method returns 'interface{}' which allows for any type beyond just
// Models.
func (s Stream) Collect(collector Collector) interface{} {
	result := collector.supplier()

	for m := range s.ch {
		result = collector.accumulator(result, m)
	}

	result = collector.finisher(result)

	return result
}

// ForEach is a terminating process that does return anything. For each
// Model in the stream, the Consumer will be called on that model.
func (s Stream) ForEach(consumer Consumer) {
	if s.ch == nil {
		return
	}

	for m := range s.ch {
		consumer(m)
	}
}
