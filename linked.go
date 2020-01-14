package gotype

import (
	"fmt"
)

type Element interface {
}

//链表定义
type LinkNode struct {
	Data Element
	Next *LinkNode
}

func CreateLinked() *LinkNode {
	return new(LinkNode)
}

//IsEmpty 是否为空链表
func (linked *LinkNode) IsEmpty() bool {
	if linked.Next == nil {
		return true
	}
	return false
}

//创建并插入元素
func (linked *LinkNode) Create(e ...Element) {
	cur := linked
	for _, item := range e {
		cur.Next = &LinkNode{}
		cur.Next.Data = item
		cur = cur.Next
	}
}

//Insert 在链表头部插入一个元素
func (linked *LinkNode) Insert(item Element) {
	per := linked.Next
	node := &LinkNode{Data: item}
	node.Next = per
	linked.Next = node
}

//链表尾插入一个元素
func (linked *LinkNode) Append(item Element) {
	cur := linked
	for cur.Next != nil {
		cur = cur.Next
	}
	node := &LinkNode{Data: item}
	cur.Next = node
}

//在链表中移除指定的item元素，多个相同的会全部被移除
func (linked *LinkNode) Remove(item Element) {
	cur := linked
	for cur.Next != nil {
		if cur.Next.Data == item {
			cur.Next = cur.Next.Next
		} else {
			cur = cur.Next
		}
	}
}

//链表长度
func (linked *LinkNode) Length() int {
	i := 1
	cur := linked.Next
	for cur.Next != nil {
		cur = cur.Next
		i++
	}
	return i
}

//Empty 清空链表
func (linked *LinkNode) Empty() {
	linked.Data = 0
	linked.Next = nil
}

func (linked *LinkNode) String() string {

	if linked.IsEmpty() {
		return "空链表"
	}
	var str string
	cur := linked.Next

	for cur.Next != nil {
		str += fmt.Sprintf(" %v ", cur.Data)
		cur = cur.Next
	}
	str += fmt.Sprintf(" %v ", cur.Data)
	return str
}
