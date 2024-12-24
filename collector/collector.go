package collector

import (
	"encoding/json"
	"fmt"
	"generator/models"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"sort"
	"sync"
)

const filePath = "out/api_responses.json" // Path for storing API responses
const numberOfNFTs = 7573                 // Total number of NFTs to process

// GetResponses reads and parses API responses from a JSON file.
func GetResponses() []*models.APIResponse {
	file, err := os.Open(filePath)
	if err != nil {
		// Consider logging errors instead of panic for more graceful error handling
		panic(fmt.Errorf("Error opening file: %v\n", err))
	}
	defer file.Close()

	data, err := ioutil.ReadAll(file)
	if err != nil {
		panic(fmt.Errorf("Error reading file: %v\n", err))
	}

	var apiResponses []*models.APIResponse
	err = json.Unmarshal(data, &apiResponses)
	if err != nil {
		panic(fmt.Errorf("Error parsing JSON response: %v\n", err))
	}

	return apiResponses
}

// GetMetadataWithError reads metadata for a given tokenID from a file and returns it.
func GetMetadataWithError(tokenID string) (*models.APIResponse, error) {
	file, err := os.Open(fmt.Sprintf("./assets/results/metadata/%s.json", tokenID))
	if err != nil {
		// Return an error to allow the caller to handle it
		return nil, fmt.Errorf("Error opening file: %v\n", err)
	}
	defer file.Close()

	data, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, fmt.Errorf("Error reading file: %v\n", err)
	}

	var apiResponses = new(models.APIResponse)
	err = json.Unmarshal(data, apiResponses)
	if err != nil {
		return nil, fmt.Errorf(" Error parsing JSON response: %v\n", err)
	}

	return apiResponses, nil
}

// PrintRarities extracts and prints unique rarity values from the API responses.
func PrintRarities() {
	apiResponses := GetResponses()

	data := make(map[string]struct{})

	for _, r := range apiResponses {
		for _, a := range r.Attributes {
			if a.TraitType != "Rarity" {
				continue
			}
			data[a.Value] = struct{}{}
			break
		}
	}

	for k := range data {
		fmt.Println(k)
	}
}

// OrderMetadata sorts the API responses by TokenID and saves the ordered data to a file.
func OrderMetadata() {
	apiResponses := GetResponses()

	sort.Slice(apiResponses, func(i, j int) bool {
		return apiResponses[i].TokenID < apiResponses[j].TokenID
	})

	orderedData, err := json.MarshalIndent(apiResponses, "", "  ")
	if err != nil {
		fmt.Printf("Error marshaling ordered metadata: %v\n", err)
		return
	}

	err = ioutil.WriteFile(filePath, orderedData, 0644)
	if err != nil {
		fmt.Printf("Error writing to file: %v\n", err)
		return
	}

	fmt.Println("Metadata ordered and saved successfully.")
}

// fetchAPIResponse fetches data for a given tokenID and sends it to the result channel.
func fetchAPIResponse(tokenID int, resultChan chan *models.APIResponse) {
	apiURL := computeURL(tokenID)

	defer func() {
		log.Println("finished fetching data from API...", apiURL)
	}() // Deferred log statement

	log.Println("Fetching data from API...", apiURL)

	resp, err := http.Get(apiURL)
	if err != nil {
		fmt.Printf("Error fetching data from API: %v\n", err)
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Error reading response body: %v\n", err)
		return
	}

	var apiResponse models.APIResponse
	err = json.Unmarshal(body, &apiResponse)
	if err != nil {
		fmt.Printf("Error parsing JSON response: %v\n", err)
		return
	}

	apiResponse.TokenID = tokenID

	resultChan <- &apiResponse
}

// saveToFile saves API response data to a file.
func saveToFile(filename string, data []*models.APIResponse) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	bytes, err := json.Marshal(data)
	if err != nil {
		return err
	}

	_, err = file.Write(bytes)
	return err
}

// computeURL computes the API URL for a given tokenID.
func computeURL(tokenID int) string {
	return fmt.Sprintf("https://ipfs.io/ipfs/QmNjK56KZFaoHwDqS8mb28kj2pWPcsDg8XevnwJ4T8mT4h/%d.json", tokenID)
}

// worker processes tokenIDs from a channel and fetches their API responses.
func worker(tokenIDChan chan int, resultChan chan *models.APIResponse, wg *sync.WaitGroup) {
	defer wg.Done()

	for tokenID := range tokenIDChan {
		fetchAPIResponse(tokenID, resultChan)
	}
}

// CollectAndSaveMetadata fetches, processes, and saves metadata for NFTs.
func CollectAndSaveMetadata() {
	workerCount := 5 // Set the number of concurrent goroutines

	var wg sync.WaitGroup
	tokenIDChan := make(chan int, 10)
	resultChan := make(chan *models.APIResponse, 10)

	// Start worker goroutines
	for i := 0; i < workerCount; i++ {
		wg.Add(1)
		go worker(tokenIDChan, resultChan, &wg)
	}

	// Feed the worker goroutines with API URLs
	go func() {
		for i := 0; i < numberOfNFTs; i++ {
			tokenIDChan <- i
		}
		close(tokenIDChan)
	}()

	done := make(chan bool)
	var apiResponses []*models.APIResponse

	// Collect results from workers
	go func() {
		for apiResponse := range resultChan {
			apiResponses = append(apiResponses, apiResponse)
		}
		done <- true
	}()

	wg.Wait()
	close(resultChan)

	<-done

	sort.Slice(apiResponses, func(i, j int) bool {
		return apiResponses[i].TokenID < apiResponses[j].TokenID
	})

	err := saveToFile(filePath, apiResponses)
	if err != nil {
		fmt.Printf("Error saving data to file: %v\n", err)
	} else {
		fmt.Println("Data saved to file successfully.")
	}
}
