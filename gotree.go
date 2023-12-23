package gotree

import (
	"fmt"

	"github.com/jwalton/gchalk"
	utils "github.com/zilayo/go-tree/internal"
)

////////////////////////////////////////////////////////////////////////////////////////////
//                                TYPES, CONSTANTS & GLOBALS                              //
////////////////////////////////////////////////////////////////////////////////////////////

// NodeCategory represents which type of node should be added to the tree
//   - Messages: 0
//   - Line break: 1
//   - Parent: 2
type NodeCategory int

const (
	// Messages represents a node with 1 or more message
	Messages NodeCategory = iota

	// Empty represents an empty node, which is used for line breaks
	Empty

	// Parent represents the tree's outer node.
	Parent
)

// Tree represents the full tree, containing all of the child nodes.
type Tree struct {
	// Nodes is a slice containing every Node in the tree.
	Nodes []*Node
}

// Node represents a single node of the tree
type Node struct {
	// Category is the type of node being added.
	//  - Messages: a slice of messages to add to the node
	//  - Empty: a line break to add to the node.
	//  - Parent: the very top level node of the tree.
	Category NodeCategory

	// messages are the messages to add to the node.
	messages []string

	// parent represents the parent node's index in the tree.
	// Used to decide which node this node should branch from.
	parent int

	// children represents the child nodes of the current node.
	children []int
}

var chalk = gchalk.New(gchalk.ForceLevel(gchalk.LevelBasic))

var (
	brightWhiteBold   = chalk.WithBold().WithBrightWhite().Sprintf
	brightYellowBold  = chalk.WithBold().WithBrightYellow().Sprintf
	brightRedBold     = chalk.WithBold().WithBrightRed().Sprintf
	brightMagentaBold = chalk.WithBold().WithBrightMagenta().Sprintf
	brightCyanBold    = chalk.WithBold().WithBrightCyan().Sprintf
)

///////////////////////////////////////////////////////////////////////////
//                                MAIN LOGIC                             //
//////////////////////////////////////////////////////////////////////////

// NewNode creates a new Tree.
// Returns an instance of the tree, and the index of the parent node
func NewTree(treeName string) (*Tree, int) {
	tree := &Tree{
		Nodes: make([]*Node, 0),
	}

	tree.AddParent(0, treeName)

	return tree, 1
}

// Add is a low level function to create and appends a node to the tree.
// Returns the index of the new node.
//   - category: the type of node to add (Parent, Messages, or Empty)
//   - parentIndex: the index of the parent node that this new node will belong to
//   - message: the message to display.
func (t *Tree) Add(category NodeCategory, parentIndex int, messages []string) int {
	// Build the new tree
	node := NewNode(category, parentIndex, messages)
	nodeIndex := len(t.Nodes) + 1

	// Add the children indices to the parent
	if node.parent != 0 {
		if len(t.Nodes) < node.parent {
			panic("Failed to build tree. Child nodes can't be added before the parent node")
		}

		parent := t.Nodes[node.parent-1]

		// Add the child index to the parent
		parent.children = append(parent.children, nodeIndex)
	}

	t.Nodes = append(t.Nodes, &node)

	// Return the index of the new tree
	return nodeIndex
}

// Display will print the tree to standard output.
func (t Tree) Display() {
	for i := 0; i < len(t.Nodes); i++ {
		node := t.Nodes[i]

		// Print only the root Nodes
		if node.parent == 0 {
			t.printNode(" ", i)
		}
	}
}

