package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

const (
	plus           = "+"
	minus          = "-"
	multiple       = "*"
	division       = "/"
	degree         = "^"
	openBracket    = "("
	closingBracket = ")"
)

func getInput() []string {
	input := make([]string, len(os.Args[1:]))
	for ind, val := range os.Args[1:] {
		input[ind] = strings.ReplaceAll(val, " ", "")
	}
	return input
}

func validate(str string) (arr []string, err error) {
	prevNum := ""

	if strings.Count(str, "(") != strings.Count(str, ")") {
		return nil, fmt.Errorf("unvalid expression %s", str)
	}

	for i := 0; i < len(str); i++ {
		if _, er := strconv.Atoi(string(str[i])); er == nil {
			prevNum += string(str[i])
		} else {
			if len(prevNum) != 0 {
				arr = append(arr, prevNum)
				prevNum = ""
			}
			switch string(str[i]) {
			case plus:
				arr = append(arr, plus)
			case minus:
				arr = append(arr, minus)
			case multiple:
				arr = append(arr, multiple)
			case division:
				arr = append(arr, division)
			case degree:
				arr = append(arr, degree)
			case openBracket:
				arr = append(arr, openBracket)
			case closingBracket:
				arr = append(arr, closingBracket)
			default:
				return arr, fmt.Errorf("unknown symbol '%s'", string(str[i]))
			}
		}
	}
	if len(prevNum) != 0 {
		arr = append(arr, prevNum)
	}
	return arr, err
}

func calculate(firstNum, secondNum float64, operation string) (result float64) {
	switch operation {
	case plus:
		result = firstNum + secondNum
	case minus:
		result = firstNum - secondNum
	case multiple:
		result = firstNum * secondNum
	case division:
		result = firstNum / secondNum
	case degree:
		result = math.Pow(firstNum, secondNum)
	}

	return result
}

func doStackOperations(numStack []float64, operStack []string, currentOper string) ([]float64, []string) {
	operWeight := map[string]int{
		plus:     1,
		minus:    1,
		multiple: 2,
		division: 2,
		degree:   3,
	}

	for len(operStack) > 0 {
		lastOper := operStack[len(operStack)-1]
		if currentOper == closingBracket && lastOper == openBracket {
			operStack = operStack[:len(operStack)-1]
			break
		} else if (currentOper == closingBracket || operWeight[lastOper] >= operWeight[currentOper]) && len(numStack) >= 2 {
			firstNum := numStack[len(numStack)-2]
			secondNum := numStack[len(numStack)-1]

			currentResult := calculate(firstNum, secondNum, lastOper)

			operStack = operStack[:len(operStack)-1]
			numStack = append(numStack[:len(numStack)-2], currentResult)

			if len(operStack) > 1 {
				lastOper = operStack[len(operStack)-1]
			}

		} else {
			break
		}
	}

	if currentOper != closingBracket {
		operStack = append(operStack, currentOper)
	}

	return numStack, operStack
}

func calc(expressions []string) (result float64) {
	var numStack []float64
	var operStack []string

	for _, val := range expressions {
		if num, err := strconv.Atoi(val); err == nil {
			numStack = append(numStack, float64(num))
		} else if val == openBracket {
			operStack = append(operStack, val)
		} else {
			numStack, operStack = doStackOperations(numStack, operStack, val)
		}
	}

	if len(numStack) > 2 {
		res := calculate(numStack[len(numStack)-2], numStack[len(numStack)-1], operStack[len(operStack)-1])
		res = calculate(numStack[0], res, operStack[0])
		result = res
	} else if len(numStack) == 2 {
		result = calculate(numStack[len(numStack)-2], numStack[len(numStack)-1], operStack[len(operStack)-1])
	} else {
		result = numStack[0]
	}

	return result
}

func main() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()

	input := getInput()
	var validInput [][]string

	for _, val := range input {
		value, err := validate(val)
		if err != nil {
			panic(fmt.Errorf("%v", err))
		} else {
			validInput = append(validInput, value)
		}
	}
	for i, expression := range validInput {
		fmt.Println(input[i], " = ", calc(expression))
	}
}
