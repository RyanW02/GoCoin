package GoCoin

import (
	"GoCoin/transaction"
	"fmt"
	"strconv"
)

type Chain struct {
	Blocks []Block
	UTXO map[string]transaction.TransactionOutput
}

func NewBlockChain() Chain {
	chain := new(Chain)
	chain.Blocks = make([]Block, 0)
	chain.UTXO = make(map[string]transaction.TransactionOutput)
	return *chain
}

func (chain *Chain) getLastBlock() Block {
	return chain.Blocks[len(chain.Blocks) - 1]
}

func (chain *Chain) AddGenesisBlock() Block {
	block := NewBlock("0")
	chain.Blocks = append(chain.Blocks, block)
	return block
}

func (chain *Chain) AddBlock() Block {
	block := NewBlock(chain.getLastBlock().PrevHash)
	chain.Blocks = append(chain.Blocks, block)
	return block
}

func (chain *Chain) UpdateUTXO(block Block) {
	for _, tr := range block.Transactions {
		delete(chain.UTXO, tr.Hash)
	}

	for _, tr := range block.Transactions {
		for _, output := range tr.Outputs {
			fmt.Println("yeet " + strconv.Itoa(int(output.Amount)))
			chain.UTXO[output.Hash] = output
		}
	}
}
