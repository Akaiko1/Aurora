package game

import (
	"fmt"
	"image/color"
	"math/rand"
	"scroller_game/internals/config"
	"scroller_game/internals/entities"
	"scroller_game/internals/inputs"
	"scroller_game/internals/physics"

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
	Player         *entities.Player
	Projectiles    []*entities.Projectile
	FrameCount     int
	Enemies        []*entities.Enemy
	SpawnedEnemies []*entities.Enemy
	State          GameState
	Scenario       *Scenario
	Phase          *Phase
	RandomSource   *rand.Rand
	Scenarios      []*Scenario
	FlagHitboxes   bool
	Background     *Background
	SpatialGrid    *physics.SpatialGrid
}

func init() {
	mplusNormalFace, _ = inputs.SetFontAndImages()
}

func (g *Game) Update() error {
	// Increment the frame count
	switch g.State {
	case Playing:
		g.FrameCount = (g.FrameCount + 1) % 120

		if len(g.Enemies) == 0 && len(g.SpawnedEnemies) == 0 {
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
		for _, enemy := range g.SpawnedEnemies {
			g.enemyActions(enemy)
		}

		// Handle collisions using optimized spatial partitioning
		g.UpdateCollisions()
	case SwitchLevel:
		if len(g.Scenarios) == 0 {
			g.State = GameOver
			return nil
		}

		g.Scenario = g.Scenarios[0]
		g.Scenarios = g.Scenarios[1:]

		// We have a scenario but need to check if it has phases
		if len(g.Scenario.Phases) == 0 {
			g.State = SwitchLevel // Stay in this state to process next scenario
			return nil
		}

		// Setup first phase and start playing
		g.Phase = g.Scenario.Phases[0]
		g.Scenario.Phases = g.Scenario.Phases[1:]
		g.Enemies = g.Phase.Enemies
		g.State = Playing

	case SwitchPhase:
		if len(g.Scenario.Phases) == 0 {
			// No more phases, move to next scenario
			g.State = SwitchLevel
			return nil
		}

		// Setup next phase and continue playing
		g.Phase = g.Scenario.Phases[0]
		g.Scenario.Phases = g.Scenario.Phases[1:]
		g.Enemies = g.Phase.Enemies
		g.State = Playing
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	// Draw the background first
	if g.Background != nil {
		g.Background.Draw(screen)
	} else {
		// Fallback if background isn't initialized
		screen.Fill(color.RGBA{0, 55, 0, 255})
	}

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
