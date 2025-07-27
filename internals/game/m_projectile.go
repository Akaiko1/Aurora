// Package game contains the core game logic and entity management for Aurora.
package game

import (
	"scroller_game/internals/config"
	"scroller_game/internals/entities"
)

// updateProjectiles handles movement and cleanup for all active projectiles.
// This function processes both enemy and player projectiles by calculating
// their trajectory-based movement and removing projectiles that have moved
// off-screen to prevent memory leaks and maintain performance.
func (g *Game) updateProjectiles() {
	// Create trajectory handler once for all projectiles to avoid repeated allocation
	trajectoryHandler := &entities.TrajectoryHandler{
		SpawnedEnemies: g.SpawnedEnemies,
		Player:         g.Player,
	}

	// Update all enemy projectiles using their assigned trajectory patterns
	for _, projectile := range g.Projectiles {
		dx, dy := trajectoryHandler.CalculateMovement(projectile, g.FrameCount)
		projectile.Move(dx, dy)
	}

	// Update all player projectiles using their assigned trajectory patterns
	for _, projectile := range g.Player.Projectiles {
		dx, dy := trajectoryHandler.CalculateMovement(projectile, g.FrameCount)
		projectile.Move(dx, dy)
	}

	// Remove all off-screen projectiles using unified bounds checking
	g.Projectiles = g.removeOffScreenProjectiles(g.Projectiles)
	g.Player.Projectiles = g.removeOffScreenProjectiles(g.Player.Projectiles)
}

// isProjectileOnScreen checks if a projectile is still within screen boundaries.
// It returns true if any part of the projectile is visible on screen, allowing
// for a small buffer zone to ensure smooth visual transitions.
func (g *Game) isProjectileOnScreen(projectile *entities.Projectile) bool {
	return projectile.X > -projectile.Width &&
		projectile.X < config.ScreenWidth+projectile.Width &&
		projectile.Y > -projectile.Height &&
		projectile.Y < config.ScreenHeight+projectile.Height
}

// removeOffScreenProjectiles filters out projectiles that have moved beyond screen boundaries.
// This function works for both enemy and player projectiles, making the cleanup logic DRY.
func (g *Game) removeOffScreenProjectiles(projectiles []*entities.Projectile) []*entities.Projectile {
	activeProjectiles := projectiles[:0]
	for _, projectile := range projectiles {
		if g.isProjectileOnScreen(projectile) {
			activeProjectiles = append(activeProjectiles, projectile)
		}
	}
	return activeProjectiles
}
