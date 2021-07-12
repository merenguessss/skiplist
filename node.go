package skiplist

type node struct {
	key   int32
	value interface{}
	up    *node
	down  *node
	next  *node
	pre   *node
}

func newNode(key int32, value interface{}) *node {
	return &node{
		key:   key,
		value: value,
	}
}
