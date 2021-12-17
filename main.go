package main

import "github.com/JairDavid/blockchain-practice/base"

func main() {
	for i := 0; i < 10; i++ {
		base.CreateTransaction()
	}
	base.ShowBlockchainData()
}
