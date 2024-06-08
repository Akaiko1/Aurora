package game

func (g *Game) handleEnemyHit(enemy *Enemy) {
	// Remove the enemy from the list
	g.Spawned, _ = deleteEnemy(g.Spawned, enemy)
	g.Player.Score++
}

func (g *Game) handlePlayerHit() {
	g.Player.Score--
	g.Player.Hits++
}

func (g *Game) setRandomDirection(enemy *Enemy) {
	x_r := g.RandomSource.Float32()
	y_r := g.RandomSource.Float32()
	enemy.SpeedX = (x_r * 4) - 2 // Speed between -2 and 2
	enemy.SpeedY = (y_r * 4) - 2 // Speed between -2 and 2
}
