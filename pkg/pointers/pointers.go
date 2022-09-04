package pointers

// ToPointer returns a pointer to `val`.
func ToPointer[T any](val T) *T {
	return &val
}

// FromPointer returns the non-pointer value of the passed pointer.
// If the passed pointer is nil,
// the zero value for type T is returned instead.
func FromPointer[T any](val *T) T {
	var zero T

	if val != nil {
		return *val
	}

	return zero
}
