package transaction

import (
	"GoCoin/utils"
	"crypto/rsa"
	"fmt"
)

type TransactionOutput struct {
	Recipient rsa.PublicKey
	Amount int64
	TransactionHash string
	Hash string
}

func NewTransactionOutput(recipient rsa.PublicKey, amount int64, transactionHash string) *TransactionOutput {
	transactionOutput := new(TransactionOutput)
	transactionOutput.Recipient = recipient
	transactionOutput.Amount = amount
	transactionOutput.TransactionHash = transactionHash

	raw := fmt.Sprintf("%s%d%s", utils.GetBase64Key(&recipient), amount, transactionHash)
	transactionOutput.Hash = utils.HashToString(raw)

	return transactionOutput
}

func (output *TransactionOutput) IsOwnedBy(key rsa.PublicKey) bool {
	return output.Recipient == key
}
