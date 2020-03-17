package main

import (
	"fmt"
	"github.com/soluty/x/qp/modules/gate"
	"time"
)

// 确保每一个连接的每一次请求都有用

func main() {
	s := &gate.Server{

	}
	go func() {
		time.Sleep(time.Second)
		s.Close()
		time.Sleep(time.Second)
		fmt.Println("listen again")
		//l, err := net.Listen("tcp", ":8181")
		//if err != nil {
		//	panic(err)
		//}
		//s.L = l
	}()
	err := s.Serve(":8181")
	fmt.Println(err)
}