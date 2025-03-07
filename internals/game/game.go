package game

import (
	"fmt"
	"image/color"
	"math/rand"
	"scroller_game/internals/config"
	"scroller_game/internals/entities"
	"scroller_game/internals/inputs"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

const (
	StartMenu GameState = iota
	Playing
	SwitchLevel
	SwitchPhase
	Paused
	GameOver
)

type GameState int

type Game struct {
	Player       *entities.Player
	Projectiles  []*entities.Projectile
	FrameCount   int
	Enemies      []*entities.Enemy
	Spawned      []*entities.Enemy
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
	Enemies []*entities.Enemy
}

func init() {
	mplusNormalFace, frames = inputs.SetFontAndImages()
}

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
		g.DrawGameplay(screen)
	case GameOver:
		text_op := &text.DrawOptions{}
		text_op.GeoM.Translate(config.ScreenWidth/2-100, config.ScreenHeight/2-50)
		text.Draw(screen, "Game Over", mplusNormalFace, text_op)
		text_op.GeoM.Translate(0, 50)
		text.Draw(screen, fmt.Sprintf("Score: %d", g.Player.Score), mplusNormalFace, text_op)
	case Paused:
		text.Draw(screen, "Paused", mplusNormalFace, &text.DrawOptions{})
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return config.ScreenWidth, config.ScreenHeight
}
