[2022-01-20](#2022-01-20)
1. ✅[232. Implement Queue using Stacks](#232-implement-queue-using-stacks)
2. ✅[94. Binary Tree Inorder Traversal](#94-binary-tree-inorder-traversal)
3. ✅[199. Binary Tree Right Side View](#199-binary-tree-right-side-view)
4. ✅[143. Reorder List](#143-reorder-list)
5. ✅[70. Climbing Stairs](#70-climbing-stairs)
6. ✅[剑指 Offer 22. 链表中倒数第k个节点](#剑指-offer-22-链表中倒数第k个节点)
7. ✅[124. Binary Tree Maximum Path Sum](#124-binary-tree-maximum-path-sum)
8. ✅[69. Sqrt(x)](#69-sqrtx)
9. ✅[56. Merge Intervals](#56-merge-intervals)
10. ✅[82. Remove Duplicates from Sorted List II](#82-remove-duplicates-from-sorted-list-ii)

-----

[2022-01-21](#2022-01-21)
1. ✅[19. Remove Nth Node From End of List](#19-remove-nth-node-from-end-of-list)
2. ✅[8. String to Integer (atoi)](#8-string-to-integer-atoi)
3. ✅[2. Add Two Numbers](#2-add-two-numbers)
4. ✅[144. Binary Tree Preorder Traversal](#144-binary-tree-preorder-traversal)
5. ✅[148. Sort List](#148-sort-list)
6. ✅[105. Construct Binary Tree from Preorder and Inorder Traversal](#105-construct-binary-tree-from-preorder-and-inorder-traversal)
7. ✅[151. Reverse Words in a String](#151-reverse-words-in-a-string)
8. ✅[31. Next Permutation](#31-next-permutation)

-----

[2022-01-21](#2022-01-22)
1. ✅[1143. Longest Common Subsequence](#1143-longest-common-subsequence)
2. ✅[4. Median of Two Sorted Arrays](#4-median-of-two-sorted-arrays)
3. ✅[104. Maximum Depth of Binary Tree](#104-maximum-depth-of-binary-tree)
4. ✅[129. Sum Root to Leaf Numbers](#129-sum-root-to-leaf-numbers)
5. ✅[110. Balanced Binary Tree](#110-balanced-binary-tree)
6. ✅[93. Restore IP Addresses](#93-restore-ip-addresses)
7. ✅[113. Path Sum II](#113-path-sum-ii)
8. ✅[155. Min Stack](#155-min-stack)

-----

[2022-01-21](#2022-01-23)
1. ✅[22. Generate Parentheses](#22-generate-parentheses)
2. ✅[98. Validate Binary Search Tree](#98-validate-binary-search-tree)
3. ✅[543. Diameter of Binary Tree](#543-diameter-of-binary-tree)
4. ✅[470. Implement Rand10() Using Rand7()](#470-implement-rand10-using-rand7)
5. ✅[64. Minimum Path Sum](#64-minimum-path-sum)
6. ✅[112. Path Sum](#112-path-sum)
7. ✅[234. Palindrome Linked List](#234-palindrome-linked-list)
8. ✅[48. Rotate Image](#48-rotate-image)
9. ✅[718. Maximum Length of Repeated Subarray](#718-maximum-length-of-repeated-subarray)
10. ✅[78. Subsets](#78-subsets)


## 2022-03-18

## 215. 数组中的第K个最大元素
给定整数数组 nums 和整数 k，请返回数组中第 k 个最大的元素。
请注意，你需要找的是数组排序后的第 k 个最大的元素，而不是第 k 个不同的元素。
```go
func findKthLargest(nums []int, k int) int {
    if len(nums) == 0 {
        return 0
    }

    return selection(nums, 0, len(nums)-1, len(nums)-k)
}

func selection(arr []int, left, right, k int) int {
    if left == right {
        return arr[left]
    }

    pivot := partition(arr, left, right)

    if k == pivot {
        return arr[pivot]
    } else if k < pivot {
        return selection(arr, left, pivot - 1, k)
    } else  {
        return selection(arr, pivot+1, right, k)
    }
}

func partition(arr []int, left, right int) int {
    pivot := arr[right]
    i := left -1 

    for j := left; j < right; j++ {
        if arr[j] < pivot {
            i++
            arr[j], arr[i] = arr[i], arr[j]
        }
    }

    arr[i+1], arr[right] = arr[right], arr[i+1]

    return i + 1
}
```

## 2022-03-22

##704. 二分查找  ⭐⭐⭐⭐
给定一个n个元素有序的（升序）整型数组nums 和一个目标值target ，写一个函数搜索nums中的 target，如果目标值存在返回下标，否则返回 -1

```go
func search(nums []int, target int) int {

    left,right := 0, len(nums)-1

   for left <= right {
       mid := left + (right-left)>>1
       if nums[mid] < target {
           left = mid + 1
       } else if nums[mid] > target {
           right = mid - 1 
       } else if nums[mid] == target {
           return mid
       }
   }

    return -1
}
```

##left_bound
```go
func left_bound(nums []int, target int) int {

    left,right := 0, len(nums)-1

   for left <= right {
       mid := left + (right-left)>>1
       if nums[mid] < target {
           left = mid + 1
       } else if nums[mid] > target {
           right = mid - 1 
       } else if nums[mid] == target {
           right = mid -1 
       }
   }
   
   if left >= len(nums) || nums[left] != target {
   	        return -1 
       }

    return left
}
```


##right_bound

```go
func left_bound(nums []int, target int) int {

    left,right := 0, len(nums)-1

   for left <= right {
       mid := left + (right-left)>>1
       if nums[mid] < target {
           left = mid + 1
       } else if nums[mid] > target {
           right = mid - 1 
       } else if nums[mid] == target {
           left = mid + 1
       }
   }
   
   if right < len(nums) || nums[right] != target {
   	        return -1 
       }

    return right
}
```

## 2022-03-23

##26. 删除有序数组中的重复项  ⭐⭐

```go
func removeDuplicates(nums []int) int {
    if len(nums) == 0 {
        return 0
    }

    slow, fast := 0,0
    for fast < len(nums) {
        if nums[fast] != nums[slow] {
            slow++
            nums[slow] = nums[fast]
        
        }
        fast++
    }

    return slow + 1
}
```

##83. 删除排序链表中的重复元素  ⭐⭐⭐

```go
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteDuplicates(head *ListNode) *ListNode {
    if head == nil {
        return nil
    }

    cur:= head
    for cur != nil && cur.Next != nil {
        if cur.Next.Val == cur.Val {
            cur.Next = cur.Next.Next
        } else {
            cur = cur.Next
        }
    }
    return head
}
```

##2022-03-24

##27. 移除元素 ⭐⭐

```go
func removeElement(nums []int, val int) int {
    slow, fast := 0,0

    for fast < len(nums) {
        if nums[fast] != val {
            //give fast value to slow
            nums[slow] = nums[fast]
            //then slow move
            slow++
        }

        //first fast move
        fast++
    }

    return slow
}
```

##283. 移动零

```go
func moveZeroes(nums []int)  {
    p := removeElt(nums, 0)
    for ; p < len(nums); p++ {
        nums[p] = 0
    }

}

func removeElt(nums []int, val int) int {
    slow, fast := 0, 0
    
    for fast < len(nums) {
        if nums[fast] != val {
            nums[slow] = nums[fast]
            slow++
        }
        fast++
    }

    return slow
}
```