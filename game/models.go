package game

type Hitbox struct {
	X, Y, Width, Height float32
}

type Player struct {
	X, Y, Width, Height, Speed float32
	Projectiles                []*Projectile
	Hitbox                     Hitbox
	Grazebox                   Hitbox
	Hits                       int
	Grazing                    *Projectile
	Score                      int
	Attacking                  bool
}

type Enemy struct {
	X, Y, Width, Height float32
	SpeedX, SpeedY      float32
	Hitbox              Hitbox
	Attacking           bool
}

type Projectile struct {
	X, Y, Width, Height float32
	Speed               float32
	Hitbox              Hitbox
}
