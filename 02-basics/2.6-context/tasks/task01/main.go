// Задание 1: Таймаут запроса
//
// Напиши функцию fetchData(ctx context.Context) (string, error),
// которая:
//   1. Имитирует долгий запрос через time.Sleep(3 * time.Second)
//   2. После ожидания проверяет, не был ли контекст отменён
//   3. Если отменён - возвращает пустую строку и ошибку из ctx.Err()
//   4. Если не отменён - возвращает строку "данные получены" и nil
//
// Подсказка: используй select с ctx.Done() и time.After()
// (как в примере из example/main.go)
//
// В main() вызови fetchData с таймаутом в 1 секунду.
// Не забудь вызвать defer cancel()!
//
// Ожидаемый вывод:
//   Запрос не успел: context deadline exceeded
//
// Запусти: go run main.go

package main

import (
	"context"
	"fmt"
	"time"
)

func fetchData(ctx context.Context) (string, error) {
	time.Sleep(3 * time.Second)

	select {
	case <-ctx.Done():
		return "", ctx.Err()
	default:
		return "данные получены", nil
	}
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	data, err := fetchData(ctx)
	if err != nil {
		fmt.Printf("Запрос СКОРЕЕ ВСЕГО не успел: %s", err) // Как тип ошибки узнать? Там же просто error. =(
	} else {
		fmt.Printf("Обрабатываю данные: %s", data)
	}
}
