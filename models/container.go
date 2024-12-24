package models

import (
	"strings"

	"github.com/samber/lo" // Utility library for filtering and other collection operations
)

// Traits represents all possible trait groups, categorized for a character or object.
type Traits struct {
	// Trait groups for various body parts, accessories, etc.
	Bodies        *Commons
	Tails         *Commons
	ElvenEars     *Commons
	Droplets      *Droplets
	Hands         *Commons
	Hairs         *Hairs
	Hats          *Hats
	StackableHats *StackableHats
	Mouths        *Commons
	Nose          *Commons
	Eyes          *Commons
	Glasses       *Commons
	Earrings      *Commons
	Clothes       *Commons
	Wings         *Commons
	Weapons       *Weapons
	Facegears     *Commons
	BG            *Commons
	BGAccent      *Commons
	Aura          *Aura
	// Default traits for specific genders
	DefaultMaleClothes        *Commons
	DefaultFemaleClothes      *Commons
	DefaultMaleMouths         *Commons
	DefaultFemaletMouths      *Commons
	DefaultMaleEyes           *Commons
	DefaultFemaleEyes         *Commons
	DefaultMaleHair           *Commons
	DefaultFemaleHair         *Commons
	DefaultMaleHat            *Hats
	DefaultFemaleHat          *Hats
	DefaultMaleStackableHat   *StackableHats
	DefaultFemaleStackableHat *StackableHats

	Final FinalTraits // Finalized traits with concrete selections
}

// FinalTraits represents the selected or finalized traits for a character.
type FinalTraits struct {
	Bodies        *Common
	Tails         *Common
	ElvenEars     *Common
	Droplets      DropletsSingle
	Hands         *Common
	Hairs         HairsSingle
	Hats          HatsSingle
	StackableHats StackableHatsSingle
	Mouths        *Common
	Nose          *Common
	Eyes          *Common
	Glasses       *Common
	Earrings      *Common
	Clothes       *Common
	Wings         *Common
	Weapons       WeaponsSingle
	Facegears     *Common
	BG            *Common
	BGAccent      *Common
	Aura          AuraSingle

	Metadata *APIResponse // Metadata related to the traits
	Category Category     // Trait category
	Gender   Gender       // Gender of the character
	Rarity   Rarity       // Rarity of the traits
	Specie   Specie       // Species of the character
	HasHair  bool         // Indicates if the character has hair
}

// ByGender filters traits by gender. If gender is unspecified, combines male and female traits.
func ByGender[T any](gender Gender, male, female []T) []T {
	switch gender {
	case GenderMale:
		return male
	case GenderFemale:
		return female
	default:
		return append(male, female...)
	}
}

// Copy creates a deep copy of a Traits object.
func (f Traits) Copy() *Traits {
	return &Traits{
		// Perform deep copies of all trait groups
		Bodies:    f.Bodies.Copy(),
		Tails:     f.Tails.Copy(),
		ElvenEars: f.ElvenEars.Copy(),
		Droplets: &Droplets{
			NA:                  f.Droplets.NA,
			Data:                append([]*Common{}, f.Droplets.Data...),
			DataBack:            append([]*Common{}, f.Droplets.DataBack...),
			DataBackTransparent: append([]*Common{}, f.Droplets.DataBackTransparent...),
		},
		// Similar deep copies for other traits...
		Final: f.Final.Copy(),
	}
}

// Copy creates a deep copy of a FinalTraits object.
func (f FinalTraits) Copy() FinalTraits {
	return FinalTraits{
		// Deep copy for all fields in FinalTraits
		Bodies:    f.Bodies.Copy(),
		Tails:     f.Tails.Copy(),
		ElvenEars: f.ElvenEars.Copy(),
		Droplets: DropletsSingle{
			DataFront:           f.Droplets.DataFront.Copy(),
			DataBack:            f.Droplets.DataBack.Copy(),
			DataBackTransparent: f.Droplets.DataBackTransparent.Copy(),
		},
		// Other fields copied similarly...
	}
}

// Filter types for default traits.
type Filter int

const (
	FilterGender Filter = iota
	FilterCategory
	FilterSpecie
)

// DefaultFilter filters Commons based on the provided filters and FinalTraits criteria.
func (f FinalTraits) DefaultFilter(data []*Common, filters ...Filter) []*Common {
	if len(data) == 0 {
		return nil
	}

	// Use lo.Filter to filter traits based on provided conditions.
	result := lo.Filter(data, func(common *Common, i int) bool {
		// Specie-specific filtering logic
		if f.Specie == SpecieOrigin && f.Rarity != COMMON {
			colors := strings.Split(f.Droplets.DataFront.OpenSeaTraitValue, " ")
			color := colors[1]
			if color != "Lavender" && color != "Teal" {
				return true
			}

			if !strings.Contains(strings.ToLower(common.OpenSeaTraitValue), strings.ToLower(color)) {
				return false
			}
		}

		// Gender-specific filtering
		if (len(filters) == 0 || lo.Contains(filters, FilterGender)) &&
			common.Gender != GenderUnisex && common.Gender != f.Gender {
			return false
		}

		// Category-specific filtering
		if (len(filters) == 0 || lo.Contains(filters, FilterCategory)) &&
			!lo.Contains(common.Category, f.Category) {
			return false
		}

		return true
	})

	// Apply species filtering if specified
	if len(filters) == 0 || lo.Contains(filters, FilterSpecie) {
		return f.FilterBySpecies(result)
	}

	return result
}

// FilterBySpecies filters Commons based on species constraints in FinalTraits.
func (f FinalTraits) FilterBySpecies(data []*Common) []*Common {
	specie := f.Specie

	return lo.Filter(data, func(common *Common, i int) bool {
		return lo.Contains(common.SpeciesLocked, specie) ||
			lo.Contains(common.SpeciesLocked, SpecieNone) &&
				specie != SpecieSoul ||
			common.RarityLocked.IsY() &&
				specie != SpecieOrigin &&
				specie != SpecieSoul
	})
}

// PopulateByGender organizes Commons by gender.
func PopulateByGender(data []*Common) map[Gender][]*Common {
	result := make(map[Gender][]*Common)

	for _, common := range data {
		result[common.Gender] = append(result[common.Gender], common)
	}

	return result
}
