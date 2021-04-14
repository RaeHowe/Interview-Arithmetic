package main

type LRUCache struct {
	size int //代表了这个缓存系统中现有的元素数量
	capacity int //代表了这个缓存系统的容量
	cache map[int]*DLinkedNode
	head, tail *DLinkedNode //缓存的头部节点和尾部节点
}

//设置一个双向链表数据结构
type DLinkedNode struct {
	key, value int //节点的kv值信息
	prev, next *DLinkedNode //前驱和后继指针
}

/* 简述 LRU 算法及其实现方式 */
func main()  {

}

//初始化双向链表的一个节点
func initDLinkedNode(key, value int) *DLinkedNode {
	return &DLinkedNode{
		key: key,
		value: value,
	}
}

//构造器(指定好容量大小)
func Constructor(capacity int) LRUCache {
	l := LRUCache{
		capacity: capacity,
		cache: map[int]*DLinkedNode{},
		head: initDLinkedNode(0, 0),
		tail: initDLinkedNode(0, 0),
	}

	//头部节点的下一个节点肯定就是尾部节点，尾部节点的前驱节点肯定就是头部节点
	l.head.next = l.tail
	l.tail.prev = l.head

	return l
}

//读取数据操作。逻辑分析，如果没有读取到数据，就返回1；如果读取到数据的话，首先要把该数据放置到链表首部位置，然后再返回结果
func (this *LRUCache) Get(key int) int {
	if _, ok := this.cache[key]; !ok{
		//如果没有读取到数据的话
		return -1 //返回-1
	}

	//如果读取到结果的话，就把该值放置到链表头部位置
	node := this.cache[key]
	this.moveToHead(node)
	return node.value //并且返回这个key对应的value信息
}

//新增数据操作。逻辑分析：首先判断目前缓存内部是否有该数据信息，如果有的话，就把该节点放置到链表首部；如果没有的话进行插入操作，如果size和容量已经相等的话，再去把链表尾部的节点进行删除操作。
func (this *LRUCache) Put(key int, value int) {
	if _, ok := this.cache[key]; !ok{
		//如果不存在的话
		node := initDLinkedNode(key, value)
		this.cache[key] = node
		this.addToHead(node)
		this.size++

		if this.size > this.capacity{
			removed := this.removeTail() //删除链表尾部节点
			delete(this.cache, removed.key)  //从cache的map对象中删除这个key
			this.size--
		}
	}else { //如果已经有了的话
		node := this.cache[key] //就把该节点放置到链表头部
		node.value = value
		this.moveToHead(node)
	}
}

//新增节点的两种情况，一种是新插入的数据，放置到链表头部位置:addToHead;另一种是链表中已经存在，在读取的时候，把这个节点移动到链表头部位置:moveToHead
func (this *LRUCache) addToHead(node *DLinkedNode) {
	node.prev = this.head //在头部插入节点
	node.next = this.head.next
	this.head.next.prev = node
	this.head.next = node
}

func (this *LRUCache) moveToHead(node *DLinkedNode) {
	this.removeNode(node) //把一个节点插入到头部节点的话，先删除该节点
	this.addToHead(node) //再把该节点插入到头部位置
}

//删除节点的情况：当新的数据插入进来了之后，那么在队尾的链表节点就会被移除掉
func (this *LRUCache) removeNode(node *DLinkedNode) {
	node.prev.next = node.next
	node.next.prev = node.prev //由于是双向链表，所以这个类似操作都得操作两次：一个前驱指针的改变，一个后继指针的改变
}

//删除队尾链表节点
func (this *LRUCache) removeTail() *DLinkedNode {
	node := this.tail.prev
	this.removeNode(node)
	return node
}