package ethdb

import (
	"github.com/ethereum/eth-go/ethutil"
)

type DbBackend struct {
	db ethutil.Database
}

func NewDbBackend(db ethutil.Database) *DbBackend {
	return &DbBackend{db: db}
}

func (self *DbBackend) GetBlock(hash []byte) *ethutil.Value {
	data, _ := self.db.Get(hash)
	if len(data) == 0 {
		return nil
	}

	return ethutil.NewValueFromBytes(data)
}

func (self *DbBackend) Get(hash []byte) *ethutil.Value {
	data, _ := self.db.Get(hash)
	if len(data) == 0 {
		return nil
	}

	return ethutil.NewValueFromBytes(data)
}

func (self *DbBackend) Put(key, value []byte) {
	self.db.Put(key, value)
}

func (self *DbBackend) Delete(key []byte) error {
	return self.db.Delete(key)
}
