package game

func (g *Game) playerProjectilesMovements() {
	// Update player projectiles
	for _, projectile := range g.player.Projectiles {
		projectile.Y -= projectile.Speed
		projectile.Hitbox.Y -= projectile.Speed
	}

	// Remove off-screen player projectiles
	activeProjectiles := g.player.Projectiles[:0]
	for _, projectile := range g.player.Projectiles {
		if projectile.Y > 0 {
			activeProjectiles = append(activeProjectiles, projectile)
		}
	}
	g.player.Projectiles = activeProjectiles
}

func (g *Game) playerShoot() {
	if len(g.player.Projectiles) >= 3 {
		return // Limit to 3 projectiles
	}
	projectile := &Projectile{
		X:     g.player.X + g.player.Width/2, // Center of the player
		Y:     g.player.Y,
		Width: 5, Height: 10,
		Speed: projectileSpeed,
	}
	projectile.Hitbox = Hitbox{X: projectile.X, Y: projectile.Y, Width: projectile.Width, Height: projectile.Height}
	g.player.Projectiles = append(g.player.Projectiles, projectile)
}
