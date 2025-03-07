package game

import (
	"scroller_game/internals/config"
	"scroller_game/internals/entities"
	"scroller_game/internals/events"
	"scroller_game/internals/physics"
)

func (g *Game) handleEnemyHit(enemy *entities.Enemy) {
	// Remove the enemy from the list
	g.SpawnedEnemies, _ = events.DeleteEnemy(g.SpawnedEnemies, enemy)
	g.Player.Score++
}

func (g *Game) setRandomDirection(enemy *entities.Enemy) {
	x_r := g.RandomSource.Float32()
	y_r := g.RandomSource.Float32()
	enemy.SpeedX = (x_r * 4) - 2 // Speed between -2 and 2
	enemy.SpeedY = (y_r * 4) - 2 // Speed between -2 and 2
}

func (g *Game) enemySpawn() {
	if len(g.SpawnedEnemies) > 1 || len(g.Enemies) == 0 {
		return
	}

	enemy := g.Enemies[0]
	enemy.Hitbox = physics.Hitbox{X: enemy.X, Y: enemy.Y, Width: enemy.Width, Height: enemy.Height}
	g.setRandomDirection(enemy)
	g.SpawnedEnemies = append(g.SpawnedEnemies, enemy)

	// Remove the enemy from the list
	g.Enemies = g.Enemies[1:]
}

func (g *Game) enemyActions(enemy *entities.Enemy) {
	if g.FrameCount%90 == 0 {
		g.setRandomDirection(enemy)
	}

	enemy.X += enemy.SpeedX
	enemy.Y += enemy.SpeedY
	enemy.Hitbox.X += enemy.SpeedX
	enemy.Hitbox.Y += enemy.SpeedY

	if enemy.X < 0 || enemy.X+32 > config.ScreenWidth {
		enemy.SpeedX = -enemy.SpeedX
	}
	if enemy.Y < 0 || enemy.Y+32 > config.ScreenHeight {
		enemy.SpeedY = -enemy.SpeedY
	}

	// Enemy shooting logic
	if g.FrameCount%config.EnemyShootInterval == 0 {
		enemy.EnemyShoot(&g.Projectiles) // Pass the address of g.projectiles
		enemy.IsAttacking = true
		enemy.AttackStartFrame = g.FrameCount
	}

	if enemy.IsAttacking && (g.FrameCount-enemy.AttackStartFrame+120)%120 >= 12 {
		enemy.IsAttacking = false
	}

	// Check for collisions with the enemy
	for _, projectile := range g.Player.Projectiles {
		if enemy.Hitbox.Intersects(&projectile.Hitbox) {
			g.handleEnemyHit(enemy) // Handle the enemy hit
			break
		}
	}
}
