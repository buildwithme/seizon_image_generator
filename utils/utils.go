package utils

import (
	"crypto/sha1"
	"encoding/binary"
	"generator/models"
	"strconv"
	"sync"
	"time"
)

const (
	timeFormat          = "2006-01-02 15:04:05"
	one_hundred_percent = 100000
	retryCount          = 10
)

func MustParseTime(timeString string) time.Time {
	t, err := time.Parse(timeFormat, timeString)
	if err != nil {
		panic(err)
	}
	return t
}

type Randomizer struct {
	Seed     string
	withTime bool
	mu       sync.Mutex
	Counter  int
}

var (
	mu   sync.RWMutex
	done bool
)

func SetRandomizer(flag bool) {
	mu.Lock()
	defer mu.Unlock()
	done = flag
}

func IsRandomizerDone() bool {
	mu.RLock()
	defer mu.RUnlock()
	return done
}

func NewRandomizer(seed string) *Randomizer {
	return &Randomizer{
		Seed: seed,
	}
}

func (r *Randomizer) RandomNumberBy(max int) int {
	return r.RandomNumber(max - 1)
}

func (r *Randomizer) RandomNumber(max int) int {
	mu.Lock()
	var data = r.Seed + strconv.Itoa(r.Counter)
	r.Counter++
	mu.Unlock()

	hash := sha1.New()
	hash.Write([]byte(data))
	hashBytes := hash.Sum(nil)
	randomNumber := binary.BigEndian.Uint64(hashBytes)
	return int(randomNumber % uint64(max+1))
}

func (r *Randomizer) RandomCategory() models.Category {
	var percentages []int

	var lastPercentage int

	list := models.CategoryList()

	for _, v := range list {
		percent := v.ToPercentage()

		percentNew := lastPercentage + percent

		percentages = append(percentages, percentNew)

		lastPercentage = percentNew
	}

	randomNumber := r.RandomNumber(100)

	for i, v := range percentages {
		if randomNumber <= int(v) {
			return list[i]
		}
	}

	panic("should not happen")
}

func (r *Randomizer) RandomGender() models.Gender {
	randomNumber := r.RandomNumber(100)

	if randomNumber < models.GenderMale.ToPercentage() {
		return models.GenderMale
	}

	return models.GenderFemale
}

func (r *Randomizer) HasHair(percentage int) bool {
	randomNumber := r.RandomNumber(100)
	return randomNumber < percentage
}

func (r *Randomizer) Random(data []*models.Common, na *models.Common) *models.Common {
	var percentages []float64

	var lastPercentage float64

	if na != nil && na.Distribution.GetPercentage() != 0 {
		randomNumber := r.RandomNumber(100000)

		if randomNumber < int(na.Distribution.GetPercentage()*1000) {
			return nil
		}
	}

	distributions := NormalizeDistribution(data)

	for i := range data {
		// TODO this can fail
		percent := distributions[i]

		percentNew := lastPercentage + percent*1000

		percentages = append(percentages, percentNew)

		lastPercentage = percentNew
	}

	percentages = append(percentages, 100*1000)

	for i := 0; i < retryCount; i++ {
		randomNumber := r.RandomNumber(100 * 1000)
		length := len(data)

		for i, v := range percentages {
			if randomNumber <= int(v) && i < length {
				return data[i]
			}
		}

		if done {
			break
		}
	}

	return nil
}

func NormalizeDistribution(data []*models.Common) []float64 {
	sum := 0.0

	for _, value := range data {
		percent := value.Distribution.GetPercentage()
		sum += percent
	}

	var percentages = make([]float64, 0, len(data))

	for _, value := range data {
		percent := value.Distribution.GetPercentage()
		percentages = append(percentages, percent/sum*100)
	}

	return percentages
}
