package game

import (
	"bytes"
	"fmt"
	"image/color"
	"log"
	"math/rand"
	"os"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

var (
	mplusNormalFace *text.GoTextFace
	frames          *ebiten.Image
)

func init() {
	// Open the font file
	fontBytes, err := os.ReadFile("assets/Jacquard12-Regular.ttf")
	if err != nil {
		log.Fatal(err)
	}

	mplusFaceSource, err := text.NewGoTextFaceSource(bytes.NewReader(fontBytes))
	if err != nil {
		log.Fatal(err)
	}

	mplusNormalFace = &text.GoTextFace{
		Source: mplusFaceSource,
		Size:   24,
	}

	frames, _ = ReadImage("assets/sprites/animations.png")
}

// Update updates the game state.
// It handles player controls, updates projectiles, changes enemy direction periodically,
// moves enemies, and bounces enemies off the screen edges.
func (g *Game) Update() error {
	// Increment the frame count
	switch g.State {
	case Playing:
		g.FrameCount = (g.FrameCount + 1) % 120

		if len(g.Enemies) == 0 && len(g.Spawned) == 0 {
			g.State = SwitchPhase
		}

		// Handle player controls
		g.playerEvents()

		// Spawn a new enemy
		g.enemySpawn()

		// Update projectiles
		g.projectilesMovements()
		g.playerProjectilesMovements()

		// Update enemies
		for _, enemy := range g.Spawned {
			g.enemyActions(enemy)
		}
	case SwitchLevel:
		// Switch to the next level
		if len(g.Scenarios) > 0 {
			g.Scenario = g.Scenarios[0]
			g.Phase = g.Scenario.Phases[0]
			g.Scenario.Phases = g.Scenario.Phases[1:]
			g.Enemies = g.Phase.Enemies
			g.State = Playing
		} else {
			g.State = GameOver
		}

	case SwitchPhase:
		// Switch to the next phase
		if len(g.Scenario.Phases) > 0 {
			g.Phase = g.Scenario.Phases[0]
			g.Enemies = g.Phase.Enemies
			g.State = Playing

			// Remove the current phase
			g.Scenario.Phases = g.Scenario.Phases[1:]
			break
		}

		// Switch to the next level if there are no more phases
		if len(g.Scenario.Phases) == 0 {
			g.Scenarios = g.Scenarios[1:]
			g.State = SwitchLevel
		}
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{0, 55, 0, 255})

	switch g.State {
	case Playing:
		g.drawGameplay(screen)
	case GameOver:
		text_op := &text.DrawOptions{}
		text_op.GeoM.Translate(ScreenWidth/2-100, ScreenHeight/2-50)
		text.Draw(screen, "Game Over", mplusNormalFace, text_op)
		text_op.GeoM.Translate(0, 50)
		text.Draw(screen, fmt.Sprintf("Score: %d", g.Player.Score), mplusNormalFace, text_op)
	case Paused:
		text.Draw(screen, "Paused", mplusNormalFace, &text.DrawOptions{})
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
		Hitbox:      Hitbox{Width: 10, Height: 10},
		Grazebox:    Hitbox{Width: 50, Height: 50},
	}

	player.Hitbox.CenterOn(player.X+player.Width/2, player.Y+player.Height/2)
	player.Grazebox.CenterOn(player.X+player.Width/2, player.Y+player.Height/2)
	randomSource := rand.New(rand.NewSource(time.Now().UnixNano()))

	return &Game{
		Player:       player,
		Projectiles:  []*Projectile{},
		Enemies:      []*Enemy{},
		RandomSource: randomSource,
		FlagHitboxes: false,
		State:        SwitchLevel,
		Scenarios:    getGameScenarios(),
	}
}
