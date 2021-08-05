package store

import (
	"time"
)

type BasicStore interface {
	Set(k, v []byte) error
	Get(k []byte) ([]byte, error)
	SetEx(k, v []byte, exp time.Duration) error
	Del(k []byte) error
	Shutdown()
}

type Seq struct {
	Max  uint64
	Next func() uint64
}

// KVStore store with custom options.
type KVStore struct {
	Store BasicStore
}

func NewKVStore(basicStore BasicStore) *KVStore {
	return &KVStore{basicStore}
}

func (s *KVStore) Set(k, v string) error {
	return s.Store.Set([]byte(k), []byte(v))
}
func (s *KVStore) Get(k string) ([]byte, error) {
	return s.Store.Get([]byte(k))
}

func (s *KVStore) SetEx(k, v string, exp time.Duration) error {
	return s.Store.SetEx([]byte(k), []byte(v), exp)
}

func (s *KVStore) Del(k string) error {
	return s.Store.Del([]byte(k))
}
