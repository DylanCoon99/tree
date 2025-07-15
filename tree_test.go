package tree

import (
	"fmt"
	"log"
	"testing"
)





func InitTreeTest(t *testing.T) {

	var tree *Tree

	tree := InitTree()

	if tree.Root != nil {
		t.Errorf("InitTree failed, expected nil got %v", tree.Root)
	}

	fmt.Println("Successfully init tree.")


}