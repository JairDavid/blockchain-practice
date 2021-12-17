package base

var blockchain []Block

func CreateTransaction() *Block {
	if len(blockchain) == 0 {
		return Genesis()
	}
	return &Block{}
}

func Generator() {

}

func Genesis() *Block {
	return &Block{}
}
