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

	Print (Display) different traversals (do some of them recursively; some iteratively)
		- DFS
			- Inorder
			- Preorder
			- Postorder
		- Level-order 

	Additional Supporting Functionality
		- Search: Determine if value is in the tree and return node (Done)
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



func (t *BinarySearchTree) InTree(val int) *Node {

	// returns a ptr to the node if it is in the tree; nil otherewise
	//The tree is a BST so we can determine by doing a binary search

	cur := t.Root

	for {
		if cur.Value == val {
			// we found the value
			return cur
		} else {
			if cur.Value > val {
				cur = cur.Left
				if cur == nil {
					return nil
				}
			} else {
				cur = cur.Right
				if cur == nil {
					return nil
				}
			}
		}
	}

}



func findSuccessor(node *Node) *Node{
	// find the next biggest node in the right subtree


	cur := node.Right


	for cur != nil && cur.Left != nil {
		cur = cur.Left
	}

	return cur

}






func (t *BinarySearchTree) Delete(val int) error {

	node := t.InTree(val)

	if node == nil {
		return errors.New("Error: Value not in tree.")
	}


	if node.Right == nil && node.Left == nil {
		// node is a leaf; easier delete


		return nil
	}


	// if the node has two children we have to find the inorder successor or predecessor of the node

	// if the node has one child tree, find the inorder node from that sub tree



	return nil


}


func InorderTraversalHelper(node *Node, vals *[]int) {

	// base case
	if node == nil {
		return
	}

	InorderTraversalHelper(node.Left, vals)
	*vals = append(*vals, node.Value)
	InorderTraversalHelper(node.Right, vals)

	return
}



func (t *BinarySearchTree) InorderTraversalRecursive() []int {

	// Maybe try creating a helper function with parameter called values []int

	vals := []int{}

	InorderTraversalHelper(t.Root, &vals)

	return vals
}



func (t *BinarySearchTree) InorderTraversalIterative() []int {

	// Maybe try creating a helper function with parameter called values []int

	// use a stack to track nodes while traversing as far left as possible

	// pop a node, process (append) its value to the vals, then move to its right subtree

	// repeat until stack is empty and no more nodes remain

	if t.Root == nil {
		return []int{}
	}


	vals := []int{}

	current := t.Root

	stack := []*Node{}

	for current != nil || len(stack) > 0 {

		// use a stack to track nodes while traversing as far left as possible
		for current != nil {
			stack = append(stack, current)
			current = current.Left
		}

		// pop a node
		current = stack[len(stack) - 1]
		stack = stack[:len(stack) - 1]

		// append the node value to the vals
		vals = append(vals, current.Value)

		current = current.Right

	}

	return vals
}



func (t *BinarySearchTree) PreorderTraversal() []int {

	// Visits Root -> Left -> Right


	if t.Root == nil {
		return []int{}
	}


	vals := []int{}

	stack := []*Node{t.Root}

	for len(stack) > 0 {

		// visit root
		current := stack[len(stack) - 1]
		stack = stack[:len(stack) - 1]

		vals = append(vals, current.Value)

		// push right node to stack so it is processed after left
		if current.Right != nil {
			stack = append(stack, current.Right)
		}

		// push left node to stack so it is processed first
		if current.Left != nil {
			stack = append(stack, current.Left)
		}

	}
	
	return vals
}



func (t *BinarySearchTree) PostorderTraversal() []int {

	// Visits in order Left -> Right -> Root

	// Traverse Left Subtree, the Right Subtree, and then the Root

	// Do it iteratively

	
	vals := []int{}
	stack := []*Node{}

	var lastVisited *Node

	current := t.Root

	for current != nil || len(stack) > 0 {

		// go the furthest left
		for current != nil {
			stack = append(stack, current)
			current = current.Left
		}

		current = stack[len(stack) - 1]

		// if the right child exists and it hasn't been visited, process it
		if current.Right != nil && current.Right != lastVisited {
			current = current.Right
		} else {
			vals = append(vals, current.Value)
			lastVisited = current
			stack = stack[:len(stack) - 1]
			current = nil
		}

	}

	return vals
}

