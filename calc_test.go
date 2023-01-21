package main

import (
	"math"
	"testing"
)

func TestValidate(t *testing.T) {
	t.Run("test without errors", func(t *testing.T) {
		str := "(2)^2+2/(3*4)"
		result := []string{"(", "2", ")", "^", "2", "+", "2", "/", "(", "3", "*", "4", ")"}
		realResult, err := validate(str)
		if len(result) != len(realResult) || err != nil {
			t.Errorf("\nExpected\n    %s\nto equal\n    %s", result, realResult)
		}

		for i := 0; i < len(result); i++ {
			if realResult[i] != result[i] {
				t.Errorf("\nExpected\n    %s\nto equal\n    %s", result, realResult)
			}
		}
	})
	t.Run("test with unknown symbol", func(t *testing.T) {
		str := "(2)^2+2/(3*4)n"
		_, err := validate(str)
		if err == nil {
			t.Errorf("Expected error != nil, but not given")
		}
	})
	t.Run("test with unknown symbol", func(t *testing.T) {
		str := "(2)^2+2/(3*4))"
		_, err := validate(str)
		if err == nil {
			t.Errorf("Expected error != nil, but not given")
		}
	})
}

func TestCalculate(t *testing.T) {
	t.Run("test plus", func(t *testing.T) {
		result := float64(3) + float64(5)
		realResult := calculate(3, 5, "+")

		if realResult != result {
			t.Errorf("\nExpected\n    %f\nto equal\n    %f", result, realResult)
		}
	})
	t.Run("test minus", func(t *testing.T) {
		result := float64(3) - float64(5)
		realResult := calculate(3, 5, "-")

		if realResult != result {
			t.Errorf("\nExpected\n    %f\nto equal\n    %f", result, realResult)
		}
	})
	t.Run("test multiple", func(t *testing.T) {
		result := float64(3) * float64(5)
		realResult := calculate(3, 5, "*")

		if realResult != result {
			t.Errorf("\nExpected\n    %f\nto equal\n    %f", result, realResult)
		}
	})
	t.Run("test division", func(t *testing.T) {
		result := float64(3) / float64(5)
		realResult := calculate(3, 5, "/")

		if realResult != result {
			t.Errorf("\nExpected\n    %f\nto equal\n    %f", result, realResult)
		}
	})
	t.Run("test degree", func(t *testing.T) {
		result := math.Pow(float64(3), float64(5))
		realResult := calculate(3, 5, "^")

		if realResult != result {
			t.Errorf("\nExpected\n    %f\nto equal\n    %f", result, realResult)
		}
	})
}

func TestDoStackOperations(t *testing.T) {
	t.Run("test ())", func(t *testing.T) {
		numStack := []float64{1, 2, 5, 1, 2}
		operStack := []string{"+", "(", "*", "(", "-", "(", "+"}

		realResNum, realResOper := doStackOperations(numStack, operStack, ")")

		numStack = []float64{1, 2, 5, 3}
		operStack = []string{"+", "(", "*", "(", "-"}

		if len(numStack) != len(realResNum) || len(operStack) != len(realResOper) {
			t.Errorf("\nExpected\n    %v, %v\nto equal\n    %v, %v", numStack, operStack, realResNum, realResOper)
		}

		for i := 0; i < len(numStack); i++ {
			if numStack[i] != realResNum[i] {
				t.Errorf("\nExpected\n    %v, %v\nto equal\n    %v, %v", numStack, operStack, realResNum, realResOper)
			}
		}

		for i := 0; i < len(operStack); i++ {
			if operStack[i] != realResOper[i] {
				t.Errorf("\nExpected\n    %v, %v\nto equal\n    %v, %v", numStack, operStack, realResNum, realResOper)
			}
		}
	})
}

func TestCalc(t *testing.T) {
	t.Run("test get answer", func(t *testing.T) {
		expression := []string{"(", "2", ")", "^", "2", "+", "2", "/", "(", "3", "*", "4", ")"}
		result := math.Pow(float64(2), float64(2)) + float64(2)/float64(3*4)

		realResult := calc(expression)

		if realResult != result {
			t.Errorf("\nExpected\n    %f\nto equal\n    %f", result, realResult)
		}
	})

}