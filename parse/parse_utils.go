package parse

import (
	"fmt"
	"generator/models"
	"strings"

	"github.com/samber/lo"
	"github.com/tealeg/xlsx"
)

// parse is a generic function to parse data from an xlsx sheet and map it to a custom model type T.
// Parameters:
// - sheet: The xlsx sheet to parse.
// - skip: A slice of row numbers to skip during parsing.
// - stop: The row number at which parsing should stop.
// - appendData: A callback function to handle appending the parsed data to the result of type T.
func parse[T any](sheet *xlsx.Sheet, skip []int, stop int, appendData func(*T, *models.Common, int)) *T {
	// Create a new instance of the result type.
	result := new(T)

	// Iterate through all rows in the sheet.
	for index, row := range sheet.Rows {
		rowNumber := index + 1 // Adjust for 1-based indexing.

		// Stop parsing if the stop row is reached.
		if rowNumber == stop {
			break
		}

		// Skip rows specified in the skip slice.
		if lo.Contains(skip, rowNumber) {
			continue
		}

		// Initialize a new Common model to store the parsed data for the current row.
		data := new(models.Common)

		// Iterate through all cells in the row.
		for i, cell := range row.Cells {
			// Trim leading and trailing spaces from the cell value.
			cellString := strings.Trim(cell.String(), " ")

			// Handle cell data based on its column index.
			switch i {
			case 0:
				// Column 0: FileName
				data.FileName = cellString
			case 1:
				// Column 1: OpenSeaTraitValue
				data.OpenSeaTraitValue = cellString
			case 2:
				// Column 2: Category (split by comma and validate each category)
				categories := strings.Split(cellString, ",")
				for _, category := range categories {
					if category == "" {
						continue
					}
					c := models.Category(strings.Trim(category, " "))
					if c.IsInvalid() {
						panic(fmt.Sprintf("invalid category: %s", c))
					}
					data.Category = append(data.Category, c)
				}
			case 3:
				// Column 3: Gender (validate if provided)
				if cellString != "" {
					gender := models.Gender(cellString)
					if gender.IsInvalid() {
						panic(fmt.Sprintf("invalid gender: %s", gender))
					}
					data.Gender = gender
				}
			case 4:
				// Column 4: Combined (validate the combined value)
				combined := models.Combined(cellString)
				if combined.IsInvalid() {
					panic(fmt.Sprintf("invalid combined: %s", combined))
				}
				data.Combined = combined
			case 5:
				// Column 5: MustNotInclude (split by comma and store)
				if cellString == "" {
					continue
				}
				values := strings.Split(cellString, ",")
				for _, value := range values {
					data.MustNotInclude = append(data.MustNotInclude, strings.Trim(value, " "))
				}
			case 6:
				// Column 6: SpeciesLocked (split by comma, trim, and validate each specie)
				values := strings.Split(cellString, ",")
				for _, value := range values {
					specie := models.Specie(strings.Trim(value, " "))
					if specie.IsInvalid() {
						panic(fmt.Sprintf("invalid specie: %s", specie))
					}
					data.SpeciesLocked = append(data.SpeciesLocked, specie)
				}
			case 7:
				// Column 7: Distribution (validate only for non-default sheets)
				if !strings.Contains(strings.ToLower(sheet.Name), "default") {
					distribution := models.Distribution(cellString)
					if distribution.IsInvalid() {
						panic(fmt.Sprintf("invalid distribution: %s at %s", distribution, sheet.Name))
					}
					data.Distribution = distribution
				}
			case 8:
				// Column 8: Skipped (No processing)
				continue
			case 9:
				// Column 9: MustInclude (split by comma and store)
				values := strings.Split(cellString, ",")
				for _, value := range values {
					data.MustInclude = append(data.MustInclude, strings.Trim(value, " "))
				}
			case 10:
				// Column 10: RarityLocked (split, trim, and validate each rarity value)
				values := strings.Split(cellString, ",")
				for _, value := range values {
					data.RarityLocked = models.RarityLocked(strings.Trim(value, " "))
					if data.RarityLocked.IsInvalid() {
						panic(fmt.Sprintf("invalid rarity locked: %s", data.RarityLocked))
					}
				}
			case 11:
				// Column 11: AbleToHaveStackableHat (set flag if value is "Y")
				data.AbleToHaveStackableHat = cellString == "Y"
			case 12:
				// Column 12: OnlyHaloAndHorns (set flag if value is "Y")
				data.OnlyHaloAndHorns = cellString == "Y"
			}
		}

		// Use the appendData callback to add the parsed data to the result.
		appendData(result, data, rowNumber)
	}

	// Return the populated result.
	return result
}
