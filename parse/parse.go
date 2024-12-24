package parse

import (
	"strings"

	"generator/models"

	"github.com/tealeg/xlsx"
)

// filePath is the location of the Excel file to parse.
const filePath = "data.xlsx"

// Do reads the Excel file, parses each sheet based on its name,
// and maps the data into the corresponding models.Traits structure.
func Do() *models.Traits {
	// Open the Excel file specified by filePath.
	xlFile, err := xlsx.OpenFile(filePath)
	if err != nil {
		// Panic if the file cannot be opened, as this is a critical error.
		panic(err)
	}

	// Initialize an empty Traits structure to hold all parsed data.
	data := new(models.Traits)

	// Iterate over all sheets in the Excel file.
	for _, sheet := range xlFile.Sheets {
		// Trim leading and trailing spaces from the sheet name and convert it to a SheetName type.
		sheetName := models.SheetName(strings.Trim(sheet.Name, " "))

		// Handle parsing logic based on the sheet name.
		switch sheetName {
		case "":
			// Panic if a sheet has an empty name, as this is invalid.
			panic("empty sheet name")
		case models.SheetGeneralRules:
			// Skip the "General Rules" sheet.
			continue
		case models.SheetBODIES:
			// Parse the "Bodies" sheet.
			data.Bodies = GetBodies(sheet)
		case models.SheetTAILS:
			// Parse the "Tails" sheet.
			data.Tails = GetTails(sheet)
		case models.SheetELVENEAR:
			// Parse the "Elven Ears" sheet.
			data.ElvenEars = GetElderEars(sheet)
		case models.SheetDROPLETS:
			// Parse the "Droplets" sheet.
			data.Droplets = GetDroplets(sheet)
		case models.SheetHANDS:
			// Parse the "Hands" sheet.
			data.Hands = GetHands(sheet)
		case models.SheetHAIR:
			// Parse the "Hair" sheet.
			data.Hairs = GetHairs(sheet)
		case models.SheetHATS:
			// Parse the "Hats" sheet.
			data.Hats = GetHats(sheet)
		case models.SheetSTACKABLEHATS:
			// Parse the "Stackable Hats" sheet.
			data.StackableHats = GetStackableHats(sheet)
		case models.SheetMOUTH:
			// Parse the "Mouth" sheet.
			data.Mouths = GetMouths(sheet)
		case models.SheetNOSE:
			// Parse the "Nose" sheet.
			data.Nose = GetNoses(sheet)
		case models.SheetEYES:
			// Parse the "Eyes" sheet.
			data.Eyes = GetEyes(sheet)
		case models.SheetGLASSES:
			// Parse the "Glasses" sheet.
			data.Glasses = GetGlasses(sheet)
		case models.SheetEARRINGS:
			// Parse the "Earrings" sheet.
			data.Earrings = GetEarrings(sheet)
		case models.SheetCLOTHES:
			// Parse the "Clothes" sheet.
			data.Clothes = GetClothes(sheet)
		case models.SheetWINGS:
			// Parse the "Wings" sheet.
			data.Wings = GetWings(sheet)
		case models.SheetWEAPONS:
			// Parse the "Weapons" sheet.
			data.Weapons = GetWeapons(sheet)
		case models.SheetFACEGEARS:
			// Parse the "Face Gears" sheet.
			data.Facegears = GetFaceGears(sheet)
		case models.SheetBG:
			// Parse the "Background" sheet.
			data.BG = GetBG(sheet)
		case models.SheetBGACCENTS:
			// Parse the "Background Accents" sheet.
			data.BGAccent = GetBGAccents(sheet)
		case models.SheetAURA:
			// Parse the "Aura" sheet.
			data.Aura = GetAura(sheet)
		case models.SheetDEFAULTMaleCLOTHES:
			// Parse the "Default Male Clothes" sheet.
			data.DefaultMaleClothes = GetDefaultMaleClothes(sheet)
		case models.SheetDEFAULTFemaleCLOTHES:
			// Parse the "Default Female Clothes" sheet.
			data.DefaultFemaleClothes = GetDefaultFemaleClothes(sheet)
		case models.SheetDEFAULTMaleMOUTHS:
			// Parse the "Default Male Mouths" sheet.
			data.DefaultMaleMouths = GetDefaultMaleMouths(sheet)
		case models.SheetDEFAULTFemaleMOUTHS:
			// Parse the "Default Female Mouths" sheet.
			data.DefaultFemaletMouths = GetDefaultFemaleMouths(sheet)
		case models.SheetDEFAULTMaleEYES:
			// Parse the "Default Male Eyes" sheet.
			data.DefaultMaleEyes = GetDefaultMaleEyes(sheet)
		case models.SheetDEFAULTFemaleEYES:
			// Parse the "Default Female Eyes" sheet.
			data.DefaultFemaleEyes = GetDefaultFemaleEyes(sheet)
		case models.SheetDEFAULTMaleHAIR:
			// Parse the "Default Male Hair" sheet.
			data.DefaultMaleHair = GetDefaultMaleHair(sheet)
		case models.SheetDEFAULTFemaleHAIR:
			// Parse the "Default Female Hair" sheet.
			data.DefaultFemaleHair = GetDefaultFemaleHair(sheet)
		case models.SheetDEFAULTMaleHATS:
			// Parse the "Default Male Hats" sheet.
			data.DefaultMaleHat = GetDefaultMaleHat(sheet)
		case models.SheetDEFAULTFemaleHATS:
			// Parse the "Default Female Hats" sheet.
			data.DefaultFemaleHat = GetDefaultFemaleHat(sheet)
		case models.SheetMALEDEFAULTSTACKABLEHAT:
			// Parse the "Male Default Stackable Hat" sheet.
			data.DefaultMaleStackableHat = GetMaleStackableHat(sheet)
		case models.SheetFEMALEDEFAULTSTACKABLEHAT:
			// Parse the "Female Default Stackable Hat" sheet.
			data.DefaultFemaleStackableHat = GetFemaleStackableHat(sheet)
		default:
			// Panic if the sheet name does not match any known value.
			panic("unknown sheet name: " + sheetName)
		}
	}

	// Return the populated Traits structure.
	return data
}
