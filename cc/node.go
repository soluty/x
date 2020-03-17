package cc

import (
	"github.com/soluty/x/emitter"
	"log"
	"reflect"
)

var uidGen uint64 = 0

// 为什么Node有一些XXX_开头的变量，因为Node需要深拷贝以便保存数据
type Node struct {
	emitter.Emitter

	XXX_active   bool // 是否激活，激活的节点其Update才触发, 私有属性

	X        float32 // x 坐标
	Y        float32 // y 坐标
	Children []*Node // 子节点列表
	// activeInHierarchy()


	components   []Component

	id           uint64
	parent       *Node
	level        uint
	siblingIndex int
}

func NewNode() *Node {
	uidGen++
	n := &Node{
		id:         uidGen,
		XXX_active: true,
	}
	return n
}

// 必须在scene主线程调用copy方法，然后通过channel给辅线程
// 深拷贝，浅拷贝？
func (n *Node) Copy() Node {
	return *n
}

func (n *Node) SetSiblingIndex(idx int) {
	if n.parent == nil {
		return
	}
	//oldSib := n.siblingIndex
	//n.siblingIndex = minmax(idx, 0, len(n.parent.Children))
	//n.parent.Children[:n.siblingIndex]
}

func (n *Node) GetSiblingIndex() int {
	return n.siblingIndex
}

func (n *Node) AddChild(child *Node, idx ...int) {

}

func (n *Node) SetParent(parent *Node) {
	if n.parent == parent {
		return
	}
	var oldParent = n.parent
	n.parent = parent
	if parent != nil {
		n.level = parent.level + 1
		n.siblingIndex = len(parent.Children)
		parent.Children = append(parent.Children, n)
	}
	if oldParent != nil {
		removeAt := getIndex(oldParent.Children, n)
		if removeAt >= 0 {
			oldParent.Children = append(oldParent.Children[:removeAt], oldParent.Children[removeAt+1:]...)
		}
	}
}

func (n *Node) GetParent() *Node {
	return n.parent
}

func (n *Node) Update(dt uint64) {
	if !n.XXX_active {
		return
	}
	for _, c := range n.components {
		if c.IsEnable() {
			if u, ok := c.(Updater); ok {
				u.Update(dt)
			}
		}
	}
	for _, childNode := range n.Children {
		childNode.Update(dt)
	}
}

func (n *Node) AddComponent(c Component, names ...string) *Node {
	name := ""
	if len(names) > 0 {
		name = names[0]
	}
	if name == "" {
		typ := reflect.TypeOf(c)
		if typ.Kind() == reflect.Ptr {
			typ = typ.Elem()
		} else if typ.Kind() == reflect.Struct {

		} else {
			log.Fatalf("Component must be a struct")
			return n
		}
		name = typ.Name()
	}
	if name == "" {
		log.Fatalf("Component must have name")
		return n
	}
	c._Init(name, n)
	if c.IsUnique() {
		if containsComponent(n.components, c) {
			log.Println("add a same component")
			return n
		}
	}
	if o, ok := c.(OnLoader); ok {
		c._Base().once.Do(func() {
			o.OnLoad()
		})
	}
	n.components = append(n.components, c)
	return n
}

func (n *Node) GetComponent(name string) Component {
	for _, value := range n.components {
		if value.Name() == name {
			return value
		}
	}
	return nil
}

func (n *Node) GetComponents(names ...string) []Component {
	if len(names) == 0 {
		return n.components
	}
	name := names[0]
	var ret []Component
	for _, value := range n.components {
		if value.Name() == name {
			ret = append(ret, value)
		}
	}
	return ret
}

func containsComponent(cs []Component, c Component) bool {
	for _, value := range cs {
		if value.Name() == c.Name() {
			return true
		}
	}
	return false
}

func getIndex(cs []*Node, c *Node) int {
	for idx, value := range cs {
		if value == c {
			return idx
		}
	}
	return -1
}

func minmax(v int, min int, max int) int {
	if v < min {
		v = min
	}
	if v > max {
		v = max
	}
	return v
}
