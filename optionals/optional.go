package optionals

import "errors"

var NoSuchElement = errors.New("there is no such element")

type Optional struct {
	value Model
	empty bool
}

func OptionalOf(val Model) Optional {
	if val == nil {
		return OptionalEmpty()
	}
	return Optional{value: val, empty: false}
}

func OptionalEmpty() Optional {
	return Optional{value: nil, empty: true}
}

func (o Optional) IsPresent() bool {
	return o.empty
}

func (o Optional) Get() (Model, error) {
	if !o.IsPresent() {
		return nil, NoSuchElement
	}
	return o.value, nil
}

func (o Optional) GetOrElse(basic Model)  Model {
	if !o.IsPresent() {
		return basic
	}
	return o.value
}

func (o Optional) NotEmptyOrElse(other Optional) Optional {
	if o.IsPresent() {
		return other
	}
	return o
}