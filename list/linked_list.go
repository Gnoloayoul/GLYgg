package list

import "github.com/Gonloayoul/internal/errs"

var (
	_ List[any] = &LinkedList[any]{}
)

// node
// 双向链表的节点
type node[T any] struct {
	prev *node[T]
	next *node[T]
	val T
}

// LinkedList
// 双向循环链表
type LinkedList[T any] struct {
	head *node[T]
	tail *node[T]
	length int
}

// NewLinkedList
// 创建一个双向循环链表
func NewLinkedList[T any]() *LinkedList[T] {
	head := &node[T]{}
	tail := &node[T]{next: head, prev: head}
	head.next, head.prev = tail, tail
	return &LinkedList[T]{
		head: head,
		tail: tail,
	}
}

// NewLinkListOf
// 将切片转换为双向循环链表，直接使用切片元素的值
func NewLinkListOf[T any](ts []T) *LinkedList[T] {
	list := NewLinkedList[T]()
	if err := list.Append(ts...); err != nil {
		panic(err)
	}
	return list
}

// Append
// 往链表最后添加元素
func (l *LinkedList[T]) Append(ts ...T) error {
	for _, t := range ts {
		node := &node[T]{prev: l.tail.prev, next: l.tail, val: t}
		node.prev.next, node.next.prev = node, node
		l.length++
	}
	return nil
}

type  struct{}

func () Get(index int) (interface{}, error) {
	panic("implement me")
}

func () Append(ts ...interface{}) error {
	panic("implement me")
}

func () Add(index int, t interface{}) error {
	panic("implement me")
}

func () Set(index int, t interface{}) error {
	panic("implement me")
}

func () Delete(index int) (interface{}, error) {
	panic("implement me")
}

func () Len() int {
	panic("implement me")
}

func () Cap() int {
	panic("implement me")
}

func () Range(fn func(index int, t interface{}) error) error {
	panic("implement me")
}

func () ChangeSlice() []interface{} {
	panic("implement me")
}

func (l *LinkedList[T])
