package _6

import (
	_ "embed"
	"strconv"
	"strings"
	"sync"

	"github.com/williamgcampbell/aoc2021/internal/scanner"
)

//go:embed input.txt
var input string

var (
	TotalDays           = 80
	AgeAfterBirth       = 8
	AgeAfterGivingBirth = 6
)

func (d *Day) SolvePart1() string {
	r := strings.NewReader(input)
	lines, _ := scanner.ScanCSVInt(r)
	return strconv.FormatInt(countAll(lines[0], TotalDays), 10)
}

func countAll(ages []int, afterDays int) int64 {
	c := int64(0)
	for _, age := range ages {
		c += count(age, afterDays)
	}
	return c
}

var cache = &birthCountCache{
	m: make(map[ageDay]int64),
}

func count(age int, afterDays int) int64 {
	if afterDays == 0 {
		return 1
	}

	ad := ageDay{
		age:           age,
		remainingDays: afterDays,
	}
	//if c, ok := AgeMap[ad]; ok {
	if c, ok := cache.Get(ad); ok {
		return c
	}

	var c int64
	if age == 0 {
		c = count(AgeAfterGivingBirth, afterDays-1) + count(AgeAfterBirth, afterDays-1)
	} else {
		c = count(age-1, afterDays-1)
	}

	cache.Add(ad, c)
	//AgeMap[ad] = c
	return c
}

// ageDay stores the age of a fish and the remaining days in the cycle.
// This struct is to be used as a key to track the number of fish a fish of this age will give birth
// to with the remaining days.
type ageDay struct {
	age           int
	remainingDays int
}

// birthCountCache stores the count of fish a particular fish will give birth to, given an age and days remaining.
type birthCountCache struct {
	mu sync.Mutex
	m  map[ageDay]int64
}

// Get retrieves the count without modifying it
func (b *birthCountCache) Get(key ageDay) (int64, bool) {
	b.mu.Lock()
	defer b.mu.Unlock()

	if c, ok := b.m[key]; ok {
		return c, true
	}

	return 0, false
}

// Add stores the key/vale in the cache
func (b *birthCountCache) Add(key ageDay, value int64) int64 {
	b.mu.Lock()
	defer b.mu.Unlock()

	b.m[key] = value
	return value
}
