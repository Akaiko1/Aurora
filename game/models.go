package game

import "math/rand"

type GameState int

const (
	StartMenu GameState = iota
	Playing
	SwitchLevel
	SwitchPhase
	Paused
	GameOver
)

type Game struct {
	Player       *Player
	Projectiles  []*Projectile
	FrameCount   int
	Enemies      []*Enemy
	Spawned      []*Enemy
	State        GameState
	Scenario     *Scenario
	Phase        *Phase
	RandomSource *rand.Rand
	Scenarios    []*Scenario
	FlagHitboxes bool
}

type Scenario struct {
	Name   string
	Phases []*Phase
}

type Phase struct {
	Name    string
	Enemies []*Enemy
}

type Hitbox struct {
	X, Y, Width, Height float32
}

type Player struct {
	X, Y, Width, Height, Speed float32
	Projectiles                []*Projectile
	Hitbox                     Hitbox
	Grazebox                   Hitbox
	Grazing                    *Projectile
	Hits                       int
	Score                      int
	IsAttacking                bool
	AttackStartFrame           int
}

type Enemy struct {
	X, Y, Width, Height float32
	SpeedX, SpeedY      float32
	Hitbox              Hitbox
	IsAttacking         bool
	AttackStartFrame    int
}

type Projectile struct {
	X, Y, Width, Height float32
	Speed               float32
	Hitbox              Hitbox
}
