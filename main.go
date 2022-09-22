package main

import (
	"fmt"
	"time"
)

func main() {
	blockchain := Chain{chain: genesis_block()}

	print(blockchain.get_last_block())

	blockchain.mine_block("this is the second block")

	print(blockchain.get_last_block())

	blockchain.mine_block("this is the third block")

	print(blockchain.get_last_block())

	blockchain.mine_block("this is the fourth block")

	print(blockchain.get_all_blocks())

	print(blockchain.checkChainValidity())

}

func print(printThings interface{}) {
	fmt.Println(printThings)
}

func genesis_block() []Blocks {
	var genesisBlock Blocks
	var initiateChain []Blocks
	genesisBlock.data = ""
	genesisBlock.index = 1
	genesisBlock.proof = 1
	genesisBlock.previous_hash = ""
	genesisBlock.timeStamp = time.Now().Nanosecond()
	initiateChain = append(initiateChain, genesisBlock)
	return initiateChain
}
