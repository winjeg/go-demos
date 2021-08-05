package store

import (
	badger "github.com/dgraph-io/badger/v3"
	"github.com/dgraph-io/badger/v3/options"
	"github.com/sirupsen/logrus"

	"fmt"
	"log"
	"time"
)

var (
	badgerOption = func(path string, level logrus.Level) badger.Options {
		return badger.Options{
			Dir:                           path,
			ValueDir:                      path,
			MemTableSize:                  4 << 20,
			BaseTableSize:                 2 << 20,
			BaseLevelSize:                 2 << 20,
			TableSizeMultiplier:           2,
			LevelSizeMultiplier:           10,
			MaxLevels:                     7,
			NumGoroutines:                 8,
			MetricsEnabled:                true,
			NumCompactors:                 2, // Run at least 2 compactors. Zero-th compactor prioritizes L0.
			NumLevelZeroTables:            1,
			NumLevelZeroTablesStall:       5,
			NumMemtables:                  1,
			BloomFalsePositive:            0.01,
			BlockSize:                     4 * 1024,
			SyncWrites:                    false,
			NumVersionsToKeep:             1,
			CompactL0OnClose:              false,
			VerifyValueChecksum:           false,
			Compression:                   options.Snappy,
			BlockCacheSize:                16 << 20,
			IndexCacheSize:                0,
			ZSTDCompressionLevel:          1,
			ValueLogFileSize:              1<<24 - 1,
			ValueLogMaxEntries:            1000,
			VLogPercentile:                0.0,
			ValueThreshold:                1 << 18,
			Logger:                        defaultLogger(),
			EncryptionKey:                 []byte{},
			EncryptionKeyRotationDuration: 10 * 24 * time.Hour, // Default 10 days.
			DetectConflicts:               true,
			NamespaceOffset:               -1,
		}
	}
)

func defaultLogger() *logrus.Logger {
	logger := logrus.New()
	logger.SetLevel(logrus.WarnLevel)
	logger.SetFormatter(&logrus.TextFormatter{
		ForceColors:   true,
		FullTimestamp: true,
	})
	return logger
}

func NewBadgerStore(dir string, logLevel logrus.Level) *badgerStore {
	db, err := badger.Open(badgerOption(dir, logLevel))
	if err != nil {
		log.Fatal(err)
		return nil
	}
	return &badgerStore{db}
}

type badgerStore struct {
	DB *badger.DB
}

func (bs *badgerStore) Set(k, v []byte) error {
	err := bs.DB.Update(func(txn *badger.Txn) error {
		err := txn.Set(k, v)
		return err
	})
	return err
}

func (bs *badgerStore) Get(k []byte) ([]byte, error) {
	defer func() {
		if r := recover(); r != nil {
			return
		}
	}()
	var result []byte
	err := bs.DB.View(func(txn *badger.Txn) error {
		item, _ := txn.Get(k)
		if nil == item {
			return nil
		}
		err := item.Value(func(val []byte) error {
			result = append([]byte{}, val...)
			return nil
		})
		return err
	})
	return result, err
}

func (bs *badgerStore) SetEx(k, v []byte, exp time.Duration) error {
	err := bs.DB.Update(func(txn *badger.Txn) error {
		e := badger.NewEntry(k, v).WithTTL(exp)
		err := txn.SetEntry(e)
		return err
	})
	return err
}

func (bs *badgerStore) Seq(k []byte, max uint64) func() (uint64, error) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Wrong key: ", r)
			return
		}
	}()
	seq, err := bs.DB.GetSequence(k, max)
	if err != nil {
		return nil
	}
	defer func(seq *badger.Sequence) {
		err := seq.Release()
		if err != nil {
			log.Println(err.Error())
		}
	}(seq)
	return func() (uint64, error) {
		return seq.Next()
	}
}

func (bs *badgerStore) Del(k []byte) error {
	err := bs.DB.Update(func(txn *badger.Txn) error {
		err := txn.Delete(k)
		return err
	})
	return err
}

func (bs *badgerStore) Shutdown() {
	if bs.DB == nil {
		return
	}
	err := bs.DB.Close()
	if err != nil {
		log.Println(err.Error())
		return
	}
}
