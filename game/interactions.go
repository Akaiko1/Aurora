package game

func (g *Game) handleEnemyHit(enemy *Enemy) {
	// Reset enemy position (or any other handling logic)
	enemy.X = 0
	enemy.Y = 0
	enemy.Hitbox.X = 0
	enemy.Hitbox.Y = 0
	g.setRandomDirection(enemy)
	g.player.Score++
}

func (g *Game) handlePlayerHit() {
	g.player.Score--
	g.player.Hits++
}

func (g *Game) setRandomDirection(enemy *Enemy) {
	x_r := g.rng.Float32()
	y_r := g.rng.Float32()
	enemy.SpeedX = (x_r * 4) - 2 // Speed between -2 and 2
	enemy.SpeedY = (y_r * 4) - 2 // Speed between -2 and 2
}
