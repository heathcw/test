package main

import ()

type Player struct {
	Health       int
	PlayerX      float32
	PlayerY      float32
	Speed        float32
	Velocity     float32
	Cooldown     int
	HurtCooldown int
	Hurt         bool
}

func (p *Player) updateX(x float32) {
	p.PlayerX += x
}

func (p *Player) updateY(y float32) {
	p.PlayerY += y
}

type Projectile struct {
	X, Y   float32
	VX, VY float32 // velocity
}
