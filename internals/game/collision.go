package game

import "scroller_game/internals/entities"

// ProjectileWrapper wraps a projectile with its index for efficient removal
type ProjectileWrapper struct {
	Projectile *entities.Projectile
	Index      int
}

// UpdateCollisions efficiently handles all collision detection using spatial partitioning
func (g *Game) UpdateCollisions() {
	// Clear spatial grid for this frame
	g.SpatialGrid.Clear()

	// Insert all player projectiles into spatial grid
	for i, projectile := range g.Player.Projectiles {
		wrapper := &ProjectileWrapper{
			Projectile: projectile,
			Index:      i,
		}
		g.SpatialGrid.Insert(
			projectile.X, projectile.Y,
			projectile.Width, projectile.Height,
			wrapper,
		)
	}

	// Check each enemy against nearby projectiles only
	enemiesToRemove := make([]*entities.Enemy, 0)
	projectileIndicesToRemove := make([]int, 0)

	for _, enemy := range g.SpawnedEnemies {
		// Query spatial grid for projectiles near this enemy
		nearbyObjects := g.SpatialGrid.Query(
			enemy.X, enemy.Y,
			enemy.Width, enemy.Height,
		)

		// Check collision with nearby projectiles only
		for _, obj := range nearbyObjects {
			wrapper := obj.(*ProjectileWrapper)
			if enemy.Hitbox.Intersects(&wrapper.Projectile.Hitbox) {
				enemiesToRemove = append(enemiesToRemove, enemy)

				// Only remove projectile if it's not piercing type
				if wrapper.Projectile.Type == entities.ProjectileNormal {
					projectileIndicesToRemove = append(projectileIndicesToRemove, wrapper.Index)
				}
				break // Enemy can only be hit once per frame
			}
		}
	}

	// Remove hit enemies
	for _, enemy := range enemiesToRemove {
		g.handleEnemyHit(enemy)
	}

	// Remove projectiles that hit enemies (only normal projectiles, in reverse order to maintain indices)
	for i := len(projectileIndicesToRemove) - 1; i >= 0; i-- {
		idx := projectileIndicesToRemove[i]
		if idx < len(g.Player.Projectiles) {
			g.Player.Projectiles = append(g.Player.Projectiles[:idx], g.Player.Projectiles[idx+1:]...)
		}
	}
}
