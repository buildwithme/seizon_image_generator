package models

import (
	"fmt"
	"log"
	"strings"
)

// Attribute represents a single attribute with a trait type and value.
type Attribute struct {
	TraitType string `json:"trait_type"` // The type of the trait (e.g., "Species", "Rarity").
	Value     string `json:"value"`      // The value of the trait (e.g., "Rare", "Dragon").
}

// APIResponse represents the response from an API for a specific token.
type APIResponse struct {
	Seed         string      `json:"seed,omitempty"`          // Optional seed value.
	TokenID      int         `json:"token_id"`                // Unique token identifier.
	Name         string      `json:"name"`                    // Name of the token.
	Description  string      `json:"description"`             // Description of the token.
	Image        string      `json:"image"`                   // URL to the image of the token.
	AnimationURL string      `json:"animation_url,omitempty"` // Optional animation URL.
	Attributes   []Attribute `json:"attributes"`              // List of attributes for the token.
}

// GetSpecie retrieves the "Species" attribute from the APIResponse.
func (a *APIResponse) GetSpecie() Specie {
	for _, attr := range a.Attributes {
		if attr.TraitType == "Species" {
			specie := Specie(strings.ToUpper(attr.Value)) // Convert value to uppercase.

			if specie.IsValid() { // Check if the species is valid.
				return specie
			}

			// Panic if an invalid species is encountered.
			panic(fmt.Errorf("invalid specie: %v", specie))
		}
	}

	// Log a warning if no species attribute is found.
	log.Println(fmt.Errorf("no specie found for token ID: %v", a.TokenID))

	return SpecieNone // Return a default value if not found.
}

// GetRarity retrieves the "Rarity" attribute from the APIResponse.
func (a *APIResponse) GetRarity() Rarity {
	for _, attr := range a.Attributes {
		if attr.TraitType == "Rarity" {
			rarity := Rarity(attr.Value)

			if rarity.IsValid() { // Check if the rarity is valid.
				return rarity
			}

			// Panic if an invalid rarity is encountered.
			panic(fmt.Errorf("invalid rarity: %v", rarity))
		}
	}

	// Panic if no rarity attribute is found (critical error).
	panic(fmt.Errorf("no rarity found for token ID: %v", a.TokenID))
}

// MakeAttributesUnique ensures that the attributes in the APIResponse are unique by both trait type and value.
func (a *APIResponse) MakeAttributesUnique() {
	data := make(map[string]map[string]struct{}) // Nested map to track unique attributes.

	// Populate the nested map with existing attributes.
	for _, x := range a.Attributes {
		if _, ok := data[x.TraitType]; !ok {
			data[x.TraitType] = make(map[string]struct{})
		}
		data[x.TraitType][x.Value] = struct{}{}
	}

	// Clear the original attributes slice.
	a.Attributes = []Attribute{}

	// Rebuild the attributes slice with unique values.
	for traitType, values := range data {
		for value := range values {
			a.Attributes = append(a.Attributes, Attribute{
				TraitType: traitType,
				Value:     value,
			})
		}
	}
}
