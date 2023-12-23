# Go-tree

[![Go Reference](https://pkg.go.dev/badge/github.com/zilayo/go-tree.svg)](https://pkg.go.dev/github.com/zilayo/go-tree)
[![Go Report Card](https://goreportcard.com/badge/github.com/zilayo/go-tree)](https://goreportcard.com/report/github.com/zilayo/go-tree)
![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/zilayo/go-tree)
[![License](https://img.shields.io/badge/License-MIT-blue.svg)](https://github.com/zilayo/gotree/blob/main/LICENSE)



![preview](example_tree.png?raw=true)



## Overview

Go-tree is a Go package that provides an intuitive and visually structured logging solution for Go applications. It allows developers to create hierarchical log messages that are displayed in an organized, tree-like format. This structure is ideal for debugging the execution flow of complex applications, and where understanding the relationship between various log messages is crucial.


## Features
- Execution Flow Mapping: Designed for tracing the intricate execution paths in modern Go applications.
- Enhanced Readability: Say goodbye to cluttered logs. Go-tree brings clarity to your log data by displaying log messages in a nested tree structure.
- Multiple Log Levels: Supports adding various log levels to the tree such as info, debug, error, and warn.
- Easy to Use: Designed for simplicity, allowing you to integrate structured logging with minimal setup.
- Flexible: Add plain or formatted messages at any level of the tree.
- Console Friendly: Designed for console output, making it perfect for development and debugging.


## Installation
Install go-tree in your package by running:
```
go get github.com/zilayo/go-tree/
```

## Usage

Firstly create a new tree: ``tree, parentIndex := gotree.NewTree("Example Tree")``

The ``NewTree``, ``AddMessage``, ``AddInfo``, ``AddDebug``, ``AddWarn``, ``AddError``, and ``AddMessages`` functions all return an ``int`` which represents the index of the newly created node within the tree. 

The index returned by ``NewTree`` represents the outermost node of the tree.

To add a branch to a specific node, you must the node's index within the Add functions.

All branches are ended with ``└─ ←``, making it easy to understand where a branch ends.

```go
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
```

Output:
```
  Example Tree
    ├─ Message to add
    ├─ info: Info message to add
    ├─ debug: Debug message to add
    ├─ warn: Warn message to add
    ├─ error: Error message to add
    ├─ First message to add
    ├─ Second message to add
    ├─ info: This node will have children
    │   ├─ Branch child 1
    │   ├─ Branch child 2
    │   └─ ←
    └─ ←
```

Here is a basic example of how Go-tree can be used in an application. More examples can be found in [gotree_test.go](gotree_test.go)

```go
package main

import "github.com/zilayo/go-tree"

func main(){
  // Name your tree. This will be displayed at the top of the tree when printed.
  tree, parentIndex := gotree.NewTree("Example Tree")

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
```

Output:
```
  Example Tree
    ├─ info: I am an info branch!
    ├─ Or we can use plain messages!
    ├─ And more plain messages
    ├─ debug: I am a debug branch!
    │   ├─ plain messages can be added to info/debug/error/warn branches too
    │   └─ ←
    └─ ←
```

## Contributing
Contributions are welcome! Feel free to submit issues or pull requests on the GitHub repository.

## Acknowledgements
Go-tree is HEAVILY inspired by [Jonathan Becker's](https://github.com/Jon-Becker) Logging module used in [heimdall-rs](https://github.com/Jon-Becker/heimdall-rs). Go-tree is essentially a minimized Go fork of ``trace`` & ``trace_factory`` in ``heimdall_common::utils::io::logging``.