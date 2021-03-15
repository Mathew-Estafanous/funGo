package stream

import (
	. "github.com/Mathew-Estafanous/funGo/model"
)

type Supplier func() Model

type Collector struct {
	supplier Supplier
	accumulator BiOperator
	finisher Operator
}

func NewCollector(supplier Supplier, accumulator BiOperator, finisher Operator) Collector {
	return Collector{
		supplier:    supplier,
		accumulator: accumulator,
		finisher:    finisher,
	}
}

func ToSlice() Collector {
	supplier := func() Model { return ModelSlice{} }

	accumulator := func(supp, model Model) Model {
		return append(supp.(ModelSlice), model)
	}

	finisher := basicFinisher

	return NewCollector(supplier, accumulator, finisher)
}

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
