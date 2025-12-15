package main

import (
	"exc9/mapred"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

// Define the path to the input file (relative to the project root)

const inputFilePath = "res/meditations.txt"

// WordCount holds a word and its final frequency, used for sorting.
type WordCount struct {
	Word  string
	Count int
}

// sortResults converts the map to a slice of WordCount structs and sorts it by the word count (descending).
func sortResults(results map[string]int) []WordCount {
	var wcSlice []WordCount
	for word, count := range results {
		wcSlice = append(wcSlice, WordCount{Word: word, Count: count})
	}

	// Sort the slice by Count in descending order (highest count first)
	sort.Slice(wcSlice, func(i, j int) bool {
		return wcSlice[i].Count > wcSlice[j].Count
	})
	return wcSlice
}

// ReadInput reads the file contents and splits them into chunks (lines), using
// filepath.Abs for robust path resolution.
func ReadInput(filePath string) ([]string, error) {

	// Resolve the relative path to an absolute path
	absolutePath, err := filepath.Abs(filePath)
	if err != nil {
		return nil, fmt.Errorf("could not get absolute path for %s: %w", filePath, err)
	}

	// Use the absolute path to read the file
	content, err := os.ReadFile(absolutePath)
	if err != nil {
		return nil, fmt.Errorf("error opening file at %s: %w", absolutePath, err)
	}

	text := string(content)
	lines := strings.Split(text, "\n")
	return lines, nil
}

// Main function
func main() {

	// todo read file
	inputChunks, err := ReadInput(inputFilePath)
	if err != nil {
		log.Fatalf("Error reading input file: %v", err)
	}

	// todo run your mapreduce algorithm
	var mr mapred.MapReduce
	results := mr.Run(inputChunks)

	// todo print your result to stdout

	// Convert the final map to a slice and sort by count
	sortedFrequencies := sortResults(results)

	fmt.Println("\n--- FINAL WORD FREQUENCY REPORT (Sorted by Count) ---")
	fmt.Printf("Total unique words counted: %d\n\n", len(results))

	fmt.Println("Top 20 Words:")
	for i, item := range sortedFrequencies {
		if i >= 20 {
			break // Stop after printing the top 20
		}
		// Print the results, using the sorted item's Word and Count
		fmt.Printf("  %2d. %-15s: %d\n", i+1, item.Word, item.Count)
	}
}
