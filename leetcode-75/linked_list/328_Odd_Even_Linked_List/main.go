package test

//
//import (
//
//	"fmt"
//	"reflect"
//	"testing"
//)
//
//func main() {
//
//	printLinkedList(deleteMiddle(arrayToLinkedList([]int{1, 2, 3, 4, 5})))
//}
//
///**
// * Definition for singly-linked list.
// * type ListNode struct {
// *     Val int
// *     Next *ListNode
// * }
// */
//
//type ListNode struct {
//	Val  int
//	Next *ListNode
//}
//
//func deleteMiddle(head *ListNode) *ListNode {
//	// has only one node => return nil
//	if head != nil && head.Next == nil {
//		return nil
//	}
//
//	dummy := &ListNode{Next: head}
//	slow, fast := dummy, head
//	for fast != nil && fast.Next != nil {
//		slow = slow.Next
//		fast = fast.Next.Next
//	}
//
//	slow.Next = slow.Next.Next
//	return dummy.Next
//}
//
//func arrayToLinkedList(nums []int) *ListNode {
//	if len(nums) == 0 {
//		return nil
//	}
//
//	head := &ListNode{Val: nums[0]}
//	curr := head
//
//	for _, val := range nums[1:] {
//		curr.Next = &ListNode{Val: val}
//		curr = curr.Next
//	}
//
//	return head
//}
//
//func linkedListToArray(head *ListNode) []int {
//	var result []int
//	for head != nil {
//		result = append(result, head.Val)
//		head = head.Next
//	}
//	return result
//}
//
//
//func printLinkedList(head *ListNode) {
//	for head != nil {
//		fmt.Print(head.Val)
//		if head.Next != nil {
//			fmt.Print(" → ")
//		}
//		head = head.Next
//	}
//	fmt.Println()
//}
//
//func TestDeleteMiddle(t *testing.T) {
//	tests := []struct {
//		input    []int
//		expected []int
//	}{
//		{[]int{1, 3, 4, 7, 1, 2, 6}, []int{1, 3, 4, 1, 2, 6}},
//		{[]int{1, 2, 3, 4}, []int{1, 2, 4}},
//		{[]int{2, 1}, []int{2}},
//		{[]int{1}, nil},
//		{[]int{}, nil},
//	}
//
//	for _, tt := range tests {
//		head := arrayToLinkedList(tt.input)
//		newHead := deleteMiddle(head)
//		got := linkedListToArray(newHead)
//
//		if !reflect.DeepEqual(got, tt.expected) {
//			t.Errorf("deleteMiddle(%v) = %v; want %v", tt.input, got, tt.expected)
//		}
//	}
//}
