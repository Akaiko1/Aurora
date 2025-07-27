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
	Player            *entities.Player
	Projectiles       []*entities.Projectile
	FrameCount        int
	Enemies           []*entities.Enemy
	SpawnedEnemies    []*entities.Enemy
	State             GameState
	Scenario          *Scenario
	Phase             *Phase
	RandomSource      *rand.Rand
	Scenarios         []*Scenario
	FlagHitboxes      bool
	Background        *Background
	SpatialGrid       *physics.SpatialGrid
	GameOverSelection int // 0 = Again, 1 = Exit
}

// InitializeAssets initializes fonts and images - called from main after embedded setup
func InitializeAssets() {
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

		// Update enemies
		for _, enemy := range g.SpawnedEnemies {
			g.enemyActions(enemy)
		}

		// Update all projectiles
		g.updateProjectiles()

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
	case GameOver:
		// Handle game over menu navigation
		g.gameOverMenuEvents()
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
		g.drawGameOverMenu(screen)
	case Paused:
		text.Draw(screen, "Paused", mplusNormalFace, &text.DrawOptions{})
	}
}

// drawGameOverMenu renders the game over screen with interactive menu options.
// It displays the final score and provides selectable options to restart or exit.
// Uses proper visual indicators to show the currently selected option.
func (g *Game) drawGameOverMenu(screen *ebiten.Image) {
	// Game Over title
	text_op := &text.DrawOptions{}
	text_op.GeoM.Translate(config.ScreenWidth/2-60, config.ScreenHeight/2-120)
	text.Draw(screen, "Game Over", mplusNormalFace, text_op)

	// Final score
	text_op.GeoM.Reset()
	text_op.GeoM.Translate(config.ScreenWidth/2-60, config.ScreenHeight/2-80)
	text.Draw(screen, fmt.Sprintf("Score: %d", g.Player.Score), mplusNormalFace, text_op)

	// Menu options with selection indicator
	againText := "Play Again"
	exitText := "Exit"

	if g.GameOverSelection == 0 {
		againText = "> " + againText + " <"
	} else {
		exitText = "> " + exitText + " <"
	}

	text_op.GeoM.Reset()
	text_op.GeoM.Translate(config.ScreenWidth/2-60, config.ScreenHeight/2-20)
	text.Draw(screen, againText, mplusNormalFace, text_op)

	text_op.GeoM.Reset()
	text_op.GeoM.Translate(config.ScreenWidth/2-60, config.ScreenHeight/2+20)
	text.Draw(screen, exitText, mplusNormalFace, text_op)
}

// restartGame resets the game to its initial state for a new playthrough.
// It reinitializes the player, clears all projectiles and enemies, and
// resets the game progression back to the first scenario.
func (g *Game) restartGame() {
	// Reset player state
	g.Player.X = config.ScreenWidth / 2
	g.Player.Y = config.ScreenHeight / 2
	g.Player.Score = 0
	g.Player.Hits = 0
	g.Player.Projectiles = []*entities.Projectile{}
	g.Player.IsAttacking = false
	g.Player.Grazing = nil
	g.Player.InitializeWeapons()
	g.Player.Hitbox.CenterOn(g.Player.X+g.Player.Width/2, g.Player.Y+g.Player.Height/2)
	g.Player.Grazebox.CenterOn(g.Player.X+g.Player.Width/2, g.Player.Y+g.Player.Height/2)

	// Clear all projectiles and enemies
	g.Projectiles = []*entities.Projectile{}
	g.Enemies = []*entities.Enemy{}
	g.SpawnedEnemies = []*entities.Enemy{}

	// Reset game state
	g.FrameCount = 0
	g.State = SwitchLevel
	g.Scenarios = GetGameScenarios()
	g.Scenario = nil
	g.Phase = nil
	g.GameOverSelection = 0

	// Reset spatial grid
	g.SpatialGrid = physics.NewSpatialGrid(config.SpatialGridCellSize)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return config.ScreenWidth, config.ScreenHeight
}
