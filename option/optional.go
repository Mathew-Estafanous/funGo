package option

const (
	NoSuchElement = "There is no such element"
)

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

func (o Optional) Get() Model {
	if o.value == nil {
		panic(NoSuchElement)
	}
	return o.value
}
