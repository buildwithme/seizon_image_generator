package generator

import (
	"errors"
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"log"
	"math"
	"os"

	"github.com/anthonynsimon/bild/imgio"
)

// Global ConcurrentMap to cache images
var m = NewConcurrentMap()

// NewImageCreator initializes an ImageCreator with an ID and a list of paths to process.
func NewImageCreator(id int, paths []string) *ImageCreator {
	return &ImageCreator{
		id:    id,
		Paths: paths,
	}
}

// ImageCreator represents an object responsible for creating images by compositing layers.
type ImageCreator struct {
	id    int         // Identifier for the image creation task
	Paths []string    // Paths to the image layers
	final *image.RGBA // The resulting composed image
}

// Process loads, composites, and prepares the final image by stacking layers.
func (c *ImageCreator) Process() *image.RGBA {
	for i, path := range c.Paths {
		imageSource, ok := m.Get(path) // Retrieve image from cache
		if !ok {
			panic("Image not found in cache " + path) // Consider better error handling
		}

		if c.final == nil {
			// Initialize the final image with the bounds of the first layer
			c.final = image.NewRGBA(imageSource.Bounds())
		}

		// Use `draw.Src` for the first layer, `draw.Over` for subsequent layers
		drawType := draw.Over
		if i == 0 {
			drawType = draw.Src
		}

		// Composite the image layers
		draw.Draw(c.final, imageSource.Bounds(), imageSource, image.Point{}, drawType)
	}

	// Ensure there is always a final image, even if no paths are provided
	if c.final == nil {
		paperImage := getPaperImage()
		c.final = image.NewRGBA(paperImage.Bounds())
	}

	return c.final
}

// BlendParameters represents parameters for blending images.
type BlendParameters struct {
	Brightness float64 // Brightness adjustment
	Contrast   float64 // Contrast adjustment
}

// blendPixels blends two pixels with specified brightness and contrast parameters.
func blendPixels(c1, c2 color.Color, t, brightness, contrast float64) color.Color {
	r1, g1, b1, a1 := c1.RGBA()
	r2, g2, b2, _ := c2.RGBA()

	// Average the color channels and apply brightness/contrast adjustments
	rb := float64((r1+r2)/2) * t
	gb := float64((g1+g2)/2) * t
	bb := float64((b1+b2)/2) * t

	ra := math.Pow(rb/0xffff, 1/contrast) * 0xffff
	ga := math.Pow(gb/0xffff, 1/contrast) * 0xffff
	ba := math.Pow(bb/0xffff, 1/contrast) * 0xffff

	// Clamp values to valid range
	ra = math.Min(math.Max(ra+brightness, 0), 0xffff)
	ga = math.Min(math.Max(ga+brightness, 0), 0xffff)
	ba = math.Min(math.Max(ba+brightness, 0), 0xffff)

	return color.RGBA64{uint16(ra), uint16(ga), uint16(ba), uint16(a1)}
}

// blendImages blends two images using the specified parameters.
func blendImages(img1, img2 *image.RGBA, params BlendParameters) (*image.RGBA, error) {
	bounds1 := img1.Bounds()
	bounds2 := img2.Bounds()

	if bounds1 != bounds2 {
		return nil, errors.New("images have different dimensions")
	}

	result := image.NewRGBA(bounds1)
	for y := bounds1.Min.Y; y < bounds1.Max.Y; y++ {
		for x := bounds1.Min.X; x < bounds1.Max.X; x++ {
			c1 := img1.At(x, y)
			c2 := img2.At(x, y)
			c := blendPixels(c1, c2, 1, params.Brightness, params.Contrast)
			result.Set(x, y, c)
		}
	}

	return result, nil
}

// WriteTo saves the final image to the specified path.
func (c ImageCreator) WriteTo(outputPath string) {
	if c.final == nil {
		log.Fatal(errors.New("final image is nil"))
	}

	finalImageOutput, err := os.Create(outputPath)
	if err != nil {
		log.Fatal(fmt.Errorf("failed to create: %s", err))
	}
	defer finalImageOutput.Close()

	// Save the image using imgio's PNG encoder
	if err := imgio.Save(outputPath, c.final, imgio.PNGEncoder()); err != nil {
		fmt.Println(err)
		return
	}
}

// getImage loads an image from the specified path.
func getImage(imagePath string) image.Image {
	imageSource, err := os.Open(imagePath)
	if err != nil {
		log.Fatalf("failed to open: %s", err)
	}
	defer imageSource.Close()

	imageResult, err := png.Decode(imageSource)
	if err != nil {
		log.Fatalf("failed to decode: %s", err)
	}

	return imageResult
}

// getPaperImage retrieves the default paper texture image.
func getPaperImage() image.Image {
	paperImagePath := "./assets/traits/TEXTURES/PAPERTEXTURE.png"
	img2, ok := m.Get(paperImagePath)
	if !ok {
		panic("Image not found in cache " + paperImagePath) // Consider logging instead of panicking
	}
	return img2
}
