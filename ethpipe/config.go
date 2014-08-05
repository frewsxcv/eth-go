package ethpipe

import "github.com/ethereum/eth-go/ethutil"

var cnfCtr = ethutil.Hex2Bytes("661005d2720d855f1d9976f88bb10c1a3398c77f")

type config struct {
	pipe *Pipe
}

func (self *config) Get(name string) *Object {
	configCtrl := self.pipe.World().safeGet(cnfCtr)
	var addr []byte

	switch name {
	case "NameReg":
		addr = []byte{0}
	default:
		addr = ethutil.RightPadBytes([]byte(name), 32)
	}

	objectAddr := configCtrl.GetStorage(ethutil.BigD(addr))

	return &Object{self.pipe.World().safeGet(objectAddr.Bytes())}
}

func (self *config) Exist() bool {
	return self.pipe.World().Get(cnfCtr) != nil
}