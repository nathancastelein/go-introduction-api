package pets

// categorizeAge returns the age category of a pet:
// - "Puppy" for ages 0-1
// - "Adult" for ages 2-7
// - "Senior" for ages 8+
func categorizeAge(age int) string {
	if age <= 1 {
		return "Puppy"
	}
	if age <= 7 {
		return "Adult"
	}
	return "Senior"
}
