package models

// Category defines a custom type for categorizing entities.
type Category string

// Constants representing valid categories.
const (
	CategoryNA     Category = "NA"     // Not applicable or undefined category
	CategoryCool   Category = "COOL"   // Cool category
	CategorySpooky Category = "SPOOKY" // Spooky category
	CategoryGoofy  Category = "GOOFY"  // Goofy category
	CategoryCute   Category = "CUTE"   // Cute category
)

// IsValid checks if the category is one of the defined valid categories.
func (c Category) IsValid() bool {
	switch c {
	case CategoryNA, CategoryCool, CategorySpooky, CategoryGoofy, CategoryCute:
		return true
	default:
		return false
	}
}

// CategoryList returns a slice of all valid categories except "NA".
func CategoryList() []Category {
	return []Category{CategoryCool, CategorySpooky, CategoryGoofy, CategoryCute}
}

// ToPercentage maps a category to its associated percentage value.
// Default is 0 if the category is invalid.
func (g Category) ToPercentage() int {
	switch g {
	case CategoryCool:
		return 50
	case CategorySpooky:
		return 20
	case CategoryGoofy:
		return 10
	case CategoryCute:
		return 20
	default:
		return 0
	}
}

// String converts the category to its string representation.
func (c Category) String() string {
	return string(c)
}

// IsEmpty checks if the category is an empty string.
func (c Category) IsEmpty() bool {
	return c == ""
}

// IsInvalid checks if the category is invalid by negating the result of IsValid.
func (c Category) IsInvalid() bool {
	return !c.IsValid()
}
