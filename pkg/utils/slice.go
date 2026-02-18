package utils

// Contains checks if a slice contains an element.
func Contains[T comparable](slice []T, item T) bool {
	for _, v := range slice {
		if v == item {
			return true
		}
	}
	return false
}

// ContainsAll checks if a slice contains all elements from another slice.
func ContainsAll[T comparable](slice, items []T) bool {
	for _, item := range items {
		if !Contains(slice, item) {
			return false
		}
	}
	return true
}

// ContainsAny checks if a slice contains any element from another slice.
func ContainsAny[T comparable](slice, items []T) bool {
	for _, item := range items {
		if Contains(slice, item) {
			return true
		}
	}
	return false
}

// IndexOf returns the index of an element in a slice, or -1 if not found.
func IndexOf[T comparable](slice []T, item T) int {
	for i, v := range slice {
		if v == item {
			return i
		}
	}
	return -1
}

// LastIndexOf returns the last index of an element in a slice, or -1 if not found.
func LastIndexOf[T comparable](slice []T, item T) int {
	for i := len(slice) - 1; i >= 0; i-- {
		if slice[i] == item {
			return i
		}
	}
	return -1
}

// Find returns the first element that matches the predicate.
func Find[T any](slice []T, predicate func(T) bool) (T, bool) {
	for _, v := range slice {
		if predicate(v) {
			return v, true
		}
	}
	var zero T
	return zero, false
}

// FindIndex returns the index of the first element that matches the predicate.
func FindIndex[T any](slice []T, predicate func(T) bool) int {
	for i, v := range slice {
		if predicate(v) {
			return i
		}
	}
	return -1
}

// FindLast returns the last element that matches the predicate.
func FindLast[T any](slice []T, predicate func(T) bool) (T, bool) {
	for i := len(slice) - 1; i >= 0; i-- {
		if predicate(slice[i]) {
			return slice[i], true
		}
	}
	var zero T
	return zero, false
}

// Remove removes the element at the given index.
func Remove[T any](slice []T, index int) []T {
	if index < 0 || index >= len(slice) {
		return slice
	}
	return append(slice[:index], slice[index+1:]...)
}

// RemoveElement removes the first occurrence of an element.
func RemoveElement[T comparable](slice []T, item T) []T {
	index := IndexOf(slice, item)
	if index == -1 {
		return slice
	}
	return Remove(slice, index)
}

// RemoveAll removes all occurrences of an element.
func RemoveAll[T comparable](slice []T, item T) []T {
	result := make([]T, 0, len(slice))
	for _, v := range slice {
		if v != item {
			result = append(result, v)
		}
	}
	return result
}

// RemoveAt removes elements at multiple indices.
func RemoveAt[T any](slice []T, indices ...int) []T {
	if len(indices) == 0 {
		return slice
	}

	// Create a set of indices to remove
	removeSet := make(map[int]bool)
	for _, idx := range indices {
		if idx >= 0 && idx < len(slice) {
			removeSet[idx] = true
		}
	}

	result := make([]T, 0, len(slice)-len(removeSet))
	for i, v := range slice {
		if !removeSet[i] {
			result = append(result, v)
		}
	}
	return result
}

// Filter returns a new slice with elements that pass the predicate.
func Filter[T any](slice []T, predicate func(T) bool) []T {
	result := make([]T, 0, len(slice))
	for _, v := range slice {
		if predicate(v) {
			result = append(result, v)
		}
	}
	return result
}

// Map transforms each element using the provided function.
func Map[T, U any](slice []T, fn func(T) U) []U {
	result := make([]U, len(slice))
	for i, v := range slice {
		result[i] = fn(v)
	}
	return result
}

// MapIndexed transforms each element with its index.
func MapIndexed[T, U any](slice []T, fn func(int, T) U) []U {
	result := make([]U, len(slice))
	for i, v := range slice {
		result[i] = fn(i, v)
	}
	return result
}

// FlatMap maps and flattens the result.
func FlatMap[T, U any](slice []T, fn func(T) []U) []U {
	result := make([]U, 0, len(slice))
	for _, v := range slice {
		result = append(result, fn(v)...)
	}
	return result
}

// Unique returns a new slice with duplicate elements removed.
func Unique[T comparable](slice []T) []T {
	seen := make(map[T]bool)
	result := make([]T, 0, len(slice))

	for _, v := range slice {
		if !seen[v] {
			seen[v] = true
			result = append(result, v)
		}
	}
	return result
}

// UniqueBy returns a new slice with duplicates removed based on a key function.
func UniqueBy[T any, K comparable](slice []T, keyFn func(T) K) []T {
	seen := make(map[K]bool)
	result := make([]T, 0, len(slice))

	for _, v := range slice {
		key := keyFn(v)
		if !seen[key] {
			seen[key] = true
			result = append(result, v)
		}
	}
	return result
}

