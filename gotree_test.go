package gotree

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_CreateTree(t *testing.T) {
	assert := assert.New(t)
	tree, _ := NewTree("Test::TestTree")

	assert.Equal(1, len(tree.Nodes))

	tree.Display()
}

func TestAddSingleMessage(t *testing.T) {
	assert := assert.New(t)
	tree, parentIndex := NewTree("Test::TestTree")

	tree.AddMessage(parentIndex, "first message")

	assert.Equal(2, len(tree.Nodes))
	tree.Display()
}

func TestAddMultipleMessages(t *testing.T) {
	assert := assert.New(t)
	tree, parentIndex := NewTree("Test::TestTree")

	tree.AddMessages(parentIndex, []string{"first message", "second message"})

	assert.Equal(2, len(tree.Nodes))
	tree.Display()
}

func TestAddChildNodes(t *testing.T) {
	assert := assert.New(t)
	tree, parentIndex := NewTree("Test::TestTree")

	tree.AddMessage(parentIndex, "first message")

	innerIndex := tree.AddParent(parentIndex, "inner branch")
	tree.AddMessage(innerIndex, "inner child")

	deeperIndex := tree.AddParent(innerIndex, "deeper branch")
	tree.AddMessage(deeperIndex, "deeper child")

	assert.Equal(6, len(tree.Nodes))

	tree.Display()
}

func TestAddInvalidParent(t *testing.T) {
	assert := assert.New(t)

	tree, parentIndex := NewTree("Test::TestTree")

	assert.Panics(func() { tree.AddMessage(parentIndex+1, "first message") })

	assert.Equal(1, len(tree.Nodes))

	tree.Display()
}

func TestAddLevelMessages(t *testing.T) {
	assert := assert.New(t)

	tree, parentIndex := NewTree("Test::TestTree")

	infoIndex := tree.AddInfo(parentIndex, "info branch")
	tree.AddMessages(infoIndex, []string{"info metadata 1", "info metadata 2"})

	tree.AddMessage(parentIndex, "regular message")

	debugIndex := tree.AddDebug(parentIndex, "debug branch")
	tree.AddMessages(debugIndex, []string{"debug metadata 1", "debug metadata 2"})

	errorIndex := tree.AddError(parentIndex, "error branch")
	tree.AddMessages(errorIndex, []string{"error metadata 1", "debug metadata 2"})

	tree.AddMessages(debugIndex, []string{"more regular messages", "and some more regular messages"})

	nestedInfoIndex := tree.AddInfo(parentIndex, "info branch 2")
	tree.AddMessage(nestedInfoIndex, "info metadata")

	nestedDebugIndex := tree.AddDebug(nestedInfoIndex, "debug branch within info branch")

	nestedWarnIndex := tree.AddWarn(nestedDebugIndex, "warn branch within debug branch, within info branch")

	tree.AddError(nestedWarnIndex, "error branch within warn branch, within debug branch, within info branch")

	assert.Equal(14, len(tree.Nodes))

	tree.Display()
}

func TestWithLineBreaks(t *testing.T) {
	assert := assert.New(t)

	tree, parentIndex := NewTree("Test::TestTree")

	infoIndex := tree.AddInfo(parentIndex, "info branch")
	tree.AddMessage(infoIndex, "message 1")
	tree.AddBreak(infoIndex)
	tree.AddMessages(infoIndex, []string{"message 2", "message 3"})
	tree.AddBreak(infoIndex)

	tree.AddBreak(infoIndex)
	tree.AddBreak(infoIndex)
	tree.AddBreak(infoIndex)
	tree.AddBreak(infoIndex)
	tree.AddBreak(infoIndex)
	tree.AddDebug(infoIndex, "debug branch")
	tree.AddBreak(infoIndex)
	lastidx := tree.AddMessages(infoIndex, []string{"message 4", "message 5"})

	assert.Equal(14, len(tree.Nodes))

	assert.Equal(2, len(tree.Nodes[lastidx-1].messages))

	tree.Display()
}

func TestExample(t *testing.T) {
	tree, parentIndex := NewTree("Example Tree")

	// Add an info message
	tree.AddInfo(parentIndex, "I am an info branch!")

	// Add some regular messages
	tree.AddMessages(parentIndex, []string{"Or we can use plain messages!", "And more plain messages"})

	// Create a debug branch, get it's index, and create another node for this branch.
	debugIdx := tree.AddDebug(parentIndex, "I am a debug branch!")

	// Add a message to the debug branch
	tree.AddMessage(debugIdx, "plain messages can be added to info/debug/error/warn branches too")

	// Print the tree to standard out
	tree.Display()
}

func TestExample2(t *testing.T) {
	tree, parentIndex := NewTree("Example Tree")

	// Adds nodes directly to the outermost node of the tree (parentIndex)
	tree.AddMessage(parentIndex, "Message to add")
	tree.AddInfo(parentIndex, "Info message to add")
	tree.AddDebug(parentIndex, "Debug message to add")
	tree.AddWarn(parentIndex, "Warn message to add")
	tree.AddError(parentIndex, "Error message to add")
	tree.AddMessages(parentIndex, []string{"First message to add", "Second message to add"})

	// Adds a node to the outermost node of the tree. Assigns a reference to the index of this new node, so that a new branch can be created.
	branchIndex := tree.AddInfo(parentIndex, "This node will have children")
	tree.AddMessages(branchIndex, []string{"Branch child 1", "Branch child 2"})

	// Print the tree to standard out
	tree.Display()
}
