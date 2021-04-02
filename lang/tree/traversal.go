package tree

import "fmt"

func (node *Node) Travers() {
	node.TraverseFunc(func(node *Node) {
		node.Print()
	})
	fmt.Println()
}

func (node *Node) TraverseFunc(f func(*Node)) {
	if node == nil {
		return
	}
	node.Left.TraverseFunc(f)
	f(node)
	node.Right.TraverseFunc(f)
}

// TODO:需要深入理解
func (node *Node) TraverseWithChannel() chan *Node {
	out := make(chan *Node)

	go func() {
		node.TraverseFunc(func(node *Node) {
			out <- node
		})
		close(out)
	}()

	return out
}
