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


	//tree.PrettyPrint()



}




func TestInsert(t *testing.T) {


	var tree *BinarySearchTree

	tree = InitTree()


	tree.Insert(4)
	tree.Insert(7)
	tree.Insert(6)
	tree.Insert(2)
	tree.Insert(3)
	tree.Insert(1)
	tree.Insert(0)

	v := 6

	err := tree.Insert(v)


	tree.PrettyPrint()

	if err != nil {
		t.Errorf("Error inserting value %v into tree: %v", v, err)
	}

}