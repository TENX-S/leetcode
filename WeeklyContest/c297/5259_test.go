package c297

import "testing"

func TestCalculateTax(t *testing.T) {
	t.Log(calculateTax([][]int{{10, 10}}, 5))
}
