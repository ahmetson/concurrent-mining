package main

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"
	"encoding/hex"
	"fmt"
	"log"
	"os"
	"sync"
	"time"
)

type Block struct {
	Previous   [32]byte
	Number     int      // block height
	MerkleRoot [32]byte // merkle root of all transactions
	Nonce      int
}

func EncodeToBytes(p interface{}) []byte {

	buf := bytes.Buffer{}
	enc := gob.NewEncoder(&buf)
	err := enc.Encode(p)
	if err != nil {
		log.Fatal(err)
	}
	return buf.Bytes()
}

func validNonce(routineId int, difficulty []byte, start int, end int, block Block, startTime *time.Time) int {
	for ; start <= end; start++ {
		block.Nonce = start

		hash := sha256.Sum256(EncodeToBytes(block))

		if bytes.Index(hash[:], difficulty) == 0 {
			duration := time.Since(*startTime)

			hashStr := hex.EncodeToString(hash[:])
			fmt.Println("Goroutine", routineId, "Nonce", start+1, "Hash:", hashStr, "Found in seconds", duration)

			os.Exit(0)
		}
	}

	return 0
}

func main() {
	var wg sync.WaitGroup

	var empty [32]byte

	var genericBlock = Block{
		Previous:   empty,
		Number:     1,
		MerkleRoot: empty,
		Nonce:      0,
	}

	fmt.Println("Trying to find a hash of Genesis block...")
	startTime := time.Now()

	difficulty := []byte{0, 0}

	// now divide to goroutines

	amount := 60000 // with 100 goroutines
	// amount := 6000000 // with 1 goroutine

	// for i := 0; i < 1; i++ {
	for i := 0; i < 100; i++ {
		start := (i) * amount
		end := start + amount

		wg.Add(1)

		go func(goroutineId int) {
			validNonce(goroutineId, difficulty, start, end, genericBlock, &startTime)
			wg.Done()
		}(i + 1)
	}

	wg.Wait()
}
