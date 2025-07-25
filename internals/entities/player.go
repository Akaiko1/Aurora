package entities

import (
	"scroller_game/internals/physics"
)

type Player struct {
	X, Y, Width, Height, Speed float32
	Projectiles                []*Projectile
	Hitbox                     physics.Hitbox
	Grazebox                   physics.Hitbox
	Grazing                    *Projectile
	Hits                       int
	Score                      int
	IsAttacking                bool
	AttackStartFrame           int
	CurrentWeapon              *Weapon   // Current weapon instance
	AvailableWeapons           []*Weapon // Weapons player has access to
	CurrentWeaponIndex         int       // Index in AvailableWeapons
}

// UpdatePosition updates the player's position and automatically syncs hitboxes
func (p *Player) UpdatePosition(x, y float32) {
	p.X = x
	p.Y = y
	// Center hitbox on player
	p.Hitbox.CenterOn(p.X+p.Width/2, p.Y+p.Height/2)
	p.Grazebox.CenterOn(p.X+p.Width/2, p.Y+p.Height/2)
}

// Move updates position by delta values
func (p *Player) Move(dx, dy float32) {
	p.UpdatePosition(p.X+dx, p.Y+dy)
}

// SwitchWeapon changes to the weapon at the given index
func (p *Player) SwitchWeapon(index int) {
	if index >= 0 && index < len(p.AvailableWeapons) {
		p.CurrentWeaponIndex = index
		p.CurrentWeapon = p.AvailableWeapons[index]
		// Reset fire time when switching weapons to allow immediate firing
		p.CurrentWeapon.LastFireTime = 0
	}
}

// SwitchToWeaponID switches to a specific weapon by ID
func (p *Player) SwitchToWeaponID(id WeaponID) {
	for i, weapon := range p.AvailableWeapons {
		if weapon.Definition.ID == id {
			p.SwitchWeapon(i)
			break
		}
	}
}

// GetMaxProjectiles returns the maximum projectiles for current weapon
func (p *Player) GetMaxProjectiles() int {
	if p.CurrentWeapon == nil {
		return 3 // Fallback
	}
	return p.CurrentWeapon.Definition.MaxProjectiles
}

// CanFire checks if the player can fire based on weapon fire rate
func (p *Player) CanFire(currentFrame int) bool {
	if p.CurrentWeapon == nil {
		return false
	}
	return len(p.Projectiles) < p.GetMaxProjectiles() && p.CurrentWeapon.CanFire(currentFrame)
}

// CreateProjectile creates a projectile using current weapon
func (p *Player) CreateProjectile(currentFrame int) *Projectile {
	if p.CurrentWeapon == nil {
		return nil
	}
	// Player shoots upward (-1 direction)
	centerX := p.X + p.Width/2 - p.CurrentWeapon.Definition.ProjectileWidth/2
	return p.CurrentWeapon.CreateProjectile(centerX, p.Y, -1, currentFrame)
}

// AddWeapon adds a weapon to the player's available weapons
func (p *Player) AddWeapon(weaponID WeaponID) {
	weapon := GetWeapon(weaponID)
	p.AvailableWeapons = append(p.AvailableWeapons, weapon)
}

// InitializeWeapons sets up the player's default weapons
func (p *Player) InitializeWeapons() {
	p.AvailableWeapons = []*Weapon{
		GetWeapon(WeaponIDNormal),   // Key 1
		GetWeapon(WeaponIDPiercing), // Key 2
		GetWeapon(WeaponIDRapid),    // Key 3
		GetWeapon(WeaponIDHeavy),    // Key 4
	}
	p.CurrentWeaponIndex = 0
	p.CurrentWeapon = p.AvailableWeapons[0]
}
