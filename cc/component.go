package cc

import (
	"sync"
)

type (
	// 生命周期方法，和cocos creator一样
	Starter interface {
		Start()
	}
	Updater interface {
		Update(dt uint64)
	}
	LateUpdater interface {
		LateUpdate(dt uint64)
	}
	OnEnabler interface {
		OnEnable()
	}
	OnDisabler interface {
		OnDisable()
	}
	OnLoader interface {
		OnLoad()
	}
	OnDestroyer interface {
		OnDestroy()
	}
)

// 修改Component需要发包给客户端
type Component interface {
	_Init(name string, n *Node)
	_Base() *BaseComponent
	IsEnable() bool

	Node() *Node
	Name() string
	SetEnable(e bool)
	IsUnique() bool
	GetComponent(string) Component
}

type BaseComponent struct {
	node   *Node
	name   string
	enable bool
	once   sync.Once
}

func (c *BaseComponent) _Init(name string, n *Node) {
	c.name = name
	c.node = n
}

func (c *BaseComponent) _Base() *BaseComponent {
	return c
}

func (c *BaseComponent) Node() *Node {
	return c.node
}

func (c *BaseComponent) Name() string {
	return c.name
}

func (c *BaseComponent) IsEnable() bool {
	return c.enable
}

func (c *BaseComponent) IsUnique() bool {
	return true
}

func (c *BaseComponent) SetEnable(e bool) {
	c.enable = e
}

func (c *BaseComponent) GetComponent(name string) Component {
	return c.node.GetComponent(name)
}
