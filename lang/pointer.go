package lang

import "strconv"

//Pointer is the basic pointer reference for hub
type Pointer struct {
	name  string
	value Object
	hash  int64
}

//GetType of Number
func (pointer Pointer) GetType() string {
	return "Pointer"
}

//GetValue of pointer
func (pointer Pointer) GetValue() Object {
	return pointer.value
}

//GetHash address of reference
func (pointer Pointer) GetHash() int64 {
	return pointer.hash
}

//Equals assert equality between two number
func (pointer Pointer) Equals(obj Object) bool {
	if obj.GetType() == "Pointer" {
		s := obj.(Pointer)
		return s.GetValue().Equals(pointer.GetValue())
	}
	return false
}

//ToString return a string representation about Number instance
func (pointer Pointer) ToString() string {
	if pointer.value == nil {
		return "[ident:" + pointer.name + " , hash: " + strconv.FormatInt(pointer.hash, 10) + ", value: nil ]"
	}
	return "[ident:" + pointer.name + " , hash: " + strconv.FormatInt(pointer.hash, 10) + ", value: " + pointer.value.ToString() + " ]"
}
