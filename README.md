**Условия задачи:**

1. Калькулятор умеет выполнять операции сложения строк, вычитания строки из
строки, умножения строки на число и деления строки на число: "a" + "b", "a" -
"b", "a" * b, "a" / b.
2. Данные передаются в одну строку.
3. Значения строк, передаваемых в выражении, выделяются двойными кавычками.
4. Результатом сложения двух строк является строка, состоящая из переданных.
5. Результатом деления строки на число n является строка в n раз короче исходной.
6. Результатом умножения строки на число n является строка, в которой переданная
строка повторяется ровно n раз.
7. Результатом вычитания строки из строки является строка, в которой удалена
переданная подстрока или сама исходная строка, если в неё нет вхождения
вычитаемой строки.
8. Калькулятор должен принимать на вход числа от 1 до 10 включительно, не более, и
строки длиной не более 10 символов.
9. Если строка, полученная в результате работы приложения, длиннее 40 символов,
то в выводе после 40 символа должны стоять три точки (...).
10. Калькулятор умеет работать только с целыми числами.
11. Первым аргументом выражения, подаваемым на вход, должна быть строка.
12. При вводе неподходящих чисел, строк или неподдерживаемых операций
приложение выдаёт панику и завершает свою работу.
13. При вводе выражения, не соответствующего одной из вышеописанных
арифметических операций, приложение выдаёт панику и завершает свою работу.

___
Пример работы программы:
``` go
Input:
"100" + "500"

Output:
"100500"

Input:
"Hi World!" - "World!"

Output:
"Hi "

Input:
"Bye-bye!" - "World!"

Output:
"Bye-bye!"

Input:
"Golang" * 5

Output:
"GolangGolangGolangGolangGolang"

Input:
"Example!!!" / 3

Output:
"Exa"
```
___

**Алгоритм работы строкового калькулятора:**

```mermaid
graph TD
    A1[1 Инициализация] --> A2[2 Чтение входных данных]
    A2 --> A3[3 Парсинг входных данных]
    A3 --> A4[4 Валидация ввода]
    A4 --> A5[4.1 Проверка первого аргумента на строку]
    A4 --> A6[4.2 Проверка операции]
    A4 --> A7[4.3 Проверка второго аргумента на строку или число]
    
    A5 --> A8[4.1.1 Проверка на наличие кавычек]
    A8 --> A9[4.1.2 Проверка длины строки]
    
    A6 --> A10[4.2.1 Проверка на поддерживаемые операции]
    
    A7 --> A11[4.3.1 Проверка второго аргумента на строку]
    A7 --> A12[4.3.2 Проверка второго аргумента на число]
    
    A11 --> A13[4.3.1.1 Проверка на наличие кавычек]
    A13 --> A14[4.3.1.2 Проверка длины строки]
    
    A12 --> A15[4.3.2.1 Проверка диапазона числа]
    
    A4 --> A16[5 Выполнение операции]
    
    A16 --> A17[5.1 Сложение]
    A16 --> A18[5.2 Вычитание]
    A16 --> A19[5.3 Умножение]
    A16 --> A20[5.4 Деление]
    
    A17 --> A21[6 Форматирование результата и вывод]
    A18 --> A21
    A19 --> A21
    A20 --> A21



```

1. Инициализация:
-Запустить приложение и подготовить консольный ввод для чтения строки
от пользователя.

3. Чтение входных данных:
- Считать строку, введенную пользователем, которая должна содержать
выражение в формате: "<строка1>" <операция> "<строка2>" или "<строка>"
<операция> <число> (условие 2).

4. Парсинг входных данных:
- Извлечь первую строку, операцию и вторую строку или число из введенной
строки.

5. Валидация ввода:
    3.1 Убедиться, что первый аргумент соответствует условиям (усл. 12, 13):
   - Это строка (условие 11) - у него есть кавычки (усл.3).
   - Первый аргумент длинной не более 10 символов (усл.8)

    3.2 Убедиться, что операция является одной из поддерживаемых: +, -, *, / (усл. 12, 13).

    3.3 Убедиться что второй аргумент соответствует условиям (усл. 12, 13):
    3.3.1 Проверка второго аргумента на строку
    - Определить, что это строка (условие 11) у аргумента есть кавычки (усл.3).
    - Убедиться, что строка не превышают 10 символов в длину (условие 8).
    3.3.2 Проверка второго аргумента на число
    - Или убедиться что это число в диапазоне от 1 до 10 включительно (условие 8).
    

5. Выполнение операции в зависимости от операции:
- Сложение (+): Объединить две строки (условие 4).
- Вычитание (-): Удалить вторую строку из первой. Если второй строки
нет в первой, вернуть нсходную первую строку (условие 7).
- Умножение (*): Повторить первую строку указанное количество раз (условие 6).
- Деление (/): Уменьшить длину первой строки на указанное число раз. Если результат меньше 0, вернуть пустую строку (условие 5).

