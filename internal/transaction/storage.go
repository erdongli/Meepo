package transaction

import (
	"sync"

	pb "github.com/erdongli/pbchain/proto"
)

// Storage is a thread safe in-memory storage of all transactions in a block chain
type Storage struct {
	mutex sync.RWMutex
	txs   map[[32]byte]*pb.Transaction
}

func NewStorage() *Storage {
	return &Storage{
		mutex: sync.RWMutex{},
		txs:   map[[32]byte]*pb.Transaction{},
	}
}

func (s *Storage) Get(id []byte) (*pb.Transaction, bool) {
	if len(id) != 32 {
		return nil, false
	}

	s.mutex.RLock()
	defer s.mutex.Unlock()

	tx, ok := s.txs[*(*[32]byte)(id)]
	return tx, ok
}

// Return true if the storage doesn't already has a transaction with the same id
func (s *Storage) Add(id []byte, tx *pb.Transaction) bool {
	if len(id) != 32 {
		return false
	}

	s.mutex.Lock()
	defer s.mutex.Unlock()

	txid := *(*[32]byte)(id)
	if _, ok := s.txs[txid]; ok {
		return false
	}

	s.txs[txid] = tx
	return true
}
