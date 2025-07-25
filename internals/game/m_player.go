package game

import (
	"scroller_game/internals/entities"
)

func (g *Game) playerProjectilesMovements() {
	// Create trajectory handler
	trajectoryHandler := &entities.TrajectoryHandler{
		SpawnedEnemies: g.SpawnedEnemies,
		Player:         g.Player, // For consistency, though player projectiles don't track player
	}

	// Update player projectiles using trajectory system
	for _, projectile := range g.Player.Projectiles {
		dx, dy := trajectoryHandler.CalculateMovement(projectile, g.FrameCount)
		projectile.Move(dx, dy)
	}

	// Remove off-screen player projectiles
	activeProjectiles := g.Player.Projectiles[:0]
	for _, projectile := range g.Player.Projectiles {
		// Player projectiles move up (negative Y), keep while Y is above top of screen
		if projectile.Y > -projectile.Height {
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
	projectile := g.Player.CreateProjectile(g.FrameCount)
	if projectile != nil {
		g.Player.Projectiles = append(g.Player.Projectiles, projectile)
	}
}
