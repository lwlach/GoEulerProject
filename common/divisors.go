package common

func SumOfDivisors(n int) int {
	var sum int
	for i := 1; i <= n/2; i++ {
		if n%i == 0 {
			sum += i
		}
	}
	return sum
}
