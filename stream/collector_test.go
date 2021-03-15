package stream

import (
	. "github.com/Mathew-Estafanous/funGo/model"
	"testing"
)

func TestNewCollector(t *testing.T) {
	supplier := func() Model {
		return ModelSlice{}
	}

	accumulator := func(m1, m2 Model) Model {
		return append(m1.(ModelSlice), m2)
	}

	finisher := func(m Model) Model {
		return m
	}

	collector := NewCollector(supplier, accumulator, finisher)
	if _, ok := collector.supplier().(ModelSlice); !ok {
		t.Error("NewCollector did not use the supplier that was passed in.")
	}

	collectorAccumulator := collector.accumulator(ModelSlice{}, ModelInt(1))
	if !collectorAccumulator.Equals(ModelSlice{ ModelInt(1) }) {
		t.Error("NewCollector did not use the accumulator that was passed in.")
	}

	collectorFinisher := collector.finisher(ModelInt(1))
	if !collectorFinisher.Equals(ModelInt(1)) {
		t.Error("NewCollector did not use the finisher that was passed in.")
	}
}

func TestToSlice(t *testing.T) {
	collector := ToSlice()

	supplierResult := collector.supplier()
	if _, ok := supplierResult.(ModelSlice); !ok {
		t.Error("ToSlice supplier does not return a valid ModelSlice type.")
	}

	accumulatorResult := collector.accumulator(ModelSlice{}, ModelInt(1))
	if !accumulatorResult.Equals(ModelSlice{ ModelInt(1) }) {
		t.Error("ToSlice accumulator should properly append Model into the given slice.")
	}

	finisherResult := collector.finisher(ModelInt(1))
	if !finisherResult.Equals(ModelInt(1)) {
		t.Error("Finished for ToSlice should return the exact same model.")
	}
}

func TestGroupingBy(t *testing.T) {
	mockCollector := Collector{
		supplier: func() Model { return ModelSlice{} },
		accumulator: func(m1, m2 Model) Model { return append(m1.(ModelSlice), m2) },
		finisher: func(m Model) Model { return m },
	}


	groupCollector := GroupingBy(func(m Model) Model {
									return m.(ModelInt)
								}, mockCollector)

	supplierResult := groupCollector.supplier()
	if !supplierResult.Equals(ModelMap{}) {
		t.Error("GroupingBy supplier did not return a valid ModelMap type.")
	}

	accumulatorResult := ModelMap{}
	for i := 0; i < 2; i++ {
		accumulatorResult = groupCollector.accumulator(accumulatorResult, ModelInt(i)).(ModelMap)
	}

	expectedMap := ModelMap{
		ModelInt(0): ModelSlice{ ModelInt(0) },
		ModelInt(1): ModelSlice{ ModelInt(1) },
	}

	if !accumulatorResult.Equals(expectedMap) {
		t.Error("GroupingBy accumulator did not properly create the correct ModelMap")
	}
}

func TestBasicFinisher(t *testing.T) {
	result := BasicFinisher(ModelInt(1))
	if !result.Equals(ModelInt(1)) {
		t.Error("Basic Finisher is not simply return exact same model as expected.")
	}
}
