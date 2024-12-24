package processor

import (
	"fmt"
	"generator/models"
	"generator/utils"
	"strings"

	"github.com/samber/lo"
)

func Process(r *utils.Randomizer, c *models.Traits) {
	c.Droplets.Data = lo.Filter(c.Droplets.Data, func(droplet *models.Common, i int) bool {
		return models.Rarity(droplet.OpenSeaTraitValue) == c.Final.Rarity
	})
	if len(c.Droplets.Data) > 0 {
		c.Final.Droplets.DataFront = c.Droplets.Data[0]

		traitValue := c.Final.Droplets.DataFront.OpenSeaTraitValue

		c.Final.Droplets.DataBack = ExtractByTraitValue(c.Droplets.DataBack, traitValue)
		c.Final.Droplets.DataBackTransparent = ExtractByTraitValue(c.Droplets.DataBackTransparent, traitValue)
	} else {
		panic("no droplets found")
	}

	originBGData := c.BG.Data
	c.BG.Data = lo.Filter(originBGData, func(common *models.Common, i int) bool {
		return !lo.Contains(common.MustNotInclude, c.Final.Specie.String())
	})
	if picked := r.Random(c.BG.Data, c.BG.NA); picked != nil {
		c.Final.BG = picked
	} else {
		fmt.Printf("no bg found for token %d and %s\n", c.Final.Metadata.TokenID, r.Seed)
	}

	c.BG.Data = c.Final.DefaultFilter(c.BG.Data, models.FilterGender, models.FilterCategory)
	if picked := r.Random(c.BGAccent.Data, c.BGAccent.NA); picked != nil {
		c.Final.BGAccent = picked
	} else {
		fmt.Println("no bg accent found")
	}

	c.Aura.Normal = c.Final.DefaultFilter(c.Aura.Normal)
	if picked := r.Random(c.Aura.Normal, c.Aura.NA); picked != nil {
		c.Final.Aura.Back = picked

		if picked.Combined.Bool() {
			c.Final.Aura.Front = ExtractByTraitValue(c.Aura.Front, picked.OpenSeaTraitValue)
		}
	}

	c.Wings.Data = c.Final.DefaultFilter(c.Wings.Data)
	if picked := r.Random(c.Wings.Data, c.Wings.NA); picked != nil {
		c.Final.Wings = picked
	}

	c.Weapons.Front = c.Final.DefaultFilter(c.Weapons.Front)
	if picked := r.Random(c.Weapons.Front, c.Weapons.NA); picked != nil {
		c.Final.Weapons.Front = picked

		if picked.Combined.Bool() {
			c.Final.Weapons.Back = ExtractByTraitValue(c.Weapons.Back, picked.OpenSeaTraitValue)
		}
	}

	c.Bodies.Data = c.Final.DefaultFilter(c.Bodies.Data)
	if c.Final.Specie == models.SpecieOrigin {
		if c.Final.Rarity != models.COMMON {
			color := strings.Split(c.Final.Droplets.DataFront.OpenSeaTraitValue, " ")[1]
			c.Bodies.Data = lo.Filter(c.Bodies.Data, func(body *models.Common, i int) bool {
				return strings.Contains(body.OpenSeaTraitValue, color)
			})
			c.Final.Bodies = c.Bodies.Data[0]
		} else if len(c.Bodies.Data) > 0 {
			c.Final.Bodies = c.Bodies.Data[0]
		}
	} else if picked := r.Random(c.Bodies.Data, c.Bodies.NA); picked != nil {
		c.Final.Bodies = picked
	} else {
		panic("no body found")
	}

	if c.Final.Specie == models.SpecieElven {
		c.Final.ElvenEars = ExtractByTraitValue(c.ElvenEars.Data, c.Final.Bodies.OpenSeaTraitValue)
	}

	if c.Final.Bodies != nil && c.Final.Bodies.Combined.Bool() {
		c.Final.Tails = OptionalExtractByTraitValue(c.Tails.Data, c.Final.Bodies.OpenSeaTraitValue)
	}

	var excludeMouth, excludeNose, excludeEarrings, excludeFacegears, excludeHairs bool
	var forceStackableHat, forceGlasses bool

	if !c.Final.HasHair && (c.Final.Specie != models.SpecieBeing ||
		strings.Contains(strings.ToLower(c.Final.Bodies.OpenSeaTraitValue), "being")) {
		originHatData := c.Hats.Data

		c.Hats.Data = c.Final.DefaultFilter(originHatData)
		c.Hats.Data = lo.Filter(c.Hats.Data, func(common *models.Common, i int) bool {
			return c.Final.Specie != models.SpecieFeline ||
				lo.Contains(common.SpeciesLocked, models.SpecieFeline)
		})
		c.Hats.Data = lo.Filter(c.Hats.Data, func(common *models.Common, i int) bool {
			case1 := lo.Contains(common.SpeciesLocked, c.Final.Specie)
			case2 := common.RarityLocked.IsY() &&
				c.Final.Specie != models.SpecieSoul && c.Final.Specie != models.SpecieOrigin
			case3 :=
				lo.Contains(common.SpeciesLocked, models.SpecieNone) &&
					c.Final.Specie != models.SpecieSoul && c.Final.Specie != models.SpecieOrigin
			return case1 || case2 || case3
		})
		if picked := r.Random(c.Hats.Data, c.Hats.NA); picked != nil {
			c.Final.Hats.Data = picked

			skiMask := "Dark Ski Mask"
			if strings.Contains(picked.OpenSeaTraitValue, skiMask) {
				excludeMouth = true
				excludeNose = true
				excludeEarrings = true
				excludeFacegears = true
				excludeHairs = true
			}
		} else {
			fmt.Println("no hat found")

			forceStackableHat = true
		}
	}

	if !excludeFacegears &&
		c.Final.Specie != models.SpecieMonkey &&
		c.Final.Specie != models.SpecieCyborg &&
		(c.Final.Hats.Data == nil || !lo.Contains(c.Final.Hats.Data.MustNotInclude, "FACEGEAR")) {

		var checkNose, checkEarless bool
		if c.Final.Hats.Data != nil {
			checkNose = lo.Contains(c.Final.Hats.Data.MustNotInclude, "NOSE")
			checkEarless = lo.Contains(c.Final.Hats.Data.MustNotInclude, "EARLESS")
		}

		if (!checkEarless || c.Final.Hats.DataEarless == nil) &&
			(!checkNose || c.Final.Nose == nil) {
			c.Facegears.Data = c.Final.DefaultFilter(c.Facegears.Data)
			c.Facegears.Data = lo.Filter(c.Facegears.Data, func(common *models.Common, i int) bool {
				return c.Final.Specie != models.SpecieFeline ||
					!lo.Contains(common.MustNotInclude, models.SpecieFeline.String())
			})
			if picked := r.Random(c.Facegears.Data, c.Facegears.NA); picked != nil {
				c.Final.Facegears = picked
			}
		}
	}

	if c.Final.Specie == models.SpecieBeing {
		if c.Final.Hats.Data == nil &&
			!strings.Contains(strings.ToLower(c.Final.Bodies.OpenSeaTraitValue), "being") {
			if c.Final.Hats.DataEarless != nil {
				excludeEarrings = true
			}

			// mandatory
			originalEarlessHats := c.Hats.DataEarless
			originalEarlessHats = c.Final.DefaultFilter(originalEarlessHats)
			originalEarlessHats = lo.Filter(originalEarlessHats, func(common *models.Common, i int) bool {
				hasNose := c.Final.Nose != nil
				hasMouth := c.Final.Mouths != nil
				hasFaceGear := c.Final.Facegears != nil

				if lo.Contains(common.MustNotInclude, "NOSE") && hasNose ||
					lo.Contains(common.MustNotInclude, "FACEGEAR") && hasFaceGear ||
					lo.Contains(common.MustNotInclude, "MOUTH") && hasMouth {
					return false
				}

				return !lo.Contains([]string{"4B", "5B", "6B", "7B", "8B"}, c.Final.Bodies.FileName)
			})
			if picked := r.Random(originalEarlessHats, nil); picked != nil {
				c.Final.Hats.DataEarless = picked
			}
		}
		if c.Final.Hats.DataEarless != nil {
			excludeHairs = true
		}
	}

	if c.Final.Hats.Data == nil || !lo.Contains(c.Final.Hats.Data.MustNotInclude, "EYES") {
		originalEyes := c.Eyes.Data
		c.Eyes.Data = c.Final.DefaultFilter(originalEyes)
		c.Eyes.Data = lo.Filter(c.Eyes.Data, func(common *models.Common, i int) bool {
			case1 := lo.Contains(common.SpeciesLocked, c.Final.Specie)
			case2 := common.RarityLocked.IsY() &&
				c.Final.Specie != models.SpecieSoul && c.Final.Specie != models.SpecieOrigin
			case3 := lo.Contains(common.SpeciesLocked, models.SpecieNone) &&
				c.Final.Specie != models.SpecieSoul && c.Final.Specie != models.SpecieOrigin
			return case1 || case2 || case3
		})

		distributionNA := c.Eyes.NA
		if c.Final.Specie == models.SpecieOrigin {
			distributionNA = nil
		}

		if picked := r.Random(c.Eyes.Data, distributionNA); picked != nil {
			if !lo.Contains(picked.MustNotInclude, "EARLESS HAT") || c.Final.Hats.DataEarless == nil {
				c.Final.Eyes = picked
				if lo.Contains(picked.MustNotInclude, "NOSE") {
					excludeNose = true
				}
				if lo.Contains(picked.MustNotInclude, "EARLESS HAT") {
					c.Final.Hats.DataEarless = nil
					excludeHairs = false
				}
			}
		} else {
			forceGlasses = true
		}
	}

	if forceGlasses {
		c.Glasses.Data = c.Final.DefaultFilter(c.Glasses.Data)
		if picked := r.Random(c.Glasses.Data, c.Glasses.NA); picked != nil {
			c.Final.Glasses = picked
			excludeNose = true
			if lo.Contains(picked.MustInclude, "EYES") {
				c.Final.Eyes = r.Random(c.Eyes.Data, nil)
			}
		}
	}

	if !excludeNose && (c.Final.Hats.Data == nil || !lo.Contains(c.Final.Hats.Data.MustNotInclude, "NOSE")) {
		c.Nose.Data = c.Final.DefaultFilter(c.Nose.Data)
		if picked := r.Random(c.Nose.Data, c.Nose.NA); picked != nil {
			c.Final.Nose = picked
		}
	}
	if c.Final.Gender == models.GenderFemale && c.Final.HasHair {
		fmt.Println()
	}

	if !excludeHairs && c.Final.HasHair && c.Final.Specie != models.SpecieOrigin {
		originalHairs := c.Hairs.Hair

		c.Hairs.Hair = c.Final.DefaultFilter(originalHairs)
		c.Hairs.Hair = lo.Filter(c.Hairs.Hair, func(common *models.Common, i int) bool {
			return c.Final.Specie != models.SpecieFeline ||
				lo.Contains(common.SpeciesLocked, models.SpecieFeline)
		})
		c.Hairs.Hair = lo.Filter(c.Hairs.Hair, func(common *models.Common, i int) bool {
			return c.Final.Specie != models.SpecieElven ||
				!lo.Contains(common.MustNotInclude, c.Final.Specie.String()) && c.Final.Glasses == nil
		})
		if c.Final.Specie == models.SpecieFeline {
			c.Hairs.Hair = lo.Filter(c.Hairs.Hair, func(common *models.Common, i int) bool {
				return lo.Contains(common.SpeciesLocked, models.SpecieFeline)
			})
		} else if c.Final.Specie == models.SpecieElven {
			c.Hairs.Hair = lo.Filter(c.Hairs.Hair, func(common *models.Common, i int) bool {
				return !lo.Contains(common.SpeciesLocked, models.SpecieFeline) &&
					!lo.Contains(common.SpeciesLocked, models.SpecieSoul)
			})
		}
		if picked := r.Random(c.Hairs.Hair, c.Hairs.NA); picked != nil {
			c.Final.Hairs.Hair = picked

			if picked.Combined.Bool() {
				c.Final.Hairs.HairBack = ExtractByTraitValue(c.Hairs.HairBack, picked.OpenSeaTraitValue)
			}
		} else {
			fmt.Println("no hair found")

			forceStackableHat = true
		}
	}

	var stackableHatDistribution = c.StackableHats.NA
	if c.Final.Hairs.Hair == nil && c.Final.Hats.Data == nil {
		forceStackableHat = true
		stackableHatDistribution = nil
	}

	originalClothes := c.Clothes.Data
	c.Clothes.Data = c.Final.DefaultFilter(originalClothes)

	c.Clothes.Data = lo.Filter(c.Clothes.Data, func(common *models.Common, i int) bool {
		case1 := lo.Contains(common.SpeciesLocked, c.Final.Specie)
		case2 := common.RarityLocked.IsY() &&
			c.Final.Specie != models.SpecieSoul && c.Final.Specie != models.SpecieOrigin
		case3 :=
			lo.Contains(common.SpeciesLocked, models.SpecieNone) &&
				c.Final.Specie != models.SpecieSoul && c.Final.Specie != models.SpecieOrigin
		return case1 || case2 || case3
	})
	if picked := r.Random(c.Clothes.Data, c.Clothes.NA); picked != nil {
		c.Final.Clothes = picked
	}

	if c.Final.Weapons.Front != nil {
		c.Final.Hands = OptionalExtractByTraitValueContains(c.Hands.Data, c.Final.Bodies.OpenSeaTraitValue)
	}

	if !excludeMouth && (c.Final.Hats.Data == nil || !lo.Contains(c.Final.Hats.Data.MustNotInclude, "NOSE")) {
		originalMouth := c.Mouths.Data
		c.Mouths.Data = c.Final.DefaultFilter(originalMouth)
		c.Mouths.Data = lo.Filter(c.Mouths.Data, func(common *models.Common, i int) bool {
			return c.Final.Hats.DataEarless == nil || !lo.Contains([]string{
				"Tongue Out",
				"Middy",
				"Shmoke",
				"Dark Bandana",
				"Light Bandana",
				"Country Road",
			}, common.OpenSeaTraitValue)
		})
		c.Mouths.Data = lo.Filter(c.Mouths.Data, func(common *models.Common, i int) bool {
			case1 := lo.Contains(common.SpeciesLocked, c.Final.Specie)
			case2 := common.RarityLocked.IsY() &&
				c.Final.Specie != models.SpecieSoul && c.Final.Specie != models.SpecieOrigin
			case3 := lo.Contains(common.SpeciesLocked, models.SpecieNone) &&
				c.Final.Specie != models.SpecieSoul && c.Final.Specie != models.SpecieOrigin
			return (case1 || case2 || case3) &&
				(!lo.Contains(common.MustNotInclude, "EARLESS HAT") || c.Final.Hats.DataEarless == nil)
		})
		if picked := r.Random(c.Mouths.Data, c.Mouths.NA); picked != nil {
			c.Final.Mouths = picked
		}
	}

	if !excludeEarrings &&
		c.Final.Hats.DataEarless == nil && c.Final.Specie != models.SpecieFeline &&
		(c.Final.Hats.Data == nil || !lo.Contains(c.Final.Hats.Data.MustNotInclude, "EARRINGS")) &&
		(c.Final.Hairs.Hair == nil || !lo.Contains(c.Final.Hairs.Hair.MustNotInclude, "EARRINGS")) {

		c.Earrings.Data = c.Final.DefaultFilter(c.Earrings.Data)
		c.Earrings.Data = lo.Filter(c.Earrings.Data, func(common *models.Common, i int) bool {
			return (c.Final.Specie == models.SpecieFeline && lo.Contains(common.SpeciesLocked, models.SpecieFeline)) ||
				(c.Final.Specie != models.SpecieFeline && c.Final.Specie != models.SpecieElven) &&
					(c.Final.Specie == models.SpecieElven && lo.Contains(common.SpeciesLocked, models.SpecieElven))
		})
		if picked := r.Random(c.Earrings.Data, c.Earrings.NA); picked != nil {
			c.Final.Earrings = picked
		}
	}

	if c.Final.Hairs.Hair != nil && c.Final.Hairs.Hair.AbleToHaveStackableHat ||
		c.Final.Hats.Data != nil && c.Final.Hats.Data.AbleToHaveStackableHat || forceStackableHat {

		if c.Final.Specie != models.SpecieFeline {

			// TODO check Goggle Gear Red for SOUL species
			c.StackableHats.Data = c.Final.DefaultFilter(c.StackableHats.Data)
			if picked := r.Random(c.StackableHats.Data, stackableHatDistribution); picked != nil {
				if c.Final.Hats.DataEarless == nil || c.Final.Hats.DataEarless.AbleToHaveStackableHat {
					c.Final.StackableHats.DataFront = picked
					if picked.Combined.Bool() {
						c.Final.StackableHats.DataBack = ExtractByTraitValue(c.StackableHats.DataBack, picked.OpenSeaTraitValue)
					}
				}
			}

			if c.Final.Hairs.Hair != nil &&
				c.Final.Hairs.Hair.OnlyHaloAndHorns &&
				(lo.Contains(halos, c.Final.StackableHats.DataFront.FileName) ||
					!lo.Contains(horns, c.Final.StackableHats.DataFront.FileName)) {
				c.Final.StackableHats.DataFront = nil
				c.Final.StackableHats.DataBack = nil
			}

			if c.Final.Hairs.Hair != nil {
				if c.Final.StackableHats.DataFront != nil {
					if !lo.Contains(halos, c.Final.StackableHats.DataFront.FileName) &&
						!lo.Contains(horns, c.Final.StackableHats.DataFront.FileName) {
						c.Final.StackableHats.DataFront = nil
						c.Final.StackableHats.DataBack = nil
					}
				}
			}
		}
	}
}

var halos = []string{
	"2HATSS",
	"2HATSSB",
	"3HATSS",
	"3HATSSB",
	"3HATSSB",
	"14HATSS",
	"14HATSSB",
	"15HATSS",
	"15HATSSB",
}
var horns = []string{
	"9HATSS",
	"10HATSS",
	"13HATSS",
	"18HATSS",
	"19HATSS",
	"20HATSS",
}

func extract(origin, temp []*models.Common) []*models.Common {
	var result []*models.Common

	for _, v := range temp {
		found := ExtractByFileName(origin, v.FileName)
		if found == nil {
			continue
		}
		result = append(result, found)
	}

	return result
}
