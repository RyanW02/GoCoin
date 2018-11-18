package main

import (
	"GoCoin"
	"GoCoin/transaction"
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	chain := GoCoin.NewBlockChain()
	wallet1 := GoCoin.NewWallet(chain)
	wallet2 := GoCoin.NewWallet(chain)

	tr := transaction.NewTransaction(wallet1.PublicKey, wallet1.PublicKey, 100)
	tr.Outputs = append(tr.Outputs, *transaction.NewTransactionOutput(wallet1.PublicKey, 100, tr.Hash))

	genBlock := chain.AddGenesisBlock()
	genBlock.AddTransaction(*tr)
	fmt.Println(len(genBlock.Transactions))
	genBlock.Mine(chain)

	fmt.Printf("Wallet 1 balance: %d\n", wallet1.GetBalance())
	fmt.Printf("Wallet 2 balance: %d\n", wallet2.GetBalance())

	tr2 := transaction.NewTransaction(wallet1.PublicKey, wallet2.PublicKey, 33)
	block2 := chain.AddBlock()
	block2.AddTransaction(*tr2)
	block2.Mine(chain)

	fmt.Printf("Wallet 1 balance: %d\n", wallet1.GetBalance())
	fmt.Printf("Wallet 2 balance: %d\n", wallet2.GetBalance())
}
