package game

import (
	"math/rand"
	"time"
)

func getGameScenarios() []*Scenario {
	Phase1 := &Phase{
		Name:    "Phase 1",
		Enemies: []*Enemy{mob(), mob(), mob(), mob(), mob(), mob()},
	}
	Phase2 := &Phase{
		Name:    "Phase 2",
		Enemies: []*Enemy{mob(), mob(), mob(), mob(), mob(), mob()},
	}
	Final := &Phase{
		Name:    "Final",
		Enemies: []*Enemy{mob(), mob(), mob(), mob(), mob(), mob()},
	}

	Scenario1 := &Scenario{
		Name:   "Scenario 1",
		Phases: []*Phase{Phase1, Phase2, Final},
	}

	scenarios := []*Scenario{Scenario1}
	return scenarios
}

func mob() *Enemy {
	rands := rand.New(rand.NewSource(time.Now().UnixNano()))
	return &Enemy{X: float32(rands.Intn(ScreenWidth - 50)), Y: float32(rands.Intn(ScreenHeight - 50)),
		Width: 32, Height: 32, SpeedX: 1, SpeedY: 1}
}
