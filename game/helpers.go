package game

import (
	"image"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
)

func (hb *Hitbox) CenterOn(centerX, centerY float32) {
	hb.X = centerX - hb.Width/2
	hb.Y = centerY - hb.Height/2
}

// Deletes projectile from slice by index
func deleteProjectile(slice []*Projectile, index int) []*Projectile {
	return append(slice[:index], slice[index+1:]...)
}

// Deletes enemy from slice by reference
func deleteEnemy(slice []*Enemy, enemy *Enemy) ([]*Enemy, bool) {
	for i, e := range slice {
		if e == enemy {
			return append(slice[:i], slice[i+1:]...), true
		}
	}
	return slice, false
}

// Check if two hitboxes intersect
func (hb *Hitbox) Intersects(other *Hitbox) bool {
	return hb.X < other.X+other.Width &&
		hb.X+hb.Width > other.X &&
		hb.Y < other.Y+other.Height &&
		hb.Y+hb.Height > other.Y
}

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
