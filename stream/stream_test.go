package stream

import (
	. "github.com/Mathew-Estafanous/funGo/model"
	"testing"
)

func TestNewStream(t *testing.T) {
	testValues := []ModelInt { 1,2,3,4,5 }
	testChan := make(chan Model)
	stream := NewStream(testChan)

	go func() {
		defer close(testChan)
		for _, m := range testValues {
			testChan <- m
		}
	}()

	for _, m := range testValues {
		if model := <- stream.ch; !model.Equals(m) {
			t.Error("New Stream did not create a stream with the correct channel.")
		}
	}
}

func TestNewStreamFromSlice(t *testing.T) {
	testSlice := ModelSlice{
		ModelInt(1), ModelInt(2), ModelInt(3), ModelInt(4),
	}
	stream := NewStreamFromSlice(testSlice)

	for _, model := range testSlice {
		if m := <- stream.ch; !m.Equals(model) {
			t.Error("New Stream from slice did not create a stream with the correct channel values.")
		}
	}
}

func TestStream_Filter(t *testing.T) {
	type test struct {
		name string
		values ModelSlice
		predicate Predicate
		want ModelSlice
	}

	filterTest := test {
		name: "Filter should properly apply predicate and filter out unwanted models based on predicate.",
		values:  ModelSlice{ ModelInt(4), ModelInt(2), ModelInt(7), ModelInt(8) },
		predicate: func(m Model) bool {
			return m.(ModelInt) > 5
		},
		want: ModelSlice{ ModelInt(7), ModelInt(8) },
	}

	result := NewStreamFromSlice(filterTest.values).
				Filter(filterTest.predicate)

	for _, model := range filterTest.want {
		if m:= <- result.ch; !m.Equals(model) {
			t.Error(filterTest.name)
		}
	}
}

func TestStream_Map(t *testing.T) {
	type test struct {
		name string
		values ModelSlice
		operator Operator
		want ModelSlice
	}

	mapTest := test {
		name: "Map should take given Model and turn it into a ModelByte.",
		values: ModelSlice{ ModelInt(5), ModelInt(8), ModelInt(3) },
		operator: func(m Model) Model {
			return ModelByte(m.(ModelInt))
		},
		want: ModelSlice{ ModelByte(5), ModelByte(8), ModelByte(3) },
	}

	result := NewStreamFromSlice(mapTest.values).
				Map(mapTest.operator)

	for _, model := range mapTest.want {
		if m:= <- result.ch; !m.Equals(model) {
			t.Error(mapTest.name)
		}
	}
}

func TestStream_Limit(t *testing.T) {
	type test struct {
		name string
		values ModelSlice
		limit int
		want ModelSlice
	}

	limitTest := []test {
		{
			name: "Stream with 4 models limited to 2 should return a stream of only 2 elements.",
			values: ModelSlice{ ModelInt(4), ModelInt(2), ModelInt(7), ModelInt(8) },
			limit: 2,
			want: ModelSlice{ ModelInt(4), ModelInt(2) },
		},
		{
			name: "Stream with 2 models limited to 3 should remain completely unchanged.",
			values: ModelSlice{ ModelInt(4), ModelInt(2) },
			limit: 3,
			want: ModelSlice{ ModelInt(4), ModelInt(2) },
		},
	}

	for _, te := range limitTest {
		result := NewStreamFromSlice(te.values).Limit(te.limit)
		count := 0
		for m := range result.ch {
			if !m.Equals(te.want[count]) {
				t.Error(te.name)
			}
			count++
		}

		if count != len(te.want) {
			t.Error(te.name)
		}
	}
}