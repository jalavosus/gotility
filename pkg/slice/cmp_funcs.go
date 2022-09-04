package slice

// Find returns the first element of a slice for which `fn` returns true,
// along with a boolean representing whether a matching value was found.
func Find[T any](data []T, fn CmpFunc[T]) (T, bool) {
	var (
		found T
		ok    bool
	)

	for i := range data {
		if fn(data[i]) {
			found = data[i]
			ok = true
			break
		}
	}

	return found, ok
}

// FindValue returns the first element of a slice which is equal to `val`,
// along with a boolean representing whether a matching value was found.
func FindValue[T comparable](data []T, val T) (T, bool) {
	return Find(data, func(v T) bool {
		return v == val
	})
}

// Filter returns all values in a slice for which `fn` returns true.
func Filter[T any](data []T, fn CmpFunc[T]) []T {
	var filtered []T

	for i := range data {
		if fn(data[i]) {
			filtered = append(filtered, data[i])
		}
	}

	return filtered
}

// FilterValue returns all values in a slice which are equal to `val`.
func FilterValue[T comparable](data []T, val T) []T {
	return Filter(data, func(v T) bool {
		return v == val
	})
}

// Contains returns a boolean representing whether `fn`
// returns true for any element in the passed slice.
func Contains[T any](data []T, fn CmpFunc[T]) bool {
	var contains bool

	for i := range data {
		if fn(data[i]) {
			contains = true
			break
		}
	}

	return contains
}

// ContainsValue returns a boolean representing whether
// any element in the passed slice == `val`.
func ContainsValue[T comparable](data []T, val T) bool {
	return Contains(data, func(v T) bool {
		return v == val
	})
}
