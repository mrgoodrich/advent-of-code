package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	"code"
)

func readInput(r io.Reader) []string {
	lines := []string{}

	reader := bufio.NewReader(r)
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			break
		}

		lines = append(lines, strings.TrimSpace(line))
	}

	return lines
}

func main() {
	outSz := 0
	for _, line := range readInput(os.Stdin) {
		decodedLen, err := code.DecodeLen(line)
		if err != nil {
			log.Fatalf("failed to decode %v: %v", line, err)
		}

		outSz += decodedLen
	}

	fmt.Printf("output size %v\n", outSz)
}
