package tasks

import (
	"sync"
)

/*
Напиши функцию на Go, которая демонстрирует работу с гонками данных. Затем исправь её,
используя мьютекс для предотвращения гонок.

Часть 1: Напиши функцию increment на Go, которая увеличивает глобальную переменную counter на 1 в 100 горутинах.
Каждая горутина должна выполнить 1000 итераций увеличения.

Часть 2: Исправь код, добавив мьютекс для синхронизации доступа к counter, чтобы избежать гонок данных.
*/
var counter int
var mu sync.Mutex

func Increment() {
	wg := sync.WaitGroup{}
	mu = sync.Mutex{}

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < 1000; j++ {
				mu.Lock()
				counter++
				mu.Unlock()
			}
		}()
	}

	wg.Wait()
}
