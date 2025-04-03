// Package algorithms содержит тесты для алгоритма поиска второго по величине элемента
package algorithms

import "testing"

func TestFindSecondMax(t *testing.T) {
	tests := []struct {
		name string
		arr  []int
		want int
	}{
		{"Test 1", []int{5, 2, 8, 1, 9, 4}, 8},
		{"Test 2", []int{1, 2, 3, 4, 5}, 4},
		{"Test 3", []int{5, 5, 5, 5}, -1},
		{"Test 4", []int{10, 5}, 5},
		{"Test 5", []int{5, 10}, 5},
		{"All same elements", []int{3, 3, 3}, -1},
		{"With negative numbers", []int{-5, -2, -8}, -5},
		{"Mixed numbers", []int{5, 5, 2, 5}, 2},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FindSecondMax(tt.arr); got != tt.want {
				t.Errorf("FindSecondMax() = %v, want %v", got, tt.want)
			}
		})
	}
}
