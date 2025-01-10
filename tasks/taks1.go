package tasks

/*
UniqueSlice Задача 1: Напиши функцию на Go, которая принимает на вход срез строк и возвращает новый срез строк,
содержащий только уникальные строки (удаляя дубликаты). Порядок строк в результирующем срезе должен быть таким же,
как и в исходном срезе.
*/
func UniqueSlice(input []string) []string {
	existWords := make(map[string]struct{}, len(input))
	result := make([]string, 0, len(input))
	for _, word := range input {
		if _, ok := existWords[word]; !ok {
			result = append(result, word)
			existWords[word] = struct{}{}
		}
	}

	return result
}
