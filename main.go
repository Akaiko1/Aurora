package main

import (
	"log"
	"math/rand"
	"scroller_game/internals/config"
	"scroller_game/internals/entities"
	"scroller_game/internals/game"
	"scroller_game/internals/inputs"
	"scroller_game/internals/physics"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
)

func NewGame() *game.Game {
	player := &entities.Player{
		X:           config.ScreenWidth / 2,
		Y:           config.ScreenHeight / 2,
		Width:       config.EntitySize,
		Height:      config.EntitySize,
		Speed:       config.PlayerSpeed,
		Projectiles: []*entities.Projectile{},
		Hitbox:      physics.Hitbox{Width: config.PlayerHitboxSize, Height: config.PlayerHitboxSize},
		Grazebox:    physics.Hitbox{Width: config.PlayerGrazeboxSize, Height: config.PlayerGrazeboxSize},
	}

	// Initialize player weapons
	player.InitializeWeapons()

	player.Hitbox.CenterOn(player.X+player.Width/2, player.Y+player.Height/2)
	player.Grazebox.CenterOn(player.X+player.Width/2, player.Y+player.Height/2)
	randomSource := rand.New(rand.NewSource(time.Now().UnixNano()))

	return &game.Game{
		Player:            player,
		Projectiles:       []*entities.Projectile{},
		Enemies:           []*entities.Enemy{},
		RandomSource:      randomSource,
		FlagHitboxes:      false,
		State:             game.SwitchLevel,
		Scenarios:         game.GetGameScenarios(),
		Background:        game.InitBackground(randomSource),
		SpatialGrid:       physics.NewSpatialGrid(config.SpatialGridCellSize),
		GameOverSelection: 0,
	}
}

func main() {
	// Set up embedded asset loading
	inputs.LoadEmbeddedImageFunc = LoadEmbeddedImage
	inputs.LoadEmbeddedFontFunc = LoadEmbeddedFont
	inputs.UseEmbeddedAssets = true // Change to false for development with file system

	// Initialize assets after setting embedded flags
	game.InitializeAssets()

	game_loop := NewGame()

	ebiten.SetWindowSize(config.ScreenWidth, config.ScreenHeight)
	ebiten.SetWindowTitle("Aurora Scroller")
	if err := ebiten.RunGame(game_loop); err != nil {
		log.Fatal(err)
	}
}
