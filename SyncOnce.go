package main

import (
	"fmt"
	"image"
	"image/png"
	"os"
	"sync"
)

// An example showing how to use sync.Once to perform initialization tasks just once.
// This solution is thread safe and is protected against race conditions.

var loadIconsOnce sync.Once
var icons map[string]image.Image

// Concurrency-safe function
func Icon(name string) image.Image {
	//lazy loading of all the image files is done just once
	loadIconsOnce.Do(loadIcons) //pass the name of the initialization function to the sync.Do() function
	return icons[name]
}

func loadIcons() {
	icons = make(map[string]image.Image)
	icons["spades.png"] = loadIcon("images/spades.png")
	icons["hearts.png"] = loadIcon("images/hearts.png")
	icons["diamonds.png"] = loadIcon("images/diamonds.png")
	icons["clubs.png"] = loadIcon("images/clubs.png")
}

func loadIcon(name string) image.Image {
	imageFile, err := os.Open(name)
	if err != nil {
		panic(err)
	}
	defer imageFile.Close()
	loadedImage, err := png.Decode(imageFile)
	if err != nil {
		panic(err)
	}
	return loadedImage
}

func main() {
	icon := Icon("spades.png")
	fmt.Println(icon)
}
