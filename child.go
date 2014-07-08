package eth

import (
	"fmt"
	"github.com/ethereum/eth-go/ethutil"
	"github.com/ethereum/eth-go/ethwire"
	"net"
	"time"
)

// TODO graceful error handling

// Child connection serves to connection to an internal node (but not necesarly)
// This connection is completely synchronous and therefor blocking.
type Child struct {
	*Peer
	conn net.Conn
	host string
}

func NewChild(eth *Ethereum, host string) (*Child, error) {
	conn, err := net.DialTimeout("tcp", host, 5*time.Second)
	if err != nil {
		return nil, err
	}

	peer := NewPeer(conn, eth, false)
	peer.childQueue = make(chan *ethwire.Msg, 5)
	peer.caps = CapTxTy

	self := &Child{peer, conn, host}

	return self, nil
}

func (self *Child) GetBlock(hash []byte) *ethutil.Value {
	self.writeMessage(ethwire.NewMessage(ethwire.MsgChildGetBlockTy, []interface{}{hash}))

	// Get response (blocking)
	msg := <-self.Peer.childQueue
	if msg.Type != ethwire.MsgChildBlockTy {
		panic(fmt.Sprintf("expected block, got %v", msg.Type))
	}

	return msg.Data
}

func (self *Child) Get(hash []byte) *ethutil.Value {
	self.writeMessage(ethwire.NewMessage(ethwire.MsgChildGetHashTy, []interface{}{hash}))

	// Get response (blocking)
	msg := <-self.Peer.childQueue

	return msg.Data.Get(0)
}

func (self *Child) Put(key, value []byte)   {}
func (self *Child) Delete(key []byte) error { return nil }
