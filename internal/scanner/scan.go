package scanner

import (
	"bufio"
	"io"
	"strconv"
	"strings"
)

// ScanCSVInt returns an array of ints, split by new line
func ScanCSVInt(handle io.Reader) ([][]int, error) {
	scanner := bufio.NewScanner(handle)
	scanner.Split(bufio.ScanLines)
	var lines [][]int

	for scanner.Scan() {
		values, err := toCSVInt(scanner.Text())
		if err != nil {
			return nil, err
		}
		lines = append(lines, values)
	}
	return lines, nil
}

func toCSVInt(s string) ([]int, error) {
	var values []int
	strCSV := strings.Split(s, ",")
	for _, sv := range strCSV {
		v, err := strconv.Atoi(sv)
		if err != nil {
			return nil, err
		}
		values = append(values, v)
	}
	return values, nil
}

// ScanIntLines returns an array of ints, split by new line
func ScanIntLines(handle io.Reader) ([]int, error) {
	scanner := bufio.NewScanner(handle)
	scanner.Split(bufio.ScanLines)
	var lines []int

	for scanner.Scan() {
		v, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return nil, err
		}
		lines = append(lines, v)
	}
	return lines, nil
}

// ScanLines returns an array of strings, split by new lines
func ScanLines(handle io.Reader) []string {
	scanner := bufio.NewScanner(handle)
	scanner.Split(bufio.ScanLines)
	var lines []string

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines
}

// ScanUntilEmptyLine returns an array of strings, split by empty lines in the reader.
// When continuing from one line to the next, newlineDelimiter will be put between the values.
func ScanUntilEmptyLine(handle io.Reader, newlineDelimiter string) []string {
	scanner := bufio.NewScanner(handle)
	scanner.Split(bufio.ScanLines)
	var lines []string
	var line string
	for scanner.Scan() {
		t := scanner.Text()
		if len(t) == 0 {
			lines = append(lines, line)
			line = ""
			continue
		}

		if len(line) == 0 {
			line = t
		} else {
			line += newlineDelimiter + t
		}
	}

	if len(line) != 0 {
		lines = append(lines, line)
	}

	return lines
}
