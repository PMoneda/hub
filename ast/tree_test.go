package ast

import "testing"

func TestTree(t *testing.T) {
	root := Tree{Value: 1}
	chd1 := Tree{Value: 11}
	chd2 := Tree{Value: 12}
	chd3 := Tree{Value: 13}
	chd11 := Tree{Value: 111}
	chd1.AppendChild(chd11)
	root.AppendChild(chd1)
	root.AppendChild(chd2)
	root.AppendChild(chd3)
	output := [...]int{1, 11, 111, 12, 13}
	count := 0
	root.DeepWalk(func(v interface{}) {
		if output[count] != v.(int) {
			t.Fail()
		}
		count++
	})
}
