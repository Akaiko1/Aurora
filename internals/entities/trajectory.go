package entities

import (
	"math"
)

// TrajectoryHandler calculates movement for different trajectory types
type TrajectoryHandler struct {
	SpawnedEnemies []*Enemy // Reference to enemies for tracking
	Player         *Player  // Reference to player for enemy projectiles to track (optional)
}

// CalculateMovement returns dx, dy for a projectile based on its trajectory type
func (th *TrajectoryHandler) CalculateMovement(projectile *Projectile, currentFrame int) (float32, float32) {
	switch projectile.TrajectoryType {
	case TrajectoryLinear:
		return th.calculateLinear(projectile)
	case TrajectoryTracking:
		return th.calculateTracking(projectile)
	case TrajectorySine:
		return th.calculateSine(projectile, currentFrame)
	case TrajectorySpiral:
		return th.calculateSpiral(projectile, currentFrame)
	case TrajectoryArc:
		return th.calculateArc(projectile, currentFrame)
	default:
		return th.calculateLinear(projectile) // Fallback to linear
	}
}

// calculateLinear - straight line movement
func (th *TrajectoryHandler) calculateLinear(projectile *Projectile) (float32, float32) {
	return 0, projectile.Speed * projectile.Direction
}

// calculateTracking - tracks closest enemy
func (th *TrajectoryHandler) calculateTracking(projectile *Projectile) (float32, float32) {
	// Find closest enemy
	var closestEnemy *Enemy
	var closestDistance float32 = math.MaxFloat32

	projectileCenterX, projectileCenterY := projectile.GetCenter()

	for _, enemy := range th.SpawnedEnemies {
		enemyCenterX := enemy.X + enemy.Width/2
		enemyCenterY := enemy.Y + enemy.Height/2

		// Calculate distance
		dx := enemyCenterX - projectileCenterX
		dy := enemyCenterY - projectileCenterY
		distance := float32(math.Sqrt(float64(dx*dx + dy*dy)))

		if distance < closestDistance {
			closestDistance = distance
			closestEnemy = enemy
		}
	}

	if closestEnemy != nil {
		// Calculate direction to enemy
		enemyCenterX := closestEnemy.X + closestEnemy.Width/2
		enemyCenterY := closestEnemy.Y + closestEnemy.Height/2

		dx := enemyCenterX - projectileCenterX
		dy := enemyCenterY - projectileCenterY

		// Normalize direction and apply speed
		distance := float32(math.Sqrt(float64(dx*dx + dy*dy)))
		if distance > 0 {
			return (dx / distance) * projectile.Speed, (dy / distance) * projectile.Speed
		}
	}

	// No enemies, move in original direction
	return 0, projectile.Speed * projectile.Direction
}

// calculateSine - sine wave pattern
func (th *TrajectoryHandler) calculateSine(projectile *Projectile, currentFrame int) (float32, float32) {
	framesSinceSpawn := currentFrame - projectile.SpawnFrame

	// Handle frame wrapping (frames wrap every 120)
	if framesSinceSpawn < 0 {
		framesSinceSpawn += 120
	}

	// Sine wave amplitude and frequency
	amplitude := float32(20.0) // How wide the wave is
	frequency := float32(0.15) // How fast it oscillates

	// Calculate sine wave offset
	currentSine := amplitude * float32(math.Sin(float64(framesSinceSpawn)*float64(frequency)))
	lastSine := amplitude * float32(math.Sin(float64(framesSinceSpawn-1)*float64(frequency)))

	// Move forward in original direction, but add sine wave to X
	dx := currentSine - lastSine
	dy := projectile.Speed * projectile.Direction

	return dx, dy
}

// calculateSpiral - spiral movement pattern
func (th *TrajectoryHandler) calculateSpiral(projectile *Projectile, currentFrame int) (float32, float32) {
	framesSinceSpawn := currentFrame - projectile.SpawnFrame

	// Handle frame wrapping (frames wrap every 120)
	if framesSinceSpawn < 0 {
		framesSinceSpawn += 120
	}

	// Spiral parameters - simpler approach
	spiralSpeed := float32(0.3) // Rotation speed
	amplitude := float32(25.0)  // Fixed spiral radius

	angle := float64(framesSinceSpawn) * float64(spiralSpeed)
	prevAngle := float64(framesSinceSpawn-1) * float64(spiralSpeed)

	// Calculate spiral offset (always same radius, just rotating)
	currentOffsetX := amplitude * float32(math.Cos(angle))
	prevOffsetX := amplitude * float32(math.Cos(prevAngle))

	// Movement is just the difference in spiral position + forward movement
	dx := currentOffsetX - prevOffsetX
	dy := projectile.Speed * projectile.Direction

	return dx, dy
}

// calculateArc - arcing movement (like a thrown projectile)
func (th *TrajectoryHandler) calculateArc(projectile *Projectile, currentFrame int) (float32, float32) {
	framesSinceSpawn := currentFrame - projectile.SpawnFrame

	// Arc parameters
	gravity := float32(0.2)
	initialVelocityX := projectile.Speed * 0.5                  // Horizontal component
	initialVelocityY := projectile.Speed * projectile.Direction // Vertical component

	// Calculate arc movement
	dx := initialVelocityX
	dy := initialVelocityY + gravity*float32(framesSinceSpawn)

	return dx, dy
}
