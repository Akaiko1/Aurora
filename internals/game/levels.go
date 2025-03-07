package game

import (
	"math/rand"
	"scroller_game/internals/config"
	"scroller_game/internals/entities"
	"time"
)

type Scenario struct {
	Name   string
	Phases []*Phase
}

type Phase struct {
	Name    string
	Enemies []*entities.Enemy
}

func GetGameScenarios() []*Scenario {
	// Create Phases
	Phase1 := &Phase{
		Name:    "Phase 1",
		Enemies: []*entities.Enemy{mob(), mob(), mob()},
	}
	Phase2 := &Phase{
		Name:    "Phase 2",
		Enemies: []*entities.Enemy{mob(), mob()},
	}
	Final := &Phase{
		Name:    "Final",
		Enemies: []*entities.Enemy{mob()},
	}

	// Create Scenarios
	Scenario1 := &Scenario{
		Name:   "Scenario 1",
		Phases: []*Phase{Phase1, Phase2, Final},
	}

	Scenario2 := &Scenario{
		Name:   "Scenario 2",
		Phases: []*Phase{Phase1, Phase2, Final},
	}

	// Set up the game
	scenarios := []*Scenario{Scenario1, Scenario2}
	return scenarios
}

func mob() *entities.Enemy {
	rands := rand.New(rand.NewSource(time.Now().UnixNano()))
	return &entities.Enemy{X: float32(rands.Intn(config.ScreenWidth - 50)), Y: float32(rands.Intn(config.ScreenHeight - 50)),
		Width: 32, Height: 32, SpeedX: 1, SpeedY: 1}
}
