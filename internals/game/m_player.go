// Package game contains the core game logic and entity management for Aurora.
package game

// handlePlayerHit processes when the player entity takes damage.
// It decrements the player's score and increments their hit counter
// for tracking game statistics.
func (g *Game) handlePlayerHit() {
	g.Player.Score--
	g.Player.Hits++
}

// playerShoot handles the player's shooting mechanism.
// It creates a new projectile using the player's current weapon and frame timing,
// then adds it to the player's active projectile collection if creation was successful.
// The projectile creation may fail if the weapon is on cooldown or other constraints.
func (g *Game) playerShoot() {
	projectile := g.Player.CreateProjectile(g.FrameCount)
	if projectile != nil {
		g.Player.Projectiles = append(g.Player.Projectiles, projectile)
	}
}
