package fizzbuzz

import "testing"

func TestFizzbuzz(t *testing.T) {
	t.Run("should return same number when it is not multiple of three or five", func(t *testing.T) {
		// Arrange
		number := 1

		// Act
		result := fizzbuzz(number)

		// Assert
		if result != "1" {
			t.Errorf("FizzBuzz(1) = %s; want 1", result)
		}
	})

	t.Run("should return Fizz when number is multiple of three", func(t *testing.T) {
		// Arrange
		number := 3

		// Act
		result := fizzbuzz(number)

		// Assert
		if result != "Fizz" {
			t.Errorf("FizzBuzz(3) = %s; want Fizz", result)
		}
	})

	t.Run("should return Buzz when number is multiple of five", func(t *testing.T) {
		// Arrange
		number := 5

		// Act
		result := fizzbuzz(number)

		// Assert
		if result != "Buzz" {
			t.Errorf("FizzBuzz(5) = %s; want Buzz", result)
		}
	})

	t.Run("should return FizzBuzz when number is multiple of three and five", func(t *testing.T) {
		// Arrange
		number := 15

		// Act
		result := fizzbuzz(number)

		// Assert
		if result != "FizzBuzz" {
			t.Errorf("FizzBuzz(15) = %s; want FizzBuzz", result)
		}
	})
}
