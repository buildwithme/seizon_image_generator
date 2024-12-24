package parse

import (
	"generator/models"

	"github.com/tealeg/xlsx"
)

// Common processor function for Commons models.
// Adds data to the result's Data slice unless the FileName is "NA",
// in which case the NA field is set.
var commonProcessor = func(result *models.Commons, data *models.Common, rowNumber int) {
	if data.FileName == "NA" {
		result.NA = data
		return
	}
	result.Data = append(result.Data, data)
}

// Parses the Droplets sheet and categorizes data into Data, DataBack, or DataBackTransparent slices
// based on the row number. Skips specific rows and stops at a defined limit.
func GetDroplets(sheet *xlsx.Sheet) *models.Droplets {
	skip := []int{1, 2, 15, 16, 29, 30, 31}
	stop := 44

	return parse(sheet, skip, stop, func(result *models.Droplets, data *models.Common, rowNumber int) {
		if data.FileName == "NA" {
			result.NA = data
			return
		}
		if rowNumber < 15 {
			result.Data = append(result.Data, data)
		} else if rowNumber < 29 {
			result.DataBack = append(result.DataBack, data)
		} else {
			result.DataBackTransparent = append(result.DataBackTransparent, data)
		}
	})
}

// Parses the Stackable Hats sheet and categorizes data into Data or DataBack slices
// based on the row number. Skips specific rows and stops at a defined limit.
func GetStackableHats(sheet *xlsx.Sheet) *models.StackableHats {
	skip := []int{1, 2, 21, 22, 23, 24}
	stop := 30

	return parse(sheet, skip, stop, func(result *models.StackableHats, data *models.Common, rowNumber int) {
		if data.FileName == "NA" {
			result.NA = data
			return
		}
		if rowNumber < 25 {
			result.Data = append(result.Data, data)
		} else {
			result.DataBack = append(result.DataBack, data)
		}
	})
}

// Parses the Weapons sheet and categorizes data into Front or Back slices
// based on the row number. Skips specific rows and stops at a defined limit.
func GetWeapons(sheet *xlsx.Sheet) *models.Weapons {
	skip := []int{1, 2, 27, 28, 29, 30}
	stop := 52

	return parse(sheet, skip, stop, func(result *models.Weapons, data *models.Common, rowNumber int) {
		if data.FileName == "NA" {
			result.NA = data
			return
		}
		if rowNumber < 30 {
			result.Front = append(result.Front, data)
		} else {
			result.Back = append(result.Back, data)
		}
	})
}

// Parses the Hats sheet, categorizing data into Data or DataEarless slices, or setting
// NA or NAEarless fields based on the FileName and row number.
// Skips specific rows and stops at a defined limit.
func GetHats(sheet *xlsx.Sheet) *models.Hats {
	skip := []int{1, 2, 68, 69}
	stop := 87

	return parse(sheet, skip, stop, func(result *models.Hats, data *models.Common, rowNumber int) {
		if data.FileName == "NA" {
			if rowNumber < 70 {
				result.NA = data
			} else {
				result.NAEarless = data
			}
			return
		}
		if rowNumber < 70 {
			result.Data = append(result.Data, data)
		} else {
			result.DataEarless = append(result.DataEarless, data)
		}
	})
}

// Parses the Aura sheet, categorizing data into Normal or Front slices
// based on the row number. Skips specific rows and stops at a defined limit.
func GetAura(sheet *xlsx.Sheet) *models.Aura {
	skip := []int{1, 2, 35, 36, 37}
	stop := 43

	return parse(sheet, skip, stop, func(result *models.Aura, data *models.Common, rowNumber int) {
		if data.FileName == "NA" {
			result.NA = data
			return
		}
		if rowNumber < 37 {
			result.Normal = append(result.Normal, data)
		} else {
			result.Front = append(result.Front, data)
		}
	})
}

// Parses a Commons sheet for BG, applying the common processor function.
// Skips specific rows and stops at a defined limit.
func GetBG(sheet *xlsx.Sheet) *models.Commons {
	skip := []int{1, 2}
	stop := 17

	return parse(sheet, skip, stop, commonProcessor)
}

// Parses a Commons sheet for BGAccents, applying the common processor function.
// Skips specific rows and stops at a defined limit.
func GetBGAccents(sheet *xlsx.Sheet) *models.Commons {
	skip := []int{1, 2}
	stop := 13

	return parse(sheet, skip, stop, commonProcessor)
}

// Parses a Commons sheet for FaceGears, applying the common processor function.
// Skips specific rows and stops at a defined limit.
func GetFaceGears(sheet *xlsx.Sheet) *models.Commons {
	skip := []int{1, 2}
	stop := 49

	return parse(sheet, skip, stop, commonProcessor)
}

