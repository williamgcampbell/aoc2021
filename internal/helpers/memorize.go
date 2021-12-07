package helpers

// FuncIntInt is a representation of a function with input and output of int
type FuncIntInt func(int) int

// MemorizedIntInt caches the results of a FuncIntInt
func MemorizedIntInt(fn FuncIntInt) FuncIntInt {
	cache := make(map[int]int)

	return func(input int) int {
		if val, found := cache[input]; found {
			return val
		}

		result := fn(input)
		cache[input] = result
		return result
	}
}
