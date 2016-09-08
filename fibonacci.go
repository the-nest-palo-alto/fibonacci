package fibonacci

/*
Fibonacci ...
*/
func Fibonacci(count int) int {
	p, q := 0, 1
	for index := 0; index < count; index++ {
		p, q = q, p+q
	}
	return q
}
