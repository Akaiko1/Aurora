package entities

import (
	"scroller_game/internals/config"
)

// WeaponID represents different weapon types
type WeaponID string

const (
	WeaponIDNormal   WeaponID = "normal"
	WeaponIDPiercing WeaponID = "piercing"
	WeaponIDEnemy    WeaponID = "enemy_basic"
	WeaponIDRapid    WeaponID = "rapid_fire"
	WeaponIDHeavy    WeaponID = "heavy_cannon"

	// Enemy-specific weapons
	WeaponIDDragonSpiral    WeaponID = "dragon_spiral"
	WeaponIDSkeletonPiercer WeaponID = "skeleton_piercer"
	// Easy to add more: WeaponIDSpread, WeaponIDLaser, etc.
)

// WeaponDefinition defines all properties of a weapon type
type WeaponDefinition struct {
	ID               WeaponID
	Name             string
	MaxProjectiles   int
	ProjectileType   ProjectileType
	ProjectileSpeed  float32
	ProjectileWidth  float32
	ProjectileHeight float32
	FireRate         int            // Frames between shots (lower = faster)
	Damage           int            // Future: damage system
	Color            uint32         // Future: different colored projectiles
	TrajectoryType   TrajectoryType // Movement pattern
}

// Weapon represents an instance of a weapon
type Weapon struct {
	Definition   *WeaponDefinition
	LastFireTime int // Track when this weapon last fired
}

// CanFire checks if weapon can fire based on fire rate
func (w *Weapon) CanFire(currentFrame int) bool {
	// Handle frame wrapping (frames wrap every 120)
	framesSinceLastFire := currentFrame - w.LastFireTime
	if framesSinceLastFire < 0 {
		framesSinceLastFire += 120 // Account for frame wrapping
	}
	return framesSinceLastFire >= int(w.Definition.FireRate)
}

// CreateProjectile creates a projectile based on weapon definition
func (w *Weapon) CreateProjectile(x, y float32, direction float32, currentFrame int) *Projectile {
	projectile := &Projectile{
		X:              x,
		Y:              y,
		Width:          w.Definition.ProjectileWidth,
		Height:         w.Definition.ProjectileHeight,
		Speed:          w.Definition.ProjectileSpeed, // Use base speed, direction handled in movement
		Type:           w.Definition.ProjectileType,
		Direction:      direction, // Store direction separately
		TrajectoryType: w.Definition.TrajectoryType,
		SpawnFrame:     currentFrame,
		InitialX:       x,
		InitialY:       y,
		Color:          w.Definition.Color,
	}
	projectile.UpdatePosition(x, y) // Initialize hitbox
	return projectile
}

// Fire marks the weapon as having fired this frame
func (w *Weapon) Fire(currentFrame int) {
	w.LastFireTime = currentFrame
}

