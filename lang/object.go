//Package lang contains all object definitions for hub
package lang

import (
	"fmt"
	"strconv"
)

//Object is the basic interface for all objects in hub
type Object interface {
	GetHash() int64
	GetType() string
	Equals(Object) bool
	ToString() string
}

//BuildString build a new hub instance of String
func BuildString(_str string) String {
	var str String
	str.value = _str
	return str
}

//BuildNumber build a new hub instance of Number
func BuildNumber(_str string) Number {
	var num Number
	number, err := strconv.ParseFloat(_str, 64)
	if err != nil {
		fmt.Println(_str + " is not a number")
		panic(err.Error())
	}
	num.value = number
	return num
}

//BuildPointer build a new hub instance of Pointer
func BuildPointer(_str string) Pointer {
	var ptr Pointer
	ptr.name = _str
	return ptr
}

//BuildBoolean build a new hub instance of Boolean
func BuildBoolean(_str string) Boolean {
	var ptr Boolean
	v, err := strconv.ParseBool(_str)
	if err != nil {
		panic(err.Error())
	}
	ptr.value = v
	return ptr
}
