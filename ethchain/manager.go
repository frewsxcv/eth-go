package ethchain

import (
	"container/list"
	"github.com/ethereum/eth-go/ethcrypto"
	"github.com/ethereum/eth-go/ethutil"
	"github.com/ethereum/eth-go/ethwire"
)

type EthManager interface {
	StateManager() *StateManager
	BlockChain() *BlockChain
	TxPool() *TxPool
	Broadcast(msgType ethwire.MsgType, data []interface{})
	Reactor() *ethutil.ReactorEngine
	PeerCount() int
	IsMining() bool
	IsListening() bool
	Peers() *list.List
	KeyManager() *ethcrypto.KeyManager
	ClientIdentity() ethwire.ClientIdentity
	Backend() ethutil.Backend
}
