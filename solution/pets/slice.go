package pets

// filterAdultAges returns only ages that are in the "Adult" range
func filterAdultAges(ages []int) []int {
	var results []int

	for _, age := range ages {
		if categorizeAge(age) == "Adult" {
			results = append(results, age)
		}
	}
	return results
}
