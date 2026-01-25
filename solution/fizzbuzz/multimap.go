package fizzbuzz

func multiFizzBuzzMap(numbers []int) map[int]string {
	results := make(map[int]string)

	for _, number := range numbers {
		results[number] = fizzbuzz(number)
	}
	return results
}
