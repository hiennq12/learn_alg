package test

import (
	"fmt"
	"reflect"
	"testing"
)

type ListNode struct {
	Val  int
	Next *ListNode
}


func TestDeleteMiddle(t *testing.T) {
	tests := []struct {
		input    []int
		expected []int
	}{
		{[]int{1, 3, 4, 7, 1, 2, 6}, []int{1, 4, 1, 6, 3, 7, 2}},
		{[]int{1, 2, 3, 4}, []int{1, 3, 2, 4}},
		{[]int{2, 1}, []int{2, 1}},
		{[]int{1}, []int{1}},
		{[]int{}, nil},
	}

	for _, tt := range tests {
		head := arrayToLinkedList(tt.input)
		newHead := oddEvenList(head)
		got := linkedListToArray(newHead)

		if !reflect.DeepEqual(got, tt.expected) {
			t.Errorf("deleteMiddle(%v) = %v; want %v", tt.input, got, tt.expected)
		}
	}
}

func oddEvenList(head *ListNode) *ListNode {
	odd := head
	even := head.Next
	evenHead := even
	for even != nil && even.Next != nil {
		odd.Next = even.Next
		odd = odd.Next
		even = odd.Next
		even = even.Next
	}

	odd.Next = evenHead

	return head
}

func arrayToLinkedList(nums []int) *ListNode {
	if len(nums) == 0 {
		return nil
	}

	head := &ListNode{Val: nums[0]}
	curr := head

	for _, val := range nums[1:] {
		curr.Next = &ListNode{Val: val}
		curr = curr.Next
	}

	return head
}

func linkedListToArray(head *ListNode) []int {
	var result []int
	for head != nil {
		result = append(result, head.Val)
		head = head.Next
	}
	return result
}

func printLinkedList(head *ListNode) {
	for head != nil {
		fmt.Print(head.Val)
		if head.Next != nil {
			fmt.Print(" → ")
		}
		head = head.Next
	}
	fmt.Println()
}

