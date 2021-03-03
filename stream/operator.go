package stream

import . "github.com/Mathew-Estafanous/funGo/model"

// Operator simply takes in a given Model and returns an altered
// or different model of either the same type or of a different
// type entirely.
//
// This is particularly helpful when using the Map()
// function that applies the given operator to every Model in the stream
type Operator func(m Model) Model

// MultiOperator is very similar to the Operator in what it does and
// its main use. The key difference is that the operator requires that it
// returns an array of models from the given model.
//
// This is especially when using the FlatMap() method in streams. It is used
// like a one to many operation.
type MultiOperator func(m Model) []Model