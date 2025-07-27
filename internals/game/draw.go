package game

import (
	"fmt"
	"image"
	"image/color"
	"scroller_game/internals/config"
	"scroller_game/internals/entities"
	"scroller_game/internals/inputs"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

// getSpriteFrame extracts a sprite frame with DRY principle
func getSpriteFrame(sprite *ebiten.Image, isAttacking bool) *ebiten.Image {
	var left, top int

	if isAttacking {
		left, top = config.SpriteAttackLeft, config.SpriteAttackTop
	} else {
		left, top = config.SpriteIdleLeft, config.SpriteIdleTop
	}

	return sprite.SubImage(image.Rect(
		left, top,
		left+config.SpriteFrameSize, top+config.SpriteFrameSize)).(*ebiten.Image)
}

// getPlayerSpriteFrame extracts a player sprite frame
func getPlayerSpriteFrame(sprite *ebiten.Image, isAttacking bool) *ebiten.Image {
	var left, top int

	if isAttacking {
		left, top = config.SpriteAttackLeft, config.SpriteAttackTop
	} else {
		left, top = config.SpriteIdleLeft, config.SpriteIdleTop
	}

	return sprite.SubImage(image.Rect(
		left, top,
		left+config.SpriteFrameSize, top+config.SpriteFrameSize)).(*ebiten.Image)
}

// createDrawOptions creates standard draw options with position translation
func createDrawOptions(x, y float32) *ebiten.DrawImageOptions {
	options := &ebiten.DrawImageOptions{}
	options.GeoM.Translate(float64(x), float64(y))
	return options
}

// uint32ToRGBA converts a uint32 color to color.RGBA
func uint32ToRGBA(c uint32) color.RGBA {
	return color.RGBA{
		R: uint8((c >> 24) & 0xFF),
		G: uint8((c >> 16) & 0xFF),
		B: uint8((c >> 8) & 0xFF),
		A: uint8(c & 0xFF),
	}
}

var (
	mplusNormalFace *text.GoTextFace
)

func (g *Game) DrawGameplay(screen *ebiten.Image) {
	// Draw Phase and Scenario
	g.TextOptions.GeoM.Reset()
	scenarioName := "Loading..."
	if g.Scenario != nil {
		scenarioName = g.Scenario.Name
	}
	text.Draw(screen, scenarioName, mplusNormalFace, g.TextOptions)
	
	g.TextOptions.GeoM.Translate(0, 20)
	phaseName := "Loading..."
	if g.Phase != nil {
		phaseName = g.Phase.Name
	}
	text.Draw(screen, phaseName, mplusNormalFace, g.TextOptions)

	// Draw game info
	g.TextOptions.GeoM.Reset()
	g.TextOptions.GeoM.Translate(config.ScreenWidth-150, 20)
	text.Draw(screen, fmt.Sprintf("You were hit: %d", g.Player.Hits), mplusNormalFace, g.TextOptions)
	g.TextOptions.GeoM.Translate(65, 30)
	text.Draw(screen, fmt.Sprintf("Score: %d", g.Player.Score), mplusNormalFace, g.TextOptions)

	// Draw weapon info
	g.TextOptions.GeoM.Translate(-125, 25)
	weaponName := "Unknown"
	if g.Player.CurrentWeapon != nil {
		weaponName = g.Player.CurrentWeapon.Definition.Name
	}
	text.Draw(screen, fmt.Sprintf("Weapon: %s (%d/%d)", weaponName, len(g.Player.Projectiles), g.Player.GetMaxProjectiles()), mplusNormalFace, g.TextOptions)

	if g.Player.Grazing != nil {
		g.TextOptions.GeoM.Translate(0, 350)
		text.Draw(screen, "Graze!", mplusNormalFace, g.TextOptions)
	}

	// Draw player using new sprite system with helper function
	options := createDrawOptions(g.Player.X, g.Player.Y)

	// Draw player using helper function
	playerFrame := getPlayerSpriteFrame(inputs.PlayerSprite, g.Player.IsAttacking)
	screen.DrawImage(playerFrame, options)

	for _, projectile := range g.Player.Projectiles {
		projectileColor := uint32ToRGBA(projectile.Color)
		vector.DrawFilledRect(screen, projectile.X, projectile.Y, projectile.Width, projectile.Height, projectileColor, true)
	}

	// Draw enemies using new sprite system
	for _, enemy := range g.SpawnedEnemies {
		options := createDrawOptions(enemy.X, enemy.Y)

		// Get sprite based on enemy type using the new hashmap system
		enemyTypeName := entities.GetEnemyTypeName(enemy.Type)
		enemySprite := inputs.GetEnemySprite(enemyTypeName)

		// Draw enemy frame using helper function
		enemyFrame := getSpriteFrame(enemySprite, enemy.IsAttacking)
		screen.DrawImage(enemyFrame, options)
	}

	for _, projectile := range g.Projectiles {
		projectileColor := uint32ToRGBA(projectile.Color)
		vector.DrawFilledRect(screen, projectile.X, projectile.Y, projectile.Width, projectile.Height, projectileColor, true)
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
