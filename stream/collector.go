package stream

import (
	. "github.com/Mathew-Estafanous/funGo/model"
)

// Supplier, very simply supplies the type of Model that the collector
// will use while it is collecting all the provided elements.
type Supplier func() Model

// NOTICE:
// This struct is heavily inspired by the Java Streams Collector library and the
// associated functionality. Credit goes to the engineers who developed
// the architecture of that entire library.
//
// Collector is a struct that outlines the reduction operation on the
// accumulated results that were given. This means that, it can be used to collect
// the resulting elements within the stream into an outlined format. This means
// collecting into Slices, Maps, or any other structure for that matter.
type Collector struct {
	supplier Supplier
	accumulator BiOperator
	finisher Operator
}

// NewCollector is used to create a new Collector struct with the given supplier,
// accumulator and finisher functions.
func NewCollector(supplier Supplier, accumulator BiOperator, finisher Operator) Collector {
	return Collector{
		supplier:    supplier,
		accumulator: accumulator,
		finisher:    finisher,
	}
}

// ToSlice builds a collector that will accumulate all elements into a
// ModelSlice type.
func ToSlice() Collector {
	supplier := func() Model { return ModelSlice{} }

	accumulator := func(supp, model Model) Model {
		return append(supp.(ModelSlice), model)
	}

	finisher := basicFinisher

	return NewCollector(supplier, accumulator, finisher)
}

// ToMap builds a collector that will accumulate all elements into a
// ModelMap type.
//
// This function is useful if the key and values aren't expected to be altered
// while being accumulated. Allowing for a simple ToMap() call instead of having
// to specify both the key and value mappers.
func ToMap() Collector {
	return ToMapSpecify(basicFinisher, basicFinisher)
}

// ToMapSpecify will build a collector that accumulates all elements into a
// ModelMap by using the passed in key and value Mappers inside the created
// accumulator.
func ToMapSpecify(keyMapper, valueMapper Operator) Collector {
	supplier := func() Model { return ModelMap{} }

	accumulator := func(supp, model Model) Model {
		k := keyMapper(model)
		v := valueMapper(model)

		supp.(ModelMap)[k] = v
		return supp
	}

	finisher := basicFinisher

	return NewCollector(supplier, accumulator, finisher)
}

// GroupingBy simply groups each element according to it's
// classifier, then placing it in the downstream collector for the value.
func GroupingBy(classifier Operator, downstream Collector) Collector {
	supplier := func() Model { return ModelMap{} }

	accumulator := func(supp, model Model) Model {
		k := classifier(model)
		container, ok := supp.(ModelMap)[k]
		if !ok {
			container = downstream.supplier()
		}

		grouped := downstream.accumulator(container, model)
		supp.(ModelMap)[k] = grouped
		return supp
	}

	finisher := func(m Model) Model {
		if downstream.finisher == nil {
			return basicFinisher(m)
		}

		s := supplier()
		for k, v := range m.(ModelMap) {
			s.(ModelMap)[k] = downstream.finisher(v)
		}
		return m
	}

	return NewCollector(supplier, accumulator, finisher)
}

func basicFinisher(m Model) Model {
	return m
}
