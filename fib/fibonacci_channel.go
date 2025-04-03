// Package routines содержит реализацию генерации чисел Фибоначчи с использованием горутин и каналов
// Название: Fibonacci Generator with Go Routines
// Описание: Генерация последовательности Фибоначчи с параллельной обработкой
// Разработчик: [Barbashev Aleksandr]
// Лицензия: GPLv3
// История изменений:
// - 2025-03-04: Первоначальная реализация

package routines

import (
	"math/rand"
	"time"
)

// GenerateFibonacci генерирует последовательность Фибоначчи до n чисел и отправляет в канал
func GenerateFibonacci(n int, ch chan<- int) {
	a, b := 0, 1
	for i := 0; i < n; i++ {
		ch <- a
		a, b = b, a+b
		time.Sleep(100 * time.Millisecond) // Для наглядности
	}
	close(ch)
}

// ProcessFibonacci получает числа из канала, добавляет случайное число и выводит
func ProcessFibonacci(ch <-chan int, done chan<- bool) {
	rand.Seed(time.Now().UnixNano())
	for num := range ch {
		randomAdd := rand.Intn(10) + 1
		result := num + randomAdd
		println("Fibonacci:", num, "+ random:", randomAdd, "=", result)
	}
	done <- true
}
