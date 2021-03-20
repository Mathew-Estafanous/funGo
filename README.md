# FunGo
[![Go Report Card](https://goreportcard.com/badge/github.com/Mathew-Estafanous/funGo)](https://goreportcard.com/report/github.com/Mathew-Estafanous/funGo)
![](http://godoc.org/github.com/Mathew-Estafanous/funGo?status.svg)
----
### What is This?
FunGo is an easy-to-use functional programming package that brings a simple
declarative approach to golang. Handle large slices of data in a descriptive
approach. It is always easier to skim and read `ForEach` than `for v := range mySlice` 
and that is exactly what this package gives you. It abstracts the specific details
and provides a simple and flexible API that **describes** what you are doing. Leave the
rest to the package to handle.

### Inspiration
Golang is a great language for its simplicity and ease of use. However, I felt that using
a declarative approach would provide more clarity in my code, instead of having many loops.
Java Streams, was what I turned to for inspiration. Their approach to such a system felt
simple, easy to use and exactly what I was looking for. So, I decided to try my hand at
developing something like that.

----
## How To Use
The main focus of this package is the `stream` package. This package is what you will use
when you are planning iterate over a slice of data and apply a variety of operations.

When using streams, there are 3 main stages that must be present:
- [Creation](#Creation) - Only ONE
- [Non-Terminal Operation](#Non Terminal Operation) - Unlimited
- [Terminal Operation](#Terminal Operation) - Only ONE

### Creation
All streams start with the creation step. Note that a `Stream` requires that all types have
the simple `Equal` behaviour defined in the `Model` interface. You are free to use any type
that relates to your domain as long as they respect the model interface rules. 

There are two ways to create a stream. The first, is `NewStreamFromSlice`, which takes in a
slice of models. The second is `NewStream` which takes in a model channel. **Note: When passing
in a channel, it is YOUR job to close the channel.**

There are several basic Model types that are provided:
- ModelInt
- ModelByte
- ModelFloat
- ModelMap
- ModelSlice

### Non Terminal Operation
This stage is where the bulk of the operation will occur. There is a wide variety of operations
such as `Filter` and `Map`. You can find a full list of them all in the [godoc.](https://pkg.go.dev/github.com/Mathew-Estafanous/funGo/stream#Stream)
These operations are best used as a series of chain operations that are combined to come to a
final result. There is no limit to how many of these operations you can chain together.


### Terminal Operation
This is the final stage of the stream, and is meant as the operation to get the final result. 
Such as `Count` which will return the total number of remaining models within the stream. There
are also iterative terminal operations like `ForEeach`.

An important terminal operation is the `Collect` operation, which takes the remaining elements and
collects them into a dataset. You can collect the remaining elements into a Map, Slice or any other
datatype that you chose.

---
## Examples

Creation & Non-Terminal Operation example
```go
// COMING SOON
```

---
## Connect or Contact

**Email** - mathewestafanous13@gmail.com

**Website** - https://mathewestafanous.com

**Github** - https://github.com/Mathew-Estafanous