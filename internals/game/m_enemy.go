// Package game contains the core game logic and entity management for Aurora.
package game

import (
	"scroller_game/internals/config"
	"scroller_game/internals/entities"
	"scroller_game/internals/physics"
)

// handleEnemyHit processes when an enemy takes damage from the player.
// It applies damage to the enemy and awards enemy-specific points if the enemy is defeated.
func (g *Game) handleEnemyHit(enemy *entities.Enemy) {
	defeated := enemy.TakeDamage(1) // Each hit does 1 damage
	if defeated {
		g.Player.Score += enemy.ScoreValue
	}
}

// removeDefeatedEnemies removes multiple enemies from the spawned enemies list.
func (g *Game) removeDefeatedEnemies(enemiesToRemove []*entities.Enemy) {
	if len(enemiesToRemove) == 0 {
		return
	}
	
	// Create a lookup map of enemies to remove
	removeMap := make(map[*entities.Enemy]bool, len(enemiesToRemove))
	for _, enemy := range enemiesToRemove {
		removeMap[enemy] = true
	}
	
	// Filter out enemies that should be removed
	activeEnemies := g.SpawnedEnemies[:0]
	for _, enemy := range g.SpawnedEnemies {
		if !removeMap[enemy] {
			activeEnemies = append(activeEnemies, enemy)
		}
	}
	g.SpawnedEnemies = activeEnemies
}

// setRandomDirection assigns random movement velocities to an enemy.
// The enemy will receive speed values between -EnemyMaxSpeed and +EnemyMaxSpeed
// for both X and Y axes, creating unpredictable movement patterns.
func (g *Game) setRandomDirection(enemy *entities.Enemy) {
	x_r := g.RandomSource.Float32()
	y_r := g.RandomSource.Float32()
	enemy.SpeedX = (x_r * config.EnemySpeedRange) - config.EnemyMaxSpeed
	enemy.SpeedY = (y_r * config.EnemySpeedRange) - config.EnemyMaxSpeed
}

// enemySpawn manages the spawning of enemies from the waiting queue.
// It limits spawned enemies to prevent overwhelming the player and ensures
// there are enemies available to spawn. The first enemy in the queue is
// spawned with a hitbox and random movement direction.
func (g *Game) enemySpawn() {
	if len(g.SpawnedEnemies) > 1 || len(g.Enemies) == 0 {
		return
	}

	enemy := g.Enemies[0]
	enemy.Hitbox = physics.Hitbox{X: enemy.X, Y: enemy.Y, Width: enemy.Width, Height: enemy.Height}
	g.setRandomDirection(enemy)
	g.SpawnedEnemies = append(g.SpawnedEnemies, enemy)

	g.Enemies = g.Enemies[1:]
}

// enemyActions handles all behavior for an active enemy during each frame.
// This includes periodic direction changes, movement updates, screen boundary
// collision detection with direction reversal, shooting logic with attack
// animations, and attack state management based on frame timing.
func (g *Game) enemyActions(enemy *entities.Enemy) {
	// Change direction periodically for unpredictable movement
	if g.FrameCount%config.EnemyDirectionChangeInterval == 0 {
		g.setRandomDirection(enemy)
	}

	enemy.Move(enemy.SpeedX, enemy.SpeedY)

	// Bounce off screen edges
	if enemy.X < 0 || enemy.X+enemy.Width > config.ScreenWidth {
		enemy.SpeedX = -enemy.SpeedX
	}
	if enemy.Y < 0 || enemy.Y+enemy.Height > config.ScreenHeight {
		enemy.SpeedY = -enemy.SpeedY
	}

	// Handle enemy weapon firing and attack animations
	if enemy.CanFire(g.FrameCount) {
		enemy.EnemyShoot(&g.Projectiles, g.FrameCount)
		enemy.IsAttacking = true
		enemy.AttackStartFrame = g.FrameCount
		enemy.Weapon.Fire(g.FrameCount)
	}

	// End attack animation after the configured duration
	if enemy.IsAttacking && (g.FrameCount-enemy.AttackStartFrame+120)%120 >= config.AttackAnimationDuration {
		enemy.IsAttacking = false
	}
}
