package inputs

import (
	"image"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
)

func ReadImage(path string) (*ebiten.Image, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	img, _, err := image.Decode(file)
	if err != nil {
		return nil, err
	}
	image := ebiten.NewImageFromImage(img)
	return image, nil
}
