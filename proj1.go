package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	for {
		RawInputData := readInput()
		fmt.Println("\n Полученное выражение:" + RawInputData)
		parsedData := parseData(RawInputData)
		fmt.Println("\n Полученные данные:\n первый аргумент: " + parsedData[0] + "\n операция: " + parsedData[1] + "\n второй аргумент: " + parsedData[2])
		if checkFirstArg(parsedData[0]) {
			fmt.Println("\n проверка пройдена, вычисляем...")
			calculatedValue := calculate(parsedData[0])
			fmt.Println("\n Вычисленное значение:" + calculatedValue)
		}

		fmt.Println("\n Программа закончила выполнение")
	}
}

func readInput() string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("\n Введите выражение в виде \n\"<строка1>\" <операция> \"<строка2>\" \n или \n\"<строка>\" <операция> <число>\n ")
	text, _ := reader.ReadString('\n') // Ждет ввода данных в формате строки
	return text
}

func parseData(input string) []string {
	// Разделяем строку по пробелам
	parts := strings.Fields(input)
	if len(parts) == 3 {

		return parts
	} else {
		fmt.Println("\n Количество аргументов не соответствует ожидаемому")
		ArgErrorPanic()
	}
	return nil
}

func ArgErrorPanic() {
	panic("Недопустимая операция")
}

func checkFirstArg(input string) bool {
	// Проверка на наличие кавычек
	if len(input) < 2 || input[0] != '"' || input[len(input)-1] != '"' {
		return false
	}

	// Убираем кавычки для проверки длины
	trimmedInput := input[1 : len(input)-1]

	// Проверка длины строки без кавычек
	if len(trimmedInput) > 10 {
		return false
	}

	return true
}

func checkOperation(op string) bool {
	// Проверяем, что операция является одной из поддерживаемых
	switch op {
	case "+", "-", "*", "/":
		return true
	default:
		return false
	}
}

func checkSecondArg(input string) bool {
	// Проверка на наличие кавычек
	if len(input) < 2 || input[0] != '"' || input[len(input)-1] != '"' {
		// Если нет кавычек, проверяем, является ли это числом
		num, err := strconv.Atoi(input)
		if err != nil {
			return false // Не число
		}
		// Проверка диапазона для чисел
		if num < 1 || num > 10 {
			return false // Число вне диапазона
		}
		return true // Число в диапазоне
	}

	// Убираем кавычки для проверки длины
	trimmedInput := input[1 : len(input)-1]

	// Проверка длины строки без кавычек
	if len(trimmedInput) > 10 {
		return false
	}

	return true
}

func calculate(expression string) string {
	return "1"
}

func calcAdd() {

}

func calcSub() {

}
