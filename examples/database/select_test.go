package database

import "testing"

func TestSimpleQuery(t *testing.T) {
	// Arrange
	expectedResult := true

	// Act
	result, err := SimpleQuery()

	// Assert
	if err != nil {
		t.Fatalf("err = %s; want nil", err)
	}

	if result != expectedResult {
		t.Fatalf("result is %t; want %t", result, expectedResult)
	}
}
