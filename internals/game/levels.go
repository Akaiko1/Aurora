package game

import (
	"math/rand"
	"scroller_game/internals/config"
	"scroller_game/internals/entities"
	"time"
)

func GetGameScenarios() []*Scenario {
	Phase1 := &Phase{
		Name:    "Phase 1",
		Enemies: []*entities.Enemy{mob(), mob(), mob(), mob(), mob(), mob()},
	}
	Phase2 := &Phase{
		Name:    "Phase 2",
		Enemies: []*entities.Enemy{mob(), mob(), mob(), mob(), mob(), mob()},
	}
	Final := &Phase{
		Name:    "Final",
		Enemies: []*entities.Enemy{mob(), mob(), mob(), mob(), mob(), mob()},
	}

	Scenario1 := &Scenario{
		Name:   "Scenario 1",
		Phases: []*Phase{Phase1, Phase2, Final},
	}

	scenarios := []*Scenario{Scenario1}
	return scenarios
}

func mob() *entities.Enemy {
	rands := rand.New(rand.NewSource(time.Now().UnixNano()))
	return &entities.Enemy{X: float32(rands.Intn(config.ScreenWidth - 50)), Y: float32(rands.Intn(config.ScreenHeight - 50)),
		Width: 32, Height: 32, SpeedX: 1, SpeedY: 1}
}
