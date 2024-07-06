package _0240127

import (
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

// 双链表
type CacheNode struct {
	key  int
	num  int
	prev *CacheNode
	next *CacheNode
}

// 2024/1/27 19:27
// https://leetcode.cn/problems/lru-cache/
type LRUCache struct {
	capacity int                // LRU缓存长度
	length   int                // 已有元素长度，length <= capacity
	head     *CacheNode         // 头部节点哨兵
	tail     *CacheNode         // 尾部节点哨兵，当capacity容量满，同时value不在其中
	nmap     map[int]*CacheNode // 值对应的节点map
}

func Constructor(capacity int) LRUCache {
	if capacity == 0 {
		panic("lru capacity is too small")
	}

	var lru = LRUCache{
		capacity: capacity,
		length:   0,
		head:     &CacheNode{},
		tail:     &CacheNode{},
		nmap:     make(map[int]*CacheNode),
	}

	// 初始头、尾
	concatNode(lru.head, lru.tail)
	return lru
}

func concatNode(p1, p2 *CacheNode) {
	p1.next = p2
	p2.prev = p1
}

// 刷新node，将node移动到首位，并更新其值为val
func (c *LRUCache) refresh(node *CacheNode, val int) {
	// 刷新node值
	if node.num != val {
		node.num = val
	}

	// 如果非首位，则需要调整查询node的顺序
	headNext := c.head.next
	if node != headNext {
		// 提取node
		nodeNext, nodePrev := node.next, node.prev
		concatNode(nodePrev, nodeNext)

		// 插入首位
		concatNode(node, headNext)
		concatNode(c.head, node)
	}
}
func (c *LRUCache) Get(key int) int {
	if node, ok := c.nmap[key]; ok {
		// 刷新node，将node移动到首位
		c.refresh(node, node.num)

		// 返回对应node值
		return node.num
	}

	return -1
}

// 插入元素
//  1. 插入元素存在:
//     1.1 找到元素，更新值
//     1.2 元素是否首位，首位则更新链表，存在提取到首部
//  2. 插入元素不存在:
//     2.1 创建新node，LRU是否已满(len < cap)：
//     * 未满: 插入到头部，len++
//     * 已满：淘汰尾部旧node，插入新node到头部
func (c *LRUCache) Put(key int, value int) {
	node, ok := c.nmap[key]
	// 插入元素存在
	if ok {
		c.refresh(node, value)
		return
	}

	// 插入元素不存在
	newNode := &CacheNode{key: key, num: value}
	c.nmap[key] = newNode
	if c.length < c.capacity { // 容量充足
		c.length++
	} else { // 容量已满, 淘汰末尾节点
		removeNode := c.tail.prev
		concatNode(removeNode.prev, c.tail)
		delete(c.nmap, removeNode.key)
	}

	// 插入首位
	concatNode(newNode, c.head.next)
	concatNode(c.head, newNode)
}

func (c *LRUCache) String() string {
	buf := strings.Builder{}
	for p := c.head.next; p != c.tail; p = p.next {
		buf.WriteString(strconv.Itoa(p.num))
		if p.next != c.tail {
			buf.WriteString("->")
		}
	}
	return buf.String()
}

func TestLRU(t *testing.T) {
	lruObj := Constructor(2)
	lru := &lruObj
	lru.Put(1, 1) // {1=1}
	t.Logf("lru str: %s, %v", lru, *lru)
	lru.Put(2, 2) // {2=2}
	t.Logf("lru str: %s, %v", lru, *lru)
	lru.Put(3, 3) // {3=3,2=2}
	t.Logf("lru str: %s, %v", lru, *lru)
	lru.Put(4, 4) // {4=4,3=3}
	t.Logf("lru str: %s, %v", lru, *lru)
	assert.Equal(t, -1, lru.Get(5))
	assert.Equal(t, 3, lru.Get(3))
	assert.Equal(t, "3->4", lru.String())
	assert.Equal(t, -1, lru.Get(1))
	assert.Equal(t, 4, lru.Get(4))
	assert.Equal(t, "4->3", lru.String())

	lruObj = Constructor(1)
	lru = &lruObj
	lru.Put(2, 1)
	t.Logf("lru %v", lruObj)
	assert.Equal(t, -1, lru.Get(1))
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
