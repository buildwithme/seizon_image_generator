package models

// RarityLocked represents whether a rarity is locked, with specific states.
type RarityLocked string

// Constants representing valid RarityLocked values.
const (
	RarityLockedNA   RarityLocked = "NA" // Not applicable or undefined rarity lock
	RarityLockedNone RarityLocked = ""   // No rarity lock
	RarityLockedY    RarityLocked = "Y"  // Rarity is locked
)

// String converts the RarityLocked value to its string representation.
func (rl RarityLocked) String() string {
	return string(rl)
}

// IsY checks if the rarity is locked (value is "Y").
func (rl RarityLocked) IsY() bool {
	return rl == RarityLockedY
}

// IsInvalid checks if the RarityLocked value is invalid by negating IsValid.
func (rl RarityLocked) IsInvalid() bool {
	return !rl.IsValid()
}

// IsValid checks if the RarityLocked value is one of the predefined valid values.
func (rl RarityLocked) IsValid() bool {
	return rl == RarityLockedNA || rl == RarityLockedNone || rl == RarityLockedY
}
