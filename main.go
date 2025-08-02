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
	data := make([]int, size) // создаем срез для хранения случайных чисел
	for i := 0; i < size; i++ {
		data[i] = rand.Int() // заполняем срез случайными числами от 0 до 999999
	}
	return data
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
func maxChunks(data []int) int {
	var wg sync.WaitGroup
	if len(data) == 0 {
		return 0
	}

	chunkMaxes := make([]int, CHUNKS)
	chunkSize := (len(data) + CHUNKS - 1) / CHUNKS

	for i := 0; i < CHUNKS; i++ {
		start := i * chunkSize
		if start >= len(data) {
			continue // нет данных — пропускаем
		}
		end := start + chunkSize
		if end > len(data) {
			end = len(data)
		}

		wg.Add(1)
		go func(i int, p []int) {
			defer wg.Done()
			chunkMaxes[i] = maximum(p)
		}(i, data[start:end])
	}

	wg.Wait()
	max := maximum(chunkMaxes)

	return max

}

func main() {
	fmt.Printf("Генерируем %d целых чисел", SIZE)
	slice := generateRandomElements(SIZE)

	fmt.Println("Ищем максимальное значение в один поток")
	start := time.Now()
	max := maximum(slice) // ваш код здесь
	elapsed := time.Since(start).Microseconds()

	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d ms\n", max, elapsed)

	fmt.Printf("Ищем максимальное значение в %d потоков", CHUNKS)
	start = time.Now()
	max = maxChunks(slice)
	elapsed = time.Since(start).Microseconds()
	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d ms\n", max, elapsed)
}
