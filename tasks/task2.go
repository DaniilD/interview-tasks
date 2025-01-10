package tasks

import "github.com/samber/lo"

/*
Задача 2: Напиши функцию на Go, которая принимает на вход срез целых чисел и параллельно обрабатывает их в несколько горутин. Каждая горутина должна умножать число на 2 и отправлять результат в канал. В конце функция должна возвращать срез с результатами.

Условия:

Обработка должна выполняться параллельно в 4 горутинах.
Порядок результатов в возвращаемом срезе должен совпадать с порядком исходных чисел.
*/

func handle(nums []int) []int {
	chunks := lo.Chunk(nums, len(nums)/4)
	resultChans := make([]chan int, 0, len(chunks))
	for _, chunk := range chunks {
		resultChan := worker(chunk)
		resultChans = append(resultChans, resultChan)
	}

	output := make([]int, 0, len(nums))
	for _, resultChan := range resultChans {
		for num := range resultChan {
			output = append(output, num)
		}
	}

	return output
}

func worker(nums []int) chan int {
	result := make(chan int, len(nums))

	go func() {
		for _, num := range nums {
			result <- num * 2
		}
		close(result)
	}()

	return result
}
