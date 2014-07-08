package ethwire

import (
	"fmt"
	"github.com/ethereum/eth-go/ethutil"
	"net"
	"time"
)

// TODO graceful error handling

// Child connection serves to connection to an internal node (but not necesarly)
// This connection is completely synchronous and therefor blocking.
type ChildConn struct {
	conn net.Conn
	host string
	msgs chan *Msg
}

func NewChildConn(host string) (*ChildConn, error) {
	conn, err := net.DialTimeout("tcp", host, 5*time.Second)
	if err != nil {
		return nil, err
	}

	self := &ChildConn{conn: conn, host: host, msgs: make(chan *Msg, 1)}

	return self, nil
}

func (self *ChildConn) GetBlock(hash []byte) *ethutil.Value {
	self.writeMessage(NewMessage(MsgChildGetBlockTy, hash))

	// Get response (blocking)
	msg := self.GetMessage()
	if msg.Type != MsgChildBlockTy {
		panic(fmt.Sprintf("expected block, got %v", msg.Type))
	}

	return msg.Data
}

func (self *ChildConn) GetMessage() *Msg {
	return <-self.msgs
}

func (self *ChildConn) writeMessage(msg *Msg) {
	WriteMessage(self.conn, msg)
}

func (self *ChildConn) readMessage() (*Msg, error) {
	b := make([]byte, 1440)

	n, _ := self.conn.Read(b)
	msg, _, _, err := ReadMessage(b[:n])
	if err != nil {
		return nil, err
	}

	return msg, nil
}

func (self *ChildConn) handleInbound() {
	for {
		msg, err := self.readMessage()
		if err != nil {
			panic(err)

			// TODO
			continue
		}

		switch msg.Type {
		case MsgHandshakeTy:
			// fine
		default:
			self.msgs <- msg
		}
	}
}
