package main

import ("testing"
		"github.com/stretchr/testify/assert"
)
func TestGenerateRandomElements(t *testing.T) {
	result0 := generateRandomElements(0)
	assert.Equal(t, 0, len(result0), "ожидался пустой срез")

	size := 10
	result := generateRandomElements(size)
	assert.Equal(t, size, len(result), "неверная длина среза")
	
}

func TestMaximum(t *testing.T) {
	assert.Equal(t, 5, maximum([]int{1, 5, 3, 2}), "неправильный максимум")
	assert.Equal(t, 42, maximum([]int{42}), "неправильный максимум для одного элемента")
	assert.Equal(t, -1, maximum([]int{-1, -5, -3}), "неправильный максимум для отрицательных")
	assert.Equal(t, 0, maximum([]int{}), "ожидался 0 на пустом срезе")


}

func TestMaxChunks(t *testing.T) {
	assert.Equal(t ,3, maxChunks([]int{1, 3, 2, 0}), "неправильный максимум в чанках")
	assert.Equal(t, 7, maxChunks([]int{7}), "неправильный результат для одного элемента")
	assert.Equal(t, 0, maxChunks([]int{}), "ожидался 0 на пустом срезе")
	
}
