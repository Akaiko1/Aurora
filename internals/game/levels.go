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
			mob(entities.EnemyTypeSpider, entities.WeaponIDEnemy, 1, 1), // 1 HP, 1 point
			mob(entities.EnemyTypeSpider, entities.WeaponIDEnemy, 1, 1), // 1 HP, 1 point
			mob(entities.EnemyTypeSpider, entities.WeaponIDEnemy, 1, 1), // 1 HP, 1 point
		},
	}
	Phase2 := &Phase{
		Name: "Phase 2",
		Enemies: []*entities.Enemy{
			mob(entities.EnemyTypeGoblin, entities.WeaponIDEnemy, 2, 2),
			mob(entities.EnemyTypeSpider, entities.WeaponIDEnemy, 1, 1),
			mob(entities.EnemyTypeSpider, entities.WeaponIDEnemy, 1, 1),
		},
	}
	Phase3 := &Phase{
		Name: "Phase 3",
		Enemies: []*entities.Enemy{
			mob(entities.EnemyTypeSkeleton, entities.WeaponIDSkeletonPiercer, 3, 5),
			mob(entities.EnemyTypeDragon, entities.WeaponIDDragonSpiral, 5, 8),
		},
	}
	Phase4 := &Phase{
		Name: "Phase 4",
		Enemies: []*entities.Enemy{
			mob(entities.EnemyTypeGoblin, entities.WeaponIDEnemy, 2, 2),
			mob(entities.EnemyTypeGoblin, entities.WeaponIDSkeletonPiercer, 2, 3),
			mob(entities.EnemyTypeGoblin, entities.WeaponIDEnemy, 2, 2),
		},
	}
	Phase5 := &Phase{
		Name: "Phase 5",
		Enemies: []*entities.Enemy{
			mob(entities.EnemyTypeGoblin, entities.WeaponIDSkeletonPiercer, 2, 3),
			mob(entities.EnemyTypeGoblin, entities.WeaponIDSkeletonPiercer, 2, 3),
			mob(entities.EnemyTypeDragon, entities.WeaponIDDragonSpiral, 5, 5),
			mob(entities.EnemyTypeDragon, entities.WeaponIDDragonSpiral, 5, 5),
		},
	}
	Final := &Phase{
		Name: "Final",
		Enemies: []*entities.Enemy{
			mob(entities.EnemyTypeDragon, entities.WeaponIDDragonSpiral, 12, 8),
		},
	}

	// Create Scenarios
	Scenario1 := &Scenario{
		Name:   "Scenario 1",
		Phases: []*Phase{Phase1, Phase2, Phase3},
	}

	Scenario2 := &Scenario{
		Name:   "Scenario 2",
		Phases: []*Phase{Phase4, Phase5, Final},
	}

	// Set up the game
	scenarios := []*Scenario{Scenario1, Scenario2}
	return scenarios
}

func mob(enemyType entities.EnemyType, weaponID entities.WeaponID, hitPoints int, scoreValue int) *entities.Enemy {
	rands := rand.New(rand.NewSource(time.Now().UnixNano()))
	x := float32(rands.Intn(config.ScreenWidth - config.EntitySize*2))
	y := float32(rands.Intn(config.ScreenHeight - config.EntitySize*2))

	// Create enemy of the specified type with the specified weapon, HP, and score
	enemy := entities.NewEnemy(x, y, enemyType, weaponID, hitPoints, scoreValue)

	return enemy
}
