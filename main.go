package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"runtime"
	"sync"
)

type Data struct {
	A int `json:"a"`
	B int `json:"b"`
}

func worker(data []Data, wg *sync.WaitGroup, resultChan chan int) {
	defer wg.Done()
	sum := 0
	for _, item := range data {
		sum += item.A + item.B
	}
	resultChan <- sum
}

func main() {
	jsonFile := "data.json"
	numGoroutines := runtime.GOMAXPROCS(0)
	fmt.Println("Enter goroutines quantity")
	_, err := fmt.Scanln(&numGoroutines)
	if err != nil {
		log.Fatalf("Failed to read goroutines quantity %v", err)
	}

	file, err := ioutil.ReadFile(jsonFile)
	if err != nil {
		log.Fatalf("Failed to read file: %v", err)
	}

	var data []Data
	if err := json.Unmarshal(file, &data); err != nil {
		log.Fatalf("Failed to parse JSON: %v", err)
	}

	chunkSize := (len(data) + numGoroutines - 1) / numGoroutines
	var wg sync.WaitGroup
	resultChan := make(chan int, numGoroutines)
	for i := 0; i < numGoroutines; i++ {
		start := i * chunkSize
		end := start + chunkSize
		if end > len(data) {
			end = len(data)
		}
		wg.Add(1)
		go worker(data[start:end], &wg, resultChan)
	}

	wg.Wait()
	close(resultChan)

	totalSum := 0
	for sum := range resultChan {
		totalSum += sum
	}

	fmt.Printf("Total Sum: %d\n", totalSum)
}
