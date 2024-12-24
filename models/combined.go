package models

// Combined represents a three-state value: "", "NA", or "YES".
type Combined string

// Constants defining valid Combined values.
const (
	CombinedNo  Combined = ""    // Represents "No" or an empty state.
	CombinedNA  Combined = "NA"  // Represents "Not Applicable".
	CombinedYes Combined = "YES" // Represents "Yes".
)

// IsValid checks if the Combined value is one of the predefined valid states.
func (c Combined) IsValid() bool {
	switch c {
	case CombinedNA, CombinedNo, CombinedYes:
		return true
	default:
		return false
	}
}

// String converts the Combined value to its string representation.
func (c Combined) String() string {
	return string(c)
}

// Bool converts the Combined value to a boolean.
// Returns true only if the value is CombinedYes.
func (c Combined) Bool() bool {
	return c == CombinedYes
}

// IsInvalid checks if the Combined value is invalid by negating IsValid.
func (c Combined) IsInvalid() bool {
	return !c.IsValid()
}
