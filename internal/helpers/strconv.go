package helpers

import "strconv"

func MustAtoI(a string) int {
	result, err := strconv.Atoi(a)
	if err != nil {
		panic(err)
	}

	return result
}

func MustParseInt(s string, base, bitSize int) int64 {
	result, err := strconv.ParseInt(s, base, bitSize)
	if err != nil {
		panic(err)
	}

	return result
}

func MustParseUint(s string, base, bitSize int) uint64 {
	v, err := strconv.ParseUint(s, base, bitSize)
	if err != nil {
		panic(err)
	}
	return v
}
