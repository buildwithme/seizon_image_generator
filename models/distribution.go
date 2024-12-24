package models

import (
	"strconv"
	"strings"
)

// Distribution represents the distribution state of an entity.
type Distribution string

// Constants representing specific distribution states.
const (
	DistributionRevealed Distribution = "ALREADY REVEALED" // Indicates that the distribution is fully revealed.
)

// String returns the string representation of the Distribution.
func (d Distribution) String() string {
	return string(d)
}

// IsRevealed checks if the distribution is revealed (100%).
func (d Distribution) IsRevealed() bool {
	return d == DistributionRevealed
}

// IsUnrevealed checks if the distribution is not revealed (less than 100%).
func (d Distribution) IsUnrevealed() bool {
	return d != DistributionRevealed
}

// IsEmpty checks if the distribution is empty.
func (d Distribution) IsEmpty() bool {
	return d == ""
}

// IsValid checks if the distribution is valid.
// A valid distribution is either "ALREADY REVEALED", contains a percentage ("%"), or is empty.
func (d Distribution) IsValid() bool {
	return d == DistributionRevealed ||
		strings.Contains(d.String(), "%") ||
		d.String() == ""
}

// IsInvalid checks if the distribution is invalid by negating IsValid.
func (d Distribution) IsInvalid() bool {
	return !d.IsValid()
}

// GetPercentage parses and returns the percentage value of the distribution.
// If the distribution is invalid, it returns 0. If it's revealed, it returns 100.
func (d Distribution) GetPercentage() float64 {
	if !d.IsValid() {
		return 0 // Return 0 for invalid distributions
	}

	if d.IsRevealed() {
		return 100 // Return 100% for revealed distributions
	}

	// Remove the percentage symbol ("%") from the string
	value := strings.Replace(d.String(), "%", "", 1)

	if value == "" {
		return 0 // Return 0 if no value is found
	}

	// Convert the remaining string to a float
	parsefloat, err := strconv.ParseFloat(value, 64)
	if err != nil {
		panic(err) // Consider handling errors more gracefully
	}

	return parsefloat
}
