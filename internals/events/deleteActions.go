package events

import "scroller_game/internals/entities"

func DeleteProjectile(slice []*entities.Projectile, index int) []*entities.Projectile {
	return append(slice[:index], slice[index+1:]...)
}

func DeleteEnemy(slice []*entities.Enemy, enemy *entities.Enemy) ([]*entities.Enemy, bool) {
	for i, e := range slice {
		if e == enemy {
			return append(slice[:i], slice[i+1:]...), true
		}
	}
	return slice, false
}
