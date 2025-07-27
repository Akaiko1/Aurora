package game

import (
	"os"
	"scroller_game/internals/config"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

func (g *Game) playerEvents() {
	if ebiten.IsKeyPressed(ebiten.KeyArrowLeft) && g.Player.X-g.Player.Speed > 0 {
		g.Player.Move(-g.Player.Speed, 0)
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowRight) && g.Player.X+g.Player.Width+g.Player.Speed < config.ScreenWidth {
		g.Player.Move(g.Player.Speed, 0)
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowUp) && g.Player.Y-g.Player.Speed > 0 {
		g.Player.Move(0, -g.Player.Speed)
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowDown) && g.Player.Y+g.Player.Height+g.Player.Speed < config.ScreenHeight {
		g.Player.Move(0, g.Player.Speed)
	}

	// Toggle hitboxes with the B key
	if inpututil.IsKeyJustPressed(ebiten.KeyB) {
		g.FlagHitboxes = !g.FlagHitboxes
	}

	// Weapon switching
	if inpututil.IsKeyJustPressed(ebiten.Key1) {
		g.Player.SwitchWeapon(0) // Normal
	}
	if inpututil.IsKeyJustPressed(ebiten.Key2) {
		g.Player.SwitchWeapon(1) // Piercing
	}
	if inpututil.IsKeyJustPressed(ebiten.Key3) {
		g.Player.SwitchWeapon(2) // Rapid Fire
	}
	if inpututil.IsKeyJustPressed(ebiten.Key4) {
		g.Player.SwitchWeapon(3) // Heavy Cannon
	}

	// Player shooting
	if ebiten.IsKeyPressed(ebiten.KeySpace) {
		if g.Player.CanFire(g.FrameCount) {
			g.playerShoot()
			g.Player.IsAttacking = true
			g.Player.AttackStartFrame = g.FrameCount
			g.Player.CurrentWeapon.Fire(g.FrameCount) // Mark weapon as fired
		}
	}

	// Reset attack state after cooldown period
	// Formula: (CurrentFrame - StartFrame + 120) % 120 >= CooldownFrames
	// The +120 handles frame counter wrapping (0-119), ensuring proper time calculation
	// Examples:
	//   - Attack at frame 110, current 5: (5-110+120)%120 = 15 frames elapsed
	//   - Attack at frame 50, current 60: (60-50+120)%120 = 10 frames elapsed
	if g.Player.IsAttacking && (g.FrameCount-g.Player.AttackStartFrame+120)%120 >= config.AttackCooldownFrames {
		g.Player.IsAttacking = false
	}

	// Player projectile interaction
	activeProjectiles := g.Projectiles[:0]
	for _, projectile := range g.Projectiles {
		hit := false
		
		if g.Player.Hitbox.Intersects(&projectile.Hitbox) {
			g.handlePlayerHit()
			hit = true
		} else if g.Player.Grazebox.Intersects(&projectile.Hitbox) && g.Player.Grazing != projectile {
			g.Player.Grazing = projectile
		}

		// Keep projectile if not hit
		if !hit {
			activeProjectiles = append(activeProjectiles, projectile)
		}
	}
	g.Projectiles = activeProjectiles

	// Handle graze ending
	if g.Player.Grazing != nil && !g.Player.Grazebox.Intersects(&g.Player.Grazing.Hitbox) {
		g.Player.Grazing = nil
		g.Player.Score++
	}
}

// gameOverMenuEvents handles input for the game over menu navigation.
func (g *Game) gameOverMenuEvents() {
	// Menu navigation using UP/DOWN arrows
	if inpututil.IsKeyJustPressed(ebiten.KeyArrowUp) {
		g.GameOverSelection = 0 // Select "Play Again"
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyArrowDown) {
		g.GameOverSelection = 1 // Select "Exit"
	}

	// Confirm selection with ENTER key
	if inpututil.IsKeyJustPressed(ebiten.KeyEnter) {
		switch g.GameOverSelection {
		case 0: // Play Again
			g.restartGame()
		case 1: // Exit
			os.Exit(0)
		}
	}
}
