package tree

import (
	"fmt"
	//"log"
	"testing"
)




func InitExampleTree() *BinarySearchTree{

	var tree *BinarySearchTree

	tree = InitTree()


	tree.Insert(4)
	tree.Insert(7)
	tree.Insert(6)
	tree.Insert(2)
	tree.Insert(3)
	tree.Insert(1)
	tree.Insert(0)

	return tree

}





func TestInitTree(t *testing.T) {

	var tree *BinarySearchTree

	tree = InitTree()

	if tree.Root != nil {
		t.Errorf("InitTree failed, expected nil got %v", tree.Root)
	}

	fmt.Println("Successfully init tree.")


}



func TestPrettyPrint(t *testing.T) {



	//tree.PrettyPrint()



}




func TestInsert(t *testing.T) {


	tree := InitExampleTree()

	v := 6

	tree.Insert(v)


	tree.PrettyPrint()

	/*
	if err != nil {
		t.Errorf("Error inserting value %v into tree: %v", v, err)
		return
	}
	*/


	fmt.Println("Passed Insert Test.")

}



func TestGetSize(t *testing.T) {

	tree := InitExampleTree()

	n := tree.getSize()

	if n != 7 {
		t.Errorf("Failed Test Get Size, expected 7 got %v", n)
		return
	}

	fmt.Println("Passed Get Size Test")

}




func TestInTree(t *testing.T) {


	tree := InitExampleTree()

	node := tree.InTree(2)

	if node == nil {
		t.Errorf("Test InTree failed; expected node got: %v", node)
		return
	}


	//fmt.Printf("Here is the right and left nodes for %v.    RIGHT: %v    LEFT: %v", node, node.Right, node.Left)


	node = tree.InTree(10)

	if node != nil {
		t.Errorf("Expected nil, got: %v", node)
		return 
	}

	fmt.Println("Passed In Tree Test")

}




func TestInorderTraversalRecursive(t *testing.T) {

	tree := InitExampleTree()

	traversal := tree.InorderTraversalRecursive()

	fmt.Println(traversal)

}


func TestInorderTraversalIterative(t *testing.T) {

	tree := InitExampleTree()

	traversal := tree.InorderTraversalIterative()

	fmt.Println(traversal)

}



func TestPreorderTraversal(t *testing.T) {

	tree := InitExampleTree()

	traversal := tree.PreorderTraversal()

	fmt.Println(traversal)

}



func TestPostorderTraversal(t *testing.T) {

	tree := InitExampleTree()

	traversal := tree.PostorderTraversal()

	fmt.Println(traversal)

}




