package list

type List struct {
	root Element
	size int
}

type Element struct {
	prev,next *Element
	Value     interface{}
	list      *List
}

func (e *Element) Next () *Element {
	if e.list != nil && e.next == &e.list.root {
		return nil
	}

	return e.next
}

func (e *Element) Prev () *Element {
	if e.list != nil && e.prev == &e.list.root {
		return nil
	}

	return e.prev
}

func New() *List {
	return new(List).Init()
}

// 初始化函数
func (this *List) Init () *List {
	this.root.prev = &this.root
	this.root.next = &this.root
	this.size = 0
	return this
}

// 插入节点到指定节点的下一位
func (this *List) insert (e, at *Element) *Element {
	n := at.next
	at.next = e
	e.prev = at
	e.next = n
	n.prev = e
	e.list = this
	this.size++
	return e
}

// 删除节点
func (this *List) remove(e *Element) *Element {
	e.prev.next = e.next
	e.next.prev = e.prev
	e.next = nil
	e.prev = nil
	e.list = nil
	this.size--
	return e
}

// 从尾部插入
func (this *List) Push (value interface{}) {
	this.insert(&Element{Value: value}, this.root.prev)
}

// 从首部插入
func (this *List) Unshift (value interface{}) {
	this.insert(&Element{Value: value}, &this.root)
}

// 从弹出尾部最后一个节点
func (this *List) Pop () interface{} {
	if this.size == 0 {
		return nil
	}

	return this.remove(this.root.prev).Value
}

// 弹出首部第一个节点
func (this *List) Shift () interface{} {
	if this.size == 0 {
		return nil
	}

	return this.remove(this.root.next).Value
}

// 获取链表长度
func (this *List) Size() int {
	return this.size
}

// 获取首位
func (this *List) Front() *Element {
	return this.root.next
}

// 获取末位
func (this *List) Back() *Element {
	return this.root.prev
}

// 将节点插入到首位
func (this *List) MoveToFront(e *Element) {
	// 检查节点是否在当前链表中,以及链表首个节点是否为当前节点
	if e.list != this || this.root.next == e {
		return
	}
	// 将节点从链表删除后再插入
	this.insert(this.remove(e), &this.root)
}

// 将节点插入到末位
func (this *List) MoveToBack(e *Element) {
	// 检查节点是否在当前链表中,以及链表尾部节点是否为当前节点
	if e.list != this || this.root.prev == e {
		return
	}
	this.insert(this.remove(e), this.root.prev)
}

// 将节点插入到指定节点的上一位
func (this *List) MoveBefore(e, mark *Element) {
	// 检查节点是否在当前链表中
	if e.list != this || e == mark || mark.list != this {
		return
	}
	this.insert(this.remove(e), mark.prev)
}

// 将节点插入到指定节点的下一位
func (this *List) MoveAfter(e, mark *Element) {
	// 检查节点是否在当前链表中
	if e.list != this || e == mark || mark.list != this {
		return
	}
	this.insert(this.remove(e), mark)
}

// 删除节点
func (this *List) Remove(e *Element) interface{} {
	return this.remove(e)
}
