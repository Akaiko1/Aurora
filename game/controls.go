package game

import "github.com/hajimehoshi/ebiten/v2"

func (g *Game) playerControls() {
	if ebiten.IsKeyPressed(ebiten.KeyArrowLeft) {
		g.player.X -= g.player.Speed
		g.player.Hitbox.X -= g.player.Speed
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowRight) {
		g.player.X += g.player.Speed
		g.player.Hitbox.X += g.player.Speed
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowUp) {
		g.player.Y -= g.player.Speed
		g.player.Hitbox.Y -= g.player.Speed
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowDown) {
		g.player.Y += g.player.Speed
		g.player.Hitbox.Y += g.player.Speed
	}

	// Player shooting
	if ebiten.IsKeyPressed(ebiten.KeySpace) {
		if g.frameCount%5 == 0 {
			g.playerShoot()
		}
	}
}
