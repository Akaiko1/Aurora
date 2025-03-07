package physics

type Hitbox struct {
	X, Y, Width, Height float32
}

func (hb *Hitbox) CenterOn(centerX, centerY float32) {
	hb.X = centerX - hb.Width/2
	hb.Y = centerY - hb.Height/2
}

// Intersects returns true if the hitbox intersects with another hitbox
func (hb *Hitbox) Intersects(other *Hitbox) bool {
	return hb.X < other.X+other.Width &&
		hb.X+hb.Width > other.X &&
		hb.Y < other.Y+other.Height &&
		hb.Y+hb.Height > other.Y
}
