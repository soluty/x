package xnet

import (
	"io"
	"net"
	"time"

	"gopkg.in/tomb.v2"
)

type Server struct {
	l net.Listener
	t tomb.Tomb

	p     Protocol
	codec Codec
}

func (s *Server) Accept(l net.Listener) {
	// rpc.NewServer()
	for {
		conn, err := l.Accept()
		if err != nil {
			if ne, ok := err.(net.Error); ok && ne.Temporary() {
				time.Sleep(10 * time.Millisecond)
				continue
			}
			//log.Print("server: accept:", err.Error())
			s.t.Kill(err)
			return
		}
		s.t.Go(func() error {
			defer func() {

			}()
			s.serveConn(conn)
			return nil
		})
	}
}

type Protocol interface {
	NewCodec(rw io.ReadWriteCloser) Codec
}

type Codec interface {
	Receive() (interface{}, error)
	Send(interface{}) error
	// Close can be called multiple times and must be idempotent.
	Close()
}

func (s *Server) Send(msg interface{}) {
	s.codec.Send(msg)
}

func (s *Server) serveConn(conn io.ReadWriteCloser) {
	//s.codec = s.p.NewCodec(conn)
	//s.serveCodec(srv)
}

func (s *Server) serveCodec(codec Codec) {
	//for  {
	//	msg, err := codec.Receive()
	//	codec.Send(1)+
	//}
	//codec.Close()
}
