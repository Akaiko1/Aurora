package main

import (
	"log"
	"math/rand"
	"scroller_game/internals/config"
	"scroller_game/internals/entities"
	"scroller_game/internals/game"
	"scroller_game/internals/physics"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
)

func NewGame() *game.Game {
	player := &entities.Player{
		X:           config.ScreenWidth / 2,
		Y:           config.ScreenHeight / 2,
		Width:       32,
		Height:      32,
		Speed:       config.PlayerSpeed,
		Projectiles: []*entities.Projectile{},
		Hitbox:      physics.Hitbox{Width: 10, Height: 10},
		Grazebox:    physics.Hitbox{Width: 50, Height: 50},
	}

	player.Hitbox.CenterOn(player.X+player.Width/2, player.Y+player.Height/2)
	player.Grazebox.CenterOn(player.X+player.Width/2, player.Y+player.Height/2)
	randomSource := rand.New(rand.NewSource(time.Now().UnixNano()))

	return &game.Game{
		Player:       player,
		Projectiles:  []*entities.Projectile{},
		Enemies:      []*entities.Enemy{},
		RandomSource: randomSource,
		FlagHitboxes: false,
		State:        game.SwitchLevel,
		Scenarios:    game.GetGameScenarios(),
		Background:   game.InitBackground(randomSource),
	}
}

func main() {
	game_loop := NewGame()

	ebiten.SetWindowSize(config.ScreenWidth, config.ScreenHeight)
	ebiten.SetWindowTitle("S-lav-o")
	if err := ebiten.RunGame(game_loop); err != nil {
		log.Fatal(err)
	}
}
