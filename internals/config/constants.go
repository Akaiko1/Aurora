package config

const (
	ScreenWidth        = 640
	ScreenHeight       = 480
	PlayerSpeed        = 4
	EnemyShootInterval = 50 // Frames between each shot
	ProjectileSpeed    = 5

	// Entity dimensions
	EntitySize         = 32 // Standard size for player/enemy sprites
	ProjectileWidth    = 5
	ProjectileHeight   = 10
	PlayerHitboxSize   = 10 // Player hitbox is smaller than sprite
	PlayerGrazeboxSize = 50 // Larger area for grazing mechanics

	// Game timing constants
	EnemyDirectionChangeInterval = 90 // Frames between enemy direction changes
	AttackAnimationDuration      = 12 // Frames for attack animation
	AttackCooldownFrames         = 18 // Frames after attack before resetting
	MaxPlayerProjectiles         = 3  // Maximum simultaneous player projectiles (legacy)

	// Weapon system constants
	MaxNormalProjectiles   = 5 // Weapon type 1: disappearing projectiles
	MaxPiercingProjectiles = 2 // Weapon type 2: piercing projectiles

	// Spatial grid optimization
	SpatialGridCellSize = 64 // Cell size for collision detection optimization
)
