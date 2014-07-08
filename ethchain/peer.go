package ethchain

import (
	"time"
)

type Peer interface {
	Inbound() bool
	LastSend() time.Time
	LastPong() int64
	Host() []byte
	Port() uint16
	Version() string
	PingTime() string
	Connected() *int32
}
