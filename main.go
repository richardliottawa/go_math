package go_math

import (
	"errors"
)

func Add(args ...int) int {
	result := 0
	for _, val := range args {
		result += val
	}
	return result
}

func Sub(a, b int) int {
	return a - b
}

func Mul(a, b int) int {
	return a * b
}

func Div(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("Divide by zero operation is prohibited.")
	}
	return a / b, nil
}

func Pow(a, b int) int {
	if b == 1 {
		return a
	}

	if b == 0 {
		return 1
	}

	if b&1 == 0 {
		half := Pow(a, b/2)
		return half * half
	}
	return Pow(a, b-1) * a
}
