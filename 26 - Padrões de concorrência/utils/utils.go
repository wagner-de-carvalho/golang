package utils

func Fibonacci(posicao int) int {
	if posicao <= 1 {
		return posicao
	}

	return Fibonacci(posicao - 2) + Fibonacci(posicao - 1)
}