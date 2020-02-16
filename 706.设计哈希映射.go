/*
 * @lc app=leetcode.cn id=706 lang=golang
 *
 * [706] 设计哈希映射
 */

// @lc code=start
const conuntBuckt = 16

type LinkNode struct {
	Data KV
	Next *LinkNode
}

type KV struct {
	Key   *int
	Value int
}

/** Initialize your data structure here. */
type MyHashMap struct {
	Bucket [conuntBuckt]*LinkNode
}

/** Initialize your data structure here. */
func Constructor() MyHashMap {
	myMap := MyHashMap{}
	for i := 0; i < conuntBuckt; i++ {
		myMap.Bucket[i] = &LinkNode{KV{nil, 0}, nil}
	}
	return myMap
}

func (this *LinkNode) AddNode(key int, value int) {
	data := KV{Key: &key, Value: value}
	node := this
	for {
		if *node.Data.Key == key {
			node.Data.Value = value
			break
		}
		if node.Next == nil {
			node.Next = &LinkNode{Data: data, Next: nil}
			break
		} else {
			node = node.Next
		}
	}
}

func HashCode(key int) int {
	return (key % conuntBuckt)
}

/** value will always be non-negative. */
func (this *MyHashMap) Put(key int, value int) {
	mapIndex := HashCode(key)
	link := this.Bucket[mapIndex]
	if link.Data.Key == nil {
		link.Data.Key = &key
		link.Data.Value = value
	} else {
		link.AddNode(key, value)
	}
}

/** Returns the value to which the specified key is mapped, or -1 if this map contains no mapping for the key */
func (this *MyHashMap) Get(key int) int {
	mapIndex := HashCode(key)
	link := this.Bucket[mapIndex]
	node := link
	ret := -1
	if link.Data.Key != nil {
		for {
			if node == nil {
				break
			}
			if *node.Data.Key == key {
				ret = node.Data.Value
				break
			} else {
				node = node.Next
			}
		}
	}
	return ret
}

/** Removes the mapping of the specified value key if this map contains a mapping for the key */
func (this *MyHashMap) Remove(key int) {
	mapIndex := HashCode(key)
	link := this.Bucket[mapIndex]
	if link.Data.Key == nil {
		return
	}
	if *link.Data.Key == key {
		if link.Next == nil {
			this.Bucket[mapIndex] = &LinkNode{KV{nil, 0}, nil}
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
		if *node.Next.Data.Key == key {
			node.Next = node.Next.Next
			break
		} else {
			node = node.Next
		}
	}
}

/**
 * Your MyHashMap object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Put(key,value);
 * param_2 := obj.Get(key);
 * obj.Remove(key);
 */
// @lc code=end

