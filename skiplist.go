package skiplist

import (
	"errors"
	"math/rand"
)

type SkipList struct {
	head          *node
	tail          *node
	level         int
	promotionRate float32
}

const MaxIntValue = int32(^uint32((0)) >> 1)
const MinIntValue = ^MaxIntValue

var ErrorNotFindNode = errors.New("can't find node")

func NewSkipList() *SkipList {
	head := &node{
		key: MinIntValue,
	}
	tail := &node{
		key: MaxIntValue,
	}
	head.next = tail
	tail.pre = head
	return &SkipList{
		head:          head,
		tail:          tail,
		level:         0,
		promotionRate: 0.5,
	}
}

func (sl *SkipList) SetPromotionRate(rate float32) {
	sl.promotionRate = rate
}

func (sl *SkipList) SearchValue(key int32) (interface{}, error) {
	n := sl.findNode(key)
	if n.key == key {
		return n.value, nil
	}
	return nil, ErrorNotFindNode
}

func (sl *SkipList) findNode(key int32) *node {
	p := sl.head
	for true {
		for p.next.key != MaxIntValue && p.next.key <= key {
			p = p.next
		}
		if p.down == nil {
			break
		}
		p = p.down
	}
	return p
}

func (sl *SkipList) Insert(key int32, value interface{}) error {
	n := sl.findNode(key)
	isPromotion := false
	currLevel := 0
	if n.key == key {
		// todo key 重复
		return nil
	}

	initNode := newNode(key, value)
	sl.appendNode(n, initNode)

	for rand.Float32() < sl.promotionRate && !isPromotion {
		if currLevel == sl.level {
			isPromotion = true
			sl.promoteLevel()
		}

		for n.key != MinIntValue && n.up == nil {
			n = n.pre
		}

		n = n.up
		newPromoteNode := newNode(key, value)
		sl.appendNode(n, newPromoteNode)
		newPromoteNode.down = initNode
		initNode.up = newPromoteNode
		initNode = newPromoteNode
		currLevel++
	}
	return nil
}

func (sl *SkipList) Remove(key int32) error {
	n := sl.findNode(key)
	if n.key != key {
		return ErrorNotFindNode
	}

	currLevel := 0
	for n != nil {
		n.pre.next = n.next
		n.next.pre = n.pre
		if currLevel != 0 && n.pre.key == MinIntValue && n.next.key == MaxIntValue {
			sl.removeLevel(n.pre, n.next)
		} else {
			currLevel++
		}
		n = n.up
	}
	return nil
}

func (sl *SkipList) appendNode(curr *node, newNode *node) {
	newNode.next = curr.next
	curr.next.pre = newNode
	curr.next = newNode
	newNode.pre = curr
}

func (sl *SkipList) removeLevel(leftNode, rightNode *node) {
	if leftNode.up == nil {
		leftNode.down.up = nil
		rightNode.down.up = nil
	} else {
		leftNode.down.up = leftNode.up
		leftNode.up.down = leftNode.down
		rightNode.down.up = rightNode.up
		rightNode.up.down = rightNode.down
	}
}

func (sl *SkipList) promoteLevel() {
	sl.level++
	p := newNode(MinIntValue, nil)
	q := newNode(MaxIntValue, nil)

	p.next = q
	q.next = p
	p.down = sl.head
	q.down = sl.tail
	sl.head.up = p
	sl.tail.up = q
	sl.head = p
	sl.tail = q
}
