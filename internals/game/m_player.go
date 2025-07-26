package game

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
