package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var romanNumerals = map[string]int{
	"I": 1, "II": 2, "III": 3, "IV": 4, "V": 5,
	"VI": 6, "VII": 7, "VIII": 8, "IX": 9, "X": 10,
}

var arabicToRoman = map[int]string{
	1: "I", 2: "II", 3: "III", 4: "IV", 5: "V",
	6: "VI", 7: "VII", 8: "VIII", 9: "IX", 10: "X",
	20: "XX", 30: "XXX", 40: "XL", 50: "L",
	60: "LX", 70: "LXX", 80: "LXXX", 90: "XC", 100: "C",
}

func romanToInt(roman string) (int, error) {
	if value, exists := romanNumerals[roman]; exists {
		return value, nil
	}
	return 0, errors.New("некорректное римское число")
}

func intToRoman(num int) (string, error) {
	if num <= 0 {
		return "", errors.New("римские числа могут быть только положительными")
	}

	result := ""
	for _, value := range []int{100, 90, 50, 40, 10, 9, 5, 4, 1} {
		for num >= value {
			result += arabicToRoman[value]
			num -= value
		}
	}

	return result, nil
}

func calculate(a, b int, operator string) (int, error) {
	switch operator {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "*":
		return a * b, nil
	case "/":
		if b == 0 {
			return 0, errors.New("деление на ноль")
		}
		return a / b, nil
	default:
		return 0, errors.New("неизвестный оператор")
	}
}

func main() {

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Введите выражение: ")
	or_input, _ := reader.ReadString('\n')
	input := strings.TrimRight(or_input, "\r\n")
	parts := strings.Split(input, " ")

	if len(parts) != 3 {
		panic("некорректный ввод, ожидается формат: число оператор число")
	}

	aStr, operator, bStr := parts[0], parts[1], parts[2]

	var a, b int
	var err error
	isRoman := false

	if _, err = strconv.Atoi(aStr); err == nil {
		a, _ = strconv.Atoi(aStr)
		b, err = strconv.Atoi(bStr)

		if err != nil {
			panic("некорректное число")
		}
		if a < 1 || a > 10 || b < 1 || b > 10 {
			panic("числа должны быть в диапазоне от 1 до 10 включительно")
		}
	} else {
		a, err = romanToInt(aStr)
		if err != nil {
			panic(err)
		}
		b, err = romanToInt(bStr)
		if err != nil {
			panic(err)
		}
		isRoman = true
	}

	result, err := calculate(a, b, operator)
	if err != nil {
		panic(err)
	}

	if isRoman {
		romanResult, err := intToRoman(result)
		if err != nil {
			panic(err)
		}
		fmt.Println("Результат:", romanResult)
	} else {
		fmt.Println("Результат:", result)
	}
}
