package main

import (
	"embed"
	"image"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

// Embed all assets into the binary
//
//go:embed assets/sprites/*.png
var SpriteFiles embed.FS

//go:embed assets/fonts/*.ttf
var FontFiles embed.FS

// Flag to control whether to use embedded assets or files
var UseEmbeddedAssets = true

// LoadEmbeddedImage loads an image from embedded filesystem
func LoadEmbeddedImage(path string) (*ebiten.Image, error) {
	file, err := SpriteFiles.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	img, _, err := image.Decode(file)
	if err != nil {
		return nil, err
	}
	return ebiten.NewImageFromImage(img), nil
}

// LoadEmbeddedFont loads a font from embedded filesystem
func LoadEmbeddedFont(path string) ([]byte, error) {
	return FontFiles.ReadFile(path)
}

// GetAvailableSprites returns list of available sprite files
func GetAvailableSprites() []string {
	entries, err := SpriteFiles.ReadDir("assets/sprites")
	if err != nil {
		log.Printf("Error reading embedded sprites: %v", err)
		return nil
	}

	var sprites []string
	for _, entry := range entries {
		if !entry.IsDir() {
			sprites = append(sprites, "assets/sprites/"+entry.Name())
		}
	}
	return sprites
}
