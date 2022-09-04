package slice

// CmpFunc is any function which takes a variable of type T
// and returns a boolean based on the passed value.
type CmpFunc[T any] func(T) bool
