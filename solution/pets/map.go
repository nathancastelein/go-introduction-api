package pets

// countByCategory counts how many times each category appears in the slice
func countByCategory(categories []string) map[string]int {
	counts := make(map[string]int)

	for _, category := range categories {
		counts[category]++
	}
	return counts
}
