package ast

//Tree implements a DataStructure
type Tree struct {
	Value    interface{}
	Children []Tree
}

//DeepWalk inside Tree
func (tree *Tree) DeepWalk(callback func(interface{})) {
	if tree.Children != nil {
		for i := 0; i < len(tree.Children); i++ {
			node := tree.Children[i]
			node.DeepWalk(callback)

		}
	}
	callback(tree.Value)
}

//AppendChild to existing Tree
func (tree *Tree) AppendChild(child Tree) {
	if tree.Children == nil {
		tree.Children = make([]Tree, 0)
	}
	tree.Children = append(tree.Children, child)
}
