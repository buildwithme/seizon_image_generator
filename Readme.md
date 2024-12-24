/*
# NFT Generator

This project is a Go-based NFT generation system designed to process metadata, traits, and layered images for NFTs. It supports both single and batch processing with efficient concurrency.

## Features

- **Metadata Processing**:
  - Parses metadata from JSON or XLSX files.
  - Validates and processes attributes like gender, rarity, species, and categories.

- **Image Generation**:
  - Combines multiple layers to create a final NFT image.
  - Saves images to designated directories.

- **Concurrency**:
  - Utilizes Goroutines to process multiple NFTs simultaneously.

- **File Management**:
  - Handles metadata and image outputs, replacing placeholders when required.

- **Customizable Traits**:
  - Easily extensible to add new traits or modify existing ones.

---

## Project Structure

- `main.go`: Entry point for the NFT generation process.
- `collector/`: Handles reading and processing metadata.
- `generator/`: Contains logic for image generation.
- `models/`: Defines data structures for metadata and traits.
- `parse/`: Parses and organizes trait data from files.
- `processor/`: Processes and randomizes trait data.
- `utils/`: Utility functions for randomization and file handling.

---

## Setup Instructions

### Prerequisites

1. Install Go (version 1.18 or later).
2. Run `go mod tidy` to install necessary dependencies.

### Folder Structure

Ensure the following directory structure exists:

assets/
├── traits/                  // Contains trait folders like BACKGROUND, CLOTHES, etc.
├── results/                 // Stores generated images and metadata
│   ├── images/              // Output images
│   ├── metadata/            // Output metadata files
│   └── rarity.json          // Summarized rarity information

---

## Usage

### Single Token Generation

- Edit the `executeSingle` function in `main.go` to specify the desired `tokenID`.
- Run the program:
  go run main.go

### Batch Token Generation

- Edit the `executeCollection` function in `main.go` to set the desired number of NFTs (`nrNFTs`).
- Run the program:
  go run main.go

### Replace Metadata Image URLs

- Uncomment the `replaceImageURLs` function in `main.go`.
- Run the program:
  go run main.go

---

## Adding New Traits

1. **Extend Models**:
   - Add the new trait as a struct or constant in `models/`.

2. **Update Parsing**:
   - Extend the `parse/` package to parse data for the new trait.

3. **Process in Generator**:
   - Modify `processToken()` in `main.go` to handle the new trait.

4. **Test**:
   - Verify metadata and image generation with the new trait.

---

## Configuration

### Adjustable Parameters

- `baseFolder`: Base directory for trait layers.
- `maxWorkers`: Number of concurrent workers for NFT processing.
- `max_NFTS`: Maximum number of NFTs to generate.

### Seed for Randomization

Modify the randomizer seed in `executeSingle()` or `executeCollection()` for reproducibility:
  seed := uuid.NewString() // Generate a unique seed

---

## Contributions

Contributions to improve functionality or add features are welcome! Please:
1. Fork the repository.
2. Implement your changes.
3. Create a pull request with a detailed description.

---

## License

This project is licensed under the MIT License. See the LICENSE file for more details.
