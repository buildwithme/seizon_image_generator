package main

import (
	"encoding/json"
	"fmt"
	"generator/collector"
	"generator/generator"
	"generator/models"
	"generator/parse"
	"generator/processor"
	"generator/utils"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
	"sync"

	"github.com/google/uuid"
)

var (
	baseFolder = "./assets/traits/"
	maxWorkers = 15
	max_NFTS   = 7573
)

var (
	mu        sync.Mutex
	muRar     sync.Mutex
	m         = make(map[string]struct{})
	responses = collector.GetResponses()
	rarities  = make(map[string]map[string]int)
	ram       = make(chan TraitData, 1000)
)

func main() {
	executeSingle(1)

	// executeCollection(7573)

	// replaceImageURLs(7573)
}

func replaceImageURLs(nrNFTs int) {
	for tokenID := 0; tokenID < nrNFTs; tokenID++ {

		metadata, err := collector.GetMetadataWithError(strconv.Itoa(tokenID))
		if err != nil {
			log.Printf("File not found: %d: %s", tokenID, err)
			continue
		}

		metadata.Image = strings.ReplaceAll(metadata.Image, "REPLACE_ME", "bafybeibh3auum3psmutucg52tlmdj4zkdyqkvlzta43k76mvgpkr72otby")

		writeToSimpleFile(fmt.Sprintf("./assets/results/metadata/%d.json", tokenID), metadata)
	}
}

func executeSingle(tokenID int) {
	maxWorkers = 1

	seed := uuid.NewString()

	mm, err := collector.GetMetadataWithError(strconv.Itoa(tokenID))
	if err == nil {
		seed = mm.Seed
	} else {
		log.Println("Error getting metadata: ", err)
	}

	tr := parse.Do()
	r := utils.NewRandomizer(seed)

	var wg sync.WaitGroup
	workers := make(chan struct{}, maxWorkers)
	workers <- struct{}{}
	wg.Add(1)

	go processToken(&wg, workers, r, tr.Copy(), tokenID)

	wg.Wait()
}

func executeCollection(nrNFTs int) {
	tr := parse.Do()

	go func() {
		for {
			select {
			case data := <-ram:
				muRar.Lock()
				if rarities[data.Folder] == nil {
					rarities[data.Folder] = make(map[string]int)
				}
				rarities[data.Folder][data.Name]++
				muRar.Unlock()
				writeToSimpleFile("./assets/results/rarity.json", rarities)
			}
		}
	}()

	var wg sync.WaitGroup
	workers := make(chan struct{}, maxWorkers)

	if nrNFTs > max_NFTS {
		nrNFTs = max_NFTS
	}

	for tokenID := 0; tokenID < nrNFTs; tokenID++ {
		workers <- struct{}{}
		wg.Add(1)

		go processToken(&wg, workers, nil, tr.Copy(), tokenID)
	}

	utils.SetRandomizer(true)

	wg.Wait()

	writeToSimpleFile("./assets/results/rarity.json", rarities)
}

type TraitData struct {
	Folder   string
	Name     string
	FileName string
}

