package C299

import (
	"math"
	"math/big"
)

func countHousePlacements(n int) int {
	num := big.NewInt(0)
	num.Mul(fib(big.NewInt(int64(n))), fib(big.NewInt(int64(n))))
	res := big.NewInt(0)
	res.Mod(num, big.NewInt(int64(math.Pow(10, 9))+7))
	return int(res.Int64())
}

func fib(n *big.Int) *big.Int {

	f2 := big.NewInt(2)
	f1 := big.NewInt(3)

	if n.Cmp(big.NewInt(1)) == 0 {
		return f2
	}

	if n.Cmp(big.NewInt(2)) == 0 {
		return f1
	}

	for i := 3; n.Cmp(big.NewInt(int64(i))) >= 0; i++ {
		next := big.NewInt(0)
		next.Add(f2, f1)
		f2 = f1
		f1 = next
	}

	return f1
}
