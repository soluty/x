package gocos

type Component interface {
	Node() *Node
}

type BaseComponent struct {

}


type Node struct {

}

// 往任意节点的任意组件发送消息
func sendMsg(nodeId int64, component string, res interface{}, req ...interface{}) error{

}



// GetComponent("Abc")
// 假设这是一个 server
func (a *Acomp) Afunc(int) string  {
	//分布式的node
}