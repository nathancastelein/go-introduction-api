package pets

import "testing"

func TestCategorizeAge(t *testing.T) {
	t.Run("should return Puppy for age 0", func(t *testing.T) {
		// Arrange
		age := 0

		// Act
		result := categorizeAge(age)

		// Assert
		if result != "Puppy" {
			t.Errorf("categorizeAge(0) = %s; want Puppy", result)
		}
	})

	t.Run("should return Puppy for age 1", func(t *testing.T) {
		// Arrange
		age := 1

		// Act
		result := categorizeAge(age)

		// Assert
		if result != "Puppy" {
			t.Errorf("categorizeAge(1) = %s; want Puppy", result)
		}
	})

	t.Run("should return Adult for age 2", func(t *testing.T) {
		// Arrange
		age := 2

		// Act
		result := categorizeAge(age)

		// Assert
		if result != "Adult" {
			t.Errorf("categorizeAge(2) = %s; want Adult", result)
		}
	})

	t.Run("should return Adult for age 7", func(t *testing.T) {
		// Arrange
		age := 7

		// Act
		result := categorizeAge(age)

		// Assert
		if result != "Adult" {
			t.Errorf("categorizeAge(7) = %s; want Adult", result)
		}
	})

	t.Run("should return Senior for age 8", func(t *testing.T) {
		// Arrange
		age := 8

		// Act
		result := categorizeAge(age)

		// Assert
		if result != "Senior" {
			t.Errorf("categorizeAge(8) = %s; want Senior", result)
		}
	})

	t.Run("should return Senior for age 15", func(t *testing.T) {
		// Arrange
		age := 15

		// Act
		result := categorizeAge(age)

		// Assert
		if result != "Senior" {
			t.Errorf("categorizeAge(15) = %s; want Senior", result)
		}
	})
}
