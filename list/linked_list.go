package list

import "github.com/Gnoloayoul/GYLgg/internal/errs"

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
func NewLinkedListOf[T any](ts []T) *LinkedList[T] {
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

func (l *LinkedList[T]) findNode(index int) *node[T] {
	var cur *node[T]
	if index <= l.Len() / 2 {
		cur = l.head
		for i := -1; i < index; i++ {
			cur = cur.next
		}
	} else {
		cur = l.tail
		for i := l.Len(); i > index; i-- {
			cur = cur.prev
		}
	}
	return cur
}

func (l *LinkedList[T]) Get(index int) (T, error) {
	if !l.checkIndex(index) {
		var zeroValue T
		return zeroValue, errs.NewErrIndexOutOfRange(l.Len(), index)
	}
	n := l.findNode(index)
	return n.val, nil
}

func (l *LinkedList[T]) checkIndex(index int) bool {
	return 0 <= index && index < l.Len()
}

// Add
// 在下标为 index 的位置插入一个元素
// 当 index 等于 LinkedList 长度等同于Append
func (l *LinkedList[T]) Add(index int, t T) error {
	if index < 0 || index > l.length {
		return errs.NewErrIndexOutOfRange(l.length, index)
	}
	if index == l.length {
		return l.Append(t)
	}
	next := l.findNode(index)
	node := &node[T]{prev: next.prev, next: next, val: t}
	node.prev.next, node.next.prev = node, node
	l.length++
	return nil
}

// Set
// 设置链表中 index 索引处的值为t
func (l *LinkedList[T]) Set(index int, t T) error {
	if !l.checkIndex(index) {
		return errs.NewErrIndexOutOfRange(l.Len(), index)
	}
	node := l.findNode(index)
	node.val = t
	return nil
}

// Delete
// 删除指定位置的元素
func (l *LinkedList[T]) Delete(index int) (T, error) {
	if !l.checkIndex(index) {
		var zeroValue T
		return zeroValue, errs.NewErrIndexOutOfRange(l.Len(), index)
	}
	node := l.findNode(index)
	node.prev.next = node.next
	node.next.prev = node.prev
	node.prev, node.next = nil, nil
	l.length--
	return node.val, nil
}

func (l *LinkedList[T]) Len() int {
	return l.length
}

func (l *LinkedList[T]) Cap() int {
	return l.Len()
}

func (l *LinkedList[T]) Range(fn func(index int, t T) error) error {
	for cur, i := l.head.next, 0; i < l.length; i++ {
		err := fn(i, cur.val)
		if err != nil {
			return err
		}
		cur = cur.next
	}
	return nil
}

// ChangeSlice
// 将 LinkedList 转化为切片。
// 是  NewLinkListOf 的逆操作
func (l *LinkedList[T]) ChangeSlice() []T {
	slice := make([]T, l.length)
	for cur, i := l.head.next, 0; i < l.length; i++ {
		slice[i] = cur.val
		cur = cur.next
	}
	return slice
}

