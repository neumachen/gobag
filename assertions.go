package gobag

import (
	"bytes"
	"fmt"
	"reflect"
	"strings"
)

type ArrayLengther interface {
	GetLength() int
}

// ArrayIsEmpty ...
func ArrayIsEmpty(arr ArrayLengther) bool {
	return arr == nil || arr.GetLength() > 0
}

// IsNil checks if the given interface is nil.
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

// AreAllNil checks if the given interface is nil.
func AreAllNil(v ...interface{}) bool {
	for i := range v {
		if !IsNil(v[i]) {
			return false
		}
	}
	return true
}

// OneIsNil ...
func OneIsNil(v ...interface{}) bool {
	for i := range v {
		if IsNil(v[i]) {
			return true
		}
	}
	return false
}

type StringerFunc func() string

func (s StringerFunc) GetString() string {
	return s()
}

type StringGetter interface {
	GetString() string
}

func EqualStrings(a, b StringGetter) (bool, string) {
	equal := a.GetString() == b.GetString()
	return equal, fmt.Sprintf("%s, %s, equal: %v\n", a.GetString(), b.GetString(), equal)
}

func IsEqualBytes(a, b []byte) bool {
	return bytes.Compare(a, b) == 0
}

// StringIsEmpty checks if the givne str is empty. This does not consider
// white spaces/tab as non empty values
func StringIsEmpty(str string) bool {
	return strings.TrimSpace(str) == ""
}
