package entities

// WeaponType represents different weapon types
type WeaponType int

const (
	WeaponNormal   WeaponType = iota // Type 1: Disappears on hit, max 5 projectiles
	WeaponPiercing                   // Type 2: Pierces through enemies, max 2 projectiles
)

// ProjectileType represents different projectile behaviors
type ProjectileType int

const (
	ProjectileNormal   ProjectileType = iota // Disappears on hit
	ProjectilePiercing                       // Pierces through enemies
	ProjectileTracking                       // Tracks closest enemy
)

// TrajectoryType represents different movement patterns
type TrajectoryType int

const (
	TrajectoryLinear   TrajectoryType = iota // Straight line movement
	TrajectoryTracking                       // Tracks closest enemy
	TrajectorySine                           // Sine wave pattern
	TrajectorySpiral                         // Spiral pattern
	TrajectoryArc                            // Arcing movement
)
