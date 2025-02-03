package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

//Уберите комментарий если необходимо включить логгирование

func Logger(input string) {
	//fmt.Println(input)
}

func main() {
	for {
		RawInputData := readInput()
		Logger("\n Полученное выражение:" + RawInputData)
		parsedData := parseData(RawInputData)
		Logger("\n Полученные данные:\n первый аргумент: " + parsedData[0] + "\n операция: " + parsedData[1] + "\n второй аргумент: " + parsedData[2])
		if checkFirstArg(parsedData[0]) && checkOperation(parsedData[1]) {
			secondArgType := checkSecondArg(parsedData[2])
			if (secondArgType == "int") || (secondArgType == "string") {
				Logger("\n проверка пройдена, вычисляем...")
				calculatedValue := calculate(parsedData, secondArgType)
				Logger("\n Вычисленное значение до форматирования:" + calculatedValue)
				formattedString := formatResult(calculatedValue)
				Logger("\n Вычисленное значение:" + `"` + formattedString + `"`)

				fmt.Println(formattedString)
			}
		}
		Logger("\n Программа закончила выполнение")
	}
}

func readInput() string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("\n Введите выражение в виде \n\"<строка1>\" <операция> \"<строка2>\" \n или \n\"<строка>\" <операция> <число>\n ")
	text, _ := reader.ReadString('\n') // Ждет ввода данных в формате строки
	return text
}

func parseData(input string) []string {
	// Регулярное выражение для поиска подстрок, игнорируя пробелы в кавычках.
	// Это выражение ищет либо текст внутри кавычек ("..."), либо непробельные символы вне кавычек.
	// - "([^"]+)" - захватывает текст в кавычках.
	// - (\S+) - захватывает последовательности непробельных символов.
	re := regexp.MustCompile(`"([^"]+)"|(\S+)`)

	// Используем регулярное выражение для поиска всех совпадений в строке input.
	matches := re.FindAllStringSubmatch(input, -1)

	// Инициализируем срез для хранения частей строки, которые будут извлечены.
	var parts []string

	// Проходим по всем найденным совпадениям.
	for _, match := range matches {
		// Если текст в кавычках найден (match[1] не пустой), добавляем его в parts с кавычками.
		if match[1] != "" {
			parts = append(parts, fmt.Sprintf("\"%s\"", match[1])) // часть в кавычках, добавляем кавычки
		} else {
			// Если текст вне кавычек найден, добавляем его в parts.
			parts = append(parts, match[2]) // часть вне кавычек
		}
	}

	// Проверяем, что количество частей соответствует ожидаемому значению (3).
	if len(parts) == 3 {
		return parts // Если все в порядке, возвращаем части.
	} else {
		// Если количество частей не соответствует ожиданиям, логируем ошибку.
		Logger("\nКоличество аргументов не соответствует ожидаемому. ожидалось: 3, найдено:" + fmt.Sprint(len(parts)))
		Logger("Элементы аргументов:")

		// Логируем каждую найденную часть для отладки.
		for i, part := range parts {
			Logger(fmt.Sprintf("Аргумент %d: %s", i+1, part))
		}

		// Вызываем панику, так как количество аргументов неверное.
		ArgErrorPanic()
	}

	// Возвращаем nil, хотя этот код никогда не будет достигнут, так как функция либо вернет parts,
	// либо вызовет панику.
	return nil
}

func ArgErrorPanic() {
	panic("Недопустимая операция")
}

func checkFirstArg(input string) bool {
	// Проверка на наличие кавычек
	if len(input) < 2 || (input[0] != '"') || (input[len(input)-1] != '"') {
		Logger("\n Первый аргумент имеет менее двух кавычек")
		ArgErrorPanic()
	} else if len(input) < 2 || (input[0] == '"') && (input[len(input)-1] == '"') {

		Logger("\n Первый аргумент имеет кавычки, значит это строка. Проверяем соответствие длинны")

		// Убираем кавычки для проверки длины
		trimmedInput := input[1 : len(input)-1]

		// Проверка длины строки без кавычек
		if len(trimmedInput) > 10 {
			Logger("\n Первый аргумент является строкой но его длина превышает допустимую. ошибка")
			ArgErrorPanic()
			return false
		}
		Logger("\n Первый аргумент является строкой и проходит проверку")

		return true
	}

	Logger("\n Непредвиденная ошибка")
	ArgErrorPanic()
	return false

}

