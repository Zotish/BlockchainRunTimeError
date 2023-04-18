package main

import (
	"fmt"

	blockchainfunction "github.com/Zotish/Blockchainfunction"
	blockstructure "github.com/Zotish/blockStructure"
)

var Blockchain []blockstructure.Block

func main() {
	fmt.Println("STart ")
	StoreGenBlock := *blockchainfunction.CreateGenBlock()
	Blockchain = append(Blockchain, StoreGenBlock)
	fmt.Println("", Blockchain)
	StoreNewBlock := blockchainfunction.CreateNewBlock(&blockstructure.Block{}, &blockstructure.TransactionDetails{})
	Blockchain = append(Blockchain, StoreNewBlock)
	fmt.Println("", Blockchain)

}
