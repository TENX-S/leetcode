package Stack

import "strconv"

type Stack []int

func (s *Stack) Pop() int {
	res := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return res
}

func (s *Stack) Push(x int) {
	*s = append(*s, x)
}

func evalRPN(tokens []string) int {
	var stack Stack
	for _, t := range tokens {
		switch t {
		case "+":
			a := stack.Pop()
			b := stack.Pop()
			stack.Push(b + a)
		case "-":
			a := stack.Pop()
			b := stack.Pop()
			stack.Push(b - a)
		case "*":
			a := stack.Pop()
			b := stack.Pop()
			stack.Push(b * a)
		case "/":
			a := stack.Pop()
			b := stack.Pop()
			stack.Push(b / a)
		default:
			val, _ := strconv.Atoi(t)
			stack.Push(val)
		}
	}
	return stack[0]
}
