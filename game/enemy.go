package game

func (g *Game) enemySpawn() {
	if len(g.Spawned) > 1 || len(g.Enemies) == 0 {
		return
	}

	enemy := g.Enemies[0]
	enemy.Hitbox = Hitbox{X: enemy.X, Y: enemy.Y, Width: enemy.Width, Height: enemy.Height}
	g.setRandomDirection(enemy)
	g.Spawned = append(g.Spawned, enemy)

	// Remove the enemy from the list
	g.Enemies = g.Enemies[1:]
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
	if g.FrameCount%90 == 0 {
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
	if g.FrameCount%enemyShootInterval == 0 {
		enemy.enemyShoot(&g.Projectiles) // Pass the address of g.projectiles
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

func (g *Game) projectilesMovements() {
	for _, projectile := range g.Projectiles {
		projectile.Y += projectile.Speed
		projectile.Hitbox.Y += projectile.Speed
	}

	activeProjectiles := g.Projectiles[:0]
	for _, projectile := range g.Projectiles {
		if projectile.Y < ScreenHeight {
			activeProjectiles = append(activeProjectiles, projectile)
		}
	}
	g.Projectiles = activeProjectiles
}
