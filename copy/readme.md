有些时候我们需要深拷贝，有两种方式实现，一种是序列化然后反序列化，一种是通过反射。
反射的方式效率一般要高一点，所以这里提供反射的方式。

两种方式默认都不会拷贝私有字段
代码拷贝自https://github.com/mohae/deepcopy 感谢。