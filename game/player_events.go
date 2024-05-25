package game

import (
	"github.com/hajimehoshi/ebiten/v2"
)

func (g *Game) playerEvents() {
	if ebiten.IsKeyPressed(ebiten.KeyArrowLeft) && g.player.X-g.player.Speed > 0 {
		g.player.X -= g.player.Speed
		g.player.Hitbox.X -= g.player.Speed
		g.player.Grazebox.X -= g.player.Speed
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowRight) && g.player.X+g.player.Width+g.player.Speed < ScreenWidth {
		g.player.X += g.player.Speed
		g.player.Hitbox.X += g.player.Speed
		g.player.Grazebox.X += g.player.Speed
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowUp) && g.player.Y-g.player.Speed > 0 {
		g.player.Y -= g.player.Speed
		g.player.Hitbox.Y -= g.player.Speed
		g.player.Grazebox.Y -= g.player.Speed
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowDown) && g.player.Y+g.player.Height+g.player.Speed < ScreenHeight {
		g.player.Y += g.player.Speed
		g.player.Hitbox.Y += g.player.Speed
		g.player.Grazebox.Y += g.player.Speed
	}

	// Player shooting
	if ebiten.IsKeyPressed(ebiten.KeySpace) {
		if g.frameCount%5 == 0 {
			g.playerShoot()
		}
	}

	// Playes projectiles interaction
	for idx, projectile := range g.projectiles {
		if g.player.Hitbox.Intersects(&projectile.Hitbox) {
			g.handlePlayerHit()
			g.projectiles = deleteProjectile(g.projectiles, idx)
			break
		}

		if g.player.Grazebox.Intersects(&projectile.Hitbox) && g.player.Grazing != projectile {
			g.player.Grazing = projectile
			break
		}

		if g.player.Grazing != nil && !g.player.Grazebox.Intersects(&g.player.Grazing.Hitbox) {
			g.player.Grazing = nil
			g.player.Score++
		}

	}
}
