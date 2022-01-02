package main

/******************************* 加一 *******************************/
func plusOne(digits []int) []int {
	num := len(digits)
	rst := make([]int, num, num+1)

	plusVal := 1
	for i := num - 1; i >= 0; i-- {
		tmp := digits[i] + plusVal
		if tmp > 9 {
			rst[i] = tmp % 10
		} else {
			rst[i] = tmp
			plusVal = 0
		}
	}

	if plusVal == 1 {
		rst = append([]int{1}, rst...)
	}

	return rst
}

/**************************** 合并两个有序链表 ****************************/
type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	protectedNode := &ListNode{} //保护节点
	lastNode := protectedNode    //链表合并处的指针

	for list1 != nil || list2 != nil {
		if list1 != nil && list2 != nil { //两个链表均未结束
			if list1.Val <= list2.Val {
				lastNode.Next = list1
				list1 = list1.Next
			} else {
				lastNode.Next = list2
				list2 = list2.Next
			}
		} else {
			if list1 == nil { //list1提前结束
				lastNode.Next = list2
			} else { //list2提前结束
				lastNode.Next = list1
			}
			break
		}

		lastNode = lastNode.Next
	}

	return protectedNode.Next
}

/**************************** 设计循环双端队列 ****************************/

//双向链表节点定义
type DequeNode struct {
	val  int        //节点值
	prev *DequeNode //前序节点
	next *DequeNode //后序节点
}

//在双向链表的某节点后插入新节点
func (node *DequeNode) InsertAfter(val int) {
	newNode := &DequeNode{
		val:  val,
		prev: node,      //本节点为新节点的前序
		next: node.next, //本节点的后序节点为新节点的后序节点
	}

	//更新原来后序节点的指针
	node.next.prev = newNode
	node.next = newNode
}

//删除双向链表的某节点
func (node *DequeNode) Remove() {
	node.prev.next = node.next
	node.next.prev = node.prev
}

type MyCircularDeque struct {
	length   int        //当前队列的长度
	capacity int        //队列最大容量
	head     *DequeNode //头保护节点
	tail     *DequeNode //尾保护节点
}

func Constructor(k int) MyCircularDeque {
	headNode := &DequeNode{}
	tailNode := &DequeNode{}
	headNode.next = tailNode
	tailNode.prev = headNode

	return MyCircularDeque{
		length:   0,
		capacity: k,
		head:     headNode,
		tail:     tailNode,
	}
}

func (this *MyCircularDeque) InsertFront(value int) bool {
	if this.IsFull() {
		return false
	}

	this.head.InsertAfter(value)
	this.length++
	return true
}

func (this *MyCircularDeque) InsertLast(value int) bool {
	if this.IsFull() {
		return false
	}

	this.tail.prev.InsertAfter(value)
	this.length++
	return true
}

func (this *MyCircularDeque) DeleteFront() bool {
	if this.IsEmpty() {
		return false
	}

	this.head.next.Remove()
	this.length--
	return true
}

func (this *MyCircularDeque) DeleteLast() bool {
	if this.IsEmpty() {
		return false
	}

	this.tail.prev.Remove()
	this.length--
	return true
}

func (this *MyCircularDeque) GetFront() int {
	if this.IsEmpty() {
		return -1
	}

	return this.head.next.val
}

func (this *MyCircularDeque) GetRear() int {
	if this.IsEmpty() {
		return -1
	}

	return this.tail.prev.val
}

func (this *MyCircularDeque) IsEmpty() bool {
	return this.length == 0
}

func (this *MyCircularDeque) IsFull() bool {
	return this.length == this.capacity
}
