package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type ListNode struct {
	Val  int
	Next *ListNode
}

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func main() {

	fmt.Println(binarySearch(3, []int{1, 2, 3, 4, 5, 6, 12, 30, 40, 55, 60, 88}))
}

// need to have a sorted array of nums
func binarySearch(needle int, nums []int) (int, bool) {
	min := 0
	max := len(nums) - 1

	for min <= max {
		mid := (min + max) / 2
		if nums[mid] < needle {
			min = mid + 1
		} else if nums[mid] > needle {
			max = mid - 1
		} else {
			return mid, true
		}
	}

	return 0, false
}
