package main

import (
	"log"
	"math/big"

	"github.com/erdongli/meepo/internal/chain"
)

func main() {
	cnt, incr := big.NewInt(0), big.NewInt(1)
	bc := chain.NewBlockChain()
	for {
		if err := bc.Append([][]byte{cnt.Bytes()}); err != nil {
			log.Fatalf("failed to append to block chain: %v", err)
		}
		cnt.Add(cnt, incr)
	}
}
