## 1
`go build -ldflags "-s -w"` 这种方式编译。

解释一下参数的意思：

-ldflags： 表示将后面的参数传给连接器（5/6/8l）
-s：去掉符号信息
-w：去掉DWARF调试信息
注意：

-s 去掉符号表（这样panic时，stack trace就没有任何文件名/行号信息了，这等价于普通C/C+=程序被strip的效果）

-w 去掉DWARF调试信息，得到的程序就不能用gdb调试了

##2
(upx)https://upx.github.io/ 压缩