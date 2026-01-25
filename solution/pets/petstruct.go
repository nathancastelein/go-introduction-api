package pets

import "fmt"

// PetInfo represents a pet with additional information
type PetInfo struct {
	Name     string
	Category string
	Age      int
}

// NewPetInfo creates a new PetInfo instance
func NewPetInfo(name, category string, age int) *PetInfo {
	return &PetInfo{
		Name:     name,
		Category: category,
		Age:      age,
	}
}

// Describe returns a formatted description of the pet
// Format: "{Name} is a {Age} year old {Category}"
func (p *PetInfo) Describe() string {
	return fmt.Sprintf("%s is a %d year old %s", p.Name, p.Age, p.Category)
}
