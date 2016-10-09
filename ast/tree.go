package ast

import "fmt"

//Tree implements a DataStructure
type Tree struct {
	Value    interface{}
	Children []*Tree
	Parent   *Tree
}

//DeepWalk inside Tree
func (tree *Tree) DeepWalk(callback func(interface{})) {
	if tree.Children != nil {
		for i := 0; i < len(tree.Children); i++ {
			nodes := tree.Children
			node := nodes[i]
			node.DeepWalk(callback)

		}
	}
	callback(tree.Value)
}

var deep int

//Print inside Tree
func (tree *Tree) Print(callback func(interface{})) {
	for i := 0; i < deep*4; i++ {
		fmt.Print("-")
	}
	callback(tree.Value)
	if tree.Children != nil {
		nodes := tree.Children
		for i := 0; i < len(nodes); i++ {
			node := nodes[i]
			deep++
			node.Print(callback)
		}
	}
	deep--
}

//Walk inside the Tree
func (tree *Tree) Walk(callback func(interface{})) {
	callback(tree.Value)
	if tree.Children != nil {
		nodes := tree.Children
		for i := 0; i < len(nodes); i++ {
			node := nodes[i]
			node.Walk(callback)
		}
	}
}

//AppendChild to existing Tree
func (tree *Tree) AppendChild(child *Tree) {
	if tree.Children == nil {
		tree.Children = make([]*Tree, 0)
	}
	(*child).Parent = tree
	tree.Children = append(tree.Children, child)

}

//RemoveChild from existing Tree
func (tree *Tree) RemoveChild(i int) {
	tree.Children = append(tree.Children[:0], tree.Children[1:]...)
}

//HasChildren verify if tree has children
func (tree *Tree) HasChildren() bool {
	return len(tree.Children) > 0
}

//ToString prints a tree node
func (tree *Tree) ToString() string {
	if tree == nil {
		return "nil"
	}
	return fmt.Sprintf("%v", tree.Value)

}
