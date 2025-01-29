package main

func main() {
	ch := make(chan int, 1) // Буферизированный канал с размером 1
	for i := 0; i < 10; i++ {
		select {
		case x := <-ch: // Если в канале есть данные, считываем их
			print(x)
		case ch <- i: // Если в канале нет данных, записываем `i`
		}
	}
}