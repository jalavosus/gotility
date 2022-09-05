# gotility

[![Go Reference](https://pkg.go.dev/badge/github.com/jalavosus/gotility.svg)](https://pkg.go.dev/github.com/jalavosus/gotility)

gotility provides various utility functions, types, and more. 

## Packages/Functions

### pkg/slice

- `First[T any](data []T) T`
  - Returns the first element of a slice, or the zero value of type T if len(data) == 0.

- `Last[T any](data []T) T`
  - Returns the last element of a slice, of the zero value of type T if len(data) == 0.

- `Head[T any](data []T) []T`
  - Returns all elements of a slice except the last.

- `Tail[T any](data []T) []T`
  - Returns all elements of a slice except the first.

- `Find[T any](data []T, fn CmpFunc[T]) (T, bool)`
  - Returns the first element of a slice for which `fn` returns true, along with a boolean representing whether a matching value was found.

- `FindValue[T comparable](data []T, val T)  (T, bool)`
  - Returns the first element of a slice which is equal to `val`, along with a boolean representing whether a matching value was found.

- `Filter[T any](data []T, fn CmpFunc[T]) []T`
  - Returns all values in a slice for which `fn` returns true.

- `FilterValue[T comparable](data []T, val T) []T`
  - Returns all values in a slice which are equal to `val`.

- `Contains[T any](data []T, fn CmpFunc[T]) bool`
  - Returns a boolean representing whether `fn` returns true for any element in the passed slice.

- `ContainsValue[T comparable](data []T, val T) bool`
  - Returns a boolean representing whether any element in the passed slice == `val`.

### pkg/pointers

- `ToPointer[T any](val T) *T`
  - Returns a pointer to `val`.

- `FromPointer[T any](val *T) T`
  - Returns the non-pointer value of the passed pointer. If the passed pointer is nil, the zero value for type T is returned instead.