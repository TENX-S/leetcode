package Stack

import "testing"

func Test_evalRPN(t *testing.T) {
	t.Log(evalRPN([]string{"2", "1", "+", "3", "*"}))
}
