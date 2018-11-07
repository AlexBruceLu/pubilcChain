package main

import (
	"fmt"
	"test/blc/blc"
	"time"
)

func main() {
	blockChain := blc.NewBlockChain()
	blockChain.AddBlock("to xiaohong 20 BTC")
	blockChain.AddBlock("to xiaohong 20 BTC")
	blockChain.AddBlock("to xiaohong 20 BTC")

	for _, block := range blockChain.Blocks {
		fmt.Printf("Data: %s\n", string(block.Data))
		fmt.Printf("PrevHash: %x\n", block.PrevHash)
		fmt.Printf("Hash: %x\n", block.Hash)
		fmt.Printf("nonce: %x\n", block.Nonce)
		fmt.Printf("DTimeStamp: %v\n", time.Unix(block.TimeStamp, 0).Format("2006-01-02 15:04:05"))
		fmt.Printf("\n\n")
	}
}
