package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	current := head

	for current != nil && current.Next != nil {
		if current.Val == current.Next.Val {
			current.Next = current.Next.Next
		} else {
			current = current.Next
		}
	}

	return head
}

func createLinkedList(values []int) *ListNode {
	if len(values) == 0 {
		return nil
	}
	head := &ListNode{Val: values[0]}
	current := head
	for _, val := range values[1:] {
		current.Next = &ListNode{Val: val}
		current = current.Next
	}
	return head
}

func printLinkedList(head *ListNode) {
	current := head
	for current != nil {
		fmt.Printf("%d ", current.Val)
		current = current.Next
	}
	fmt.Println()
}

func main() {
	input1 := []int{1, 1, 2}
	head1 := createLinkedList(input1)
	fmt.Println("Original list:")
	printLinkedList(head1)

	result1 := deleteDuplicates(head1)
	fmt.Println("After removing duplicates:")
	printLinkedList(result1)

	input2 := []int{1, 1, 2, 3, 3}
	head2 := createLinkedList(input2)
	fmt.Println("\nOriginal list:")
	printLinkedList(head2)

	result2 := deleteDuplicates(head2)
	fmt.Println("After removing duplicates:")
	printLinkedList(result2)
}
