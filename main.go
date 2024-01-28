package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strconv"
)

func OutputQuad(input string, x, y int) string {
	cmd := exec.Command(input, strconv.Itoa(x), strconv.Itoa(y))

	// Capture the standard output of the command
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return fmt.Sprintf("Error creating StdoutPipe: %v", err)
	}

	// Start the command
	if err := cmd.Start(); err != nil {
		return fmt.Sprintf("Error starting command: %v", err)
	}

	// Read and store all lines of the output
	var outputLines string
	scanner := bufio.NewScanner(stdout)
	for scanner.Scan() {
		line := scanner.Text()
		outputLines += line
	}

	return outputLines
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	y := 0
	var x int
	var principalString string
	for scanner.Scan() {
		line := scanner.Text()
		x = len(line)
		principalString += line
		y++
	}
	quadA := OutputQuad("./quadA", x, y)
	quadB := OutputQuad("./quadB", x, y)
	quadC := OutputQuad("./quadC", x, y)
	quadD := OutputQuad("./quadD", x, y)
	quadE := OutputQuad("./quadE", x, y)
	stringsToCheck := []string{quadA, quadB, quadC, quadD, quadE}
	s := []string{"quadA", "quadB", "quadC", "quadD", "quadE"}

	matchingStrings := []string{}
	for i, str := range stringsToCheck {
		if str == principalString {
			matchingStrings = append(matchingStrings, s[i])
		}
	}

	for i, v := range matchingStrings {
		fmt.Printf("[%s] [%d] [%d]", v, x, y)
		if i != len(matchingStrings)-1 {
			fmt.Print(" || ")
		}
	}
	fmt.Println()
}
