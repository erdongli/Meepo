package transaction

import (
	"github.com/erdongli/pbchain/internal/script"
	pb "github.com/erdongli/pbchain/proto"
)

func Validate(tx *pb.Transaction) bool {
	for i, txIn := range tx.TxIns {
		prevTx, err := Get(txIn.PrevOutput.Txid)
		if err != nil {
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
