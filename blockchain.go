package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"strconv"
	"time"
)

type Block struct {
	Index     int
	Timestamp string
	Data      string
	Hash      string
	PrevHash  string
}

var Blockchain []Block

func calculateHash(block Block) string {
	record := strconv.Itoa(block.Index) + block.Timestamp + block.Data + block.PrevHash
	h := sha256.New()
	h.Write([]byte(record))
	hashed := h.Sum(nil)
	return hex.EncodeToString(hashed)
}

func generateBlock(oldBlock Block, data string) Block {
	var newBlock Block

	newBlock.Index = oldBlock.Index + 1
	newBlock.Timestamp = time.Now().String()
	newBlock.Data = data
	newBlock.PrevHash = oldBlock.Hash
	newBlock.Hash = calculateHash(newBlock)

	return newBlock
}

func isBlockValid(newBlock, oldBlock Block) bool {
	if oldBlock.Index+1 != newBlock.Index {
		return false
	}

	if oldBlock.Hash != newBlock.PrevHash {
		return false
	}

	if calculateHash(newBlock) != newBlock.Hash {
		return false
	}

	return true
}

func main() {
	genesisBlock := Block{0, time.Now().String(), "Genesis Block", "", ""}
	genesisBlock.Hash = calculateHash(genesisBlock)
	Blockchain = append(Blockchain, genesisBlock)

	fmt.Println(Blockchain[0])

	block1 := generateBlock(Blockchain[0], "Block 1 Data")
	if isBlockValid(block1, Blockchain[0]) {
		Blockchain = append(Blockchain, block1)
		fmt.Println(Blockchain[1])
	}

	block2 := generateBlock(Blockchain[1], "Block 2 Data")
	if isBlockValid(block2, Blockchain[1]) {
		Blockchain = append(Blockchain, block2)
		fmt.Println(Blockchain[2])
	}
}
