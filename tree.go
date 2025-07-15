package tree 


type Node struct {
	Value  int
	Right  *Node
	Left   *Node
}

type Tree struct {
	Root *Node
}




func InitTree() *Tree {

	return &Tree{
		Root: nil,
	}

}