package luhn

import (
	"log"
	"math/rand"
	"strings"
	"time"
)

// GenerateString Just generate 10 chars randomly
// mod62 to get the
func GenerateString(l int) string {
	// log.Println("Generate a random string")
	seed := time.Now().UnixNano()
	log.Println("Seed used", seed)

	r := rand.New(rand.NewSource(seed))

	var str strings.Builder

	for i := 0; i < l; i++ {
		str.WriteString(InverseCodeMap[r.Intn(CodeLen)])
	}

	return str.String()
}
