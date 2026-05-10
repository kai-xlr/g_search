package main

import (
	"fmt"
	"os"
	"strings"
)

func search(query string, contents string) bool {
	query = strings.ToLower(query)
	found := false

	lines := strings.Split(contents, "\n")

	for i, line := range lines {
		if strings.Contains(strings.ToLower(line), query) {
			fmt.Printf("%d: %s\n", i+1, line)
			found = true
		}
	}
	return found
}

func main() {
	args := os.Args

	if len(args) < 3 {
		fmt.Fprintln(os.Stderr, "Usage: go run main.go <query> <file_path>")
		os.Exit(1)
	}
	query := args[1]
	filePath := args[2]

	contents, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: Could not read file '%s': %v\n", filePath, err)
		os.Exit(1)
	}

	search(query, string(contents))

}
