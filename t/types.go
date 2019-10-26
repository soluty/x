package t

// 主要是方便双层map或者list的书写，单层可用可不用
// 有更加有需要的数据结构，使用https://github.com/emirpasic/gods，注意这里的数据结构都非线程安全
// 数据结构学习https://visualgo.net/zh
type Map = map[string]interface{}
type MapString = map[string]string
type MapInt = map[string]int
type MapInt32 = map[string]int32
type MapUint = map[string]uint

type MapStringAny = Map
type MapStringString = MapString
type MapStringInt = MapInt
type MapStringInt32 = MapInt32
type MapStringUint = MapUint

type MapIntAny = map[int]interface{}
type MapIntString = map[int]string
type MapIntInt = map[int]int
type MapIntInt32 = map[int]int32
type MapIntUint = map[int]uint

type Slice = []interface{}
type SliceString = []string
type SliceInt = []int
type SliceInt32 = []int32
type SliceInt64 = []int64
type SliceUint = []uint
type SliceUint32 = []uint32
type SliceUint64 = []uint64
