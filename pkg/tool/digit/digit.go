package digit

// Count - counts number of digits in number
func Count(number int) int {
	count := 0
	for number != 0 {
		number /= 10
		count += 1
	}
	return count
}