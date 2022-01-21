package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func preOrderWalk(root *TreeNode, ch chan int) {
	if root == nil {
		return
	}

	ch <- root.Val
	preOrderWalk(root.Left, ch)
	preOrderWalk(root.Right, ch)
}

func preOrder(root *TreeNode) <-chan int {
	ch := make(chan int)
	go func() {
		preOrderWalk(root, ch)
		close(ch)
	}()

	return ch
}

func main() {

}
