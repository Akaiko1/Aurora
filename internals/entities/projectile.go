package entities

import "scroller_game/internals/physics"

type Projectile struct {
	X, Y, Width, Height float32
	Speed               float32
	Hitbox              physics.Hitbox
}
