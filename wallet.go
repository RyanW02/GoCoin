package GoCoin

import (
	"GoCoin/transaction"
	"GoCoin/utils"
	"crypto/rsa"
	"github.com/pkg/errors"
)

type Wallet struct {
	PublicKey rsa.PublicKey
	PrivateKey rsa.PrivateKey
	BlockChain Chain
}

func NewWallet(chain Chain) Wallet {
	priv, pub := utils.GenerateRSAPair()

	return Wallet{
		PublicKey: pub,
		PrivateKey: priv,
		BlockChain: chain,
	}
}

func (wallet *Wallet) GetBalance() int64 {
	balance := int64(0)

	for _, transaction := range wallet.GetTransactions() {
		balance += transaction.Amount
	}

	return balance
}

func (wallet *Wallet) GetTransactions() []transaction.TransactionOutput {
	var userTransactions []transaction.TransactionOutput

	for _, transaction := range wallet.BlockChain.UTXO {
		if transaction.IsOwnedBy(wallet.PublicKey) {
			userTransactions = append(userTransactions, transaction)
		}
	}

	return userTransactions
}

func (wallet *Wallet) Transfer(recipient rsa.PublicKey, amount int64) (transaction.Transaction, error) {
	var tr transaction.Transaction
	if amount > wallet.GetBalance() {
		return tr, errors.New("Not enough funds")
	}

	tr = *transaction.NewTransaction(wallet.PublicKey, recipient, amount)
	tr.Outputs = append(tr.Outputs, *transaction.NewTransactionOutput(recipient, amount, tr.Hash))

	sourceAmount := int64(0)
	for _, input := range wallet.GetTransactions() {
		sourceAmount += input.Amount
		tr.Inputs = append(tr.Inputs, input)

		if sourceAmount > amount {
			diff := sourceAmount - amount
			tr.Outputs = append(tr.Outputs, *transaction.NewTransactionOutput(wallet.PublicKey, diff, tr.Hash))
		}

		if sourceAmount == amount {
			break
		}
	}

	tr.Signature = tr.Sign(wallet.PrivateKey)
	return tr, nil
}
