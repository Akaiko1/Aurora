package entities

import (
	"scroller_game/internals/physics"
)

type Enemy struct {
	X, Y, Width, Height float32
	SpeedX, SpeedY      float32
	Hitbox              physics.Hitbox
	IsAttacking         bool
	AttackStartFrame    int
	Weapon              *Weapon // Enemy weapon
}

// UpdatePosition updates the enemy's position and automatically syncs hitbox
func (e *Enemy) UpdatePosition(x, y float32) {
	e.X = x
	e.Y = y
	e.Hitbox.X = x
	e.Hitbox.Y = y
}

// Move updates position by delta values
func (e *Enemy) Move(dx, dy float32) {
	e.UpdatePosition(e.X+dx, e.Y+dy)
}

// CanFire checks if enemy can fire based on weapon fire rate
func (e *Enemy) CanFire(currentFrame int) bool {
	if e.Weapon == nil {
		return false
	}
	return e.Weapon.CanFire(currentFrame)
}

// CreateProjectile creates a projectile using enemy's weapon
func (e *Enemy) CreateProjectile(currentFrame int) *Projectile {
	if e.Weapon == nil {
		return nil
	}
	// Enemy shoots downward (+1 direction)
	centerX := e.X + e.Width/2 - e.Weapon.Definition.ProjectileWidth/2
	shootY := e.Y + e.Height
	return e.Weapon.CreateProjectile(centerX, shootY, 1, currentFrame)
}

// InitializeWeapon sets up the enemy's weapon
func (e *Enemy) InitializeWeapon(weaponID WeaponID) {
	e.Weapon = GetWeapon(weaponID)
}

// EnemyShoot creates and adds projectile to the provided slice (legacy compatibility)
func (enemy *Enemy) EnemyShoot(projectiles *[]*Projectile, currentFrame int) {
	if enemy.Weapon == nil {
		enemy.InitializeWeapon(WeaponIDEnemy) // Default weapon
	}

	projectile := enemy.CreateProjectile(currentFrame)
	if projectile != nil {
		*projectiles = append(*projectiles, projectile)
	}
}
