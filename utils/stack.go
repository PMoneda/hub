//Package utils for utilities types
package utils

import "fmt"

//Stack type is a basic stack
type Stack []interface{}

//Push element at the end of stack
func (s *Stack) Push(v interface{}) {
	*s = append(*s, v)
}

//Pop the last element
func (s *Stack) Pop() interface{} {
	res := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return res
}

//Top shows the last element without pop
func (s *Stack) Top() interface{} {
	res := (*s)[len(*s)-1]
	return res
}

//IsEmpty returns if the stack is empty
func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}

//Print prints the stack
func (s *Stack) Print() {
	for i := 0; i < len(*s); i++ {
		fmt.Print((*s)[i])
		fmt.Print("  ")
	}
	fmt.Println()
}
