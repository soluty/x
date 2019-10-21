package t

// 主要是方便双层map或者list的书写，单层可用可不用
// 有更加有需要的数据结构，使用https://github.com/emirpasic/gods，注意这里的数据结构都非线程安全
// 数据结构学习https://visualgo.net/zh
type Map = map[string]interface{}
type MapStr = map[string]string
type MapInt = map[string]int
type MapInt32 = map[string]int32
type MapUint = map[string]uint

type IntMap = map[int]interface{}
type IntMapStr = map[int]string
type IntMapInt = map[int]int
type IntMapInt32 = map[int]int32
type IntMapUint = map[int]uint

type List = []interface{}
type StrList = []string
type IntList = []int
type Int32List = []int32
type UintList = []uint

func _() {
	_ = Map{
		"a": "a",
		"b": "cc",
	}
}
