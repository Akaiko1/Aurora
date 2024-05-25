package game

import (
	"image/color"
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type Game struct {
	player       *Player
	projectiles  []*Projectile
	enemies      []*Enemy
	frameCount   int
	rng          *rand.Rand
	flagHitboxes bool
}

type Hitbox struct {
	X, Y, Width, Height float32
}

// Check if two hitboxes intersect
func (hb *Hitbox) Intersects(other *Hitbox) bool {
	return hb.X < other.X+other.Width &&
		hb.X+hb.Width > other.X &&
		hb.Y < other.Y+other.Height &&
		hb.Y+hb.Height > other.Y
}

type Player struct {
	X, Y, Width, Height float32
	Speed               float32
	Projectiles         []*Projectile
	Hitbox              Hitbox
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

// Update updates the game state.
// It handles player controls, updates projectiles, changes enemy direction periodically,
// moves enemies, and bounces enemies off the screen edges.
func (g *Game) Update() error {
	// Increment the frame count
	g.frameCount++

	// Handle player controls
	g.playerControls()

	// Spawn a new enemy
	g.enemySpawn()

	// Update projectiles
	// Remove off-screen projectiles
	g.projectilesMovements()
	g.playerProjectilesMovements()

	// Change enemy direction periodically
	// Change direction every 120 frames (2 seconds)
	// Move enemy
	// Bounce enemy off the screen edges
	for _, enemy := range g.enemies {
		g.enemyActions(enemy)
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{0, 0, 0, 255})

	// Draw the player
	vector.DrawFilledRect(screen, g.player.X, g.player.Y, g.player.Width, g.player.Height,
		color.RGBA{0, 255, 0, 255}, true)
	// Draw player projectiles
	for _, projectile := range g.player.Projectiles {
		vector.DrawFilledRect(screen, projectile.X, projectile.Y, 5, 10, color.RGBA{0, 255, 255, 255}, true)
	}
	// Draw the enemies
	for _, enemy := range g.enemies {
		vector.DrawFilledRect(screen, enemy.X, enemy.Y, 32, 32, color.RGBA{255, 0, 0, 255}, true)
	}
	// Draw projectiles
	for _, projectile := range g.projectiles {
		vector.DrawFilledRect(screen, projectile.X, projectile.Y, 5, 10, color.RGBA{255, 255, 0, 255}, true)
	}

	if g.flagHitboxes {
		vector.StrokeRect(screen, g.player.Hitbox.X, g.player.Hitbox.Y,
			g.player.Hitbox.Width, g.player.Hitbox.Height, 2, color.RGBA{255, 255, 255, 255}, true)
		for _, projectile := range g.player.Projectiles {
			vector.StrokeRect(screen, projectile.Hitbox.X, projectile.Hitbox.Y,
				projectile.Hitbox.Width, projectile.Hitbox.Height, 2, color.RGBA{255, 255, 255, 255}, true)
		}
		for _, enemy := range g.enemies {
			vector.StrokeRect(screen, enemy.Hitbox.X, enemy.Hitbox.Y,
				enemy.Hitbox.Width, enemy.Hitbox.Height, 2, color.RGBA{255, 255, 255, 255}, true)
		}
		for _, projectile := range g.projectiles {
			vector.StrokeRect(screen, projectile.Hitbox.X, projectile.Hitbox.Y,
				projectile.Hitbox.Width, projectile.Hitbox.Height, 2, color.RGBA{255, 255, 255, 255}, true)
		}
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return ScreenWidth, ScreenHeight
}

func NewGame() *Game {
	player := &Player{
		X:           ScreenWidth / 2,
		Y:           ScreenHeight / 2,
		Width:       32,
		Height:      32,
		Speed:       playerSpeed,
		Projectiles: []*Projectile{},
		Hitbox:      Hitbox{Width: 20, Height: 20}}

	player.Hitbox.CenterOn(player.X+player.Width/2, player.Y+player.Height/2)

	return &Game{
		player:       player,
		projectiles:  []*Projectile{},
		enemies:      []*Enemy{},
		rng:          rand.New(rand.NewSource(time.Now().UnixNano())),
		flagHitboxes: true,
	}
}
