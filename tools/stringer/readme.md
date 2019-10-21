go get -u golang.org/x/tools/cmd/stringer

推荐在go generate下使用，自动生成int枚举的string方法

//go:generate stringer -type ErrCode

