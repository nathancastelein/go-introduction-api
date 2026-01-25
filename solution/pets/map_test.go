package pets

import "testing"

func TestCountByCategory(t *testing.T) {
	// Arrange
	categories := []string{"Dog", "Cat", "Dog", "Cat", "Bird"}

	// Act
	counts := countByCategory(categories)

	// Assert
	if len(counts) != 3 {
		t.Fatalf("Number of categories = %d; want 3", len(counts))
	}

	if counts["Dog"] != 2 {
		t.Errorf("Dog count = %d; want 2", counts["Dog"])
	}

	if counts["Cat"] != 2 {
		t.Errorf("Cat count = %d; want 2", counts["Cat"])
	}

	if counts["Bird"] != 1 {
		t.Errorf("Bird count = %d; want 1", counts["Bird"])
	}
}
