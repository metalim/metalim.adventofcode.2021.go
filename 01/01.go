package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/pkg/errors"
)

func catch(err error, args ...interface{}) {
	if err == nil {
		return
	}
	if len(args) > 0 {
		panic(errors.Wrap(err, fmt.Sprintf(args[0].(string), args[1:]...)))
	}
	panic(err)
}

func Input() string {
	file := "input.txt"
	if len(os.Args) > 1 {
		file = os.Args[1]
	}
	input, err := os.ReadFile(file)
	catch(err)
	return strings.TrimSpace(string(input))
}

func Lines(input string) []string {
	return strings.Split(input, "\n")
}

func Ints(lines []string) []int {
	var err error
	ints := make([]int, len(lines))
	for i, line := range lines {
		ints[i], err = strconv.Atoi(line)
		catch(err, "line %d", i)
	}
	return ints
}

func main() {
	ints := Ints(Lines(Input()))

	// Part 1
	var increased int
	for i, v := range ints[1:] {
		if v > ints[i] {
			increased++
		}
	}
	fmt.Println(increased)

	// Part 2
	increased = 0
	sum := ints[0] + ints[1] + ints[2]
	for i, v := range ints[3:] {
		newSum := sum + v - ints[i]
		if newSum > sum {
			increased++
		}
		sum = newSum
	}
	fmt.Println(increased)
}
