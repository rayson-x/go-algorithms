package list

type List struct {
	root *Element
}

type Element struct {
	prev,next *Element
	value interface{}
}

func New() *List {
	// 初始化根节点
	root := new(Element)
	root.prev = root
	root.next = root

	list := &List {root: root}

	return list
}

// 从尾部插入
func (this *List) Append (value interface{}) {
	element := &Element {
		value: value,
		prev : this.root.prev,
		next : this.root,
	}

	// 之前最后一个的下一节点设置为插入节点
	element.prev.next = element
	this.root.prev = element
}

// 从首部插入
func (this *List) PushFront (value interface{}) {
	element := &Element {
		value: value,
		prev : this.root,
		next : this.root.next,
	}

	// 之前首部的上一节点设置为插入节点
	element.next.prev = element
	this.root.next = element
}

// 从弹出尾部最后一个节点
func (this *List) Pop () interface{} {
	// 取出尾部节点
	element := this.root.prev
	// 如果尾部节点的上一节点为root节点,说明链表内数据清空
	if (element == this.root) {
		return element.value
	}

	this.root.prev = element.prev
	element.prev.next = element.next

	return element.value
}

// 弹出首部第一个节点
func (this *List) Shift () interface{} {
	// 取出首部节点
	element := this.root.next
	// 如果首部节点的下一节点为root节点,说明链表内数据清空
	if (element == this.root) {
		return element.value
	}

	this.root.next = element.next
	element.next.prev = element.prev

	return element.value
}

func (this *List) Top () interface{} {
	return this.root.next.value
}

func (this *List) End () interface{} {
	return this.root.prev.value
}
