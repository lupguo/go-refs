package _0240306

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

// LRUCache /**
// /**
//   - Your LRUCache object will be instantiated and called as such:
//   - obj := Constructor(capacity);
//   - param_1 := obj.Get(key);
//   - obj.Put(key,value);
//     */
//
// Get 函数 get 和 put 必须以 O(1) 的平均时间复杂度运行
// 用到HashMap
type LRUCache struct {
	cap  int
	len  int
	head *DoubleLink         // 头节点
	tail *DoubleLink         // 尾节点
	kmap map[int]*DoubleLink // 通过key健快速找到节点位置
}

type DoubleLink struct {
	data int // 节点元素值
	prev *DoubleLink
	next *DoubleLink
}

// Constructor LRUCache(int capacity) 以 正整数 作为容量 capacity 初始化 LRU 缓存
func Constructor(capacity int) LRUCache {
	cache := LRUCache{
		cap:  capacity,
		len:  0,
		kmap: make(map[int]*DoubleLink),
		head: &DoubleLink{
			data: math.MaxInt,
		},
		tail: &DoubleLink{
			data: math.MaxInt,
		},
	}
	concatNode(cache.head, cache.tail)
	return cache
}

// Get int get(int key) 如果关键字 key 存在于缓存中，则返回关键字的值，否则返回 -1
func (this *LRUCache) Get(key int) int {
	kmap := this.kmap
	if v, ok := kmap[key]; ok {
		return v.data
	}
	return NotExistKey
}

// NotExistKey 是否不存在的key
const NotExistKey = -1

// Put void put(int key, int value) 如果关键字 key 已经存在，则变更其数据值 value ；
// 如果不存在，则向缓存中插入该组 key-value
// 如果插入操作导致关键字数量超过 capacity ，则应该 逐出 最久未使用的关键字。
func (this *LRUCache) Put(key int, value int) {
	// key不存在，初始化Node，并加入head头节点，检测len是否超过了cap，超过则淘汰tail节点
	head := this.head

	if this.Get(key) == NotExistKey {
		// 初始化节点
		node := &DoubleLink{
			data: value,
		}

		// 将node插入到头部
		concatNode(node, head.next)
		concatNode(head, node)

		// kmap更新
		this.kmap[key] = node

		// 容量是否超过了，超过则淘汰tail
		if this.len > this.cap {
			this.tail.prev.next = nil
			this.tail.prev = nil
		} else {
			this.len++
		}

		return
	}

	// key已存在，需要更新
	node := this.kmap[key]

	// 将node插入链表头部，组合node.prev，node.next
	concatNode(node.prev, node.next)
	concatNode(node, head.next)
	concatNode(head, node)
}

// 连接双链表到两个节点
func concatNode(a *DoubleLink, b *DoubleLink) {
	a.next = b
	b.prev = a
}

func TestLRUCache(t *testing.T) {
	cache := Constructor(2)
	cache.Put(1, 1) // 1:1
	cache.Put(2, 2) // 2:2,1:1
	cache.Put(3, 3) //  3:3, 2:2

	assert.Equal(t, 3, cache.Get(3))  // 3
	assert.Equal(t, -1, cache.Get(1)) // -1
	assert.Equal(t, -1, cache.Get(4)) // 4

	cache.Put(2, 2)                   // 2:2, 3:3
	cache.Put(4, 4)                   // 4:4, 2:2
	assert.Equal(t, -1, cache.Get(3)) // -1
	assert.Equal(t, 4, cache.Get(4))  //
}
