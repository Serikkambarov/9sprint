package main

import "testing"

func TestGenerateRandomElements(t *testing.T) {
	result0 := generateRandomElements(0)
	if len(result0) != 0 {
		t.Errorf("при size=0 ожидался пустой срез, но получили len=%d", len(result0))
	}

	size := 10
	result := generateRandomElements(size)
	if len(result) != size {
		t.Errorf("Ожидался срез длины %d, но получен %d", size, len(result))
	}
	for i, v := range result {
		if v < 0 || v >= 1000000 {
			t.Errorf("Элемент на позиции %d вне диапазона: %d", i, v)
		}
	}
}

func TestMaximum(t *testing.T) {
	if maximum([]int{1, 5, 3, 2}) != 5 {
		t.Errorf("Ожидали максимум 5, получили другое значение")
	}
	if maximum([]int{42}) != 42 {
		t.Errorf("максимум одного элемента неверен")
	}
	if maximum([]int{-1, -5, -3}) != -1 {
		t.Errorf("максимум для отрицательных неверен")
	}

	defer func() {
		if r := recover(); r == nil {
			t.Errorf("ожидался panic при пустом срезе, но его не было")
		}
	}()
	_ = maximum([]int{}) // проверка паники
}

func TestMaxChunks(t *testing.T) {
	result := maxChunks([]int{1, 3, 2, 0})
	if result != 3 {
		t.Errorf("maxChunks failed: got %d want %d", result, 3)
	}

	if maxChunks([]int{7}) != 7 {
		t.Errorf("для одного элемента ожидали 7")
	}

	defer func() {
		if r := recover(); r == nil {
			t.Errorf("ожидался panic на пустом срезе в maxChunks, но его не было")
		}
	}()
	_ = maxChunks([]int{}) // проверка паники
}
