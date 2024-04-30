package main

import (
	"bufio"
	"fmt"
	"os"
)

func calculate() float64 {
	var first, second, res float64
	var operation string

	fmt.Println("Введите первое число:")
	fmt.Scan(&first)
	fmt.Println("Введите второе число:")
	fmt.Scan(&second)
	fmt.Println("Введите знак операции:")
	fmt.Scan(&operation)

	switch operation {
	case "+":
		res = first + second
	case "-":
		res = first - second
	case "*":
		res = first * second
	case "/":
		res = first / second

	default:
		fmt.Println("Неверный знак операции!")
	}

	fmt.Println("Итоговое значение = " + fmt.Sprint(res))

	return res
}

func saveResultToFile(result float64) {
	textfile, err := os.Create("result.txt")

	if err != nil {
		fmt.Println(err)
		return
	}

	textfile.WriteString("Итоговое значение = " + fmt.Sprint(result))
	textfile.Close()
}

func exitOnEnterKeyPress() {
	fmt.Print("Нажмите любую клавишу для завершения...")

	s := bufio.NewScanner(bufio.NewReader(os.Stdin))
	skipLast := true

	for s.Scan() {
		if skipLast {
			skipLast = false
			continue
		}

		if s.Text() == "" {
			os.Exit(0)
		}
	}
}

func main() {
	var calculatedValue = calculate()
	saveResultToFile(calculatedValue)

	exitOnEnterKeyPress()
}
