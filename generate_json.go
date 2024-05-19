package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"time"
)

type JsonData struct {
	A int `json:"a"`
	B int `json:"b"`
}

func main() {
	const numObjects = 1000000
	const minValue = -10
	const maxValue = 10

	rand.Seed(time.Now().UnixNano())

	data := make([]JsonData, numObjects)
	for i := 0; i < numObjects; i++ {
		data[i] = JsonData{
			A: rand.Intn(maxValue-minValue+1) + minValue,
			B: rand.Intn(maxValue-minValue+1) + minValue,
		}
	}

	file, err := os.Create("data.json")
	if err != nil {
		fmt.Printf("Failed to create file: %v\n", err)
		return
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "    ")

	if err := encoder.Encode(data); err != nil {
		fmt.Printf("Failed to encode data to JSON: %v\n", err)
		return
	}

	fmt.Println("JSON file generated successfully.")
}
