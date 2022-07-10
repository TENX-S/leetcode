package CPC

import (
	"strconv"
	"strings"
)

func InversePolish(expr string) int {
	var stack []int

	pieces := strings.Split(expr, " ")
	for _, p := range pieces {
		switch p {
		case "+":
			a := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			b := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			stack = append(stack, a+b)
		case "-":
			a := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			b := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			stack = append(stack, b-a)
		case "*":
			a := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			b := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			stack = append(stack, a*b)
		case "/":
			a := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			b := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			stack = append(stack, b/a)
		default:
			val, _ := strconv.Atoi(p)
			stack = append(stack, val)
		}
	}
	return stack[0]
}
