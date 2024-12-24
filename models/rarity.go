package models

// Rarity represents the rarity of an item, with predefined values.
type Rarity string

// Constants representing valid rarity values.
const (
	ONE_OF_ONE       Rarity = "1 Of 1"           // Unique rarity
	UNKNOWN_COLOR1   Rarity = "?1"               // Placeholder for unknown color 1
	UNKNOWN_COLOR2   Rarity = "?2"               // Placeholder for unknown color 2
	UNKNOWN_COLOR3   Rarity = "?3"               // Placeholder for unknown color 3
	COMMON           Rarity = "Common"           // Common rarity
	RARE_PURPLE      Rarity = "Rare Purple"      // Rare purple rarity
	RARE_ORANGE      Rarity = "Rare Orange"      // Rare orange rarity
	MYTHIC_LAVENDER  Rarity = "Mythic Lavender"  // Mythic lavender rarity
	RARE_YELLOW      Rarity = "Rare Yellow"      // Rare yellow rarity
	ULTRA_BLUE       Rarity = "Ultra Blue"       // Ultra blue rarity
	ULTRA_GREEN      Rarity = "Ultra Green"      // Ultra green rarity
	LEGENDARY_SILVER Rarity = "Legendary Silver" // Legendary silver rarity
	ULTRA_PINK       Rarity = "Ultra Pink"       // Ultra pink rarity
	LEGENDARY_VIOLET Rarity = "Legendary Violet" // Legendary violet rarity
	LEGENDARY_BLUE   Rarity = "Legendary Blue"   // Legendary blue rarity
	MYTHIC_TEAL      Rarity = "Mythic Teal"      // Mythic teal rarity
)

// IsValid checks if the Rarity is one of the predefined valid values.
func (s Rarity) IsValid() bool {
	switch s {
	case RARE_PURPLE, RARE_ORANGE, MYTHIC_LAVENDER, UNKNOWN_COLOR1, RARE_YELLOW,
		ONE_OF_ONE, COMMON, ULTRA_BLUE, ULTRA_GREEN, LEGENDARY_SILVER,
		UNKNOWN_COLOR2, ULTRA_PINK, LEGENDARY_VIOLET, LEGENDARY_BLUE, MYTHIC_TEAL, UNKNOWN_COLOR3:
		return true
	default:
		return false
	}
}

// IsInvalid checks if the Rarity is invalid by negating IsValid.
func (s Rarity) IsInvalid() bool {
	return !s.IsValid()
}

// String converts the Rarity value to its string representation.
func (s Rarity) String() string {
	return string(s)
}
