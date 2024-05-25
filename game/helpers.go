package game

func (hb *Hitbox) CenterOn(centerX, centerY float32) {
	hb.X = centerX - hb.Width/2
	hb.Y = centerY - hb.Height/2
}

// Deletes projectile from slice by index
func deleteProjectile(slice []*Projectile, index int) []*Projectile {
	return append(slice[:index], slice[index+1:]...)
}
