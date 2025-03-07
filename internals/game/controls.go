package game

import (
	"scroller_game/internals/config"
	"scroller_game/internals/events"

	"github.com/hajimehoshi/ebiten/v2"
)

func (g *Game) playerEvents() {
	if ebiten.IsKeyPressed(ebiten.KeyArrowLeft) && g.Player.X-g.Player.Speed > 0 {
		g.Player.X -= g.Player.Speed
		g.Player.Hitbox.X -= g.Player.Speed
		g.Player.Grazebox.X -= g.Player.Speed
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowRight) && g.Player.X+g.Player.Width+g.Player.Speed < config.ScreenWidth {
		g.Player.X += g.Player.Speed
		g.Player.Hitbox.X += g.Player.Speed
		g.Player.Grazebox.X += g.Player.Speed
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowUp) && g.Player.Y-g.Player.Speed > 0 {
		g.Player.Y -= g.Player.Speed
		g.Player.Hitbox.Y -= g.Player.Speed
		g.Player.Grazebox.Y -= g.Player.Speed
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowDown) && g.Player.Y+g.Player.Height+g.Player.Speed < config.ScreenHeight {
		g.Player.Y += g.Player.Speed
		g.Player.Hitbox.Y += g.Player.Speed
		g.Player.Grazebox.Y += g.Player.Speed
	}

	// Toggle hitboxes with the B key
	if ebiten.IsKeyPressed(ebiten.KeyB) {
		if g.FrameCount%6 == 0 {
			g.FlagHitboxes = !g.FlagHitboxes
		}
	}

	// Player shooting
	if ebiten.IsKeyPressed(ebiten.KeySpace) {
		if g.FrameCount%5 == 0 {
			g.playerShoot()
			g.Player.IsAttacking = true
			g.Player.AttackStartFrame = g.FrameCount
		}
	}

	if g.Player.IsAttacking && (g.FrameCount-g.Player.AttackStartFrame+120)%120 >= 18 {
		g.Player.IsAttacking = false
	}

	// Playes projectiles interaction
	for idx, projectile := range g.Projectiles {
		if g.Player.Hitbox.Intersects(&projectile.Hitbox) {
			g.handlePlayerHit()
			g.Projectiles = events.DeleteProjectile(g.Projectiles, idx)
			break
		}

		if g.Player.Grazebox.Intersects(&projectile.Hitbox) && g.Player.Grazing != projectile {
			g.Player.Grazing = projectile
			break
		}

		if g.Player.Grazing != nil && !g.Player.Grazebox.Intersects(&g.Player.Grazing.Hitbox) {
			g.Player.Grazing = nil
			g.Player.Score++
		}

	}
}