// Parses a Commons sheet for Wings, applying the common processor function.
// Skips specific rows and stops at a defined limit.
func GetWings(sheet *xlsx.Sheet) *models.Commons {
	skip := []int{1, 2}
	stop := 33

	return parse(sheet, skip, stop, commonProcessor)
}

// Parses a Commons sheet for Clothes, applying the common processor function.
// Skips specific rows and stops at a defined limit.
func GetClothes(sheet *xlsx.Sheet) *models.Commons {
	skip := []int{1, 2}
	stop := 149

	return parse(sheet, skip, stop, commonProcessor)
}

// Parses a Commons sheet for Earrings, applying the common processor function.
// Skips specific rows and stops at a defined limit.
func GetEarrings(sheet *xlsx.Sheet) *models.Commons {
	skip := []int{1, 2}
	stop := 24

	return parse(sheet, skip, stop, commonProcessor)
}

// Parses a Commons sheet for Glasses, applying the common processor function.
// Skips specific rows and stops at a defined limit.
func GetGlasses(sheet *xlsx.Sheet) *models.Commons {
	skip := []int{1, 2}
	stop := 27

	return parse(sheet, skip, stop, commonProcessor)
}

// Parses a Commons sheet for Eyes, applying the common processor function.
// Skips specific rows and stops at a defined limit.
func GetEyes(sheet *xlsx.Sheet) *models.Commons {
	skip := []int{1, 2}
	stop := 211

	return parse(sheet, skip, stop, commonProcessor)
}

// Parses a Commons sheet for Mouths, applying the common processor function.
// Skips specific rows and stops at a defined limit.
func GetMouths(sheet *xlsx.Sheet) *models.Commons {
	skip := []int{1, 2}
	stop := 115

	return parse(sheet, skip, stop, commonProcessor)
}

// Parses a Commons sheet for Noses, applying the common processor function.
// Skips specific rows and stops at a defined limit.
func GetNoses(sheet *xlsx.Sheet) *models.Commons {
	skip := []int{1, 2}
	stop := 18

	return parse(sheet, skip, stop, commonProcessor)
}

// Parses a Commons sheet for Tails, applying the common processor function.
// Skips specific rows and stops at a defined limit.
func GetTails(sheet *xlsx.Sheet) *models.Commons {
	skip := []int{1, 2}
	stop := 8

	return parse(sheet, skip, stop, commonProcessor)
}

// Parses a Commons sheet for Bodies, applying the common processor function.
// Skips specific rows and stops at a defined limit.
func GetBodies(sheet *xlsx.Sheet) *models.Commons {
	skip := []int{1, 2}
	stop := 40

	return parse(sheet, skip, stop, commonProcessor)
}

// Parses a Commons sheet for ElderEars, applying the common processor function.
// Skips specific rows and stops at a defined limit.
func GetElderEars(sheet *xlsx.Sheet) *models.Commons {
	skip := []int{1, 2}
	stop := 10

	return parse(sheet, skip, stop, commonProcessor)
}

// Parses a Commons sheet for Hands, applying the common processor function.
// Skips specific rows and stops at a defined limit.
func GetHands(sheet *xlsx.Sheet) *models.Commons {
	skip := []int{1, 2}
	stop := 33

	return parse(sheet, skip, stop, commonProcessor)
}

// Parses the Hairs sheet, categorizing data into Hair or HairBack slices
// based on the row number. Skips specific rows and stops at a defined limit.
func GetHairs(sheet *xlsx.Sheet) *models.Hairs {
	skip := []int{1, 2, 111, 112, 113}
	stop := 127

	return parse(sheet, skip, stop, func(result *models.Hairs, data *models.Common, rowNumber int) {
		if data.FileName == "NA" {
			result.NA = data
			return
		}
		if rowNumber < 113 {
			result.Hair = append(result.Hair, data)
		} else {
			result.HairBack = append(result.HairBack, data)
		}
	})
}

// Parses a Commons sheet for DefaultMaleClothes, applying the common processor function.
// Skips specific rows and stops at a defined limit.
func GetDefaultMaleClothes(sheet *xlsx.Sheet) *models.Commons {
	skip := []int{1, 2}
	stop := 55

	return parse(sheet, skip, stop, commonProcessor)
}

// Parses a Commons sheet for DefaultFemaleClothes, applying the common processor function.
// Skips specific rows and stops at a defined limit.
func GetDefaultFemaleClothes(sheet *xlsx.Sheet) *models.Commons {
	skip := []int{1, 2}
	stop := 24

	return parse(sheet, skip, stop, commonProcessor)
}

