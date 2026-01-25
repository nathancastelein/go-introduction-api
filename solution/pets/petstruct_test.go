package pets

import "testing"

func TestPetDescribe(t *testing.T) {
	// Arrange
	pet := NewPetInfo("Buddy", "Dog", 3)

	// Act
	description := pet.Describe()

	// Assert
	expected := "Buddy is a 3 year old Dog"
	if description != expected {
		t.Errorf("Describe() = %s; want %s", description, expected)
	}
}
