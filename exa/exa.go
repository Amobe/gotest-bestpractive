package exa

func add(a, b int) int {
	return a + b
}

func mul(a, b int) int {
	return a * b
}

func Times(a, times int) int {
	res := 0
	for i := 0; i < times; i++ {
		res = add(a, res)
	}
	return res
}

func Pow(a, pow int) int {
	res := 1
	for i := 0; i < pow; i++ {
		res = mul(a, res)
	}
	return res
}
