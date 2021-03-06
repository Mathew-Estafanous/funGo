package stream

import (
	. "github.com/Mathew-Estafanous/funGo/model"
	"testing"
)


func createStream(slice ModelSlice) Stream {
	openChannel := make(chan Model)

	go func() {
		defer close(openChannel)
		for _, m := range slice {
			openChannel <- m
		}
	}()

	return Stream{ openChannel }
}

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
		error     string
		values    ModelSlice
		predicate Predicate
		want      ModelSlice
	}

	filterTest := test {
		error:  "Filter should properly apply predicate and filter out unwanted models based on predicate.",
		values: ModelSlice{ ModelInt(4), ModelInt(2), ModelInt(7), ModelInt(8) },
		predicate: func(m Model) bool {
			return m.(ModelInt) > 5
		},
		want: ModelSlice{ ModelInt(7), ModelInt(8) },
	}

	result := createStream(filterTest.values).
				Filter(filterTest.predicate)

	for _, model := range filterTest.want {
		if m:= <- result.ch; !m.Equals(model) {
			t.Error(filterTest.error)
		}
	}
}

func TestStream_Map(t *testing.T) {
	type test struct {
		error    string
		values   ModelSlice
		operator Operator
		want     ModelSlice
	}

	mapTest := test {
		error:  "Map should take given Model and turn it into a ModelByte.",
		values: ModelSlice{ ModelInt(5), ModelInt(8), ModelInt(3) },
		operator: func(m Model) Model {
			return ModelByte(m.(ModelInt))
		},
		want: ModelSlice{ ModelByte(5), ModelByte(8), ModelByte(3) },
	}

	result := createStream(mapTest.values).
				Map(mapTest.operator)

	for _, model := range mapTest.want {
		if m:= <- result.ch; !m.Equals(model) {
			t.Error(mapTest.error)
		}
	}
}

func TestStream_FlatMap(t *testing.T) {
	type test struct {
		error    string
		values   ModelSlice
		operator MultiOperator
		result   ModelSlice
	}

	flatMapTest := test {
		error: "Given a Stream the multi-operator should double each element and add it to next stream.",
		values: ModelSlice{ ModelInt(3), ModelInt(4) },
		operator: func(m Model) []Model {
			return ModelSlice{ m, m}
		},
		result: ModelSlice{ ModelInt(3), ModelInt(3), ModelInt(4), ModelInt(4) },
	}

	result := createStream(flatMapTest.values).FlatMap(flatMapTest.operator)
	index := 0
	for m := range result.ch {
		if !flatMapTest.result[index].Equals(m) {
			t.Error(flatMapTest.error)
		}
		index++
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
		result := createStream(te.values).Limit(te.limit)
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

func TestStream_Distinct(t *testing.T) {
	type test struct {
		error string
		value ModelSlice
		want ModelSlice
	}

	distinctTests := []test {
		{
			error: "Duplicates in model slice should be removed from stream and only 1 of each element be found.",
			value: ModelSlice{ ModelInt(1), ModelInt(1), ModelInt(2) },
			want: ModelSlice{ ModelInt(1), ModelInt(2) },
		},
		{
			error: "Stream with no duplicate elements should not be altered.",
			value: ModelSlice{ ModelInt(1), ModelInt(2) },
			want: ModelSlice{ ModelInt(1), ModelInt(2) },
		},
	}

	for _, te := range distinctTests {
		result := createStream(te.value).Distinct()
		index := 0
		for m := range result.ch {
			if !m.Equals(te.want[index]) {
				t.Error(te.error)
			}
			index++
		}
	}
}

func TestStream_ForEach(t *testing.T) {
	type test struct {
		error string
		value ModelSlice
		consumer Consumer
		want int
	}

	count := 0
	forEachTest := test{
		error: "Stream with three models should have the consumer called on all three.",
		value: ModelSlice{ ModelInt(1), ModelInt(2), ModelInt(3) },
		consumer: func(m Model) {
			count++
		},
		want: 3,
	}

	createStream(forEachTest.value).ForEach(forEachTest.consumer)

	if count != forEachTest.want {
		t.Error(forEachTest.error)
	}
}