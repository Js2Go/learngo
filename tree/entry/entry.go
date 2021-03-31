package main

import (
	"fmt"
	"learngo/tree"
)

type myTreeNode struct {
	node *tree.Node
}

func (myNode *myTreeNode) posOrder() {
	if myNode == nil || myNode.node == nil {
		return
	}
	left := myTreeNode{myNode.node.Left}
	left.posOrder()
	right := myTreeNode{myNode.node.Right}
	right.posOrder()
	myNode.node.Print()
}

func main() {
	var root tree.Node

	root = tree.Node{Value: 3}
	root.Left = &tree.Node{}
	root.Right = &tree.Node{5, nil, nil}
	root.Right.Left = new(tree.Node)
	root.Left.Right = tree.CreateNode(2)
	root.Right.Left.SetValue(4)

	root.Travers()

	//nodeCount := 0
	//root.TraverseFunc(func(*tree.Node) {
	//	nodeCount++
	//})
	//fmt.Println("Node count: ", nodeCount)

	//fmt.Println()
	//myRoot := myTreeNode{&root}
	//myRoot.posOrder()
	//fmt.Println()

	//root.setValue(4)
	//root.print()

	//pRoot := &root
	//var pRoot *treeNode
	//pRoot.print()
	//pRoot.setValue(100)
	//pRoot = &root
	//pRoot.setValue(1200)
	//pRoot.print()

	//nodes := []treeNode{
	//	{value: 3},
	//	{},
	//	{6, nil, &root},
	//}
	//fmt.Println(nodes)

	c := root.TraverseWithChannel()

	maxNode := 0
	for node := range c {
		if node.Value > maxNode {
			maxNode = node.Value
		}
	}
	fmt.Println("Max node value: ", maxNode)
}
