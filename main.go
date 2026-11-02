package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const (
	SIZE   = 100_000_000
	CHUNKS = 8
)

// generateRandomElements generates random elements.
func generateRandomElements(size int) []int {
	numbers := make([]int, size)
	for i := 0; i < size; i++ {
		numbers[i] = rand.Int()
	}
	return numbers
}

// maximum returns the maximum number of elements.
func maximum(data []int) int {
	if len(data) == 0 {
		return 0
	}

	max := data[0]

	for _, v := range data {
		if v > max {
			max = v
		}
	}

	return max
}

// maxChunks returns the maximum number of elements in a chunks.
func maxChunks(data [][]int) int {
	count := len(data)
	var wg sync.WaitGroup
	maxGroup := make([]int, count)
	wg.Add(count)
	for i, chunk := range data {
		go func(c []int) {
			defer wg.Done()
			max := maximum(c)
			maxGroup[i] = max
		}(chunk)
	}
	wg.Wait()
	return maximum(maxGroup)
}

func splitTochanks(data []int, count int) [][]int {
	result := make([][]int, count)
	chunkSize := len(data) / count
	for i := range count {
		start := i * chunkSize
		end := start + chunkSize
		if i == count - 1 {
			end = len(data)
		}
		result[i] = data[start:end]
	}
	return result
}

func main() {
	fmt.Printf("Генерируем %d целых чисел", SIZE)
	randomNumbers := generateRandomElements(SIZE)

	fmt.Println("Ищем максимальное значение в один поток")
	startTime := time.Now()
	max := maximum(randomNumbers)
	elapsed := time.Since(startTime).Microseconds()

	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d ms\n", max, elapsed)

	fmt.Printf("Ищем максимальное значение в %d потоков\n", CHUNKS)
	chunks := splitTochanks(randomNumbers, CHUNKS)
	startTime = time.Now()
	max = maxChunks(chunks)
	elapsed = time.Since(startTime).Microseconds()

	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d ms\n", max, elapsed)
}
