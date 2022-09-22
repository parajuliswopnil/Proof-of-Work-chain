package main

import (
	"crypto/sha256"
	"encoding/hex"
	"math"
	"strconv"
)

func proofOfWork(previous_proof int, data string, index int) int {
	new_proof := 1
	check_proof := false
	for !check_proof {
		work := calculation(new_proof, previous_proof, index, data)
		if work[:6] == "000000" {
			check_proof = true
		} else {
			new_proof = new_proof + 1
		}
	}
	return new_proof
}

func calculation(new_proof int, previous_proof int, index int, data string) string {
	formula := strconv.FormatFloat(math.Pow(float64(new_proof), 3)-math.Pow(float64(previous_proof), 3), 'f', 0, 64) + data
	hash := sha256.New()
	hash.Write([]byte(formula))
	calculationString := hex.EncodeToString(hash.Sum(nil))
	return calculationString
}
