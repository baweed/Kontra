// Основной исполняемый файл проекта
// Название: Control Work Variant 1
// Описание: Реализация контрольной работы по алгоритмам и горутинам
// Разработчик: [Barbashev Aleksandr]
// Лицензия: GPLv3
// История изменений:
// - 2025-03-04: Первоначальная реализация

package main

import (
	"fmt"
	algorithms "kontro/algo"
	routines "kontro/fib"
)

func main() {
	// Задание 1: Алгоритмы
	arr := []int{5, 2, 15, 1, 9, 13}
	secondMax := algorithms.FindSecondMax(arr)
	fmt.Printf("Second max in %v is %d\n", arr, secondMax)

	// Задание 2: Горутины и каналы
	ch := make(chan int)
	done := make(chan bool)

	go routines.GenerateFibonacci(10, ch)
	go routines.ProcessFibonacci(ch, done)

	<-done
}