// Reverse returns a new slice with elements in reverse order.
func Reverse[T any](slice []T) []T {
	result := make([]T, len(slice))
	for i, v := range slice {
		result[len(slice)-1-i] = v
	}
	return result
}

// ReverseInPlace reverses a slice in place.
func ReverseInPlace[T any](slice []T) {
	for i, j := 0, len(slice)-1; i < j; i, j = i+1, j-1 {
		slice[i], slice[j] = slice[j], slice[i]
	}
}

// Take returns the first n elements.
func Take[T any](slice []T, n int) []T {
	if n <= 0 {
		return nil
	}
	if n >= len(slice) {
		return slice
	}
	return slice[:n]
}

// TakeWhile takes elements while the predicate is true.
func TakeWhile[T any](slice []T, predicate func(T) bool) []T {
	result := make([]T, 0)
	for _, v := range slice {
		if !predicate(v) {
			break
		}
		result = append(result, v)
	}
	return result
}

// Drop removes the first n elements.
func Drop[T any](slice []T, n int) []T {
	if n <= 0 {
		return slice
	}
	if n >= len(slice) {
		return nil
	}
	return slice[n:]
}

// DropWhile drops elements while the predicate is true.
func DropWhile[T any](slice []T, predicate func(T) bool) []T {
	for i, v := range slice {
		if !predicate(v) {
			return slice[i:]
		}
	}
	return nil
}

// First returns the first element or an error if empty.
func First[T any](slice []T) (T, error) {
	if len(slice) == 0 {
		var zero T
		return zero, ErrSliceEmpty
	}
	return slice[0], nil
}

// FirstOr returns the first element or a default value.
func FirstOr[T any](slice []T, defaultValue T) T {
	if len(slice) == 0 {
		return defaultValue
	}
	return slice[0]
}

// Last returns the last element or an error if empty.
func Last[T any](slice []T) (T, error) {
	if len(slice) == 0 {
		var zero T
		return zero, ErrSliceEmpty
	}
	return slice[len(slice)-1], nil
}

// LastOr returns the last element or a default value.
func LastOr[T any](slice []T, defaultValue T) T {
	if len(slice) == 0 {
		return defaultValue
	}
	return slice[len(slice)-1]
}

// ErrSliceEmpty is returned when accessing an empty slice.
var ErrSliceEmpty = &sliceError{"slice is empty"}

type sliceError struct {
	msg string
}

func (e *sliceError) Error() string {
	return e.msg
}

// Chunk splits a slice into chunks of the given size.
func Chunk[T any](slice []T, size int) [][]T {
	if size <= 0 {
		return nil
	}

	chunks := make([][]T, 0, (len(slice)+size-1)/size)
	for i := 0; i < len(slice); i += size {
		end := i + size
		if end > len(slice) {
			end = len(slice)
		}
		chunks = append(chunks, slice[i:end])
	}
	return chunks
}

// Partition splits a slice into two based on a predicate.
func Partition[T any](slice []T, predicate func(T) bool) (trueSlice, falseSlice []T) {
	for _, v := range slice {
		if predicate(v) {
			trueSlice = append(trueSlice, v)
		} else {
			falseSlice = append(falseSlice, v)
		}
	}
	return
}

// Concat concatenates multiple slices.
func Concat[T any](slices ...[]T) []T {
	totalLen := 0
	for _, s := range slices {
		totalLen += len(s)
	}

	result := make([]T, 0, totalLen)
	for _, s := range slices {
		result = append(result, s...)
	}
	return result
}

// Flatten flattens a slice of slices.
func Flatten[T any](slices [][]T) []T {
	totalLen := 0
	for _, s := range slices {
		totalLen += len(s)
	}

	result := make([]T, 0, totalLen)
	for _, s := range slices {
		result = append(result, s...)
	}
	return result
}

// ForEach applies a function to each element.
func ForEach[T any](slice []T, fn func(T)) {
	for _, v := range slice {
		fn(v)
	}
}

// ForEachIndexed applies a function to each element with its index.
func ForEachIndexed[T any](slice []T, fn func(int, T)) {
	for i, v := range slice {
		fn(i, v)
	}
}

// All returns true if all elements pass the predicate.
func All[T any](slice []T, predicate func(T) bool) bool {
	for _, v := range slice {
		if !predicate(v) {
			return false
		}
	}
	return true
}

// Any returns true if any element passes the predicate.
func Any[T any](slice []T, predicate func(T) bool) bool {
	for _, v := range slice {
		if predicate(v) {
			return true
		}
	}
	return false
}

