package main

import "fmt"

type Node struct {
	data string
	next *Node
}

var head *Node = new(Node)

func initial() {
	head.data = "Daqing"
	head.next = nil

	var nodeHarbin *Node = &Node{
		data: "Harbin",
		next: nil,
	}
	head.next = nodeHarbin

	var nodeJilin *Node = &Node{
		data: "Jilin",
		next: nil,
	}
	nodeHarbin.next = nodeJilin

	var tail *Node = &Node{
		data: "Beijing",
		next: nil,
	}
	nodeJilin.next = tail
}

func output(node *Node) {
	p := node
	for {
		if p == nil {
			break
		}

		fmt.Printf("%s--> ", p.data)
		p = p.next
	}

	fmt.Printf("End\n\n")
}

func main() {
	initial()
	output(head)

}
