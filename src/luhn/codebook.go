package luhn

// Define Code Point
// Our code points are going to be base62
// exclude + and /
// 0 - 9 10 - 35 36 - 61
// 0 - 0 a - z A - Z

// CodeMap containing  char representation to int
var CodeMap map[string]int

// InverseCodeMap containing  int representation to char
var InverseCodeMap map[int]string

// CodeInitialised setup codemap being ready
var CodeInitialised bool = false

// CodeLen indicates the codebook length
const CodeLen = 62

// For Codelen 10
// const CodeLen = 10

func init() {
	CodeMap = map[string]int{}
	InverseCodeMap = map[int]string{}

	for i := 0; i < CodeLen; i++ {
		CodeMap[getCharForIndex62(int(i))] = int(i)
		InverseCodeMap[int(i)] = getCharForIndex62(int(i))
	}

	// log.Println(CodeMap)
	// log.Println(InverseCodeMap)

	CodeInitialised = true
}

func getCharForIndex62(i int) string {
	if i >= 0 && i < 10 {
		return string(i + int('0'))
	} else if i < 36 {
		return string(i - 10 + int('a'))
	} else if i < CodeLen {
		return string(i - 36 + int('A'))
	} else {
		panic("Shit broke")
	}
}

func getCharForIndexCodeLen10(i int) string {
	if i >= 0 && i < 10 {
		return string(i + int('0'))
	}

	panic("Shit broke")
}
func getMinMaxCharZ1() (int, int) {
	return int('0'), int('9')
}

func getMinMaxCharZ2() (int, int) {
	return int('a'), int('z')
}

func getMinMaxCharZ3() (int, int) {
	return int('A'), int('Z')
}
