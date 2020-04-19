package Week_01

/*
双端队列：允许两端都可以入队出队的队列
数组实现
TODO GetRear
*/
type MyCircularDeque struct {
	buf   []int
	head  int
	tail  int
	count int
}

/** Initialize your data structure here. Set the size of the deque to be k. */
func Constructor(k int) MyCircularDeque {
	return MyCircularDeque{
		buf:   make([]int, k),
		head:  0,
		tail:  0,
		count: 0,
	}
}

/** Adds an item at the front of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) InsertFront(value int) bool {
	if this.IsFull() {
		return false
	}
	this.head = this.prev(this.head)
	this.buf[this.head] = value
	this.count++
	return true
}

/** Adds an item at the rear of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) InsertLast(value int) bool {
	if this.IsFull() {
		return false
	}
	this.tail = this.next(this.tail)
	this.buf[this.tail] = value
	this.count++
	return true
}

/** Deletes an item from the front of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) DeleteFront() bool {
	if this.IsEmpty() {
		return false
	}
	this.buf[this.head] = 0
	this.head = this.next(this.head)
	this.count--
	return true
}

/** Deletes an item from the rear of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) DeleteLast() bool {
	if this.IsEmpty() {
		return false
	}
	this.tail = this.prev(this.tail)
	this.buf[this.tail] = 0
	this.count--
	return true
}

/** Get the front item from the deque. */
func (this *MyCircularDeque) GetFront() int {
	if this.IsEmpty() {
		return -1
	}
	return this.buf[this.head]
}

/** Get the last item from the deque. */
func (this *MyCircularDeque) GetRear() int {
	if this.IsEmpty() {
		return -1
	}
	return this.buf[this.prev(this.tail)]
}

/** Checks whether the circular deque is empty or not. */
func (this *MyCircularDeque) IsEmpty() bool {
	return this.count <= 0
}

/** Checks whether the circular deque is full or not. */
func (this *MyCircularDeque) IsFull() bool {
	return this.count == len(this.buf)
}

func (this *MyCircularDeque) prev(i int) int {
	return (i - 1) & (len(this.buf) - 1)
}
func (this *MyCircularDeque) next(i int) int {
	return (i + 1) & (len(this.buf) - 1)
}

/**
 * Your MyCircularDeque object will be instantiated and called as such:
 * obj := Constructor(k);
 * param_1 := obj.InsertFront(value);
 * param_2 := obj.InsertLast(value);
 * param_3 := obj.DeleteFront();
 * param_4 := obj.DeleteLast();
 * param_5 := obj.GetFront();
 * param_6 := obj.GetRear();
 * param_7 := obj.IsEmpty();
 * param_8 := obj.IsFull();
 */
