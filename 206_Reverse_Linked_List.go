package main

import "fmt"

func main206() {

	five := &ListNode{
		Val:  5,
		Next: nil,
	}
	four := &ListNode{
		Val:  4,
		Next: five,
	}
	three := &ListNode{
		Val:  3,
		Next: four,
	}
	two := &ListNode{
		Val:  2,
		Next: three,
	}
	one := &ListNode{
		Val:  1,
		Next: two,
	}

	f := func(head *ListNode) {
		for head != nil {
			fmt.Printf("%d->", head.Val)
			head = head.Next
		}
		fmt.Println()
	}
	f(one)
	n := reverseList2(one)
	f(n)
}

type ListNode struct {
	Val  int
	Next *ListNode
}

// 遍历并新建一个list
func reverseList1(head *ListNode) *ListNode {
	var newHead *ListNode
	for head != nil {
		node := &ListNode{
			Val:  head.Val,
			Next: newHead,
		}
		newHead = node
		head = head.Next
	}
	return newHead
}

// 直接修改原有链表的next指向
func reverseList2(head *ListNode) *ListNode {
	var tmp *ListNode
	for head != nil {
		next := head.Next
		head.Next = tmp
		tmp = head
		head = next
	}
	return tmp
}
