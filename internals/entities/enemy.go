package entities

import (
	"scroller_game/internals/config"
	"scroller_game/internals/physics"
)

type Enemy struct {
	X, Y, Width, Height float32
	SpeedX, SpeedY      float32
	Hitbox              physics.Hitbox
	IsAttacking         bool
	AttackStartFrame    int
}

// TODO: move to manager
func (enemy *Enemy) EnemyShoot(projectiles *[]*Projectile) {
	projectile := &Projectile{
		X:     enemy.X + 13, // Center of the enemy
		Y:     enemy.Y + 32,
		Width: 5, Height: 10,
		Speed: config.ProjectileSpeed,
	}
	projectile.Hitbox = physics.Hitbox{X: projectile.X, Y: projectile.Y, Width: projectile.Width, Height: projectile.Height}
	*projectiles = append(*projectiles, projectile)
}
