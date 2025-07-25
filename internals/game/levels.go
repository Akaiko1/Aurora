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
	// Create Phases with different enemy types and weapons
	Phase1 := &Phase{
		Name: "Phase 1",
		Enemies: []*entities.Enemy{
			mob(entities.EnemyTypeSpider, entities.WeaponIDEnemy),
			mob(entities.EnemyTypeSpider, entities.WeaponIDEnemy),
			mob(entities.EnemyTypeSpider, entities.WeaponIDEnemy),
		},
	}
	Phase2 := &Phase{
		Name: "Phase 2",
		Enemies: []*entities.Enemy{
			mob(entities.EnemyTypeGoblin, entities.WeaponIDEnemy),
			mob(entities.EnemyTypeDragon, entities.WeaponIDDragonSpiral),
		},
	}
	Final := &Phase{
		Name: "Final",
		Enemies: []*entities.Enemy{
			mob(entities.EnemyTypeSkeleton, entities.WeaponIDSkeletonPiercer),
		},
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

func mob(enemyType entities.EnemyType, weaponID entities.WeaponID) *entities.Enemy {
	rands := rand.New(rand.NewSource(time.Now().UnixNano()))
	x := float32(rands.Intn(config.ScreenWidth - config.EntitySize*2))
	y := float32(rands.Intn(config.ScreenHeight - config.EntitySize*2))

	// Create enemy of the specified type with the specified weapon
	enemy := entities.NewEnemy(x, y, enemyType, weaponID)

	return enemy
}
