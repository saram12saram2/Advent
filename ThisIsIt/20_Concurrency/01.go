package main

import (
	"fmt"
	"sync"
)

func exit() {
	for i := 0; i < 10; i++ {
		go fmt.Println(i) // Запускаем 10 горутин
	}
	fmt.Println("Exit") // Программа может завершиться до вывода всех чисел
}

func main() {

	// exit()

	var wg sync.WaitGroup

	wg.Add(10) // Указываем, что будем ждать 10 горутин

	for i := 0; i < 10; i++ {
		go func(n int) {
			defer wg.Done() // Горутина завершилась - уменьшаем счётчик
			fmt.Println(n)
		}(i) // Передаём i как аргумент, чтобы избежать ошибки
	}

	wg.Wait() // Ждём завершения всех горутин
	fmt.Println("Exit")
}
