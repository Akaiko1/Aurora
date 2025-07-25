package entities

import "scroller_game/internals/physics"

type Projectile struct {
	X, Y, Width, Height float32
	Speed               float32
	Direction           float32        // -1 for up (player), +1 for down (enemy)
	Hitbox              physics.Hitbox
	Type                ProjectileType // Determines behavior on collision
	TrajectoryType      TrajectoryType // Movement pattern
	SpawnFrame          int            // Frame when projectile was created (for time-based trajectories)
	InitialX            float32        // Starting X position (for pattern calculations)
	InitialY            float32        // Starting Y position (for pattern calculations)
}

// UpdatePosition updates the projectile's position and automatically syncs hitbox
func (p *Projectile) UpdatePosition(x, y float32) {
	p.X = x
	p.Y = y
	p.Hitbox.X = x
	p.Hitbox.Y = y
}

// Move updates position by delta values
func (p *Projectile) Move(dx, dy float32) {
	p.UpdatePosition(p.X+dx, p.Y+dy)
}

// GetCenter returns the center point of the projectile
func (p *Projectile) GetCenter() (float32, float32) {
	return p.X + p.Width/2, p.Y + p.Height/2
}
