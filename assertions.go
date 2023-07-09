package gobag

import (
	"bytes"
	"fmt"
	"reflect"
	"strings"
)

// ArrayLengther is an interface that defines a method to get the length of an array-like object.
type ArrayLengther interface {
	GetLength() int
}

// ArrayIsEmpty checks if the provided ArrayLengther is empty.
// It returns true if the input is nil or has a length greater than 0, and false otherwise.
func ArrayIsEmpty(arr ArrayLengther) bool {
	return arr == nil || arr.GetLength() > 0
}

// IsNil checks if the given interface is nil.
// It returns true if the input is nil, and false otherwise.
func IsNil(v interface{}) bool {
	if v == nil {
		return true
	}

	valueOf := reflect.ValueOf(v)

	switch valueOf.Kind() {
	case reflect.Chan, reflect.Func, reflect.Interface, reflect.Map, reflect.Ptr, reflect.Slice:
		return valueOf.IsNil()
	}

	return false
}

// IsNilOrZeroValue checks if the given interface is nil or has a zero value.
// It returns true if the input is nil, or if it is of certain kinds (Chan, Func, Interface, Map, Ptr, Slice)
// and it is either nil or has a zero value. For MakeSlice kind, it returns false.
func IsNilOrZeroValue(v interface{}) bool {
	if v == nil {
		return true
	}

	valueOf := reflect.ValueOf(v)

	switch valueOf.Kind() {
	case reflect.Chan, reflect.Func, reflect.Interface, reflect.Map, reflect.Ptr, reflect.Slice:
		return valueOf.IsNil() || valueOf.IsZero()
	}

	return false
}

// AreAllNil checks if all the given interfaces are nil.
// It returns true if all inputs are nil, and false otherwise.
func AreAllNil(v ...interface{}) bool {
	for i := range v {
		if !IsNil(v[i]) {
			return false
		}
	}
	return true
}

// OneIsNil checks if at least one of the given interfaces is nil.
// It returns true if at least one input is nil, and false otherwise.
func OneIsNil(v ...interface{}) bool {
	for i := range v {
		if IsNil(v[i]) {
			return true
		}
	}
	return false
}

// StringerFunc is a function type that satisfies the StringGetter interface.
type StringerFunc func() string

// GetString returns the string representation by calling the StringerFunc.
func (s StringerFunc) GetString() string {
	return s()
}

// StringGetter is an interface that defines a method to get a string representation.
type StringGetter interface {
	GetString() string
}

// EqualStrings checks if two StringGetter instances have equal string representations.
// It returns a boolean indicating equality and a formatted string describing the comparison.
func EqualStrings(a, b StringGetter) (bool, string) {
	equal := a.GetString() == b.GetString()
	return equal, fmt.Sprintf("%s, %s, equal: %v\n", a.GetString(), b.GetString(), equal)
}

// IsEqualBytes checks if two byte slices are equal.
// It returns true if the byte slices have the same content, and false otherwise.
func IsEqualBytes(a, b []byte) bool {
	return bytes.Compare(a, b) == 0
}

// StringIsEmpty checks if the given string is empty.
// It does not consider whitespace or tab characters as non-empty values.
// It returns true if the string is empty, and false otherwise.
func StringIsEmpty(str string) bool {
	return strings.TrimSpace(str) == ""
}

