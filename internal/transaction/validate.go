package transaction

import (
	"github.com/erdongli/pbchain/internal/script"
	pb "github.com/erdongli/pbchain/proto"
)

type Validator struct {
	storage *Storage
}

func NewValidator() *Validator {
	return &Validator{
		storage: NewStorage(),
	}
}

func (v *Validator) Validate(tx *pb.Transaction) bool {
	for i, txIn := range tx.TxIns {
		prevTx, ok := v.storage.Get(txIn.PrevOutput.Txid)
		if !ok {
			return false
		}
		if int(txIn.PrevOutput.Index) >= len(prevTx.TxOuts) {
			return false
		}
		txOut := prevTx.TxOuts[txIn.PrevOutput.Index]
		if !script.ValidateTxIn(tx, txOut, i) {
			return false
		}
	}
	return true
}
