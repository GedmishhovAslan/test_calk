package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func romanToArabic(roman string) int {
	romanNumerals := map[rune]int{'I': 1, 'V': 5, 'X': 10, 'L': 50, 'C': 100, 'D': 500, 'M': 1000}
	arabic := 0
	lastValue := 0

	for i := len(roman) - 1; i >= 0; i-- {
		value := romanNumerals[rune(roman[i])]
		if value < lastValue {
			arabic -= value
		} else {
			arabic += value
		}
		lastValue = value
	}
	return arabic

}

func arabicToRoman(arabic int) string {
	numerals := []struct {
		Value  int
		Symbol string
	}{
		{1000, "M"}, {900, "CM"}, {500, "D"}, {400, "CD"},
		{100, "C"}, {90, "XC"}, {50, "L"}, {40, "XL"},
		{10, "X"}, {9, "IX"}, {5, "V"}, {4, "IV"}, {1, "I"},
	}

	var roman strings.Builder
	for _, numeral := range numerals {
		for arabic >= numeral.Value {
			roman.WriteString(numeral.Symbol)
			arabic -= numeral.Value
		}
	}
	return roman.String()
}

func calculate(num1, num2 int, operator string) int {
	switch operator {
	case "+":
		return num1 + num2
	case "-":
		return num1 - num2
	case "*":
		return num1 * num2
	case "/":
		return num1 / num2
	default:
		err := errors.New("Не корректный оператор")
		fmt.Println("Ошибка:", err)
		return 0

	}

}

func main() {
	var firstElement, operator, secondElement string

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Введите выражение (например, 1 + 1):")
	scanner.Scan()
	input := scanner.Text()
	parts := strings.Fields(input)
	if len(parts) != 3 {
		err := errors.New("Не корректное значение")
		fmt.Println("Ошибка:", err)
		return
	}

	firstElement = parts[0]
	operator = parts[1]
	secondElement = parts[2]
	flag1 := false
	flag2 := false

	num1, err1 := strconv.Atoi(firstElement)
	if err1 != nil {
		flag1 = true
		num1 = romanToArabic(firstElement)
	}

	num2, err2 := strconv.Atoi(secondElement)
	if err2 != nil {
		flag2 = true
		num2 = romanToArabic(secondElement)
	}
	if flag1 != flag2 {
		fmt.Println("Оба значения должны быть в одинаковом формате")
		return
	}

	_, err := checkNumbers(num1, num2)
	if err != nil {
		fmt.Println("Ошибка:", err)
		return
	}

	result := calculate(num1, num2, operator)
	if flag1 == true && flag2 == true {
		if result < 0 {
			err := errors.New("Римское число не может быть отрицательным")
			fmt.Println("Ошибка:", err)
			return
		}
		fmt.Printf("Result (Roman): %s\n", arabicToRoman(result))

	} else {
		fmt.Printf("Result (Arabic): %d\n", result)

	}

}

func checkNumbers(num1, num2 int) (bool, error) {
	if num1 < 0 || num1 > 10 && num2 < 0 || num2 > 10 {
		return false, errors.New("число вне диапазона 1-10")
	}
	return true, nil
}
