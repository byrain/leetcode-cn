/*
 * @lc app=leetcode.cn id=705 lang=golang
 *
 * [705] 设计哈希集合
 */

const conuntBuckt = 16

type LinkNode struct {
	Key  *int
	Next *LinkNode
}

/** Initialize your data structure here. */
type MyHashSet struct {
	Bucket [conuntBuckt]*LinkNode
}

/** Initialize your data structure here. */
func Constructor() MyHashSet {
	mySet := MyHashSet{}
	for i := 0; i < conuntBuckt; i++ {
		mySet.Bucket[i] = &LinkNode{nil, nil}
	}
	return mySet
}

func (this *LinkNode) AddNode(key int) {
	node := this
	for {
		if *node.Key == key {
			break
		}
		if node.Next == nil {
			node.Next = &LinkNode{Key: &key, Next: nil}
			break
		} else {
			node = node.Next
		}
	}
}

func HashCode(key int) int {
	return (key % conuntBuckt)
}

func (this *MyHashSet) Add(key int) {
	mapIndex := HashCode(key)
	link := this.Bucket[mapIndex]
	if link.Key == nil {
		link.Key = &key
	} else {
		link.AddNode(key)
	}
}

func (this *MyHashSet) Remove(key int) {
	mapIndex := HashCode(key)
	link := this.Bucket[mapIndex]
	if link.Key == nil {
		return
	}
	if *link.Key == key {
		if link.Next == nil {
			this.Bucket[mapIndex] = &LinkNode{nil, nil}
		} else {
			this.Bucket[mapIndex] = link.Next
		}
		return
	}
	node := link
	for {
		if node.Next == nil {
			return
		}
		if *node.Next.Key == key {
			node.Next = node.Next.Next
			break
		} else {
			node = node.Next
		}
	}
}

/** Returns true if this set Contains the specified element */
func (this *MyHashSet) Contains(key int) bool {
	mapIndex := HashCode(key)
	link := this.Bucket[mapIndex]
	node := link
	ret := false
	if node.Key != nil {
		for {
			if node == nil {
				break
			}
			if *node.Key == key {
				ret = true
				break
			} else {
				node = node.Next
			}
		}
	}
	return ret
}

/**
 * Your MyHashSet object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Add(key);
 * obj.Remove(key);
 * param_3 := obj.Contains(key);
 */
// @lc code=end

