package inputs

import (
	"image"
	"log"
	"os"
	"scroller_game/internals/entities"

	"github.com/hajimehoshi/ebiten/v2"
)

var (
	PlayerSprite *ebiten.Image
	EnemySprites map[string]*ebiten.Image // Map enemy type names to their sprites
	BgTiles      *ebiten.Image
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
	return ebiten.NewImageFromImage(img), nil
}

// LoadSprites loads all game sprites
func LoadSprites() {
	var err error

	// Initialize the enemy sprites map
	EnemySprites = make(map[string]*ebiten.Image)

	// Player sprite is REQUIRED - game stops if it can't load
	PlayerSprite, err = ReadImage("assets/sprites/player.png")
	if err != nil {
		log.Fatalf("CRITICAL: Could not load required player.png: %v", err)
	}

	// Load enemy sprites automatically by discovering all enemy types
	allEnemyTypes := entities.GetAllEnemyTypes()

	// First, load spider sprite as the fallback (it's required)
	spiderSprite, err := ReadImage("assets/sprites/spider.png")
	if err != nil {
		log.Fatalf("CRITICAL: Could not load required spider.png (fallback sprite): %v", err)
	}
	EnemySprites["spider"] = spiderSprite

	// Load other enemy sprites, using spider as fallback if they fail
	for _, enemyType := range allEnemyTypes {
		enemyTypeName := entities.GetEnemyTypeName(enemyType)

		// Skip spider since we already loaded it above
		if enemyTypeName == "spider" {
			continue
		}

		spritePath := "assets/sprites/" + enemyTypeName + ".png"
		sprite, err := ReadImage(spritePath)
		if err != nil {
			log.Printf("Warning: Could not load %s, using spider sprite as fallback: %v", spritePath, err)
			sprite = spiderSprite // Use spider sprite as fallback
		}
		EnemySprites[enemyTypeName] = sprite
	}

	BgTiles, err = ReadImage("assets/sprites/bg_tiles.png")
	if err != nil {
		log.Printf("Warning: Could not load bg_tiles.png: %v", err)
	}
}

// GetEnemySprite returns the sprite for a given enemy type
func GetEnemySprite(enemyType string) *ebiten.Image {
	if sprite, exists := EnemySprites[enemyType]; exists {
		return sprite
	}
	// Return spider as default fallback (guaranteed to exist)
	return EnemySprites["spider"]
}
