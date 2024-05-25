package main

import (
	"log"
	"scroller_game/game"

	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	game_loop := game.NewGame()

	ebiten.SetWindowSize(game.ScreenWidth, game.ScreenHeight)
	ebiten.SetWindowTitle("Side Scroller Game")
	if err := ebiten.RunGame(game_loop); err != nil {
		log.Fatal(err)
	}
}
