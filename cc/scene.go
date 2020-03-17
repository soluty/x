package cc

// 框架层需要解决存储和发包的过程
// 假设服务器只有一个scene
// 序列化场景
// game 是一个scene   有account管理
// 每个player 是一个scene
// map是一个scene
type Scene struct {
	root *Node
}

func (s *Scene) Save() {

}

func (s *Scene) Load() {

}



