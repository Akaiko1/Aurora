package game

import (
	"fmt"
	"image"
	"image/color"
	"scroller_game/internals/config"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

var (
	mplusNormalFace *text.GoTextFace
	frames          *ebiten.Image
)

func (g *Game) DrawGameplay(screen *ebiten.Image) {
	//Draw Phase and Scenario
	level_info_op := &text.DrawOptions{}
	text.Draw(screen, g.Scenario.Name, mplusNormalFace, level_info_op)
	level_info_op.GeoM.Translate(0, 20)
	text.Draw(screen, g.Phase.Name, mplusNormalFace, level_info_op)

	// Draw game info
	text_op := &text.DrawOptions{}
	text_op.GeoM.Translate(config.ScreenWidth-150, 20)
	text.Draw(screen, fmt.Sprintf("You were hit: %d", g.Player.Hits), mplusNormalFace, text_op)
	text_op.GeoM.Translate(65, 30)
	text.Draw(screen, fmt.Sprintf("Score: %d", g.Player.Score), mplusNormalFace, text_op)
	if g.Player.Grazing != nil {
		text_op.GeoM.Translate(0, 390)
		text.Draw(screen, "Graze!", mplusNormalFace, text_op)
	}

	options := &ebiten.DrawImageOptions{}
	options.GeoM.Translate(float64(g.Player.X), float64(g.Player.Y))
	if g.Player.IsAttacking {
		screen.DrawImage(frames.SubImage(image.Rect(32, 0, 64, 32)).(*ebiten.Image), options)
	} else {
		screen.DrawImage(frames.SubImage(image.Rect(0, 0, 32, 32)).(*ebiten.Image), options)
	}

	for _, projectile := range g.Player.Projectiles {
		vector.DrawFilledRect(screen, projectile.X, projectile.Y, 5, 10, color.RGBA{0, 255, 255, 255}, true)
	}

	for _, enemy := range g.SpawnedEnemies {
		options := &ebiten.DrawImageOptions{}
		options.GeoM.Translate(float64(enemy.X), float64(enemy.Y))
		if !enemy.IsAttacking {
			screen.DrawImage(frames.SubImage(image.Rect(0, 32, 32, 64)).(*ebiten.Image), options)
		} else {
			screen.DrawImage(frames.SubImage(image.Rect(32, 32, 64, 64)).(*ebiten.Image), options)
		}
	}

	for _, projectile := range g.Projectiles {
		vector.DrawFilledRect(screen, projectile.X, projectile.Y, 5, 10, color.RGBA{255, 255, 0, 255}, true)
	}

	if g.FlagHitboxes {
		g.DrawHitboxes(screen)
	}
}

func (g *Game) DrawHitboxes(screen *ebiten.Image) {
	vector.StrokeRect(screen, g.Player.Hitbox.X, g.Player.Hitbox.Y,
		g.Player.Hitbox.Width, g.Player.Hitbox.Height, 2, color.RGBA{255, 255, 255, 255}, true)
	vector.StrokeRect(screen, g.Player.Grazebox.X, g.Player.Grazebox.Y,
		g.Player.Grazebox.Width, g.Player.Grazebox.Height, 2, color.RGBA{255, 125, 255, 255}, true)

	for _, projectile := range g.Player.Projectiles {
		vector.StrokeRect(screen, projectile.Hitbox.X, projectile.Hitbox.Y,
			projectile.Hitbox.Width, projectile.Hitbox.Height, 2, color.RGBA{255, 255, 255, 255}, true)
	}
	for _, enemy := range g.Enemies {
		vector.StrokeRect(screen, enemy.Hitbox.X, enemy.Hitbox.Y,
			enemy.Hitbox.Width, enemy.Hitbox.Height, 2, color.RGBA{255, 255, 255, 255}, true)
	}
	for _, projectile := range g.Projectiles {
		vector.StrokeRect(screen, projectile.Hitbox.X, projectile.Hitbox.Y,
			projectile.Hitbox.Width, projectile.Hitbox.Height, 2, color.RGBA{255, 255, 255, 255}, true)
	}
}
