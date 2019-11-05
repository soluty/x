## reflect.TypeOf()如何获得接口类型的
```
// 假设定义接口IA
type IA interface{}
// 通过这段代码获取IA的反射类型信息
reflect.TypeOf(reflect.TypeOf((*IA)(nil)).Elem())
```

