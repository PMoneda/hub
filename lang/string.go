package lang

//String is a basic object in hub lang
type String struct {
	name  string
	value string
	hash  int64
}

//GetType of String
func (str *String) GetType() string {
	return "String"
}

//GetValue of string
func (str *String) GetValue() string {
	return str.value
}

//GetHash address of reference
func (str *String) GetHash() int64 {
	return str.hash
}

//Equals assert equality between two strings
func (str *String) Equals(obj Object) bool {
	if obj.GetType() == "String" {
		s := obj.(*String)
		return s.GetValue() == str.GetValue()
	}
	return false
}

//ToString return a string representation about String instance
func (str *String) ToString() string {
	return str.value
}
