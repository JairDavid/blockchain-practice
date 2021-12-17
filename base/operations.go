package base

import (
	"crypto/rand"
	"crypto/sha256"
	"fmt"
	"math/big"
	"time"
)

var Blockchain []Block

func CreateTransaction() {
	Blockchain = append(Blockchain, *Genesis())
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

func Generator() (time.Time, []string, []byte) {
	mount, _ := rand.Int(rand.Reader, big.NewInt(60000))
	data := " A stranger sent " + mount.String() + " coins to another stranger"

	gtime := time.Now()
	gtransaction := []string{data}

	if len(Blockchain) == 0 {
		return gtime, gtransaction, []byte{}
	}
	gprevhash := Blockchain[len(Blockchain)-1].hash
	return gtime, gtransaction, gprevhash
}

func Hash(time time.Time, transactions []string, prevhash []byte) []byte {
	payload := append(prevhash, time.String()...)
	for transaction := range transactions {
		payload = append(payload, string(rune(transaction))...)
	}
	hash := sha256.Sum256(payload)
	return hash[:]
}

func ShowBlockchainData() {
	for i, block := range Blockchain {
		fmt.Println("----------------------------------")
		fmt.Printf("Data Block N° %d", i)
		fmt.Println("")
		fmt.Println("Time: ", block.timeStamp)
		fmt.Println("Transactions: ", block.transactions)
		fmt.Println("")
		fmt.Printf("Prev.Hash: %x", block.prevhash)
		fmt.Println("")
		fmt.Printf("Hash: %x", block.hash)
		fmt.Println("")
		fmt.Println("----------------------------------")
	}
}
