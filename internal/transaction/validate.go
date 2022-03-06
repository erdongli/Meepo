package transaction

import (
	"github.com/erdongli/pbchain/internal/script"
	pb "github.com/erdongli/pbchain/proto"
)

type Validator struct {
	utxos *UTXOStorage
}

func NewValidator(utxos *UTXOStorage) *Validator {
	return &Validator{
		utxos: utxos,
	}
}

func (v *Validator) Validate(tx *pb.Transaction) (uint64, bool) {
	fee := uint64(0)
	spent := map[*pb.TxOut]bool{}
	for i, txIn := range tx.TxIns {
		id := txIn.PrevOutput.Txid
		if len(id) != 32 {
			return 0, false
		}
		txOut, ok := v.utxos.Get(*(*[32]byte)(id), txIn.PrevOutput.Index)
		if !ok {
			return 0, false
		}

		if spent[txOut] {
			return 0, false
		}
		spent[txOut] = true

		if !script.ValidateTxIn(tx, txOut, i) {
			return 0, false
		}

		fee += txOut.Amount
	}

	for _, txOut := range tx.TxOuts {
		fee -= txOut.Amount
	}

	if fee <= 0 {
		return 0, false
	}

	return fee, true
}
