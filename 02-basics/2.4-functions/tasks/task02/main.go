// Задание 2: Счётчик с начальным значением
//
// Напиши функцию makeCounter(start int) func() int, которая:
//   - принимает начальное значение
//   - возвращает функцию-счётчик
//   - каждый вызов возвращает следующее число, начиная с start
//
// Ожидаемый вывод:
//   Счётчик от 5:
//   5
//   6
//   7
//   Счётчик от 100:
//   100
//   101
//   Счётчики независимы - счётчик от 5 продолжает:
//   8
//
// Запусти: go run main.go

package main

import "fmt"

// TODO: напиши функцию makeCounter(start int) func() int
// Подсказка: текущее значение храни в переменной внутри makeCounter,
// и обращайся к ней из возвращаемой функции (это и есть замыкание)
func makeCounter(start int) func() int {
	count := start - 1
	return func() int {
		count += 1
		return count
	}
}

func main() {
	// TODO: создай два независимых счётчика и проверь их работу

	counter1 := makeCounter(5)
	counter2 := makeCounter(100)

	fmt.Println("Счётчик от 5:")
	fmt.Println(counter1())
	fmt.Println(counter1())
	fmt.Println(counter1())

	fmt.Println("Счётчик от 100:")
	fmt.Println(counter2())
	fmt.Println(counter2())

	fmt.Println("Счётчики независимы - счётчик от 5 продолжает:")
	fmt.Println(counter1())
}
