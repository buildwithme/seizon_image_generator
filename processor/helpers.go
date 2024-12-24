package processor

import (
	"generator/models"
	"strings"
)

// ExtractByFileName searches for a *models.Common object in the given slice
// where the FileName field matches the specified fileName.
// Returns the matching object or nil if not found.
func ExtractByFileName(data []*models.Common, fileName string) *models.Common {
	for _, v := range data {
		if v.FileName == fileName {
			return v
		}
	}
	return nil
}

// ExtractByTraitValue searches for a *models.Common object in the given slice
// where the OpenSeaTraitValue field matches the specified traitValue.
// Returns the matching object, or nil if no match is found.
func ExtractByTraitValue(data []*models.Common, traitValue string) *models.Common {
	common := OptionalExtractByTraitValue(data, traitValue)
	if common != nil {
		return common
	}

	return nil
}

// OptionalExtractByTraitValue searches for a *models.Common object in the given slice
// where the OpenSeaTraitValue field matches the specified traitValue.
// Returns the matching object or nil if not found.
func OptionalExtractByTraitValue(data []*models.Common, traitValue string) *models.Common {
	for _, v := range data {
		if v.OpenSeaTraitValue == traitValue {
			return v
		}
	}
	return nil
}

// OptionalExtractByTraitValueContains searches for a *models.Common object in the given slice
// where the OpenSeaTraitValue field contains the specified substring traitValue.
// Returns the matching object or nil if no match is found.
func OptionalExtractByTraitValueContains(data []*models.Common, traitValue string) *models.Common {
	for _, v := range data {
		if strings.Contains(v.OpenSeaTraitValue, traitValue) {
			return v
		}
	}
	return nil
}
