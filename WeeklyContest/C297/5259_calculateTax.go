package C297

func calculateTax(brackets [][]int, income int) float64 {
	if income == 0 {
		return 0
	}
	salary := float64(income)
	var tax float64

	for i, b := range brackets {
		upper, percent := float64(b[0]), float64(b[1])/100.0
		if i == 0 {
			if salary < upper {
				return salary * percent
			}
			tax += upper * percent
			continue
		}
		if salary <= upper {
			tax += (salary - float64(brackets[i-1][0])) * percent
			return tax
		}
		tax += (upper - float64(brackets[i-1][0])) * percent
	}
	return tax
}
