package gate

import (
	"fmt"
	"github.com/soluty/x/qp/config"
	"github.com/soluty/x/qp/pb"
	"net"
	"sync"
	"time"
)

type Server struct {
	L net.Listener

	playerConns sync.Map // playerId -> playerConn
	connPlayers sync.Map // conn -> playerConn

	loginLock sync.Mutex
}

// 玩家和连接的
type playerConn struct {
	playerId string
	conn     net.Conn
}

func (s *Server) Close() error {
	return s.L.Close()
}

func (s *Server) getByConn(conn net.Conn) *playerConn {
	r, ok := s.connPlayers.Load(conn)
	if !ok {
		return nil
	}
	return r.(*playerConn)
}

func (s *Server) getByPlayerId(playerId string) *playerConn {
	r, ok := s.connPlayers.Load(playerId)
	if !ok {
		return nil
	}
	return r.(*playerConn)
}

func (s *Server) read(conn net.Conn) (msg *pb.C2SMsg, err error) {
	// todo
	return nil, nil
}

func (s *Server) Serve(listener net.Listener) error {
	l := &onceCloseListener{Listener: listener}
	s.L = l
	defer l.Close()
	var tempDelay time.Duration
	for {
		conn, err := l.Accept()
		if err != nil {
			select {
			// 等待关闭信号
			default:
			}
			if ne, ok := err.(net.Error); ok && ne.Temporary() {
				if tempDelay == 0 {
					tempDelay = 50 * time.Millisecond
				} else {
					tempDelay *= 2
				}
				if max := 1 * time.Second; tempDelay > max {
					tempDelay = max
				}
				time.Sleep(tempDelay)
				continue
			}
			return err
		}
		tempDelay = 0
		fmt.Println("accept ok")
		go func() {
			req, err := s.read(conn)
			head := int32(req.Head)
			if head <= config.ServerMsgStart {
				// disconnect
				return
			}
			if req.Head < config.GateServerMsgEnd {
				if req.Head == pb.C2SMsg_Ping {

				} else if req.Head == pb.C2SMsg_Login {

				} else {
					// disconnect ??  unknown msg, other version?
					return
				}
			} else {
				// need login
				player := s.getByConn(conn)
				if player == nil {
					// connect
					return
				}
				req.PlayerId = player.playerId
				// 转发给其它服
				switch {
				case head >= config.GateServerMsgEnd && head < config.GameServerMsgEnd: // in room msg
					//req.RoomId = "get room"
					//gameClient := pb.NewGameProtobufClient("", http.DefaultClient)
					//res, err := gameClient.ReceiveMsg(context.Background(), req)
				case head >= config.GameServerMsgEnd && head < config.OtherServerMsgEnd: // not in room msg
				default:
				}
			}
		}()
	}
}

type onceCloseListener struct {
	net.Listener
	once     sync.Once
	closeErr error
}

func (oc *onceCloseListener) Close() error {
	oc.once.Do(oc.close)
	return oc.closeErr
}

func (oc *onceCloseListener) close() { oc.closeErr = oc.Listener.Close() }
