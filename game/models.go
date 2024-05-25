package game

import "github.com/hajimehoshi/ebiten/v2"

type Hitbox struct {
	X, Y, Width, Height float32
}

type Player struct {
	X, Y, Width, Height float32
	Speed               float32
	Projectiles         []*Projectile
	Hitbox              Hitbox
	Grazebox            Hitbox
	Hits                int
	Grazing             *Projectile
	Score               int
	Sprite              *ebiten.Image
}

type Enemy struct {
	X, Y, Width, Height float32
	SpeedX, SpeedY      float32
	Hitbox              Hitbox
}

type Projectile struct {
	X, Y, Width, Height float32
	Speed               float32
	Hitbox              Hitbox
}
