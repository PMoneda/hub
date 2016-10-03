package lang

//Operator is the basic operator reference for hub
type Operator struct {
	symbol string
}

//GetType of Operator
func (operator Operator) GetType() string {
	return "Operator"
}

//GetValue of operator
func (operator Operator) GetValue() string {
	return operator.symbol
}

//GetHash address of reference
func (operator Operator) GetHash() int64 {
	panic("Not support")
}

//Equals assert equality between two operators
func (operator Operator) Equals(obj Object) bool {
	if obj.GetType() == "Operator" {
		s := obj.(Operator)
		return s.symbol == operator.symbol
	}
	return false
}

//ToString return a string representation about Operator instance
func (operator Operator) ToString() string {
	return operator.symbol
}
