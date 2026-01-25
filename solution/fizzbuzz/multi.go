package fizzbuzz

func multiFizzBuzz(numbers []int) []string {
	results := make([]string, len(numbers))

	for i, number := range numbers {
		results[i] = fizzbuzz(number)
	}
	return results
}
