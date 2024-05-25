package game

func (g *Game) enemySpawn() {
	// Spawn a new enemy
	if len(g.enemies) > 1 {
		return
	}

	enemy := &Enemy{X: 50, Y: 50, Width: 32, Height: 32, SpeedX: 1, SpeedY: 1}
	enemy.Hitbox = Hitbox{X: enemy.X, Y: enemy.Y, Width: enemy.Width, Height: enemy.Height}
	g.setRandomDirection(enemy)
	g.enemies = append(g.enemies, enemy)
}

func (enemy *Enemy) enemyShoot(projectiles *[]*Projectile) {
	projectile := &Projectile{
		X:     enemy.X + 13, // Center of the enemy
		Y:     enemy.Y + 32,
		Width: 5, Height: 10,
		Speed: projectileSpeed,
	}
	projectile.Hitbox = Hitbox{X: projectile.X, Y: projectile.Y, Width: projectile.Width, Height: projectile.Height}
	*projectiles = append(*projectiles, projectile)
}

// enemyActions handles the actions of the enemy in the game.
// It sets a random direction for the enemy every 120 frames,
// updates the enemy's position based on its speed,
// and reverses the speed if the enemy reaches the boundaries of the screen.
func (g *Game) enemyActions(enemy *Enemy) {
	if g.frameCount%120 == 0 {
		g.setRandomDirection(enemy)
	}

	enemy.X += enemy.SpeedX
	enemy.Y += enemy.SpeedY
	enemy.Hitbox.X += enemy.SpeedX
	enemy.Hitbox.Y += enemy.SpeedY

	if enemy.X < 0 || enemy.X+32 > ScreenWidth {
		enemy.SpeedX = -enemy.SpeedX
	}
	if enemy.Y < 0 || enemy.Y+32 > ScreenHeight {
		enemy.SpeedY = -enemy.SpeedY
	}

	// Enemy shooting logic
	if g.frameCount%enemyShootInterval == 0 {
		enemy.enemyShoot(&g.projectiles) // Pass the address of g.projectiles
	}

	// Check for collisions with the enemy
	for _, projectile := range g.player.Projectiles {
		if enemy.Hitbox.Intersects(&projectile.Hitbox) {
			g.handleEnemyHit(enemy) // Handle the enemy hit
			break
		}
	}
}

func (g *Game) projectilesMovements() {
	for _, projectile := range g.projectiles {
		projectile.Y += projectile.Speed
		projectile.Hitbox.Y += projectile.Speed
	}

	activeProjectiles := g.projectiles[:0]
	for _, projectile := range g.projectiles {
		if projectile.Y < ScreenHeight {
			activeProjectiles = append(activeProjectiles, projectile)
		}
	}
	g.projectiles = activeProjectiles
}
