package skiplist

import (
	"testing"
)

func TestSkipList(t *testing.T) {
	sl := NewSkipList()
	for i := 0; i < 100; i++ {
		sl.Insert(int32(i), i)
	}

	res := print(sl)
	if len(res) != 100 {
		t.Error("length error ", len(res))
	}
	for k, v := range res {
		if k != int(v.key) {
			t.Errorf("kv error: %d %d\n", k, v.key)
		}
	}

	for i := 0; i < 100; i++ {
		value, err := sl.SearchValue(int32(i))
		if err != nil {
			t.Error("search error")
		}
		if value != i {
			t.Error("value error")
		}
	}

	err := sl.Remove(7)
	if err != nil {
		t.Error("remove error")
	}
	value, err := sl.SearchValue(7)
	if err == nil || value != nil {
		t.Error("remove search error")
	}
}

func print(sl *SkipList) []*node {
	n := sl.head
	for n.down != nil {
		n = n.down
	}

	res := make([]*node, 0)
	for n.next.key != MaxIntValue {
		n = n.next
		res = append(res, n)
	}
	return res
}
