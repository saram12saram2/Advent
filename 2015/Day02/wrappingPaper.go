package main

import (
	"bufio" // для чтения ввода
	"fmt"   // для вывода
	"os"
	"strconv"
	"strings"
)

// Function to calculate the total wrapping paper needed (Ф-ция для расчёта общего количества обёрточной бумаги)
func totalWrappingPaper(dimensions []string) int {
	total := 0
	for _, dimension := range dimensions {
		l, w, h := parseDimensions(dimension) // Разбор размеров
		surfaceArea := 2*l*w + 2*w*h + 2*h*l  // Расчёт площади поверхности
		slack := min(l*w, w*h, h*l)           // Расчёт дополнительной бумаги
		total += surfaceArea + slack          // Добавление к общему количеству
	}
	return total
}

// Helper function to parse dimensions and convert them to integers (Ф-ция для разбора размеров)
func parseDimensions(dimension string) (int, int, int) {
	parts := strings.Split(dimension, "x") // Разделение строки на части
	l, _ := strconv.Atoi(parts[0])         // Преобразование первой части в целое число
	w, _ := strconv.Atoi(parts[1])         // Преобразование второй части в целое число
	h, _ := strconv.Atoi(parts[2])         // Преобразование третьей части в целое число
	return l, w, h
}

// Функция для нахождения минимального значения из трёх
func min(a, b, c int) int {
	if a < b {
		if a < c {
			return a
		}
		return c
	}
	if b < c {
		return b
	}
	return c
}

func main() {
	scanner := bufio.NewScanner(os.Stdin) // Создание сканера для чтения из стандартного ввода
	var dimensions []string

	fmt.Println("Введите размеры (например, 25x14x4), затем нажмите Enter. Напишите 'done' для завершения:")

	for scanner.Scan() {
		input := scanner.Text() // Чтение введённой строки
		if input == "done" {    // Проверка на слово "done" для завершения
			break
		}
		dimensions = append(dimensions, input) // Добавление размеров в список
	}

	// Calculate and print the total wrapping paper needed (Вычисление и вывод общего количества обёрточной бумаги)
	fmt.Println("Общее количество квадратных футов обёрточной бумаги:", totalWrappingPaper(dimensions))
}
