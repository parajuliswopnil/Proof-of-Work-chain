package main

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"time"
)

type Chain struct {
	chain []Blocks
}

func (c *Chain) mine_block(data string) {
	var block Blocks
	var lastBlock map[string]interface{}
	lastBlock = c.get_last_block()
	index := lastBlock["index"].(int) + 1
	block.data = data
	block.index = len(c.chain) + 1
	block.proof = proofOfWork(lastBlock["proof"].(int), data, index)
	block.timeStamp = time.Now().Nanosecond()
	block.previous_hash = c.get_previous_hash(lastBlock)

	c.create_block(block)
}

func (c *Chain) create_block(block Blocks) {
	c.chain = append(c.chain, block)
}

func (c *Chain) get_last_block() map[string]interface{} {
	var lastBlock Blocks
	lastBlock = c.chain[len(c.chain)-1]
	block := make(map[string]interface{})
	block["data"] = lastBlock.data
	block["index"] = lastBlock.index
	block["proof"] = lastBlock.proof
	block["time_stamp"] = lastBlock.timeStamp
	block["previous_hash"] = lastBlock.previous_hash

	return block
}

func (c *Chain) get_previous_hash(previous_block map[string]interface{}) string {
	fmt.Println("**********Last block **************")
	jsonString, _ := json.Marshal(previous_block)
	hash := sha256.New()
	hash.Write([]byte(string(jsonString)))

	hexString := hex.EncodeToString(hash.Sum(nil))
	fmt.Println("hex_string: ", hexString)
	return hexString
}

func (c *Chain) get_all_blocks() []map[string]interface{} {
	var lastBlock Blocks
	var blockList []map[string]interface{}
	for i := 0; i < len(c.chain); i++ {
		lastBlock = c.chain[i]
		block := make(map[string]interface{})
		block["data"] = lastBlock.data
		block["index"] = lastBlock.index
		block["proof"] = lastBlock.proof
		block["time_stamp"] = lastBlock.timeStamp
		block["previous_hash"] = lastBlock.previous_hash
		blockList = append(blockList, block)
	}
	return blockList
}

func (c *Chain) checkChainValidity() bool {
	var lastBlock Blocks
	var hash string
	valid := false
	for i := 0; i < len(c.chain)-1; i++ {
		lastBlock = c.chain[i]
		block := make(map[string]interface{})
		block["data"] = lastBlock.data
		block["index"] = lastBlock.index
		block["proof"] = lastBlock.proof
		block["time_stamp"] = lastBlock.timeStamp
		block["previous_hash"] = lastBlock.previous_hash
		hash = c.get_previous_hash(block)
		if hash == c.chain[i+1].previous_hash {
			valid = true
		}
	}
	return valid
}
