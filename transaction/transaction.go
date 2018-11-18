package transaction

import (
	"GoCoin/utils"
	"crypto/rsa"
	"fmt"
)

type Transaction struct {
	Sender rsa.PublicKey
	Recipient rsa.PublicKey
	Amount int64
	Signature string
	Hash string
	Inputs []TransactionOutput
	Outputs []TransactionOutput
}

func NewTransaction(sender rsa.PublicKey, recipient rsa.PublicKey, amount int64) *Transaction {
	transaction := new(Transaction)
	transaction.Sender = sender
	transaction.Recipient = recipient
	transaction.Amount = amount
	transaction.Inputs = []TransactionOutput{}
	transaction.Outputs = []TransactionOutput{}

	raw := fmt.Sprintf("%s%s%d%s", utils.GetBase64Key(&sender), utils.GetBase64Key(&recipient), amount, utils.RandomString(16))
	transaction.Hash = utils.HashToString(raw)

	return transaction
}

func (transaction *Transaction) Sign(privKey rsa.PrivateKey) string {
	raw := fmt.Sprintf("%s%s%d", utils.GetBase64Key(&transaction.Sender), utils.GetBase64Key(&transaction.Recipient), transaction.Amount)
	fmt.Println(raw)
	return utils.Sign(&privKey, raw)
}

func (transaction *Transaction) ValidateSignature() bool {
	raw := fmt.Sprintf("%s%s%d", utils.GetBase64Key(&transaction.Sender), utils.GetBase64Key(&transaction.Recipient), transaction.Amount)
	fmt.Println(raw)
	return utils.VerifySignature(&transaction.Sender, raw, []byte(transaction.Signature))
}
