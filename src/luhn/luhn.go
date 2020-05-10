package luhn

import (
	"log"
)

// GenerateCheckCodePoint takes in a given string that is in the codespace
// and generate a luhn mod n checkcodepoint for the same.
// For this sameple the codebooklength is 62
func GenerateCheckCodePoint(s string) string {
	// validate space
	ValidateStringSpace(s)

	// generatesum
	len := len(s)
	double := true
	sum := 0

	for i := len - 1; i >= 0; i-- {
		currChar := s[i]

		currVal := getIntForStringCode(string(currChar))

		hereVal := currVal

		if double {
			hereVal = hereVal * 2

			if hereVal >= CodeLen {
				hereVal = hereVal + 1 - CodeLen
			}
		}
		double = !double
		sum = sum + hereVal
	}
	return getStringForIntCode(sum * (CodeLen - 1) % CodeLen)
}

func getStringForIntCode(i int) string {
	return InverseCodeMap[i]
}

func getIntForStringCode(s string) int {
	return CodeMap[s]
}

// ValidateString takes in a given string that is in the codespace and
// validates if the string is a valid as per the luhn mod n algo
func ValidateString(s string) bool {
	// validate space
	ValidateStringSpace(s)

	checkPoint := s[len(s)-1]

	expectedCheckPoint := GenerateCheckCodePoint(s[:len(s)-1])

	return expectedCheckPoint == string(checkPoint)
}

// ValidateStringSpace check if string is in our codebook space
func ValidateStringSpace(s string) bool {
	for _, c := range s {
		if _, ok := CodeMap[string(c)]; !ok {
			log.Fatalf("String %s isn't valid", s)
			return false
		}
	}

	return true
}
