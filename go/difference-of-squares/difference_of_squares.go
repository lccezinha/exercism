package diffsquares

const testVersion = 1

// SquareOfSums exported
func SquareOfSums(number int) int {
	sum := 0

	for i := 1; i <= number; i++ {
		sum += i
	}

	return sum * sum
}

// SumOfSquares exported
func SumOfSquares(number int) int {
	sum := 0

	for i := 1; i <= number; i++ {
		sum += (i * i)
	}

	return sum
}

// Difference exported
func Difference(number int) int {
	return SquareOfSums(number) - SumOfSquares(number)
}
