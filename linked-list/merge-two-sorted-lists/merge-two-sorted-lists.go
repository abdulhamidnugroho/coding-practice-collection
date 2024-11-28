package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	// Base cases: If one of the lists is empty, return the other list
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}

	// Compare the values of list1 and list2
	if list1.Val <= list2.Val {
		// Recursively merge the next of list1 with list2
		list1.Next = mergeTwoLists(list1.Next, list2)
		return list1
	} else {
		// Recursively merge list1 with the next of list2
		list2.Next = mergeTwoLists(list1, list2.Next)
		return list2
	}
}

// Helper function to create a linked list from a slice of integers
func createList(nums []int) *ListNode {
	if len(nums) == 0 {
		return nil
	}
	head := &ListNode{Val: nums[0]}
	current := head
	for _, num := range nums[1:] {
		current.Next = &ListNode{Val: num}
		current = current.Next
	}
	return head
}

// Helper function to print a linked list
func printList(node *ListNode) {
	for node != nil {
		fmt.Print(node.Val, " ")
		node = node.Next
	}
	fmt.Println()
}

func main() {
	// Test case
	list1 := createList([]int{1, 2, 4, 5, 5})
	list2 := createList([]int{1, 3, 4})

	mergedList := mergeTwoLists(list1, list2)
	printList(mergedList) // Output: 1 1 2 3 4 4
}
