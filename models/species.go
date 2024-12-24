package models

// Specie represents the species of an entity, with predefined values.
type Specie string

// Constants defining valid species.
const (
	SpecieNone   Specie = ""       // No specific species
	SpecieNA     Specie = "NA"     // Not applicable or undefined species
	SpecieBeing  Specie = "BEING"  // General being
	SpecieCyborg Specie = "CYBORG" // Cyborg species
	SpecieElven  Specie = "ELVEN"  // Elven species
	SpecieFeline Specie = "FELINE" // Feline species
	SpecieMonkey Specie = "MONKEY" // Monkey species
	SpecieOrigin Specie = "ORIGIN" // Origin species
	SpecieSoul   Specie = "SOUL"   // Soul species
)

// IsValid checks if the Specie is one of the predefined valid values.
func (s Specie) IsValid() bool {
	switch s {
	case SpecieNA, SpecieNone, SpecieBeing, SpecieCyborg, SpecieElven, SpecieFeline, SpecieMonkey, SpecieOrigin, SpecieSoul:
		return true
	default:
		return false
	}
}

// IsInvalid checks if the Specie is invalid by negating IsValid.
func (s Specie) IsInvalid() bool {
	return !s.IsValid()
}

// String returns the string representation of the Specie.
func (s Specie) String() string {
	return string(s)
}