func checkOperation(op string) bool {
	// Проверяем, что операция является одной из поддерживаемых
	switch op {
	case "+", "-", "*", "/":
		return true
	default:
		Logger("\n Операция является недопустимой")
		ArgErrorPanic()
		return false
	}
}

func checkSecondArg(input string) string {
	// Проверка на наличие кавычек
	if len(input) < 2 || (input[0] != '"') || (input[len(input)-1] != '"') {
		Logger("\n Второй аргумент имеет менее двух кавычек")
		// Если нет кавычек, проверяем, является ли это числом
		num, err := strconv.Atoi(input)
		if err != nil { // проверка на число
			Logger("\n Второй аргумент не является int числом и имеет менее двух кавычек. ошибка")
			ArgErrorPanic()
		} else if num < 1 || num > 10 {
			Logger("\n Второй аргумент является int числом но не входит в диапазон. ошибка")
			ArgErrorPanic()
		} else {
			return "int" // Число в диапазоне
		}
		return ""
	} else if len(input) < 2 || (input[0] == '"') && (input[len(input)-1] == '"') {

		Logger("\n Второй аргумент имеет кавычки, значит это строка. Проверяем соответствие длинны")

		// Убираем кавычки для проверки длины
		trimmedInput := input[1 : len(input)-1]

		// Проверка длины строки без кавычек
		if len(trimmedInput) > 10 {
			Logger("\n Второй аргумент является строкой но его длина превышает допустимую. ошибка")
			ArgErrorPanic()
			return ""
		}
		Logger("\n Второй аргумент является строкой и проходит проверку")

		return "string"
	}

	Logger("\n Непредвиденная ошибка")
	ArgErrorPanic()
	return ""

}

func calculate(parsedData []string, secondArgType string) string {
	firstArg := parsedData[0][1 : len(parsedData[0])-1] // Убираем кавычки
	operation := parsedData[1]
	secondArg := parsedData[2]

	var result string

	if secondArgType == "string" {
		//если второй аргумент строка
		secondArg := parsedData[2][1 : len(parsedData[2])-1] // Убираем кавычки

		switch operation {
		case "+":
			result = calcAdd(firstArg, secondArg)
		case "-":
			result = calcSub(firstArg, secondArg)

		case "*":
			Logger("\n Попытка умножения строки на строку. ошибка")
			ArgErrorPanic()
		case "/":
			Logger("\n Попытка деления строки на строку. ошибка")
			ArgErrorPanic()
		}

	} else if secondArgType == "int" {
		// Если второй аргумент - число
		count, _ := strconv.Atoi(secondArg) // Преобразуем строку в число
		switch operation {
		case "+":
			Logger("\n Попытка сложения строки с числом. ошибка")
			ArgErrorPanic()
		case "-":
			Logger("\n Попытка вычитания числа из строки. ошибка")
			ArgErrorPanic()
		case "*":
			result = calcMult(firstArg, count)
		case "/":
			result = calcDiv(firstArg, count)
		}
	} else {
		Logger("\n Непредвиденная ошибка")
		ArgErrorPanic() //
	}
	return result
}

// Функция для сложения строк
func calcAdd(str1, str2 string) string {
	Logger("\n Складываем строку и строку")
	return str1 + str2 // Объединяем две строки
}

// Функция для вычитания строк
func calcSub(str1, str2 string) string {
	Logger("\n Вычитаем строку из строки")
	if strings.Contains(str1, str2) {
		return strings.Replace(str1, str2, "", 1) // Удаляем вторую строку из первой
	}
	return str1 // Если второй строки нет в первой, возвращаем первую строку
}

// Функция для умножения строк
func calcMult(str string, count int) string {
	Logger("\n Умножаем строку на число")
	if count <= 0 {
		return "" // Если количество 0 или меньше, возвращаем пустую строку
	}
	return strings.Repeat(str, count) // Повторяем первую строку указанное количество раз
}

// Функция для деления строк
func calcDiv(str string, count int) string {
	Logger("\n Делим строку на число")
	if count <= 0 {
		return "" // Если делитель 0 или меньше, возвращаем пустую строку
	}
	newLength := len(str) / count // Уменьшаем длину первой строки на указанный раз
	if newLength < 0 {
		return "" // Если результат меньше 0, возвращаем пустую строку
	}
	return str[:newLength] // Возвращаем строку до новой длины
}

func formatResult(result string) string {
	const maxLength = 40

	if len(result) > maxLength {
		return result[:maxLength] + "..."
	}
	return result
}
