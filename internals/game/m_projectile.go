package game

import (
	"scroller_game/internals/config"
	"scroller_game/internals/entities"
)

func (g *Game) updateProjectiles() {
	// Create trajectory handler once for all projectiles
	trajectoryHandler := &entities.TrajectoryHandler{
		SpawnedEnemies: g.SpawnedEnemies,
		Player:         g.Player,
	}

	// Update all enemy projectiles
	for _, projectile := range g.Projectiles {
		dx, dy := trajectoryHandler.CalculateMovement(projectile, g.FrameCount)
		projectile.Move(dx, dy)
	}

	// Update all player projectiles
	for _, projectile := range g.Player.Projectiles {
		dx, dy := trajectoryHandler.CalculateMovement(projectile, g.FrameCount)
		projectile.Move(dx, dy)
	}

	// Remove off-screen enemy projectiles (move down, positive Y)
	activeEnemyProjectiles := g.Projectiles[:0]
	for _, projectile := range g.Projectiles {
		if projectile.Y < config.ScreenHeight+projectile.Height {
			activeEnemyProjectiles = append(activeEnemyProjectiles, projectile)
		}
	}
	g.Projectiles = activeEnemyProjectiles

	// Remove off-screen player projectiles (move up, negative Y)
	activePlayerProjectiles := g.Player.Projectiles[:0]
	for _, projectile := range g.Player.Projectiles {
		if projectile.Y > -projectile.Height {
			activePlayerProjectiles = append(activePlayerProjectiles, projectile)
		}
	}
	g.Player.Projectiles = activePlayerProjectiles
}
