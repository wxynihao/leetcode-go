package main

import "fmt"

func main641() {
	circularDeque := Constructor(3)
	fmt.Println(circularDeque.InsertLast(1))  // 返回 true
	fmt.Println(circularDeque.InsertLast(2))  // 返回 true
	fmt.Println(circularDeque.InsertFront(3)) // 返回 true
	fmt.Println(circularDeque.InsertFront(4)) // 已经满了，返回 false
	fmt.Println(circularDeque.GetRear())      // 返回 2
	fmt.Println(circularDeque.IsFull())       // 返回 true
	fmt.Println(circularDeque.DeleteLast())   // 返回 true
	fmt.Println(circularDeque.InsertFront(4)) // 返回 true
	fmt.Println(circularDeque.GetFront())     // 返回 4
}

//leetcode submit region begin(Prohibit modification and deletion)
type MyCircularDeque struct {
	size  int
	front int
	last  int
	data  []int
}

/** Initialize your data structure here. Set the size of the deque to be k. */
func Constructor(k int) MyCircularDeque {
	return MyCircularDeque{
		size:  k,
		front: -1,
		last:  -1,
		data:  make([]int, k),
	}
}

/** Adds an item at the front of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) InsertFront(value int) bool {
	if this.IsFull() {
		return false
	}
	if this.last == -1 {
		this.front = 0
		this.last = 0
	} else {
		this.front = this.front - 1
		if this.front < 0 {
			this.front += this.size
		}
	}
	this.data[this.front] = value
	return true
}

/** Adds an item at the rear of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) InsertLast(value int) bool {
	if this.IsFull() {
		return false
	}
	if this.front == -1 {
		this.front = 0
		this.last = 0
	} else {
		this.last = (this.last + 1) % this.size
	}
	this.data[this.last] = value

	return true
}

/** Deletes an item from the front of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) DeleteFront() bool {
	if this.IsEmpty() {
		return false
	}
	if this.front == this.last {
		this.front = -1
		this.last = -1
	} else {
		this.front = (this.front + 1) % this.size
	}
	return true
}

/** Deletes an item from the rear of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) DeleteLast() bool {
	if this.IsEmpty() {
		return false
	}
	if this.front == this.last {
		this.front = -1
		this.last = -1
	} else {
		this.last = this.last - 1
		if this.last < 0 {
			this.last += this.size
		}
	}
	return true
}

/** Get the front item from the deque. */
func (this *MyCircularDeque) GetFront() int {
	if this.IsEmpty() {
		return -1
	}
	return this.data[this.front]
}

/** Get the last item from the deque. */
func (this *MyCircularDeque) GetRear() int {
	if this.IsEmpty() {
		return -1
	}
	return this.data[this.last]
}

/** Checks whether the circular deque is empty or not. */
func (this *MyCircularDeque) IsEmpty() bool {
	if this.front == -1 && this.last == -1 {
		return true
	}
	return false
}

/** Checks whether the circular deque is full or not. */
func (this *MyCircularDeque) IsFull() bool {
	if (this.last+1)%this.size == this.front {
		return true
	}
	return false
}
