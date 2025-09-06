package main

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
				Cooldown: r1.Intn(300) + 100,
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
				Velocity: float32(1 - 2*(r1.Intn(2))),
				Cooldown: r1.Intn(100) + 100,
			}
			g.Enemies = append(g.Enemies, newEnemy)
		}
		for i := range 10 {
			newEnemy := Player{
				Health:   15,
				PlayerX:  screenWidth - 10,
				PlayerY:  float32(20 + i*((screenHeight-40)/9)),
				Speed:    (r1.Float32() / 2),
				Cooldown: r1.Intn(100) + 100,
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
			Cooldown: r1.Intn(250) + 250,
		}
		g.Enemies = append(g.Enemies, newEnemy)
	case 5:
		g.Player.Health += 30

		newPowerUp := PowerUp{
			X:     0,
			Y:     screenHeight / 2,
			VX:    .5,
			VY:    0,
			Power: "Speed",
			Got:   false,
			Draw:  true,
		}
		g.Powerups["Speed"] = newPowerUp

		for i := range 10 {
			newEnemy := Player{
				Health:   20,
				PlayerX:  screenWidth - 10,
				PlayerY:  float32(20 + i*((screenHeight-40)/9)),
				Speed:    (r1.Float32() / 3),
				Cooldown: r1.Intn(250) + 50,
			}
			g.Enemies = append(g.Enemies, newEnemy)
		}
	case 6:
		for i := range 30 {
			newEnemy := Player{
				Health:   20,
				PlayerX:  float32(r1.Intn(screenWidth-100) + 50),
				PlayerY:  float32(screenHeight - (i * r1.Intn(10))),
				Speed:    (r1.Float32() / 2),
				Velocity: float32(1 - 2*(r1.Intn(2))),
				Cooldown: r1.Intn(300) + 100,
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
				Cooldown: r1.Intn(250) + 50,
			}
			g.Enemies = append(g.Enemies, newEnemy)
		}
		for i := range 20 {
			newEnemy := Player{
				Health:   20,
				PlayerX:  screenWidth - (r1.Float32() * 100),
				PlayerY:  float32(screenHeight - (i * r1.Intn(10))),
				Speed:    (r1.Float32() / 2),
				Velocity: float32(1 - 2*(r1.Intn(2))),
				Cooldown: r1.Intn(300) + 100,
			}
			g.Enemies = append(g.Enemies, newEnemy)
		}
		newEnemy := Player{
			Health:   30,
			PlayerX:  screenWidth - 50,
			PlayerY:  screenHeight / 2,
			Speed:    1,
			Velocity: g.Player.Velocity,
			Cooldown: r1.Intn(50) + 50,
		}
		g.Enemies = append(g.Enemies, newEnemy)
	case 8:
		newEnemy := Player{
			Health:   200,
			PlayerX:  screenWidth - 50,
			PlayerY:  screenHeight / 2,
			Speed:    .75,
			Velocity: g.Player.Velocity,
			Cooldown: r1.Intn(50) + 50,
		}
		g.Enemies = append(g.Enemies, newEnemy)
	case 9:
		g.Player.Health += 30

		newPowerUp := PowerUp{
			X:     0,
			Y:     screenHeight / 2,
			VX:    .5,
			VY:    0,
			Power: "Big",
			Got:   false,
			Draw:  true,
		}
		g.Powerups["Big"] = newPowerUp

		newEnemy := Player{
			Health:   30,
			PlayerX:  screenWidth - 50,
			PlayerY:  screenHeight - 50,
			Speed:    .75,
			Velocity: g.Player.Velocity,
			Cooldown: r1.Intn(50) + 50,
		}
		g.Enemies = append(g.Enemies, newEnemy)
		newEnemy = Player{
			Health:   30,
			PlayerX:  screenWidth - 50,
			PlayerY:  50,
			Speed:    .5,
			Velocity: -g.Player.Velocity,
			Cooldown: r1.Intn(50) + 50,
		}
		g.Enemies = append(g.Enemies, newEnemy)
	case 10:
		newEnemy := Player{
			Health:   30,
			PlayerX:  screenWidth - 50,
			PlayerY:  screenHeight - 50,
			Speed:    .75,
			Velocity: g.Player.Velocity,
			Cooldown: r1.Intn(50) + 50,
		}
		g.Enemies = append(g.Enemies, newEnemy)
		newEnemy = Player{
			Health:   30,
			PlayerX:  screenWidth - 50,
			PlayerY:  100,
			Speed:    .5,
			Velocity: -g.Player.Velocity,
			Cooldown: r1.Intn(50) + 50,
		}
		g.Enemies = append(g.Enemies, newEnemy)
		newEnemy = Player{
			Health:   30,
			PlayerX:  screenWidth - 50,
			PlayerY:  screenHeight - 100,
			Speed:    .25,
			Velocity: g.Player.Velocity,
			Cooldown: r1.Intn(50) + 50,
		}
		g.Enemies = append(g.Enemies, newEnemy)
		newEnemy = Player{
			Health:   30,
			PlayerX:  screenWidth - 50,
			PlayerY:  50,
			Speed:    1,
			Velocity: -g.Player.Velocity,
			Cooldown: r1.Intn(50) + 50,
		}
		g.Enemies = append(g.Enemies, newEnemy)
	case 11:
		for i := range 2 {
			newEnemy := Player{
				Health:   30,
				PlayerX:  float32(screenWidth - (50 * (i + 1))),
				PlayerY:  screenHeight - 50,
				Speed:    r1.Float32(),
				Velocity: g.Player.Velocity,
				Cooldown: r1.Intn(50) + 50,
			}
			g.Enemies = append(g.Enemies, newEnemy)
			newEnemy = Player{
				Health:   30,
				PlayerX:  float32(screenWidth - (50 * (i + 1))),
				PlayerY:  100,
				Speed:    r1.Float32(),
				Velocity: -g.Player.Velocity,
				Cooldown: r1.Intn(50) + 50,
			}
			g.Enemies = append(g.Enemies, newEnemy)
			newEnemy = Player{
				Health:   30,
				PlayerX:  float32(screenWidth - (50 * (i + 1))),
				PlayerY:  screenHeight - 100,
				Speed:    r1.Float32(),
				Velocity: g.Player.Velocity,
				Cooldown: r1.Intn(50) + 50,
			}
			g.Enemies = append(g.Enemies, newEnemy)
			newEnemy = Player{
				Health:   30,
				PlayerX:  float32(screenWidth - (50 * (i + 1))),
				PlayerY:  50,
				Speed:    r1.Float32(),
				Velocity: -g.Player.Velocity,
				Cooldown: r1.Intn(50) + 50,
			}
			g.Enemies = append(g.Enemies, newEnemy)
		}
	case 12:
		newEnemy := Player{
			Health:   300,
			PlayerX:  screenWidth - 50,
			PlayerY:  screenHeight / 2,
			Speed:    .9,
			Velocity: -g.Player.Velocity,
			Cooldown: r1.Intn(50) + 20,
		}
		g.Enemies = append(g.Enemies, newEnemy)
	case 13:
		newPowerUp := PowerUp{
			X:     0,
			Y:     screenHeight / 2,
			VX:    .5,
			VY:    0,
			Power: "Spread",
			Got:   false,
			Draw:  true,
		}
		g.Powerups["Spread"] = newPowerUp

		g.Player.Health += 30
		newEnemy := Player{
			Health:   30,
			PlayerX:  screenWidth / 2,
			PlayerY:  screenHeight - 50,
			Speed:    2,
			Velocity: -1.0,
			Cooldown: r1.Intn(50) + 20,
		}
		g.Enemies = append(g.Enemies, newEnemy)

		newEnemy = Player{
			Health:   30,
			PlayerX:  screenWidth / 2,
			PlayerY:  50,
			Speed:    2,
			Velocity: 1.0,
			Cooldown: r1.Intn(50) + 20,
		}
		g.Enemies = append(g.Enemies, newEnemy)
	case 14:
		newEnemy := Player{
			Health:   30,
			PlayerX:  screenWidth - 50,
			PlayerY:  screenHeight - 50,
			Speed:    2,
			Velocity: -1.0,
			Cooldown: r1.Intn(50) + 20,
		}
		g.Enemies = append(g.Enemies, newEnemy)

		newEnemy = Player{
			Health:   30,
			PlayerX:  screenWidth - 50,
			PlayerY:  50,
			Speed:    2,
			Velocity: 1.0,
			Cooldown: r1.Intn(50) + 20,
		}
		g.Enemies = append(g.Enemies, newEnemy)

		newEnemy = Player{
			Health:   30,
			PlayerX:  screenWidth / 2,
			PlayerY:  screenHeight - 50,
			Speed:    1.5,
			Velocity: -1.0,
			Cooldown: r1.Intn(50) + 20,
		}
		g.Enemies = append(g.Enemies, newEnemy)

		newEnemy = Player{
			Health:   30,
			PlayerX:  screenWidth / 4,
			PlayerY:  50,
			Speed:    1.5,
			Velocity: 1.0,
			Cooldown: r1.Intn(50) + 20,
		}
		g.Enemies = append(g.Enemies, newEnemy)
	case 15:
		newEnemy := Player{
			Health:   30,
			PlayerX:  screenWidth - 50,
			PlayerY:  screenHeight - 50,
			Speed:    2,
			Velocity: -1.0,
			Cooldown: r1.Intn(50) + 20,
		}
		g.Enemies = append(g.Enemies, newEnemy)

		newEnemy = Player{
			Health:   30,
			PlayerX:  screenWidth - 55,
			PlayerY:  50,
			Speed:    2,
			Velocity: 1.0,
			Cooldown: r1.Intn(50) + 20,
		}
		g.Enemies = append(g.Enemies, newEnemy)

		newEnemy = Player{
			Health:   30,
			PlayerX:  screenWidth / 2,
			PlayerY:  screenHeight - 50,
			Speed:    2,
			Velocity: -1.0,
			Cooldown: r1.Intn(50) + 20,
		}
		g.Enemies = append(g.Enemies, newEnemy)

		newEnemy = Player{
			Health:   30,
			PlayerX:  screenWidth/2 + 5,
			PlayerY:  50,
			Speed:    2,
			Velocity: 1.0,
			Cooldown: r1.Intn(50) + 20,
		}
		g.Enemies = append(g.Enemies, newEnemy)

		newEnemy = Player{
			Health:   30,
			PlayerX:  110,
			PlayerY:  50,
			Speed:    1.5,
			Velocity: 1.0,
			Cooldown: r1.Intn(50) + 20,
		}
		g.Enemies = append(g.Enemies, newEnemy)

		newEnemy = Player{
			Health:   30,
			PlayerX:  90,
			PlayerY:  screenHeight - 50,
			Speed:    1.5,
			Velocity: -1.0,
			Cooldown: r1.Intn(50) + 20,
		}
		g.Enemies = append(g.Enemies, newEnemy)

		newEnemy = Player{
			Health:   30,
			PlayerX:  70,
			PlayerY:  50,
			Speed:    1.5,
			Velocity: 1.0,
			Cooldown: r1.Intn(50) + 20,
		}
		g.Enemies = append(g.Enemies, newEnemy)

		newEnemy = Player{
			Health:   30,
			PlayerX:  50,
			PlayerY:  screenHeight - 50,
			Speed:    1.5,
			Velocity: -1.0,
			Cooldown: r1.Intn(50) + 20,
		}
		g.Enemies = append(g.Enemies, newEnemy)
	case 16:
		newEnemy := Player{
			Health:   400,
			PlayerX:  screenWidth,
			PlayerY:  screenHeight / 2,
			Speed:    .9,
			Velocity: -g.Player.Velocity,
			Cooldown: 30,
		}
		g.Enemies = append(g.Enemies, newEnemy)
	case 17:
		g.Player.Health += 30
		newPowerUp := PowerUp{
			X:     0,
			Y:     screenHeight / 2,
			VX:    .5,
			VY:    0,
			Power: "Blast",
			Got:   false,
			Draw:  true,
		}
		g.Powerups["Blast"] = newPowerUp
		for i := range 5 {
			newEnemy := Player{
				Health:   100,
				PlayerX:  float32(screenWidth - r1.Intn(150) - 50),
				PlayerY:  float32(20 + i*((screenHeight-40)/4)),
				Speed:    .5,
				Velocity: float32(1 - 2*(r1.Intn(2))),
				Cooldown: r1.Intn(250) + 50,
			}
			g.Enemies = append(g.Enemies, newEnemy)
		}
	case 18:
		for i := range 2 {
			newEnemy := Player{
				Health:   100,
				PlayerX:  float32(screenWidth - (50 * (i + 1))),
				PlayerY:  screenHeight - 50,
				Speed:    r1.Float32() * .5,
				Velocity: g.Player.Velocity,
				Cooldown: r1.Intn(50) + 50,
			}
			g.Enemies = append(g.Enemies, newEnemy)
			newEnemy = Player{
				Health:   100,
				PlayerX:  float32(screenWidth - (50 * (i + 1))),
				PlayerY:  100,
				Speed:    r1.Float32() * .5,
				Velocity: -g.Player.Velocity,
				Cooldown: r1.Intn(50) + 50,
			}
			g.Enemies = append(g.Enemies, newEnemy)
			newEnemy = Player{
				Health:   100,
				PlayerX:  float32(screenWidth - (50 * (i + 1))),
				PlayerY:  screenHeight - 100,
				Speed:    r1.Float32() * .5,
				Velocity: g.Player.Velocity,
				Cooldown: r1.Intn(50) + 50,
			}
			g.Enemies = append(g.Enemies, newEnemy)
			newEnemy = Player{
				Health:   100,
				PlayerX:  float32(screenWidth - (50 * (i + 1))),
				PlayerY:  50,
				Speed:    r1.Float32() * .5,
				Velocity: -g.Player.Velocity,
				Cooldown: r1.Intn(50) + 50,
			}
			g.Enemies = append(g.Enemies, newEnemy)
		}
	case 19:
		newEnemy := Player{
			Health:   100,
			PlayerX:  screenWidth - 50,
			PlayerY:  screenHeight - 50,
			Speed:    2,
			Velocity: -1.0,
			Cooldown: r1.Intn(50) + 20,
		}
		g.Enemies = append(g.Enemies, newEnemy)

		newEnemy = Player{
			Health:   100,
			PlayerX:  screenWidth - 55,
			PlayerY:  50,
			Speed:    2,
			Velocity: 1.0,
			Cooldown: r1.Intn(50) + 20,
		}
		g.Enemies = append(g.Enemies, newEnemy)

		newEnemy = Player{
			Health:   100,
			PlayerX:  screenWidth / 2,
			PlayerY:  screenHeight - 50,
			Speed:    2,
			Velocity: -1.0,
			Cooldown: r1.Intn(50) + 20,
		}
		g.Enemies = append(g.Enemies, newEnemy)

		newEnemy = Player{
			Health:   100,
			PlayerX:  screenWidth/2 + 5,
			PlayerY:  50,
			Speed:    2,
			Velocity: 1.0,
			Cooldown: r1.Intn(50) + 20,
		}
		g.Enemies = append(g.Enemies, newEnemy)

		newEnemy = Player{
			Health:   100,
			PlayerX:  110,
			PlayerY:  50,
			Speed:    1.5,
			Velocity: 1.0,
			Cooldown: r1.Intn(50) + 20,
		}
		g.Enemies = append(g.Enemies, newEnemy)

		newEnemy = Player{
			Health:   100,
			PlayerX:  90,
			PlayerY:  screenHeight - 50,
			Speed:    1.5,
			Velocity: -1.0,
			Cooldown: r1.Intn(50) + 20,
		}
		g.Enemies = append(g.Enemies, newEnemy)

		newEnemy = Player{
			Health:   100,
			PlayerX:  70,
			PlayerY:  50,
			Speed:    1.5,
			Velocity: 1.0,
			Cooldown: r1.Intn(50) + 20,
		}
		g.Enemies = append(g.Enemies, newEnemy)

		newEnemy = Player{
			Health:   100,
			PlayerX:  50,
			PlayerY:  screenHeight - 50,
			Speed:    1.5,
			Velocity: -1.0,
			Cooldown: r1.Intn(50) + 20,
		}
		g.Enemies = append(g.Enemies, newEnemy)
	case 20:
		newPowerUp := PowerUp{
			X:     0,
			Y:     screenHeight / 2,
			VX:    0.5,
			VY:    0,
			Power: "Super",
			Got:   false,
			Draw:  true,
		}
		g.Powerups["Super"] = newPowerUp

		newEnemy := Player{
			Health:   5000,
			PlayerX:  screenWidth - 50,
			PlayerY:  screenHeight / 2,
			Speed:    .25,
			Velocity: -1.0,
			Cooldown: r1.Intn(50) + 50,
		}
		g.Enemies = append(g.Enemies, newEnemy)

		newEnemy = Player{
			Health:   1,
			PlayerX:  screenWidth - 10,
			PlayerY:  screenHeight - 10,
			Speed:    0,
			Velocity: 0,
			Cooldown: r1.Intn(50) + 50,
		}
		g.Enemies = append(g.Enemies, newEnemy)

		newEnemy = Player{
			Health:   1,
			PlayerX:  screenWidth - 10,
			PlayerY:  screenHeight / 3,
			Speed:    0,
			Velocity: 0,
			Cooldown: r1.Intn(50) + 50,
		}
		g.Enemies = append(g.Enemies, newEnemy)

		newEnemy = Player{
			Health:   1,
			PlayerX:  screenWidth - 10,
			PlayerY:  10,
			Speed:    0,
			Velocity: 0,
			Cooldown: r1.Intn(50) + 50,
		}
		g.Enemies = append(g.Enemies, newEnemy)

		newEnemy = Player{
			Health:   1,
			PlayerX:  screenWidth - 10,
			PlayerY:  (screenHeight / 3) * 2,
			Speed:    0,
			Velocity: 0,
			Cooldown: r1.Intn(50) + 50,
		}
		g.Enemies = append(g.Enemies, newEnemy)
	}
	g.NextWave = false
}
