package models

// Gender represents the gender of an entity.
type Gender string

// Constants representing valid gender values.
const (
	GenderNA     Gender = "NA" // Not applicable or undefined gender
	GenderMale   Gender = "M"  // Male gender
	GenderFemale Gender = "F"  // Female gender
	GenderUnisex Gender = "U"  // Unisex or applicable to both genders
)

// IsValid checks if the gender is one of the predefined valid values.
func (g Gender) IsValid() bool {
	switch g {
	case GenderNA, GenderMale, GenderFemale, GenderUnisex:
		return true
	default:
		return false
	}
}

// ToPercentage returns a predefined percentage based on the gender.
// Defaults to 0 for invalid or unrecognized gender values.
func (g Gender) ToPercentage() int {
	switch g {
	case GenderMale:
		return 82
	case GenderFemale:
		return 18
	case GenderUnisex:
		return 100
	default:
		return 0
	}
}

// IsInvalid checks if the gender is invalid by negating IsValid.
func (g Gender) IsInvalid() bool {
	return !g.IsValid()
}
