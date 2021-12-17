package base

import "time"

type Block struct {
	timeStamp    time.Time
	transactions []string
	prevhash     []byte
	hash         []byte
}