// WeaponRegistry holds all weapon definitions
var WeaponRegistry = map[WeaponID]*WeaponDefinition{
	WeaponIDNormal: {
		ID:               WeaponIDNormal,
		Name:             "Basic",
		MaxProjectiles:   5,
		ProjectileType:   ProjectileNormal,
		ProjectileSpeed:  config.ProjectileSpeed,
		ProjectileWidth:  config.ProjectileWidth,
		ProjectileHeight: config.ProjectileHeight,
		FireRate:         3, // Can fire every 3 frames (faster)
		Damage:           1,
		Color:            0xFFFFFFFF, // White
		TrajectoryType:   TrajectoryLinear,
	},
	WeaponIDPiercing: {
		ID:               WeaponIDPiercing,
		Name:             "Piercing",
		MaxProjectiles:   2,
		ProjectileType:   ProjectilePiercing,
		ProjectileSpeed:  config.ProjectileSpeed * 1.2, // Slightly faster
		ProjectileWidth:  config.ProjectileWidth,
		ProjectileHeight: config.ProjectileHeight * 1.5, // Taller projectiles
		FireRate:         5,                             // Slower fire rate (was 8)
		Damage:           2,
		Color:            0xFF00FFFF,     // Cyan
		TrajectoryType:   TrajectorySine, // Sine wave pattern
	},
	WeaponIDEnemy: {
		ID:               WeaponIDEnemy,
		Name:             "Enemy Basic",
		MaxProjectiles:   10, // Enemies can have more projectiles
		ProjectileType:   ProjectileNormal,
		ProjectileSpeed:  config.ProjectileSpeed * 0.8, // Slower than player
		ProjectileWidth:  config.ProjectileWidth,
		ProjectileHeight: config.ProjectileHeight,
		FireRate:         config.EnemyShootInterval,
		Damage:           1,
		Color:            0xFF0000FF, // Red
		TrajectoryType:   TrajectoryLinear,
	},
	WeaponIDRapid: {
		ID:               WeaponIDRapid,
		Name:             "Rapid",
		MaxProjectiles:   8,
		ProjectileType:   ProjectileNormal,
		ProjectileSpeed:  config.ProjectileSpeed * 0.9, // Slightly slower
		ProjectileWidth:  config.ProjectileWidth * 0.8, // Thinner projectiles
		ProjectileHeight: config.ProjectileHeight,
		FireRate:         2, // Very fast fire rate!
		Damage:           1,
		Color:            0xFFFF00FF, // Yellow
		TrajectoryType:   TrajectorySpiral,
	},
	WeaponIDHeavy: {
		ID:               WeaponIDHeavy,
		Name:             "Heavy!",
		MaxProjectiles:   1,                            // Only one shot at a time
		ProjectileType:   ProjectileNormal,             // Back to normal, tracking handled by trajectory
		ProjectileSpeed:  config.ProjectileSpeed * 0.6, // Slower projectiles
		ProjectileWidth:  config.ProjectileWidth * 2,   // Much wider
		ProjectileHeight: config.ProjectileHeight * 2,  // Much taller
		FireRate:         20,                           // Very slow fire rate
		Damage:           5,                            // High damage
		Color:            0xFF8000FF,                   // Orange
		TrajectoryType:   TrajectoryTracking,           // Tracks enemies
	},
	WeaponIDDragonSpiral: {
		ID:               WeaponIDDragonSpiral,
		Name:             "Dragon Spiral",
		MaxProjectiles:   8,
		ProjectileType:   ProjectileNormal,
		ProjectileSpeed:  config.ProjectileSpeed * 0.7,  // Slower than basic enemy
		ProjectileWidth:  config.ProjectileWidth * 1.2,  // Slightly wider
		ProjectileHeight: config.ProjectileHeight * 1.2, // Slightly taller
		FireRate:         10,
		Damage:           2,          // More damage than basic
		Color:            0xFF4000FF, // Dark red/orange
		TrajectoryType:   TrajectorySpiral,
	},
	WeaponIDSkeletonPiercer: {
		ID:               WeaponIDSkeletonPiercer,
		Name:             "Skeleton Piercer",
		MaxProjectiles:   3,
		ProjectileType:   ProjectilePiercing,            // Pierces through player!
		ProjectileSpeed:  config.ProjectileSpeed * 0.9,  // Slightly faster than basic
		ProjectileWidth:  config.ProjectileWidth * 0.8,  // Thinner for piercing
		ProjectileHeight: config.ProjectileHeight * 1.8, // Much taller
		FireRate:         35,                            // Slower fire rate (powerful shots)
		Damage:           3,                             // High damage
		Color:            0xFFFFFFFF,                    // White/bone color
		TrajectoryType:   TrajectorySine,                // Sine wave makes it harder to dodge
	},
}

// GetWeapon creates a new weapon instance from registry
func GetWeapon(id WeaponID) *Weapon {
	def, exists := WeaponRegistry[id]
	if !exists {
		def = WeaponRegistry[WeaponIDNormal] // Fallback to normal
	}
	return &Weapon{
		Definition:   def,
		LastFireTime: 0,
	}
}

// NewWeapon creates a custom weapon (for future modding/scripting)
func NewWeapon(def *WeaponDefinition) *Weapon {
	return &Weapon{
		Definition:   def,
		LastFireTime: 0,
	}
}
