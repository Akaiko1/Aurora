package entities

import (
	"scroller_game/internals/physics"
)

type Player struct {
	X, Y, Width, Height, Speed float32
	Projectiles                []*Projectile
	Hitbox                     physics.Hitbox
	Grazebox                   physics.Hitbox
	Grazing                    *Projectile
	Hits                       int
	Score                      int
	IsAttacking                bool
	AttackStartFrame           int
}
