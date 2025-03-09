package game

import "scroller_game/internals/config"

func (g *Game) projectilesMovements() {
	for _, projectile := range g.Projectiles {
		projectile.Y += projectile.Speed
		projectile.Hitbox.Y += projectile.Speed
	}

	activeProjectiles := g.Projectiles[:0]
	for _, projectile := range g.Projectiles {
		if projectile.Y < config.ScreenHeight {
			activeProjectiles = append(activeProjectiles, projectile)
		}
	}
	g.Projectiles = activeProjectiles
}
