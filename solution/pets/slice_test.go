package pets

import "testing"

func TestFilterAdultAges(t *testing.T) {
	// Arrange
	ages := []int{0, 1, 2, 5, 7, 8, 12}

	// Act
	adultAges := filterAdultAges(ages)

	// Assert
	expected := []int{2, 5, 7}
	if len(adultAges) != len(expected) {
		t.Fatalf("Number of adult ages = %d; want %d", len(adultAges), len(expected))
	}

	for i, age := range expected {
		if adultAges[i] != age {
			t.Errorf("Adult age %d = %d; want %d", i, adultAges[i], age)
		}
	}
}
