package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("Введите значение")
		text, _ := reader.ReadString('\n') // Ждет ввода данных в формате строки
		text = strings.TrimSpace(text)     // Очищает все пустоты (пробелы, табуляцию)
		toNumber, _ := strconv.Atoi(text) // Преобразование строки в число
		fmt.Println(toNumber + 8)        // Вывод результата
	}
}


func main() {
	if len(os.Args) < 2 {
		panic("Необходимо передать строку с выражением.")
	}

	input := os.Args[1]
	result := calculate(input)
	fmt.Println(result)
}

func calculate(expression string) string {
	// Удаляем пробелы и разбиваем строку на компоненты
	expression = strings.ReplaceAll(expression, " ", "")
	if !strings.Contains(expression, "+") && !strings.Contains(expression, "-") && !strings.Contains(expression, "*") && !strings.Contains(expression, "/") {
		panic("Недопустимая операция. Поддерживаемые операции: +, -, *, /.")
	}

	var str1, str2 string
	var operator string

	// Определяем оператор и разбиваем строку
	if strings.Contains(expression, "+") {
		parts := strings.Split(expression, "+")
		operator = "+"
		str1, str2 = parts[0], parts[1]
	} else if strings.Contains(expression, "-") {
		parts := strings.Split(expression, "-")
		operator = "-"
		str1, str2 = parts[0], parts[1]
	} else if strings.Contains(expression, "*") {
		parts := strings.Split(expression, "*")
		operator = "*"
		str1, str2 = parts[0], parts[1]
	} else if strings.Contains(expression, "/") {
		parts := strings.Split(expression, "/")
		operator = "/"
		str1, str2 = parts[0], parts[1]
	}

	// Удаляем кавычки
	str1 = strings.Trim(str1, "\"")
	str2 = strings.Trim(str2, "\"")

	// Проверяем, является ли str2 числом
	if operator == "*" || operator == "/" {
		num, err := strconv.Atoi(str2)
		if err != nil || num < 1 || num > 10 {
			panic("Недопустимое число. Должно быть от 1 до 10.")
		}
		return performStringOperation(str1, str2, operator, num)
	}

	// Операции сложения и вычитания
	if operator == "+" {
		return "\"" + str1 + str2 + "\""
	} else if operator == "-" {
		return "\"" + performSubtraction(str1, str2) + "\""
	}

	panic("Недопустимая операция.")
}

func performStringOperation(str string, strNum string, operator string, num int) string {
	if operator == "*" {
		return "\"" + strings.Repeat(str, num) + "\""
	} else if operator == "/" {
		if num > len(str) {
			return "\"\""
		}
		return "\"" + str[:len(str)/num] + "\""
	}
	return ""
}

func performSubtraction(str1 string, str2 string) string {
	// Удаляем str2 из str1
	result := strings.ReplaceAll(str1, str2, "")
	if result == "" {
		return str1
	}
	return result
}