// printNode prints a single node in the tree with the given prefix.
// This is a recursive function to print all of a node's children.
func (t *Tree) printNode(prefix string, index int) {
	node := t.Nodes[index]

	switch node.Category {
	case Empty:
		fmt.Println(
			brightWhiteBold(utils.ReplaceLast(prefix, "│ ", " │ ")),
		)

	case Parent:
		// Print the tree title
		fmt.Printf(
			"%s %s\n",
			brightWhiteBold(utils.ReplaceLast(prefix, "│ ", " ├─")),
			node.messages[0],
		)

		// Print the child nodes
		for _, child := range node.children {
			t.printNode(fmt.Sprintf("%s   │", prefix), child-1)
		}

		if len(node.children) > 0 {
			// Print the closing statement
			fmt.Printf(
				"%s ←\n",
				brightWhiteBold("%s   └─", prefix),
			)
		}

	case Messages:
		for i := 0; i < len(node.messages); i++ {
			message := node.messages[i]

			if i == 0 {
				prefix = utils.ReplaceLast(prefix, "│ ", " ├─")
			} else if i != len(node.messages)-1 {
				prefix = utils.ReplaceLast(prefix, "│ ", " │ ")
			}

			fmt.Printf(
				"%s %s\n",
				brightWhiteBold(prefix),
				message,
			)
		}

		// Print the children
		for i, child := range node.children {
			if i == len(node.children)-1 {
				t.printNode(
					fmt.Sprintf("%s   └─", prefix),
					child-1,
				)
			} else {
				t.printNode(
					fmt.Sprintf("%s   │", prefix),
					child-1,
				)
			}
		}

	default:
	}
}

///////////////////////////////////////////////////////////////////////////
//                                HELPERS                                //
//////////////////////////////////////////////////////////////////////////

// AddParent creates and adds a Parent node to the tree.
//   - parentIndex: the index of the parent node that this new node will belong to
//   - message: the message to display.
func (t *Tree) AddParent(parentIndex int, message string) int {
	return t.Add(Parent, parentIndex, []string{message})
}

// AddInfo adds an info message to the tree.
// Returns the index of the new tree.
//   - parentIndex: the index of the parent node that this new node will belong to
//   - message: the message to display.
func (t *Tree) AddInfo(parentIndex int, message string) int {
	msg := fmt.Sprintf(
		"%s %s",
		brightCyanBold("info:"),
		message,
	)
	return t.Add(Parent, parentIndex, []string{msg})
}

// AddDebug adds a debug message to the tree.
// Returns the index of the new tree.
//   - parentIndex: the index of the parent node that this new node will belong to
//   - message: the message to display.
func (t *Tree) AddDebug(parentIndex int, message string) int {
	msg := fmt.Sprintf(
		"%s %s",
		brightMagentaBold("debug:"),
		message,
	)
	return t.Add(Parent, parentIndex, []string{msg})
}

// AddError adds an error message to the tree.
// Returns the index of the new tree.
//   - parentIndex: the index of the parent node that this new node will belong to
//   - message: the message to display.
func (t *Tree) AddError(parentIndex int, message string) int {
	msg := fmt.Sprintf(
		"%s %s",
		brightRedBold("error:"),
		message,
	)
	return t.Add(Parent, parentIndex, []string{msg})
}

// AddWarn adds a warn message to the tree.
// Returns the index of the new tree.
//   - parentIndex: the index of the parent node that this new node will belong to
//   - message: the message to display.
func (t *Tree) AddWarn(parentIndex int, message string) int {
	msg := fmt.Sprintf(
		"%s %s",
		brightYellowBold("warn:"),
		message,
	)
	return t.Add(Parent, parentIndex, []string{msg})
}

// AddMessage adds a single message to the tree
//   - parentIndex: the index of the parent node that this new node will belong to
//   - message: the message to display.
func (t *Tree) AddMessage(parentIndex int, message string) int {
	return t.Add(Messages, parentIndex, []string{message})
}

// AddMessages adds a slice of messages to the tree
//   - parentIndex: the index of the parent node that this new node will belong to
//   - messages: a slice of messages to display.
func (t *Tree) AddMessages(parentIndex int, messages []string) int {
	return t.Add(Messages, parentIndex, messages)
}

// AddBreak adds a line break to the tree
//   - parentIndex: the index of the parent node that this new node will belong to
func (t *Tree) AddBreak(parentIndex int) int {
	return t.Add(Empty, parentIndex, []string{""})
}

// NewNode creates a new node with the given parameters
//   - category: the type of node (Parent, Messages, or Empty)
//   - parentIndex: the index of the parent node that this new node will belong to
//   - messages: a slice of messages to display
func NewNode(category NodeCategory, parentIndex int, message []string) Node {
	return Node{
		Category: category,
		messages: message,
		parent:   parentIndex,
		children: []int{},
	}
}