// None returns true if no element passes the predicate.
func None[T any](slice []T, predicate func(T) bool) bool {
	return !Any(slice, predicate)
}

// Count counts elements that pass the predicate.
func Count[T any](slice []T, predicate func(T) bool) int {
	count := 0
	for _, v := range slice {
		if predicate(v) {
			count++
		}
	}
	return count
}

// CountOf counts occurrences of a specific element.
func CountOf[T comparable](slice []T, item T) int {
	count := 0
	for _, v := range slice {
		if v == item {
			count++
		}
	}
	return count
}

// IsEmpty checks if a slice is empty.
func IsEmpty[T any](slice []T) bool {
	return len(slice) == 0
}

// IsNotEmpty checks if a slice is not empty.
func IsNotEmpty[T any](slice []T) bool {
	return len(slice) > 0
}

// Len returns the length of a slice.
func Len[T any](slice []T) int {
	return len(slice)
}

// Insert inserts an element at the given index.
func Insert[T any](slice []T, index int, item T) []T {
	if index < 0 {
		index = 0
	}
	if index > len(slice) {
		index = len(slice)
	}

	result := make([]T, 0, len(slice)+1)
	result = append(result, slice[:index]...)
	result = append(result, item)
	result = append(result, slice[index:]...)
	return result
}

// InsertAll inserts multiple elements at the given index.
func InsertAll[T any](slice []T, index int, items ...T) []T {
	if len(items) == 0 {
		return slice
	}
	if index < 0 {
		index = 0
	}
	if index > len(slice) {
		index = len(slice)
	}

	result := make([]T, 0, len(slice)+len(items))
	result = append(result, slice[:index]...)
	result = append(result, items...)
	result = append(result, slice[index:]...)
	return result
}

// Clone returns a shallow copy of the slice.
func Clone[T any](slice []T) []T {
	result := make([]T, len(slice))
	copy(result, slice)
	return result
}

// Append appends elements to a slice, creating a new slice.
func Append[T any](slice []T, items ...T) []T {
	result := make([]T, len(slice)+len(items))
	copy(result, slice)
	copy(result[len(slice):], items)
	return result
}

// Prepend prepends elements to a slice.
func Prepend[T any](slice []T, items ...T) []T {
	result := make([]T, len(slice)+len(items))
	copy(result, items)
	copy(result[len(items):], slice)
	return result
}

// Difference returns elements in slice1 that are not in slice2.
func Difference[T comparable](slice1, slice2 []T) []T {
	seen := make(map[T]bool)
	for _, v := range slice2 {
		seen[v] = true
	}

	result := make([]T, 0)
	for _, v := range slice1 {
		if !seen[v] {
			result = append(result, v)
		}
	}
	return result
}

// Intersection returns elements that are in both slices.
func Intersection[T comparable](slice1, slice2 []T) []T {
	seen := make(map[T]bool)
	for _, v := range slice2 {
		seen[v] = true
	}

	result := make([]T, 0)
	for _, v := range slice1 {
		if seen[v] {
			result = append(result, v)
		}
	}
	return result
}

// Union returns unique elements from both slices.
func Union[T comparable](slice1, slice2 []T) []T {
	return Unique(Concat(slice1, slice2))
}

// Equal checks if two slices are equal.
func Equal[T comparable](slice1, slice2 []T) bool {
	if len(slice1) != len(slice2) {
		return false
	}
	for i := range slice1 {
		if slice1[i] != slice2[i] {
			return false
		}
	}
	return true
}

// EqualIgnoringOrder checks if two slices contain the same elements (order-independent).
func EqualIgnoringOrder[T comparable](slice1, slice2 []T) bool {
	if len(slice1) != len(slice2) {
		return false
	}
	return len(Difference(slice1, slice2)) == 0 && len(Difference(slice2, slice1)) == 0
}

// Fill creates a slice filled with the given value.
func Fill[T any](value T, count int) []T {
	if count <= 0 {
		return nil
	}
	result := make([]T, count)
	for i := range result {
		result[i] = value
	}
	return result
}

// Range creates a slice of integers from start to end (exclusive).
func Range(start, end int) []int {
	if start >= end {
		return nil
	}
	result := make([]int, end-start)
	for i := range result {
		result[i] = start + i
	}
	return result
}

// RangeWithStep creates a slice of integers with a step.
func RangeWithStep(start, end, step int) []int {
	if step == 0 || (step > 0 && start >= end) || (step < 0 && start <= end) {
		return nil
	}

	result := make([]int, 0)
	for i := start; (step > 0 && i < end) || (step < 0 && i > end); i += step {
		result = append(result, i)
	}
	return result
}
