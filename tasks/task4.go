package tasks

import (
	"errors"
	"fmt"
)

/*
Напиши функцию на Go, которая выполняет безопасное удаление элемента из среза.
Функция должна принимать срез целых чисел и индекс элемента, который нужно удалить,
а затем возвращать новый срез без этого элемента. Обработай возможные ошибки, такие как передача недопустимого индекса.
*/

func RemoveElement(nums []int, index int) ([]int, error) {
	if index < 0 || index > (len(nums)-1) {
		return nums, errors.New("недопустимый индекс")
	}

	output := make([]int, 0, len(nums)-1)
	for i, num := range nums {
		if i != index {
			output = append(output, num)
		}
	}

	return output, nil
}

// RemoveElement2 Вариант от нейросети
func RemoveElement2(nums []int, index int) ([]int, error) {
	if index < 0 || index >= len(nums) {
		return nil, fmt.Errorf("index %d out of range", index)
	}

	output := append(nums[:index], nums[index+1:]...)
	return output, nil
}
