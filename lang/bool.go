package lang

//Boolean is the basic boolean reference for hub
type Boolean struct {
	value bool
	hash  int64
}

//GetType of Boolean
func (boolean Boolean) GetType() string {
	return "Boolean"
}

//GetValue of pointer
func (boolean Boolean) GetValue() bool {
	return boolean.value
}

//GetHash address of reference
func (boolean Boolean) GetHash() int64 {
	return boolean.hash
}

//Equals assert equality between two number
func (boolean Boolean) Equals(obj Object) bool {
	if obj.GetType() == "Boolean" {
		s := obj.(Boolean)
		return s.value == boolean.value
	}
	return false
}

//ToString return a string representation about Number instance
func (boolean Boolean) ToString() string {
	if boolean.value {
		return "true"
	}
	return "false"
}
