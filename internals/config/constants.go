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

	// Enemy movement constants
	EnemyMaxSpeed   = 2.0 // Maximum enemy speed in any direction
	EnemySpeedRange = 4.0 // Total speed range (from -2 to +2)
	EnemyBaseSpeedX = 1.0 // Base enemy speed X
	EnemyBaseSpeedY = 1.0 // Base enemy speed Y

	// Weapon system constants
	MaxNormalProjectiles   = 5 // Weapon type 1: disappearing projectiles
	MaxPiercingProjectiles = 2 // Weapon type 2: piercing projectiles

	// Spatial grid optimization
	SpatialGridCellSize = 64 // Cell size for collision detection optimization

	// Sprite frame constants
	SpriteFrameSize   = 32 // Standard sprite frame size (32x32)
	SpriteSheetWidth  = 64 // Standard sprite sheet width (2 frames side by side)
	SpriteSheetHeight = 32 // Standard sprite sheet height for single row

	// Sprite frame coordinates
	SpriteIdleLeft   = 0  // Left frame (idle) starts at x=0
	SpriteIdleTop    = 0  // Top row for new sprite format
	SpriteAttackLeft = 32 // Right frame (attack) starts at x=32
	SpriteAttackTop  = 0  // Top row for new sprite format
)
