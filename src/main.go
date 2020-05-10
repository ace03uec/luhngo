package main

// Luhn mod 64 calculation

import (
	"log"
	"time"

	luhn "github.com/ace03uec/luhngo/src/luhn"
)

func main() {
	if !luhn.CodeInitialised {
		time.Sleep(1 * time.Second)
	}

	start := time.Now()

	for i := 0; i < 10; i++ {
		gen := luhn.GenerateString(15)
		log.Println("the string gen ", gen, " needs a checkbit ", luhn.GenerateCheckCodePoint(gen))
	}

	elapsed := time.Since(start)
	log.Printf("This took %s", elapsed)
}
