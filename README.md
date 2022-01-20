1. [](#)

-----

- [X] 206 反转链表

```go
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func reverseList(head *ListNode) *ListNode {
    var pre *ListNode
    cur := head
    for cur != nil {
        next := cur.Next
        cur.Next = pre
        pre = cur
        cur = next
    }
    return pre
}
```
复杂度分析

- 时间复杂度：O(n)，其中 n 是链表的长度。需要遍历链表一次。

- 空间复杂度：O(1)。


![](img/206.jpg)

- [X] 24 两两交换链表中的节点

```go
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func swapPairs(head *ListNode) *ListNode {
    
    if head == nil || head.Next == nil {
        return head
    }

    dummy := head.Next
    head.Next = swapPairs(dummy.Next)
    dummy.Next = head

    return dummy
}
```

复杂度分析

- 时间复杂度：O(n)，其中 n 是链表的长度。需要遍历链表一次。

- 空间复杂度：O(1)。

![](img/24.jpg)


- [X] 3. 无重复字符的最长子串

```go
func lengthOfLongestSubstring(s string) int {
    start , res := 0,0
    m := map[byte]int{}

    for i := 0; i <len(s); i++ {
        if _, exists := m[s[i]] ; exists {
            start = max(start, m[s[i]]+1)
        }
        m[s[i]] = i
        res = max(res, i-start + 1)
    }

    return res 
}

func max(a,b int) int {
    if a > b {
        return a
    } else {
        return b
    }
}

```
-----

[2022-01-20](#2022-01-20)
1. ✅[232. Implement Queue using Stacks](#232-implement-queue-using-stacks)
2. ✅[94. Binary Tree Inorder Traversal](#94-binary-tree-inorder-traversal)
3. ✅[199. Binary Tree Right Side View](#199-binary-tree-right-side-view)
4. ✅[143. Reorder List](#143-reorder-list)
5. ✅[70. Climbing Stairs](#70-climbing-stairs)

## [232. Implement Queue using Stacks](https://leetcode-cn.com/problems/implement-queue-using-stacks/submissions/)

```go
type MyQueue struct {
	input  []int
	output []int
}

func Constructor() MyQueue {
	return MyQueue{}
}

func (this *MyQueue) Push(x int) {
	this.input = append(this.input, x)
}

func (this *MyQueue) Pop() int {

	temp := this.Peek()
	this.output = this.output[:len(this.output)-1]
	return temp
}

func (this *MyQueue) Peek() int {
	if len(this.output) == 0 {
		for len(this.input) > 0 {
			this.output = append(this.output, this.input[len(this.input)-1])
			this.input = this.input[:len(this.input)-1]
		}
	}
	return this.output[len(this.output)-1]
}

func (this *MyQueue) Empty() bool {
	return len(this.input) == 0 && len(this.output) == 0
}

```

Runtime: 0 ms, faster than 100.00% of Go online submissions for Implement Queue using Stacks.
Memory Usage: 1.9 MB, less than 100.00% of Go online submissions for Implement Queue using Stacks.

## [94. Binary Tree Inorder Traversal](https://leetcode-cn.com/problems/binary-tree-inorder-traversal/)

```go
func inorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	left := inorderTraversal(root.Left)
	right := inorderTraversal(root.Right)
	left = append(left, root.Val)
	left = append(left, right...)

	return left
}
```

Accepted
- 70/70 cases passed (0 ms)
- Your runtime beats 100 % of golang submissions
- Your memory usage beats 16.83 % of golang submissions (2 MB)

## [199. Binary Tree Right Side View](https://leetcode-cn.com/problems/binary-tree-right-side-view/)

```go
func rightSideView(root *TreeNode) []int {
	res := []int{}
	if root == nil {
		return res
	}

	queue := []*TreeNode{root}
	for len(queue) != 0 {
		levelSize := len(queue)
		res = append(res, queue[levelSize-1].Val)

		for i := 0; i < levelSize; i++ {
			node := queue[0]
			queue = queue[1:]

			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}

	}

	return res
}
```

Accepted
- 215/215 cases passed (0 ms)
- Your runtime beats 100 % of golang submissions
- Your memory usage beats 76.91 % of golang submissions (2.3 MB)


## [143. Reorder List](https://leetcode-cn.com/problems/reorder-list/)
```go
func reorderList(head *ListNode) {
	if head == nil {
		return
	}

	slow, fast := head, head
	var slc []*ListNode

	for slow != nil {
		if fast == nil {
			slc = append(slc, slow)
		} else {
			if (*fast).Next == nil || (*fast).Next.Next == nil {
				fast = nil
			} else {
				fast = (*fast).Next.Next
			}
		}
		slow = (*slow).Next
	}

	cur := head
	for i := len(slc) - 1; i >= 0; i-- {
		temp := (*cur).Next
		(*cur).Next = slc[i]
		(*cur).Next.Next = temp
		cur = temp
	}

	(*cur).Next = nil
}
```

Accepted
- 12/12 cases passed (8 ms)
- Your runtime beats 87.49 % of golang submissions
- Your memory usage beats 6.71 % of golang submissions (6.4 MB)


## [70. Climbing Stairs](https://leetcode-cn.com/problems/climbing-stairs/)

```go
func climbStairs(n int) int {
	if n <= 2 {
		return n
	}

	n1, n2 := 1, 2
	for i := 3; i <= n; i++ {
		temp := n2
		n2 = n2 + n1
		n1 = temp
	}

	return n2
}
```
Accepted
- 45/45 cases passed (0 ms)
- Your runtime beats 100 % of golang submissions
- Your memory usage beats 76.72 % of golang submissions (1.9 MB)
