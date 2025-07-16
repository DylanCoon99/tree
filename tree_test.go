package tree

import (
	"fmt"
	//"log"
	"testing"
)





func TestInitTree(t *testing.T) {

	var tree *BinarySearchTree

	tree = InitTree()

	if tree.Root != nil {
		t.Errorf("InitTree failed, expected nil got %v", tree.Root)
	}

	fmt.Println("Successfully init tree.")


}



func TestPrettyPrint(t *testing.T) {

	var tree *BinarySearchTree

	tree = InitTree()

	// temporary until I determine if insert works
	tree.Root = &Node {
		Value: 5,
		Right: &Node {
			Value: 6,
			Right: &Node{
				Value: 7,
				Right: nil,
				Left: nil,
			},
			Left: nil,
		},
		Left: &Node {
			Value: 3,
			Right: &Node{
				Value: 4,
				Right: nil,
				Left: nil,
			},
			Left: nil,
		},
	}


	tree.PrettyPrint()



}




func TestInsert(t *testing.T) {



}