package mapred

// todo implement mapreduce

import (
	"regexp"
	"strings"
	"sync"
)

// MapReduce is the struct that implements the MapReduceInterface.
type MapReduce struct {
}

// Ensure MapReduce implements MapReduceInterface by including all methods.
var _ MapReduceInterface = (*MapReduce)(nil)

func (mr *MapReduce) wordCountMapper(text string) []KeyValue {
	// Regex to filter out special chars and numericals:
	re := regexp.MustCompile(`[^a-zA-Z\\s]+`)

	// Clean the text
	cleanedText := re.ReplaceAllString(text, " ")

	// Split into individual words
	words := strings.Fields(cleanedText)

	//  Emit (word, 1) for each word
	var results []KeyValue
	for _, word := range words {
		if word != "" {
			// Convert to lowercase and Value is 1 (int)
			results = append(results, KeyValue{Key: strings.ToLower(word), Value: 1})
		}
	}
	return results
}

// The Reducer Method
func (mr *MapReduce) wordCountReducer(key string, values []int) KeyValue {
	totalCount := 0
	// Sum up all the individual counts (which are all 1s from the mapper)
	for _, count := range values {
		totalCount += count
	}
	// Return the final key-value pair with the total count (int)
	return KeyValue{Key: key, Value: totalCount}
}

// The Coordinator/Runner Method
func (mr *MapReduce) Run(inputChunks []string) map[string]int {
	// Configuration for concurrency (can be constants or fields on MapReduce)
	const nMap = 8    // Number of concurrent mappers
	const nReduce = 4 // Number of concurrent reducers

	// Channels for communication between phases
	mapInputChan := make(chan string)    // Input to mappers
	mapOutputChan := make(chan KeyValue) // Output from mappers

	var mapWG sync.WaitGroup

	// MAP PHASE: Start N mapper goroutines
	for i := 0; i < nMap; i++ {
		mapWG.Add(1)
		go func() {
			defer mapWG.Done()
			for chunk := range mapInputChan {
				// Call the method on the struct
				results := mr.wordCountMapper(chunk)
				for _, kv := range results {
					mapOutputChan <- kv
				}
			}
		}()
	}

	// Feed input to mappers
	go func() {
		for _, chunk := range inputChunks {
			mapInputChan <- chunk
		}
		close(mapInputChan)
	}()

	// Wait for mappers to finish, then close the output channel
	go func() {
		mapWG.Wait()
		close(mapOutputChan)
	}()

	// SHUFFLE PHASE: Group all intermediate results by key
	intermediate := make(map[string][]int) // Key is string, values are []int

	// Read all results from the mappers and group them by key
	for kv := range mapOutputChan {
		intermediate[kv.Key] = append(intermediate[kv.Key], kv.Value)
	}

	// REDUCE PHASE: Setup
	finalResult := make(map[string]int)
	var reduceWG sync.WaitGroup
	var resultMutex sync.Mutex // Mutex to protect finalResult map

	reduceInputChan := make(chan string) // Keys to be reduced

	// Start N reducer goroutines
	for i := 0; i < nReduce; i++ {
		reduceWG.Add(1)
		go func() {
			defer reduceWG.Done()
			for key := range reduceInputChan {
				values := intermediate[key]
				// Call the method on the struct
				finalKV := mr.wordCountReducer(key, values)

				resultMutex.Lock()
				finalResult[finalKV.Key] = finalKV.Value
				resultMutex.Unlock()
			}
		}()
	}

	// Feed unique keys to reducers
	go func() {
		for key := range intermediate {
			reduceInputChan <- key
		}
		close(reduceInputChan)
	}()

	// Wait for all reducers to finish
	reduceWG.Wait()

	return finalResult
}
