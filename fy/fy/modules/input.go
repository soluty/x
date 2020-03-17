package modules

import (
	"context"
	"github.com/soluty/x/fy/fy"
	"net"
	"time"
)

var _ fy.Module = new(InputModule)

type InputModule struct {
}

func (*InputModule) Name() string {
	return "input"
}

func (*InputModule) OnCreate() {
	// tcp input
	// websocket input
	// readline input
}

func (*InputModule) OnStart(ctx context.Context) error {
	l, _ := net.Listen("tcp", ":7070")

	go func() {
		var tempDelay time.Duration
		for {
			switch ll := l.(type) {
			case *net.TCPListener:
				ll.SetDeadline(time.Time{})
				//case *testutils.PipeListener:
			}
			_, err := l.Accept()
			if err != nil {
				if ne, ok := err.(net.Error); ok {
					if ne.Timeout() {
						continue
					}
					if ne.Temporary() {
						if tempDelay == 0 {
							tempDelay = 5 * time.Millisecond
						} else {
							tempDelay *= 2
						}
						if max := 1 * time.Second; tempDelay > max {
							tempDelay = max
						}
						time.Sleep(tempDelay)
						continue
					}
				}
				// todo log
			}
			tempDelay = 0
		}
	}()
	<-ctx.Done()
	return ctx.Err()
}

func (*InputModule) OnStop(reason error) {

}
