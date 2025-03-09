package game

import (
	"image"
	"image/color"
	"math/rand"
	"scroller_game/internals/config"
	"scroller_game/internals/inputs"

	"github.com/hajimehoshi/ebiten/v2"
)

var (
	bgGrass    *ebiten.Image
	bgNoGrass  *ebiten.Image
	bgTileSize = 64 // Size of each background tile in pixels
)

type BackgroundTile struct {
	X, Y       float64
	HasGrass   bool
	TileWidth  int
	TileHeight int
}

type Background struct {
	Tiles      []BackgroundTile
	TileSize   int
	GridWidth  int
	GridHeight int
}

func InitBackground(randomSource *rand.Rand) *Background {
	// Load the single spritesheet containing both textures
	spritesheet, err := inputs.ReadImage("assets/sprites/bg_tiles.png")
	if err != nil {
		// Fallback: Create simple textures if image can't be loaded
		bgGrass = ebiten.NewImage(bgTileSize, bgTileSize)
		bgGrass.Fill(color.RGBA{0, 120, 0, 255}) // Medium green

		bgNoGrass = ebiten.NewImage(bgTileSize, bgTileSize)
		bgNoGrass.Fill(color.RGBA{101, 67, 33, 255}) // Brown
	} else {
		// Extract the two textures from the spritesheet
		// First texture (grass) is at x=0
		bgGrass = spritesheet.SubImage(image.Rect(0, 0, 64, 64)).(*ebiten.Image)
		// Second texture (no grass) is at x=64
		bgNoGrass = spritesheet.SubImage(image.Rect(64, 0, 128, 64)).(*ebiten.Image)
	}

	// Calculate how many tiles we need to fill the screen
	gridWidth := (config.ScreenWidth + bgTileSize - 1) / bgTileSize
	gridHeight := (config.ScreenHeight + bgTileSize - 1) / bgTileSize

	// Create the background tiles
	tiles := make([]BackgroundTile, 0, gridWidth*gridHeight)

	for y := 0; y < gridHeight; y++ {
		for x := 0; x < gridWidth; x++ {
			// Randomly decide if a tile has grass (70% chance for grass)
			hasGrass := randomSource.Float32() < 0.7

			tile := BackgroundTile{
				X:          float64(x * bgTileSize),
				Y:          float64(y * bgTileSize),
				HasGrass:   hasGrass,
				TileWidth:  bgTileSize,
				TileHeight: bgTileSize,
			}
			tiles = append(tiles, tile)
		}
	}

	return &Background{
		Tiles:      tiles,
		TileSize:   bgTileSize,
		GridWidth:  gridWidth,
		GridHeight: gridHeight,
	}
}

func (b *Background) Draw(screen *ebiten.Image) {
	for _, tile := range b.Tiles {
		opts := &ebiten.DrawImageOptions{}
		opts.GeoM.Translate(tile.X, tile.Y)

		if tile.HasGrass {
			screen.DrawImage(bgGrass, opts)
		} else {
			screen.DrawImage(bgNoGrass, opts)
		}
	}
}
