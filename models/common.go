package models

// Commons represents a collection of Common objects, along with a special NA value.
type Commons struct {
	Data []*Common // List of Common objects
	NA   *Common   // Special NA (Not Applicable) value
}

// Copy creates a deep copy of the Commons object.
func (c *Commons) Copy() *Commons {
	return &Commons{
		NA:   c.NA,                           // Copy reference to the NA field
		Data: append([]*Common{}, c.Data...), // Create a new slice with copied references
	}
}

// Common represents an individual item with multiple attributes.
type Common struct {
	FileName               string       // File name associated with the Common
	OpenSeaTraitValue      string       // Trait value used on OpenSea
	Category               []Category   // Categories associated with the Common
	Gender                 Gender       // Gender for this Common
	Combined               Combined     // Combined state for this Common
	MustNotInclude         []string     // List of strings that must not be included
	SpeciesLocked          []Specie     // Locked species for this Common
	Distribution           Distribution // Distribution information
	Notes                  string       // Notes or additional information
	MustInclude            []string     // List of strings that must be included
	RarityLocked           RarityLocked // Locked rarity information
	AbleToHaveStackableHat bool         // Indicates if stackable hats are allowed
	OnlyHaloAndHorns       bool         // Indicates if only halo and horns are allowed
}

// Copy creates a deep copy of a Common object.
func (c *Common) Copy() *Common {
	if c == nil {
		return nil
	}
	return &Common{
		FileName:               c.FileName,
		OpenSeaTraitValue:      c.OpenSeaTraitValue,
		Category:               append([]Category{}, c.Category...), // Create a new slice for categories
		Gender:                 c.Gender,
		Combined:               c.Combined,
		MustNotInclude:         append([]string{}, c.MustNotInclude...), // Create a new slice for MustNotInclude
		SpeciesLocked:          c.SpeciesLocked,
		Distribution:           c.Distribution,
		Notes:                  c.Notes,
		MustInclude:            append([]string{}, c.MustInclude...), // Create a new slice for MustInclude
		RarityLocked:           c.RarityLocked,
		AbleToHaveStackableHat: c.AbleToHaveStackableHat,
		OnlyHaloAndHorns:       c.OnlyHaloAndHorns,
	}
}

// Aura represents collections of Commons for normal and front layers, along with an NA value.
type Aura struct {
	Normal []*Common // Normal aura
	Front  []*Common // Front aura
	NA     *Common   // Not Applicable aura
}

// Hairs represents collections of Commons for hair and back hair, along with an NA value.
type Hairs struct {
	Hair     []*Common // Main hair
	HairBack []*Common // Back hair
	NA       *Common   // Not Applicable hair
}

// Hats represents collections of Commons for hats, earless hats, and their NA values.
type Hats struct {
	Data        []*Common // Hat data
	DataEarless []*Common // Earless hat data
	NA          *Common   // Not Applicable hat
	NAEarless   *Common   // Not Applicable earless hat
}

// StackableHats represents collections of Commons for stackable hats and their backs, along with an NA value.
type StackableHats struct {
	Data     []*Common // Stackable hats
	DataBack []*Common // Back of stackable hats
	NA       *Common   // Not Applicable stackable hat
}

// Droplets represents collections of Commons for droplets in different layers.
type Droplets struct {
	Data                []*Common // Main droplets
	DataBack            []*Common // Back droplets
	DataBackTransparent []*Common // Transparent back droplets
	NA                  *Common   // Not Applicable droplet
}

// Weapons represents collections of Commons for front and back weapons, along with an NA value.
type Weapons struct {
	Front []*Common // Front weapons
	Back  []*Common // Back weapons
	NA    *Common   // Not Applicable weapon
}

// Single item variants of the above types:

// AuraSingle represents a single aura with front and back components.
type AuraSingle struct {
	Back  *Common // Back aura
	Front *Common // Front aura
}

// HairsSingle represents a single hair item with main and back components.
type HairsSingle struct {
	Hair     *Common // Main hair
	HairBack *Common // Back hair
}

// HatsSingle represents a single hat item with earless variants.
type HatsSingle struct {
	Data        *Common // Hat data
	DataEarless *Common // Earless hat data
}

// StackableHatsSingle represents a single stackable hat with front and back components.
type StackableHatsSingle struct {
	DataFront *Common // Front of stackable hat
	DataBack  *Common // Back of stackable hat
}

// DropletsSingle represents a single droplet with front, back, and transparent back components.
type DropletsSingle struct {
	DataFront           *Common // Front droplet
	DataBack            *Common // Back droplet
	DataBackTransparent *Common // Transparent back droplet
}

// WeaponsSingle represents a single weapon with front and back components.
type WeaponsSingle struct {
	Front *Common // Front weapon
	Back  *Common // Back weapon
}
