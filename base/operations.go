package base

import (
	"crypto/sha256"
	"fmt"
	"time"
)

var Blockchain []Block

func CreateTransaction() *Block {
	if len(Blockchain) == 0 {
		b := Genesis()
		Blockchain = append(Blockchain, *b)
		return b
	}
	return &Block{}
}

func Genesis() *Block {
	Gtime, Gtransaction, GprevHash := Generator()

	return &Block{
		timeStamp:    Gtime,
		transactions: Gtransaction,
		prevhash:     GprevHash,
		hash:         Hash(Gtime, Gtransaction, GprevHash),
	}
}

func Hash(time time.Time, transactions []string, prevhash []byte) []byte {
	payload := append(prevhash, time.String()...)
	for transaction := range transactions {
		payload = append(payload, string(rune(transaction))...)
	}
	hash := sha256.Sum256(payload)
	return hash[:]
}

func Generator() (time.Time, []string, []byte) {
	gtime := time.Now()
	gtransaction := []string{"test"}
	if len(Blockchain) == 0 {
		return gtime, gtransaction, []byte{}
	}
	gprevhash := Blockchain[len(Blockchain)-1].prevhash
	return gtime, gtransaction, gprevhash
}

func ShowBlockchainData() {
	for i, block := range Blockchain {
		fmt.Println("----------------------------------")
		fmt.Printf("Data Block N° %d", i)
		fmt.Println("")
		fmt.Println("Time: ", block.timeStamp)
		fmt.Println("Transactions: ", block.transactions)
		fmt.Println("Prev. Hash: ", block.prevhash)
		fmt.Printf("Hash: %x", block.hash)
		fmt.Println("")
		fmt.Println("----------------------------------")
	}
}