6. Форматирование результата и вывод:
- Если длина полученной строки превышает 40 символов, обрезать строку до 40 символов и добавить три точки (...) в конце (условие 9).
- Вывести результат в консоль.

7. Обработка ошибок:
- Если на любом этапе происходит ошибка (например, неверный ввод,
неподходящие операции или числа), приложение должно вызвать панику и
завершить свою работу с соответствующим сообщением об ошибке
(условия 12, 13).

____
**Примеры ввод-вывод**
Группа 1: Сложение строк
___

(1.1) Описание: Сложение двух строк с пробелами.
```
Input: "Hello" + " World!"
Output: "Hello World!"

Input: "Hello " + "World!"
Output: "Hello World!"

Input: " " + " "
Output: "  " (двойной пробел)
```

(1.2) Описание: Сложение двух строк чисел, результат — одна длинная строка.
```
Input: "1234567890" + "0987654321"
Output: "12345678900987654321"
```

(1.3) Описание: Сложение двух длинных строк, результат - паника (аргумент слишком длинный).
```
Input: "A very long string" + " that is too long!"
Output: "A very long string that is too long!"

Input: "A very string" + " that is tooooooooooo long!"
Output: "A very long string that is too long!"
```

(1.4) Описание: Сложение пустой строки с непустой, результат — непустая строка.
```
Input: "" + "Non-empty"
Output: "Non-empty"

Input: "Non-empty" + ""
Output: "Non-empty"

Input: "" + ""
Output: "" (пустая строка)
```

(1.5) Описание: Сложение строк, содержащих специальные символы.
```
Input: "Hello@123" + "!$%^&*"
Output: "Hello@123!$%^&*"
```

_____
**Примеры ввод-вывод**
Группа 2: Вычитание строк
____

(2.1) Описание: Удаление подстроки из строки, результат — оставшаяся часть.
```
Input: "Powerbank" - "Power"
Output: "bank"
```


(2.2) Описание: Вычитание строки, которая не содержится в исходной, результат — исходная строка.
```
Input: "Hello" - "World"
Output: "Hello"
```


(2.3) Описание: Удаление подстроки из середины строки, результат — модифицированная строка.
```
Input: "i love u" - "love "
Output: "i u"
```


(2.4) Описание: Вычитание строки из самой себя, результат — пустая строка.
```
Input: "Test" - "Test"
Output: "" (пустая строка)
```

_____
**Примеры ввод-вывод**
Группа 3: Умножение строк
____

(3.1) Описание: Умножение строки на число, результат — повторение строки.
```
Input: "abc" * 3
Output: "abcabcabc"
```


(3.2) Описание: Умножение строки на 10, результат — строка из 10 символов 'X'.
```
Input: "X" * 10
Output: "XXXXXXXXXX"
```


(3.3) Описание: Умножение строки на 1, результат — исходная строка.
```
Input: "Repeat" * 1
Output: "Repeat"
```

(3.4) Описание: Умножение строки на float, результат — паника.
```
Input: "Repeat" * 1.4
Output: "Repeat"

Input: "Repeat" * 1,4
Output: "Repeat"
```


(3.5) Описание: Умножение строки на 0, вызывает ошибку (паника).
```
Input: "String" * 0
Output: (Паника)
```


(3.6) Описание: Умножение строки на число больше 10, вызывает ошибку 
```
Input: "Test" * 11
Output: (Паника)
```

(3.7) Описание: Умножение строки длиннее 6 символов на число больше 9, результат - обрезка до 40 симв и многоточие в конце при выводе
```
Input: "hexagon" * 9
Output: (Паника)
```

____
**Примеры ввод-вывод**
Группа 4: Деление строк
____

(4.1) Описание: Деление строки на 2, результат — первая половина строки.
```
Input: "abcdefghij" / 2
Output: "abcde"
```


(4.2) Описание: Деление длинной строки, результат — обрезка до допустимой длины.
```
Input: "This is a very long string that exceeds forty characters!" / 2
Output: "This is a very long string that ex..."
```


(4.3) Описание: Деление строки на 10, результат — пустая строка, так как длина строки равна 9.
```
Input: "abcdefghi" / 10
Output: "" (пустая строка)
```


(4.4) Описание: Деление строки на 1, результат — исходная строка.
```
Input: "No change" / 1
Output: "No change"
```

(4.5) Описание: Деление строки на float, результат — паника.
```
Input: "Repeat" / 1.4
Output: "Repeat"

Input: "Repeat" / 1,4
Output: "Repeat"
```
____
**Примеры ввод-вывод**
Группа 5: Граничные случаи и проверки
____

(5.1) Описание: Неправильная операция, вызывает ошибку (паника).
```
Input: "Hello" ^ "World"
Output: (Паника)
```


(5.2) Описание: Сложение строки с символами и числом, результат — объединение.
```
Input: "1/n\n+-/* " + "+-/* "
Output: "1/n\n+-/* +-/* "
```


(5.3) Описание: Сложение строк с одинарными кавычками.
```
Input: `biba ` + "boba"
Output: паника

Input: 'biba ' + 'boba'
Output: паника
```