func processToken(wg *sync.WaitGroup, done <-chan struct{}, r *utils.Randomizer, traits *models.Traits, tokenID int) {
	defer wg.Done()
	defer func() {
		log.Printf("Done processing token %d", tokenID)
		<-done
	}()

	var paths []string
	var index int
bigfor:
	for {
		if utils.IsRandomizerDone() && index != 0 || index > 10 {
			if !utils.IsRandomizerDone() {
				log.Printf("Failed to generate token %d", tokenID)
			}
			break
		}

		index++

		c := traits.Copy()

		metadata := responses[tokenID]

		if r == nil {
			seed := uuid.NewString()
			r = utils.NewRandomizer(seed)
		}
		metadata.Seed = r.Seed

		rarity := metadata.GetRarity()
		switch rarity {
		case models.ONE_OF_ONE, models.UNKNOWN_COLOR1, models.UNKNOWN_COLOR2, models.UNKNOWN_COLOR3:
			break bigfor
		}

		c.Final.Rarity = rarity
		c.Final.Specie = metadata.GetSpecie()
		c.Final.Metadata = metadata
		if c.Final.Specie == models.SpecieNone {
			log.Printf("Specie not found for token %d", tokenID)
			break
		}
		c.Final.Category = r.RandomCategory()
		c.Final.Gender = r.RandomGender()

		if c.Final.Specie == models.SpecieMonkey {
			c.Final.HasHair = false
		} else if c.Final.Gender == models.GenderFemale {
			c.Final.HasHair = r.HasHair(100)
		} else {
			c.Final.HasHair = r.HasHair(50)
		}

		processor.Process(r, c)

		var key string
		add := func(common *models.Common, folderName, traitName string) {
			if common == nil {
				return
			}

			if traitName != "" {
				c.Final.Metadata.Attributes = append(c.Final.Metadata.Attributes, models.Attribute{
					TraitType: traitName,
					Value:     common.OpenSeaTraitValue,
				})
			}

			key += common.OpenSeaTraitValue

			paths = append(paths, baseFolder+folderName+"/"+common.FileName+".png")

			ram <- TraitData{
				Folder:   folderName,
				Name:     common.OpenSeaTraitValue,
				FileName: common.FileName,
			}
		}

		add(c.Final.BG, "BACKGROUND", "Background")
		add(c.Final.BGAccent, "BACKGROUND ACCENT", "Background")
		add(c.Final.Droplets.DataBack, "DROPLETS (BACK)", "Rarity")
		add(c.Final.Aura.Back, "AURA (BACK)", "Background")
		add(c.Final.Tails, "TAILS", "")
		add(c.Final.Wings, "WINGS", "Wings")
		add(c.Final.Weapons.Back, "WEAPONS (BACK)", "Weapon")
		add(c.Final.Droplets.DataBackTransparent, "DROPLET (BACK TRANSPARENT)", "Rarity")
		add(c.Final.StackableHats.DataBack, "HATS (STACKABLE)", "Hat")
		add(c.Final.Hairs.HairBack, "HAIR (BACK)", "Hair")
		add(c.Final.Bodies, "BODIES", "Body")
		add(c.Final.Facegears, "FACE GEAR", "Face")
		add(c.Final.Clothes, "CLOTHES", "Clothes")
		add(c.Final.Hands, "HANDS", "")
		add(c.Final.Weapons.Front, "WEAPONS (FRONT)", "Weapons")
		add(c.Final.Eyes, "EYES", "Eyes")
		add(c.Final.Mouths, "MOUTH", "Mouth")
		add(c.Final.Nose, "NOSE", "")
		add(c.Final.Hairs.Hair, "HAIR", "Hair")
		add(c.Final.Hats.Data, "HATS", "Hat")
		add(c.Final.Hats.DataEarless, "HATS", "Hat")
		add(c.Final.StackableHats.DataFront, "HATS (STACKABLE)", "Hat")
		add(c.Final.ElvenEars, "ELVEN EARS", "")
		add(c.Final.Earrings, "EARRINGS", "")
		add(c.Final.Glasses, "GLASSES", "Glasses")
		add(c.Final.Droplets.DataFront, "DROPLETS", "Rarity")
		add(c.Final.Aura.Front, "AURA (FRONT)", "Background")

		mu.Lock()
		if _, ok := m[key]; ok {
			// log.Println("Duplicate found: ", key)
			mu.Unlock()
			continue
		}
		m[key] = struct{}{}
		mu.Unlock()

		break
	}

	g := generator.NewImageCreator(tokenID, paths)

	g.Process()

	g.WriteTo(fmt.Sprintf("./assets/results/images/%d.png", tokenID))

	metadata := responses[tokenID]
	metadata.MakeAttributesUnique()
	metadata.AnimationURL = ""
	metadata.Image = fmt.Sprintf("https://ipfs.io/ipfs/REPLACE_ME/%d.jpg", tokenID)
	writeToSimpleFile(fmt.Sprintf("./assets/results/metadata/%d.json", tokenID), metadata)
}

func writeToSimpleFile(name string, data interface{}) {
	body, err := json.MarshalIndent(data, "", "\t")
	if err != nil {
		fmt.Println(err)
	}
	err = ioutil.WriteFile(name, body, 0777)
	if err != nil {
		fmt.Println(err)
		return
	}
}
