package main

import (
	"bufio"
	"fmt"
	"math"
	"math/rand"
	"os"
	//"strings"
	"time"
)

func main() {
  //initialize the characterset that we are using
	const characters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	//to achieve 64 bits of entropy, we calculate the entropy per symbol of our set
	character_entropy := math.Log2(float64(len(characters)))
	min_chars := math.Floor((64/character_entropy) + 0.5)

	//initialize random number generator
	source := rand.NewSource(time.Now().UnixNano())
	gen := rand.New(source)

	fmt.Printf("Entropy of Characterset: %.2f\n", character_entropy)
	fmt.Printf("Desired Password Length? (minimum %v):\n", min_chars)
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		var length int
		fmt.Sscan(scanner.Text(), &length)
		fmt.Printf("Entered Length: %v\n", length)
		if length >= 12 {
			password := make([]byte, length)
			for i := 0; i < length; i++ {
				password[i] = characters[gen.Intn(len(characters))]
			}
		}
	}

//		fmt.Println("Random Keys:")
//		key := gen.Intn(26)
}
