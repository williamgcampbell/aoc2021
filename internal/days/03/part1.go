// Copyright 2020-present (c) Care.com, Inc.
//
// All rights reserved.
//
// This software is the confidential and proprietary information of
// Care.com, Inc.

package _3

import (
	_ "embed"
	"strings"
)

//go:embed input.txt
var input string

// SolvePart1 solves day three, part one, of the 2021 advent of code question.
// https://adventofcode.com/2021/day/3
func SolvePart1() string {
	_ = strings.NewReader(input)
	return ""
}
