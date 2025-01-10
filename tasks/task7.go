package tasks

import "sort"

/*
Напиши функцию на Go, которая принимает срез чисел и возвращает новый срез,
содержащий элементы с наибольшими значениями, но не более, чем 3 элемента.
Если элементов меньше трех, то вернуть все элементы. Функция должна работать за время O(n), где n — длина среза.
*/

func TopThreeNumbers(nums []int) []int {
	if len(nums) < 2 {
		return nums
	}

	sort.Slice(nums, func(i, j int) bool {
		return nums[i] > nums[j]
	})

	if len(nums) <= 3 {
		return nums
	}

	return nums[:3]
}