// Parses a Commons sheet for DefaultMaleMouths, applying the common processor function.
// Skips specific rows and stops at a defined limit.
func GetDefaultMaleMouths(sheet *xlsx.Sheet) *models.Commons {
	skip := []int{1, 2}
	stop := 28

	return parse(sheet, skip, stop, commonProcessor)
}

// Parses a Commons sheet for DefaultFemaleMouths, applying the common processor function.
// Skips specific rows and stops at a defined limit.
func GetDefaultFemaleMouths(sheet *xlsx.Sheet) *models.Commons {
	skip := []int{1, 2}
	stop := 26

	return parse(sheet, skip, stop, commonProcessor)
}

// Parses a Commons sheet for DefaultMaleEyes, applying the common processor function.
// Skips specific rows and stops at a defined limit.
func GetDefaultMaleEyes(sheet *xlsx.Sheet) *models.Commons {
	skip := []int{1, 2}
	stop := 17

	return parse(sheet, skip, stop, commonProcessor)
}

// Parses a Commons sheet for DefaultFemaleEyes, applying the common processor function.
// Skips specific rows and stops at a defined limit.
func GetDefaultFemaleEyes(sheet *xlsx.Sheet) *models.Commons {
	skip := []int{1, 2}
	stop := 18

	return parse(sheet, skip, stop, commonProcessor)
}

// Parses a Commons sheet for DefaultMaleHair, applying the common processor function.
// Skips specific rows and stops at a defined limit.
func GetDefaultMaleHair(sheet *xlsx.Sheet) *models.Commons {
	skip := []int{1, 2}
	stop := 25

	return parse(sheet, skip, stop, commonProcessor)
}

// Parses a Commons sheet for DefaultFemaleHair, applying the common processor function.
// Skips specific rows and stops at a defined limit.
func GetDefaultFemaleHair(sheet *xlsx.Sheet) *models.Commons {
	skip := []int{1, 2}
	stop := 13

	return parse(sheet, skip, stop, commonProcessor)
}

// Parses the DefaultMaleHat sheet, categorizing data into Data or DataEarless slices
// and panics if "NA" is encountered. Skips specific rows and stops at a defined limit.
func GetDefaultMaleHat(sheet *xlsx.Sheet) *models.Hats {
	skip := []int{1, 2}
	stop := 25

	return parse(sheet, skip, stop, func(result *models.Hats, data *models.Common, rowNumber int) {
		if data.FileName == "NA" {
			panic("NA is not allowed in hats")
		}
		if rowNumber < 21 {
			result.Data = append(result.Data, data)
		} else {
			result.DataEarless = append(result.DataEarless, data)
		}
	})
}

// Parses the DefaultFemaleHat sheet, categorizing data into Data or DataEarless slices
// and panics if "NA" is encountered. Skips specific rows and stops at a defined limit.
func GetDefaultFemaleHat(sheet *xlsx.Sheet) *models.Hats {
	skip := []int{1, 2}
	stop := 33

	return parse(sheet, skip, stop, func(result *models.Hats, data *models.Common, rowNumber int) {
		if data.FileName == "NA" {
			panic("NA is not allowed in hats")
		}
		if rowNumber < 25 {
			result.Data = append(result.Data, data)
		} else {
			result.DataEarless = append(result.DataEarless, data)
		}
	})
}

// Parses the MaleStackableHat sheet, categorizing data into Data or DataBack slices
// and panics if "NA" is encountered. Skips specific rows and stops at a defined limit.
func GetMaleStackableHat(sheet *xlsx.Sheet) *models.StackableHats {
	skip := []int{1, 2, 13, 14, 15, 16, 17, 18, 19}
	stop := 25

	return parse(sheet, skip, stop, func(result *models.StackableHats, data *models.Common, rowNumber int) {
		if data.FileName == "NA" {
			panic("NA is not allowed in stackable hats")
		}
		if rowNumber < 19 {
			result.Data = append(result.Data, data)
		} else {
			result.DataBack = append(result.DataBack, data)
		}
	})
}

// Parses the FemaleStackableHat sheet, categorizing data into Data or DataBack slices
// and panics if "NA" is encountered. Skips specific rows and stops at a defined limit.
func GetFemaleStackableHat(sheet *xlsx.Sheet) *models.StackableHats {
	skip := []int{1, 2, 13, 14, 15, 16, 17, 18}
	stop := 24

	return parse(sheet, skip, stop, func(result *models.StackableHats, data *models.Common, rowNumber int) {
		if data.FileName == "NA" {
			panic("NA is not allowed in stackable hats")
		}
		if rowNumber < 18 {
			result.Data = append(result.Data, data)
		} else {
			result.DataBack = append(result.DataBack, data)
		}
	})
}
