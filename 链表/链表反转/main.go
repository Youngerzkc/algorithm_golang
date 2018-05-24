// package main

// import (
// 	"fmt"
// )

// type ListNode struct {
// 	Data int
// 	Next *ListNode
// }

// func InsertListNode(ptr *ListNode, newNode *ListNode, data int) {

// 	if ptr.Next == nil {
// 		ptr.Next = newNode
// 		return
// 	}
// 	for ptr.Next != nil {
// 		ptr = ptr.Next
// 		if ptr.Next == nil {
// 			ptr.Data = newNode.Data
// 			newNode.Next = nil
// 			return
// 		}
// 	}

// }
// func DisPlayListNode(ptr *ListNode) {
// 	for ptr.Next != nil {
// 		fmt.Println("data is ", ptr.Data)
// 		ptr = ptr.Next
// 	}
// }
// func main() {
// 	var head ListNode

// 	for i := 0; i < 10; i++ {
// 		var newNode ListNode
// 		newNode.Data = i
// 		InsertListNode(&head, &newNode, 0)
// 	}
// 	DisPlayListNode(&head)
// }
