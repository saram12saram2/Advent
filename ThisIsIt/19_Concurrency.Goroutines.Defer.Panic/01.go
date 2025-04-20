package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	//printCPU()
	//printLIFO()
	//gosched()
	//waitGroup2()
	numGreeters()
}

func printCPU() {
	fmt.Println("CPU: ", runtime.NumCPU())
}

func printLIFO() {
	defer fmt.Println("Это выполняется последним")
	fmt.Println("Это выполняется первым")

	defer fmt.Println(1)
	defer fmt.Println(2)
	defer fmt.Println(3)
	fmt.Println("Main завершён")
}

func gosched() {
	go func() {
		for i := 0; i < 5; i++ {
			fmt.Println("Горутина:", i)
			runtime.Gosched() // переключение на другую горутину
		}
	}()
	fmt.Println("Главная горутина")
}

func waitGroup() {
	var wg sync.WaitGroup

	wg.Add(10)

	for i := 0; i < 10; i++ {
		go func(n int) {
			defer wg.Done()
			fmt.Println(n)
		}(i)
	}
	wg.Wait()
}

func waitGroup2() { // the buggy version важно передавать i как аргумент (n) в анонимную функцию.

	var wg sync.WaitGroup

	wg.Add(3)

	for i := 0; i < 3; i++ {
		go func() {
			defer wg.Done()
			fmt.Println(i) // shared variable, not safe! Это называется ловушка замыкания (closure trap).
		}()
	}
	wg.Wait()
}

/*
Here, you don’t pass i into the function, and goroutines reference i from outside.
тут гпт утверждает что переменная i уходит в замыкание, и как n — это независимая копия?
Это и есть ловушка замыкания (замыкание переменной цикла)— горутины "захватывают" внешнюю переменную, которая изменилась, пока они запускались.

Всегда передавай переменные цикла как аргументы в горутину.

*/

func numGreeters() {
	//объявляем анонимную ф-цию и сохраняем ее в переменную hello.
	// ф-ция принимает два аргумента: указатель на группу ожидания и номер "приветствующего"
	hello := func(wg *sync.WaitGroup, id int) {
		defer wg.Done() // говорит: "Когда эта горутина завершится, уменьши счётчик ожидания на 1"
		fmt.Printf("Hello from %v!\n", id)
	}

	const numGreeters = 5
	var wg sync.WaitGroup

	wg.Add(numGreeters) // Говорим: "Я собираюсь запустить 5 горутин — дождись всех".

	for i := 0; i < numGreeters; i++ {
		go hello(&wg, i+1) //  запускаем асинхронно = параллельно
	}

	wg.Wait()
}
