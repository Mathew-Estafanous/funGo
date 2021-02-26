package optionals

const (
	NoModelFound    = "There is no such element."
	NilNotPermitted = "Nil values are not permitted"
)

// Optional is a simple struct that contains unexported Model
// and boolean types that keep track of the optional.
//
// Optionals can be used when a return value is not guaranteed
// and is not necessarily an error. This allows a simple functional
// approach to handling values that are not guaranteed.
//
// Optionals require all values to implement Model, meaning you will
// need to implement this interface for your own project types.
type Optional struct {
	model Model
	empty bool
}

// OptionalOf simply takes in a Model and creates an Optional that
// contains that model and the value associated with it.
//
// This function does not permit any nil values to be passed in or
// else a panic of NilNotPermitted will be thrown. If a non-nil value
// is not guaranteed, look into using OptionalOfNillable.
func OptionalOf(m Model) Optional {
	if m == nil {
		panic(NilNotPermitted)
	}
	return Optional{model: m, empty: false}
}

// OptionalOfNillable is very similar to OptionalOf, however this
// function can handle values that can possibly nil. In such a case
// an empty optional will be returned. If a non-nil value is passed,
// then an optional with the value will be returned.
func OptionalOfNillable(m Model) Optional {
	if m == nil {
		return OptionalEmpty()
	}
	return Optional{model: m, empty: false}
}

// OptionalEmpty very simply returns an empty Optional that contains
// not related values.
func OptionalEmpty() Optional {
	return Optional{model: nil, empty: true}
}

// IsEmpty simply returns whether the optional contains a value or
// not in a boolean return type.
func (o Optional) IsEmpty() bool {
	return o.empty
}

// Get is meant to return the Model value that is associated with the
// optional. Use this if you can guarantee that the optional is not
// empty. If the optional is empty, then a panic of NoModelFound will
// be thrown.
func (o Optional) Get() Model {
	if o.IsEmpty() {
		panic(NoModelFound)
	}
	return o.model
}

// GetOrElse returns a different Model depending on whether or not the
// optional contains a value or not. Empty optionals will return the passed
// in model and non-empty optionals will return the value it contains.
//
// This can be useful in case where a default value should be used, in the
// case that the given optional is currently empty.
func (o Optional) GetOrElse(other Model)  Model {
	if o.IsEmpty() {
		return other
	}
	return o.model
}

// IfPresent runs the passed in function on the Model when the optional does
// contain a value. If the value is not present, then the function will not run.
func (o Optional) IfPresent(f func(value Model))  {
	if o.IsEmpty() {
		return
	}
	f(o.model)
}