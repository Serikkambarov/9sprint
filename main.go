package main

import (
	"fmt"
	"math/rand"
	"time"
	"sync"

)

const (
	SIZE   = 100_000_000
	CHUNKS = 8
)

// generateRandomElements generates random elements.
func generateRandomElements(size int) []int {
	data := make([]int, size) // создаем срез для хранения случайных чисел
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < size; i++ {
		data[i] = rand.Intn(1000000) // заполняем срез случайными числами от 0 до 999999
	}
	 return data 
}

// maximum returns the maximum number of elements.
func maximum(data []int) int {
	if len(data) == 0 {
		panic("пустой слайс")
	}
	max := data[0] 
	for _, v := range data {
		if v > max {
			max = v
		}
	}
	 return  max 
}

// maxChunks returns the maximum number of elements in a chunks.
func maxChunks(data []int) int {
	var wg sync.WaitGroup
	if len(data) == 0 {
		panic("пустой слайс")
	}

	results := make(chan int, CHUNKS)
	chunkSize := (len(data) + CHUNKS - 1) / CHUNKS

	activeChunks := 0 // считаем количество реально запущенных горутин

	for i := 0; i < CHUNKS; i++ {
		start := i * chunkSize
		if start >= len(data) {
			continue // нет данных — пропускаем
		}
		end := start + chunkSize
		if end > len(data) {
			end = len(data)
		}
		part := data[start:end]
		if len(part) == 0 {
			continue
		}

		activeChunks++ // учитываем этот кусок
		wg.Add(1)
		go func(part []int) {
			defer wg.Done()
			max := part[0]
			for _, v := range part {
				if v > max {
					max = v
				}
			}
			results <- max
		}(part)
	}

	// закрытие канала после завершения всех горутин
	
	wg.Wait()
	close(results)
	

	// читаем ровно столько максимумов, сколько горутин запускалось
	var chunkMaxes []int
	for i := 0; i < activeChunks; i++ {
		chunkMaxes = append(chunkMaxes, <-results)
	}

	max := chunkMaxes[0]
	for _, v := range chunkMaxes {
		if v > max {
			max = v
		}
	}
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
