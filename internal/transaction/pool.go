package transaction

import (
	"sync"

	pb "github.com/erdongli/pbchain/proto"
)

type Pool struct {
	mutex sync.Mutex
	txs   []*pb.Transaction
}

func NewPool(v *Validator) *Pool {
	return &Pool{
		mutex: sync.Mutex{},
		txs:   []*pb.Transaction{},
	}
}

func (p *Pool) CheckIn(tx *pb.Transaction) {
	p.mutex.Lock()
	defer p.mutex.Unlock()

	p.txs = append(p.txs)
}

func (p *Pool) CheckOut() []*pb.Transaction {
	txs := []*pb.Transaction{}

	p.mutex.Lock()
	defer p.mutex.Unlock()

	for i, tx := range p.txs {
		txs = append(txs, tx)
		p.txs[i] = nil // Avoid memory leak
	}
	p.txs = p.txs[:0]

	return txs
}
