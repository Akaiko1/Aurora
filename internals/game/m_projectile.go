package game

import "scroller_game/internals/config"

func (g *Game) projectilesMovements() {
	for _, projectile := range g.Projectiles {
		projectile.Move(0, projectile.Speed*projectile.Direction)
	}

	activeProjectiles := g.Projectiles[:0]
	for _, projectile := range g.Projectiles {
		// Enemy projectiles move down (positive Y), remove when Y > screen height
		if projectile.Y < config.ScreenHeight+projectile.Height {
			activeProjectiles = append(activeProjectiles, projectile)
		}
	}
	g.Projectiles = activeProjectiles
}
