package game

import (
	"github.com/hajimehoshi/ebiten/v2"
)

func (g *Game) playerEvents() {
	if ebiten.IsKeyPressed(ebiten.KeyArrowLeft) && g.Player.X-g.Player.Speed > 0 {
		g.Player.X -= g.Player.Speed
		g.Player.Hitbox.X -= g.Player.Speed
		g.Player.Grazebox.X -= g.Player.Speed
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowRight) && g.Player.X+g.Player.Width+g.Player.Speed < ScreenWidth {
		g.Player.X += g.Player.Speed
		g.Player.Hitbox.X += g.Player.Speed
		g.Player.Grazebox.X += g.Player.Speed
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowUp) && g.Player.Y-g.Player.Speed > 0 {
		g.Player.Y -= g.Player.Speed
		g.Player.Hitbox.Y -= g.Player.Speed
		g.Player.Grazebox.Y -= g.Player.Speed
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowDown) && g.Player.Y+g.Player.Height+g.Player.Speed < ScreenHeight {
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
			g.Projectiles = deleteProjectile(g.Projectiles, idx)
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

func (g *Game) playerProjectilesMovements() {
	// Update player projectiles
	for _, projectile := range g.Player.Projectiles {
		projectile.Y -= projectile.Speed
		projectile.Hitbox.Y -= projectile.Speed
	}

	// Remove off-screen player projectiles
	activeProjectiles := g.Player.Projectiles[:0]
	for _, projectile := range g.Player.Projectiles {
		if projectile.Y > 0 {
			activeProjectiles = append(activeProjectiles, projectile)
		}
	}
	g.Player.Projectiles = activeProjectiles
}

func (g *Game) playerShoot() {
	if len(g.Player.Projectiles) >= 3 {
		return // Limit to 3 projectiles
	}
	projectile := &Projectile{
		X:     g.Player.X + g.Player.Width/2, // Center of the player
		Y:     g.Player.Y,
		Width: 5, Height: 10,
		Speed: projectileSpeed,
	}
	projectile.Hitbox = Hitbox{X: projectile.X, Y: projectile.Y, Width: projectile.Width, Height: projectile.Height}
	g.Player.Projectiles = append(g.Player.Projectiles, projectile)
}
