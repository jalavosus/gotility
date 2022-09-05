package slice

// First returns the first element of a slice.
func First[T any](data []T) T {
	var first T

	if len(data) > 0 {
		first = data[0]
	}

	return first
}

// Last returns the last element of a slice.
func Last[T any](data []T) T {
	var last T

	if len(data) > 0 {
		last = data[len(data)-1]
	}

	return last
}

// Head returns all elements of a slice except the last.
func Head[T any](data []T) []T {
	var head = make([]T, len(data)-1)

	for i := range data[0 : len(data)-1] {
		head[i] = data[i]
	}

	return head
}

// Tail returns all elements of a slice except the first.
func Tail[T any](data []T) []T {
	var tail = make([]T, len(data)-1)

	var idx int
	for _, d := range data[1:] {
		tail[idx] = d
		idx++
	}

	return tail
}
