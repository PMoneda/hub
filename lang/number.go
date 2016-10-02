package lang

import "strconv"

//Number is the basic number class for hub
type Number struct {
	name  string
	value float64
	hash  int64
}

//GetType of Number
func (number *Number) GetType() string {
	return "Number"
}

//GetValue of string
func (number *Number) GetValue() float64 {
	return number.value
}

//GetHash address of reference
func (number *Number) GetHash() int64 {
	return number.hash
}

//Equals assert equality between two number
func (number *Number) Equals(obj Object) bool {
	if obj.GetType() == "Number" {
		s := obj.(*Number)
		return s.GetValue() == number.GetValue()
	}
	return false
}

//ToString return a string representation about Number instance
func (number *Number) ToString() string {
	return strconv.FormatFloat(number.value, 'E', -1, 64)
}
