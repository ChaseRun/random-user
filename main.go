// Package randomuser provides embedded random user portrait images.
// It contains 200 high-quality portrait images (100 men, 100 women) from randomuser.me
// that can be accessed programmatically without external API calls.
package randomuser

import (
	"embed"
	"io/fs"
	"math/rand"
	"path"
	"sort"
	"time"
)

// Embed all portrait images
//
//go:embed images/men/*.jpg
var menFS embed.FS

//go:embed images/women/*.jpg
var womenFS embed.FS

var (
	menImages   [][]byte
	womenImages [][]byte
	allImages   [][]byte
)

func init() {
	// Load all men images dynamically
	menImages = loadImages(menFS, "images/men")

	// Load all women images dynamically
	womenImages = loadImages(womenFS, "images/women")

	// Combine all images for random selection
	allImages = append(allImages, menImages...)
	allImages = append(allImages, womenImages...)

	// Seed random number generator
	rand.Seed(time.Now().UnixNano())
}

// loadImages dynamically loads all .jpg files from the embedded filesystem
func loadImages(fsys embed.FS, dir string) [][]byte {
	var images [][]byte
	var filenames []string

	// Read directory
	fs.WalkDir(fsys, dir, func(filepath string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		// Skip directories and non-jpg files
		if d.IsDir() || path.Ext(filepath) != ".jpg" {
			return nil
		}
		filenames = append(filenames, filepath)
		return nil
	})

	// Sort filenames for consistent ordering
	sort.Strings(filenames)

	// Load each image
	for _, filename := range filenames {
		data, _ := fsys.ReadFile(filename)
		images = append(images, data)
	}

	return images
}

// RandomImage returns a random portrait image (either man or woman).
// The returned bytes are JPEG image data.
func RandomImage() []byte {
	return allImages[rand.Intn(len(allImages))]
}

// RandomManImage returns a random man's portrait image.
// The returned bytes are JPEG image data.
func RandomManImage() []byte {
	return menImages[rand.Intn(len(menImages))]
}

// RandomWomanImage returns a random woman's portrait image.
// The returned bytes are JPEG image data.
func RandomWomanImage() []byte {
	return womenImages[rand.Intn(len(womenImages))]
}
