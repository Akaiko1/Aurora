package inputs

import (
	"bytes"
	"log"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

func SetFontAndImages() (*text.GoTextFace, *ebiten.Image) {
	fontBytes, err := os.ReadFile("assets/fonts/Jacquard12-Regular.ttf")
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

	frames, _ := ReadImage("assets/sprites/animations.png")

	return mplusNormalFace, frames
}
