package tree 


import (
	"fmt"
	"errors"
	"math"
)


/*
	Tree Functionality:

	Add (Insert) while preserving properties of BST (Done)

	Remove (Delete) while preserving properties of BST 

	Print (Display) different traversals (do some of them recursively )
		- Inorder
		- Preorder
		- Postorder
		- Level-order 

	Additional Supporting Functionality
		- Search: Determine if value is in the tree (Done)
		- Size: return number of nodes in the tree (Done)
		- Height/Depth: returns the max depth of the tree (Done)


*/




type Node struct {
	Value  int
	Right  *Node
	Left   *Node
}


type BinarySearchTree struct {
	Root *Node
	Size int
}


func InitTree() *BinarySearchTree {

	return &BinarySearchTree{
		Root: nil,
		Size: 0,
	}

}


func (t *BinarySearchTree) getSize() int {

	return t.Size

}



func getHeight(node *Node) int {

	if node == nil {
		return 0
	}

	return 1 + max(getHeight(node.Left), getHeight(node.Right))

}


func (t *BinarySearchTree) PrettyPrint() {
    if t.Root == nil {
        fmt.Println("Empty tree")
        return
    }

    // Get tree height for buffer sizing
    height := getHeight(t.Root)
    // Width needed: for each node at max depth, we need space for value and padding
    width := int(math.Pow(2, float64(height))) * 4

    // Create a 2D buffer to store the output
    lines := make([][]rune, height*2-1) // Each node level + connector level
    for i := range lines {
        lines[i] = make([]rune, width)
        for j := range lines[i] {
            lines[i][j] = ' '
        }
    }

    // Fill the buffer with nodes and connectors
    placeNode(t.Root, 0, width/2, width/4, lines)

    // Print the buffer
    for _, line := range lines {
        fmt.Println(string(line))
    }
}
// placeNode places a node and its connectors in the buffer
func placeNode(node *Node, row int, col int, offset int, lines [][]rune) {
    if node == nil {
        return
    }

    // Place node value (centered)
    valueStr := fmt.Sprintf("%d", node.Value)
    for i, ch := range valueStr {
        lines[row][col-len(valueStr)/2+i] = ch
    }

    // Place connectors and children
    if node.Left != nil || node.Right != nil {
        // Place connectors (/)
        if node.Left != nil {
            lines[row+1][col-offset/2] = '/'
        }
        // Place connectors (\)
        if node.Right != nil {
            lines[row+1][col+offset/2] = '\\'
        }
        // Recursively place left and right children
        placeNode(node.Left, row+2, col-offset, offset/2, lines)
        placeNode(node.Right, row+2, col+offset, offset/2, lines)
    }
}


func (t *BinarySearchTree) Insert(val int) error {

	// inserts a value while preserving BST balance


	cur := t.Root


	node := &Node{
		Value: val,
		Right: nil,
		Left: nil,
	}



	if t.Root == nil {
		t.Root = node
		t.Size += 1
		return nil
	}


	for {
		if val > cur.Value {
			// insert to the right subtree
			if cur.Right == nil {
				cur.Right = node
				t.Size += 1
				return nil
			}
			cur = cur.Right
		} else if val < cur.Value {
			if cur.Left == nil {
				cur.Left = node
				t.Size += 1
				return nil
			}
			cur = cur.Left
		} else {
			return errors.New("Value is already in tree.")
		}
	}



}



func (t *BinarySearchTree) InTree(val int) bool {

	// returns if a value is in the tree; The tree is a BST so we can determine by doing a binary search

	cur := t.Root

	for {
		if cur.Value == val {
			// we found the value
			return true
		} else {
			if cur.Value > val {
				cur = cur.Left
				if cur == nil {
					return false
				}
			} else {
				cur = cur.Right
				if cur == nil {
					return false
				}
			}
		}
	}

}



func (t *BinarySearchTree) InorderTraversal() []int {

	// Maybe try creating a helper function with parameter called values []int


	return nil
}



func (t *BinarySearchTree) PreorderTraversal() []int {

	
	return nil	
}



func (t *BinarySearchTree) PostorderTraversal() []int {

	
	return nil
}

