package transaction

import (
	"sync"

	pb "github.com/erdongli/pbchain/proto"
)

type UTXOStorage struct {
	mutex sync.RWMutex
	utxos map[[32]byte]map[uint32]*pb.TxOut
}

func NewUTXOStorage() *UTXOStorage {
	return &UTXOStorage{
		mutex: sync.RWMutex{},
		utxos: map[[32]byte]map[uint32]*pb.TxOut{},
	}
}

func (s *UTXOStorage) Get(id [32]byte, idx uint32) (*pb.TxOut, bool) {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	txOut, ok := s.utxos[id][idx]
	return txOut, ok
}

func (s *UTXOStorage) Update(block *pb.Block) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	for _, tx := range block.Txs {
		// Remove spent outputs
		for _, txIn := range tx.TxIns {
			prevOut := txIn.PrevOutput
			if prevOut == nil {
				continue
			}

			if len(prevOut.Txid) != 32 {
				continue
			}
			delete(s.utxos[*(*[32]byte)(prevOut.Txid)], prevOut.Index)
		}

		// Add unspent outputs
		id, err := Id(tx)
		if err != nil {
			continue
		}
		for i, txOut := range tx.TxOuts {
			if _, ok := s.utxos[id]; !ok {
				s.utxos[id] = map[uint32]*pb.TxOut{}
			}
			s.utxos[id][uint32(i)] = txOut
		}
	}
}
