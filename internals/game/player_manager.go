package game

import (
	"scroller_game/internals/config"
	"scroller_game/internals/entities"
	"scroller_game/internals/physics"
)

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

func (g *Game) handlePlayerHit() {
	g.Player.Score--
	g.Player.Hits++
}

func (g *Game) playerShoot() {
	if len(g.Player.Projectiles) >= 3 {
		return // Limit to 3 projectiles
	}
	projectile := &entities.Projectile{
		X:     g.Player.X + g.Player.Width/2, // Center of the player
		Y:     g.Player.Y,
		Width: 5, Height: 10,
		Speed: config.ProjectileSpeed,
	}
	projectile.Hitbox = physics.Hitbox{X: projectile.X, Y: projectile.Y, Width: projectile.Width, Height: projectile.Height}
	g.Player.Projectiles = append(g.Player.Projectiles, projectile)
}
