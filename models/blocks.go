package models

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"log"
)

// const Difficulty = 1
const TargetBits = 3

// my machine limit 3
type Block struct {
	ID           uint   `gorm:"primaryKey"`
	Index        int    `json:"index"`
	Timestamp    string `json:"timestamp"`
	Data         string `json:"data"`
	PreviousHash string `json:"previous_hash"`
	Hash         string `json:"hash"`
	Nonce        int    `json:"nonce"`
}

func CalculateHash(block Block) string {
	record := string(block.Index) + block.Timestamp + block.Data + block.PreviousHash
	hash := sha256.Sum256([]byte(record))
	return hex.EncodeToString(hash[:])
}

func ProofOfWork(block *Block) {
	target := fmt.Sprintf("%0*x", TargetBits, "") // Generate target string (e.g., "000000000000000000")
	log.Printf("Mining block with target: %s\n", target)

	for {
		block.Hash = CalculateHash(*block)
		if block.Hash[:TargetBits/4] == target[:TargetBits/4] { // Check first few bytes
			break
		}
		block.Nonce++
	}
}
