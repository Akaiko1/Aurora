package game

func (hb *Hitbox) CenterOn(centerX, centerY float32) {
	hb.X = centerX - hb.Width/2
	hb.Y = centerY - hb.Height/2
}
