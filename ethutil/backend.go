package ethutil

type Backend interface {
	GetBlock(hash []byte) *Value
	Get(hash []byte) *Value
	Put(key, value []byte)
	Delete([]byte) error
}
