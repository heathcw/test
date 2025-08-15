package main

import (
	"math/rand"
)

func wave(g *Game) {
	switch g.Wave {
	case 1:
		for i := range 10 {
			newEnemy := Player{
				Health:   10,
				PlayerX:  screenWidth - 10,
				PlayerY:  float32(20 + i*((screenHeight-40)/9)),
				Speed:    (r1.Float32() / 2),
				Cooldown: r1.Intn(300) + 100,
			}
			g.Enemies = append(g.Enemies, newEnemy)
		}
	case 2:
		for i := range 30 {
			newEnemy := Player{
				Health:   10,
				PlayerX:  screenWidth - (r1.Float32() * 100),
				PlayerY:  float32(screenHeight - (i * r1.Intn(10))),
				Speed:    (r1.Float32() / 2),
				Velocity: float32(1 - 2*(r1.Intn(2))),
				Cooldown: rand.Intn(300) + 100,
			}
			g.Enemies = append(g.Enemies, newEnemy)
		}
	case 3:
		for i := range 30 {
			newEnemy := Player{
				Health:   10,
				PlayerX:  screenWidth - (r1.Float32() * 100),
				PlayerY:  float32(screenHeight - (i * r1.Intn(10))),
				Speed:    (r1.Float32() / 2),
				Velocity: float32(1 - 2*(rand.Intn(2))),
				Cooldown: rand.Intn(100) + 100,
			}
			g.Enemies = append(g.Enemies, newEnemy)
		}
		for i := range 10 {
			newEnemy := Player{
				Health:   15,
				PlayerX:  screenWidth - 10,
				PlayerY:  float32(20 + i*((screenHeight-40)/9)),
				Speed:    (r1.Float32() / 2),
				Cooldown: rand.Intn(100) + 100,
			}
			g.Enemies = append(g.Enemies, newEnemy)
		}
	case 4:
		newEnemy := Player{
			Health:   100,
			PlayerX:  screenWidth - 50,
			PlayerY:  screenHeight / 2,
			Speed:    .5,
			Velocity: 1.0,
			Cooldown: rand.Intn(250) + 250,
		}
		g.Enemies = append(g.Enemies, newEnemy)
	case 5:
		g.Player.Health += 30
		for i := range 10 {
			newEnemy := Player{
				Health:   20,
				PlayerX:  screenWidth - 10,
				PlayerY:  float32(20 + i*((screenHeight-40)/9)),
				Speed:    (r1.Float32() / 3),
				Cooldown: rand.Intn(250) + 50,
			}
			g.Enemies = append(g.Enemies, newEnemy)
		}
	case 6:
		for i := range 30 {
			newEnemy := Player{
				Health:   20,
				PlayerX:  float32(rand.Intn(screenWidth-100) + 50),
				PlayerY:  float32(screenHeight - (i * r1.Intn(10))),
				Speed:    (r1.Float32() / 2),
				Velocity: float32(1 - 2*(rand.Intn(2))),
				Cooldown: rand.Intn(300) + 100,
			}
			g.Enemies = append(g.Enemies, newEnemy)
		}
	case 7:
		for i := range 10 {
			newEnemy := Player{
				Health:   20,
				PlayerX:  screenWidth - 10,
				PlayerY:  float32(20 + i*((screenHeight-40)/9)),
				Speed:    (r1.Float32() / 3),
				Cooldown: rand.Intn(250) + 50,
			}
			g.Enemies = append(g.Enemies, newEnemy)
		}
		for i := range 20 {
			newEnemy := Player{
				Health:   20,
				PlayerX:  screenWidth - (r1.Float32() * 100),
				PlayerY:  float32(screenHeight - (i * r1.Intn(10))),
				Speed:    (r1.Float32() / 2),
				Velocity: float32(1 - 2*(rand.Intn(2))),
				Cooldown: rand.Intn(300) + 100,
			}
			g.Enemies = append(g.Enemies, newEnemy)
		}
		newEnemy := Player{
			Health:   30,
			PlayerX:  screenWidth - 50,
			PlayerY:  screenHeight / 2,
			Speed:    1,
			Velocity: g.Player.Velocity,
			Cooldown: rand.Intn(50) + 50,
		}
		g.Enemies = append(g.Enemies, newEnemy)
	case 8:
		newEnemy := Player{
			Health:   200,
			PlayerX:  screenWidth - 50,
			PlayerY:  screenHeight / 2,
			Speed:    .75,
			Velocity: g.Player.Velocity,
			Cooldown: rand.Intn(50) + 50,
		}
		g.Enemies = append(g.Enemies, newEnemy)
	}
	g.NextWave = false
}
