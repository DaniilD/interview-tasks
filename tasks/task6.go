package tasks

import "sort"

/*
Напиши функцию на Go, которая принимает срез строк и возвращает срез строк, отсортированных по длине.
Если несколько строк имеют одинаковую длину, они должны быть отсортированы в лексикографическом порядке.
*/

func SortStringSlice(input []string) []string {
	if len(input) < 2 {
		return input
	}
	sort.Slice(input, func(i, j int) bool {
		one := []rune(input[i])
		two := []rune(input[j])

		return len(one) < len(two)
	})

	return input
}

// SortStringSlice2 Вариант от нейросети
func SortStringSlice2(input []string) []string {
	if len(input) < 2 {
		return input
	}

	sort.Slice(input, func(i, j int) bool {
		// Сначала сортируем по длине строки
		if len(input[i]) != len(input[j]) {
			return len(input[i]) < len(input[j])
		}
		// Если длины одинаковые, сортируем лексикографически
		return input[i] < input[j]
	})

	return input
}
