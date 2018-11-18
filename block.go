package GoCoin

import (
	"GoCoin/transaction"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"strings"
	"time"
)

type Block struct {
	PrevHash string
	Transactions []transaction.Transaction
	Nonce int
	Timestamp int64
	Hash string
}

var(
	difficulty = 2
)

func (block *Block) CalcHash() string {
	md := sha256.New()
	data := fmt.Sprintf("%s%s%d%d", block.PrevHash, EncodeTransactionArray(block.Transactions), block.Timestamp, block.Nonce)
	md.Write([]byte(data))
	return hex.EncodeToString(md.Sum(nil))
}

func (block *Block) Mine(chain Chain) {
	fmt.Println(len(block.Transactions))
	for !block.IsMined() {
		block.Nonce++
		block.Hash = block.CalcHash()
	}

	chain.UpdateUTXO(*block)
}

func (block *Block) IsMined() bool {
	return strings.HasPrefix(block.Hash, strings.Repeat("0", difficulty))
}

func (block *Block) AddTransaction(tr transaction.Transaction) {
	if tr.ValidateSignature() {
		block.Transactions = append(block.Transactions, tr)
		block.Hash = block.CalcHash()
	}
}

func NewBlock(prevHash string) Block {
	block := new(Block)
	block.PrevHash = prevHash
	block.Transactions = make([]transaction.Transaction, 0)
	block.Nonce = 0
	block.Timestamp = time.Now().Unix() * 1000
	block.Hash = block.CalcHash()

	return *block
}

func EncodeTransactionArray(array []transaction.Transaction) string {
	str := ""
	for _, tr := range array {
		str += tr.Hash
	}

	return base64.StdEncoding.EncodeToString([]byte(str))
}
