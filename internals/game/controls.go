package game

import (
	"scroller_game/internals/config"
	"scroller_game/internals/events"

	"github.com/hajimehoshi/ebiten/v2"
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
	if ebiten.IsKeyPressed(ebiten.KeyB) {
		if g.FrameCount%6 == 0 {
			g.FlagHitboxes = !g.FlagHitboxes
		}
	}

	// Weapon switching
	if ebiten.IsKeyPressed(ebiten.Key1) {
		g.Player.SwitchWeapon(0) // Normal
	}
	if ebiten.IsKeyPressed(ebiten.Key2) {
		g.Player.SwitchWeapon(1) // Piercing
	}
	if ebiten.IsKeyPressed(ebiten.Key3) {
		g.Player.SwitchWeapon(2) // Rapid Fire
	}
	if ebiten.IsKeyPressed(ebiten.Key4) {
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

	if g.Player.IsAttacking && (g.FrameCount-g.Player.AttackStartFrame+120)%120 >= config.AttackCooldownFrames {
		g.Player.IsAttacking = false
	}

	// Playes projectiles interaction
	for idx, projectile := range g.Projectiles {
		if g.Player.Hitbox.Intersects(&projectile.Hitbox) {
			g.handlePlayerHit()
			g.Projectiles = events.DeleteProjectile(g.Projectiles, idx)
			break
		}

		if g.Player.Grazebox.Intersects(&projectile.Hitbox) && g.Player.Grazing != projectile {
			g.Player.Grazing = projectile
			break
		}

		if g.Player.Grazing != nil && !g.Player.Grazebox.Intersects(&g.Player.Grazing.Hitbox) {
			g.Player.Grazing = nil
			g.Player.Score++
		}

	}
}
