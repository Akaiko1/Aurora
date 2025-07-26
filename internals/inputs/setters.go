package inputs

import (
	"bytes"
	"log"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

var (
	// External function to be set by main package for embedded font loading
	LoadEmbeddedFontFunc func(path string) ([]byte, error)
)

// tryLoadFont tries embedded first, then fallback to file
func tryLoadFont(path string) ([]byte, error) {
	if UseEmbeddedAssets && LoadEmbeddedFontFunc != nil {
		if fontBytes, err := LoadEmbeddedFontFunc(path); err == nil {
			return fontBytes, nil
		}
		log.Printf("Embedded font load failed for %s, trying file system", path)
	}
	return os.ReadFile(path)
}

func SetFontAndImages() (*text.GoTextFace, *ebiten.Image) {
	fontBytes, err := tryLoadFont("assets/fonts/Jacquard12-Regular.ttf")
	if err != nil {
		log.Fatal(err)
	}

	mplusFaceSource, err := text.NewGoTextFaceSource(bytes.NewReader(fontBytes))
	if err != nil {
		log.Fatal(err)
	}

	mplusNormalFace := &text.GoTextFace{
		Source: mplusFaceSource,
		Size:   24,
	}

	// Load new sprites
	LoadSprites()

	// Return nil for second parameter since we no longer use frames
	return mplusNormalFace, nil
}
