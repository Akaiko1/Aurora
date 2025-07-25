package game

import (
	"scroller_game/internals/config"
	"scroller_game/internals/entities"
)

func (g *Game) projectilesMovements() {
	// Create trajectory handler for enemy projectiles
	trajectoryHandler := &entities.TrajectoryHandler{
		SpawnedEnemies: g.SpawnedEnemies, // For potential future enemy-enemy interactions
		Player:         g.Player,         // For enemy projectiles to track player
	}

	// Update enemy projectiles using trajectory system
	for _, projectile := range g.Projectiles {
		dx, dy := trajectoryHandler.CalculateMovement(projectile, g.FrameCount)
		projectile.Move(dx, dy)
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
