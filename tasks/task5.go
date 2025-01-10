package tasks

/*
Напиши функцию на Go, которая принимает два среза целых чисел и возвращает новый срез, содержащий элементы,
которые присутствуют в обоих срезах (пересечение). Срезы могут содержать дубликаты, и они могут быть не отсортированы.
Постарайся сделать решение эффективным.
*/

func IntersectSlices(slice1, slice2 []int) []int {

	if len(slice1) < len(slice2) {
		slice1, slice2 = slice2, slice1
	}

	existsNums := make(map[int]struct{}, len(slice1))
	for _, num := range slice1 {
		if _, ok := existsNums[num]; !ok {
			existsNums[num] = struct{}{}
		}
	}

	output := make([]int, 0, len(slice2))
	for _, num := range slice2 {
		if _, ok := existsNums[num]; ok {
			output = append(output, num)
			delete(existsNums, num)
		}
	}

	return output
}
