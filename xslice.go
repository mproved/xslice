package xslice

import (
	"unsafe"
)

func CopyToType[T any, U any](source []T) []U {
	if len(source) == 0 {
		return []U{}
	}

	return unsafe.Slice((*U)(unsafe.Pointer(&source[0])), len(source))
}

func AppendIfNotNil[T any](slice *[]T, item *T) {
	if item != nil {
		*slice = append(*slice, *item)
	}
}

func AppendIfNew[T comparable](slice *[]T, item T) {
	if !Has(*slice, item) {
		*slice = append(*slice, item)
	}
}

func Find[T any](slice []T, filter func(item *T) bool) (bool, int) {
	for index, item := range slice {
		if filter(&item) {
			return true, index
		}
	}

	return false, -1
}

func Has[T comparable](slice []T, item T) bool {
	return HasFilter(
		slice,
		func(thisItem *T) bool {
			return *thisItem == item
		},
	)
}

func HasFilter[T any](slice []T, filter func(item *T) bool) bool {
	for _, item := range slice {
		if filter(&item) {
			return true
		}
	}

	return false
}

func Remove[T comparable](slice *[]T, item T) []T {
	return RemoveFilter(
		slice,
		func(this *T) bool {
			return *this == item
		},
	)
}

func RemoveFilter[T any](slice *[]T, filter func(item *T) bool) []T {
	keptItems := []T{}
	removedItems := []T{}

	for _, item := range *slice {
		if filter(&item) {
			removedItems = append(removedItems, item)
		} else {
			keptItems = append(keptItems, item)
		}
	}

	*slice = keptItems

	return removedItems
}

func EqualUnordered[T comparable](slice1 []T, slice2 []T) bool {
	if len(slice1) != len(slice2) {
		return false
	}

	m := make(map[T]struct{}, len(slice1))

	for _, item := range slice1 {
		m[item] = struct{}{}
	}

	for _, item := range slice2 {
		_, ok := m[item]

		if !ok {
			return false
		}
	}

	return true
}

func Merge[T any](slices ...[]T) []T {
	maxSize := 0

	for _, s := range slices {
		if len(s) > maxSize {
			maxSize = len(s)
		}
	}

	r := make([]T, maxSize)

	for _, s := range slices {
		for k, v := range s {
			r[k] = v
		}
	}

	return r
}
