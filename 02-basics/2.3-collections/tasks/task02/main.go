// Задание 2: Инвертировать словарь
//
// Напиши функцию invertMap, которая:
//   - принимает map[string]int (название -> число)
//   - возвращает map[int]string (число -> название)
//
// Потом выведи результат: отсортируй ключи по возрастанию и выведи каждую пару.
//
// Используй пакет slices для сортировки ключей.
//
// Ожидаемый вывод:
//   1 -> яблоко
//   2 -> банан
//   3 -> апельсин
//
// Запусти: go run main.go

package main

import (
	"fmt"
	"maps"
	"slices"
)

// TODO: напиши функцию invertMap(m map[string]int) map[int]string
func invertMap(m map[string]int) map[int]string {
	inv := make(map[int]string)

	for k, v := range m {
		inv[v] = k
	}

	return inv
}

func main() {
	fruits := map[string]int{
		"яблоко":   1,
		"банан":    2,
		"апельсин": 3,
	}

	inverted := invertMap(fruits)

	keys := slices.Sorted(maps.Keys(inverted))
	for _, k := range keys {
		fmt.Printf("%d -> %s\n", k, inverted[k])
	}
}
