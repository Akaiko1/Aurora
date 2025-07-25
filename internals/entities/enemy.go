package entities

import (
	"scroller_game/internals/config"
	"scroller_game/internals/physics"
)

// EnemyType defines different types of enemies
type EnemyType int

const (
	EnemyTypeSpider EnemyType = iota
	EnemyTypeGoblin
	EnemyTypeDragon
	EnemyTypeSkeleton
	// Add new enemy types here and they'll be automatically discovered!
	EnemyTypeCount // This should always be last - used for iteration
)

// GetAllEnemyTypes returns all enemy types for automatic discovery
func GetAllEnemyTypes() []EnemyType {
	var types []EnemyType
	for i := range EnemyTypeCount {
		types = append(types, i)
	}
	return types
}

// GetEnemyTypeName returns the string name for an enemy type
func GetEnemyTypeName(enemyType EnemyType) string {
	switch enemyType {
	case EnemyTypeSpider:
		return "spider"
	case EnemyTypeGoblin:
		return "goblin"
	case EnemyTypeDragon:
		return "dragon"
	case EnemyTypeSkeleton:
		return "skeleton"
	default:
		return "spider" // Default fallback
	}
}

type Enemy struct {
	X, Y, Width, Height float32
	SpeedX, SpeedY      float32
	Hitbox              physics.Hitbox
	IsAttacking         bool
	AttackStartFrame    int
	Weapon              *Weapon // Enemy weapon
	Type                EnemyType
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

// NewEnemy creates a new enemy of the specified type with the given weapon
func NewEnemy(x, y float32, enemyType EnemyType, weaponID WeaponID) *Enemy {
	enemy := &Enemy{
		X:      x,
		Y:      y,
		Width:  config.EntitySize, // Use constant instead of hard-coded 32
		Height: config.EntitySize, // Use constant instead of hard-coded 32
		SpeedX: config.EnemyBaseSpeedX,
		SpeedY: config.EnemyBaseSpeedY,
		Type:   enemyType,
	}

	// Initialize weapon with the provided weapon ID
	enemy.InitializeWeapon(weaponID)

	return enemy
}
