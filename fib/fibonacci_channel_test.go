// Package routines содержит тесты для генератора чисел Фибоначчи
package routines

import (
	"testing"
	"time"
)

func TestGenerateFibonacci(t *testing.T) {
	ch := make(chan int)
	done := make(chan bool)
	count := 10

	go GenerateFibonacci(count, ch)
	go ProcessFibonacci(ch, done)

	select {
	case <-done:
		return
	case <-time.After(2 * time.Second):
		t.Errorf("Fibonacci generation timed out")
	}
}
