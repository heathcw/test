package main

import (
	"image/color"
	"log"
	"math/rand"
	"strconv"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"

	//"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"golang.org/x/image/font/basicfont"
)

const (
	screenWidth  = 320
	screenHeight = 240
	playerSpeed  = 1
	weaponSpeed  = 2
)

var s1 = rand.NewSource(time.Now().UnixNano())
var r1 = rand.New(s1)

type Game struct {
	Player      Player
	Projectiles []Projectile
	Enemies     []Player
	Wave        int
	NextWave    bool
	Score       int
	Lose        bool
}

func (g *Game) Update() error {
	if g.Player.Health <= 0 {
		g.Lose = true
		return nil
	}

	if g.Player.Cooldown > 0 {
		g.Player.Cooldown--
	} else {
		g.Player.Hurt = false
	}
	// Input handling
	if ebiten.IsKeyPressed(ebiten.KeyD) {
		g.Player.updateX(playerSpeed)
	}
	if ebiten.IsKeyPressed(ebiten.KeyA) {
		g.Player.updateX(-playerSpeed)
	}
	if ebiten.IsKeyPressed(ebiten.KeyS) {
		g.Player.updateY(playerSpeed)
		g.Player.Velocity = 1.0
	}
	if ebiten.IsKeyPressed(ebiten.KeyW) {
		g.Player.updateY(-playerSpeed)
		g.Player.Velocity = -1.0
	}
	// Shoot (spacebar)
	if ebiten.IsKeyPressed(ebiten.KeySpace) && g.Player.Cooldown == 0 {
		// Add a projectile moving to the right
		g.Projectiles = append(g.Projectiles, Projectile{
			X:  g.Player.PlayerX + 8, // center of player
			Y:  g.Player.PlayerY + 8,
			VX: weaponSpeed, // pixels per frame
			VY: 0,
		})
		g.Player.Cooldown = 15
	}

	//update enemies
	var newEnemies []Player
	for _, enemy := range g.Enemies {
		enemy.Hurt = false
		//calculate movement
		if g.Wave == 1 || g.Wave == 5 {
			enemy.updateX(-enemy.Speed)
		} else if g.Wave == 2 || g.Wave == 6 {
			if enemy.PlayerY <= 10 {
				enemy.Velocity = 1.0
			}
			if enemy.PlayerY >= screenHeight-10 {
				enemy.Velocity = -1.0
			}
			enemy.updateY(enemy.Velocity * enemy.Speed)
		} else if g.Wave == 3 || g.Wave == 7 {
			if enemy.Velocity == 0 {
				enemy.updateX(-enemy.Speed)
			} else if enemy.Speed == 1 {
				if enemy.PlayerY != g.Player.PlayerY {
					enemy.Velocity = g.Player.Velocity
					enemy.updateY(enemy.Velocity * enemy.Speed)
				}
			} else {
				if enemy.PlayerY <= 10 {
					enemy.Velocity = 1.0
				}
				if enemy.PlayerY >= screenHeight-10 {
					enemy.Velocity = -1.0
				}
				enemy.updateY(enemy.Velocity * enemy.Speed)
			}
		} else if g.Wave == 4 {
			var xVelocity float32
			if enemy.PlayerY <= 100 {
				enemy.Velocity = 1.0
			}
			if enemy.PlayerY >= screenHeight-100 {
				enemy.Velocity = -1.0
			}
			if enemy.PlayerX < 50 {
				xVelocity = 0
			} else if enemy.PlayerX <= screenWidth-60 {
				xVelocity = -.5
			} else if enemy.PlayerX > screenWidth-60 {
				xVelocity = -.5
			}
			enemy.updateY(enemy.Velocity * enemy.Speed)
			enemy.updateX(xVelocity * enemy.Speed)
		} else if g.Wave == 8 {
			if enemy.PlayerY < g.Player.PlayerY {
				enemy.Velocity = 1.0
				enemy.updateY(enemy.Velocity * enemy.Speed)
			} else if enemy.PlayerY > g.Player.PlayerY {
				enemy.Velocity = -1.0
				enemy.updateY(enemy.Velocity * enemy.Speed)
			}
		}

		//shoot projectile
		if enemy.Cooldown == 0 && g.Wave < 4 {
			g.Projectiles = append(g.Projectiles, Projectile{
				X:  enemy.PlayerX - 8, // center of player
				Y:  enemy.PlayerY + 8,
				VX: -weaponSpeed, // pixels per frame
				VY: 0,
			})
			enemy.Cooldown = rand.Intn(300) + 100
		} else if enemy.Cooldown == 0 && g.Wave < 8 {
			g.Projectiles = append(g.Projectiles, Projectile{
				X:  enemy.PlayerX - 8, // center of player
				Y:  enemy.PlayerY + 8,
				VX: -1, // pixels per frame
				VY: 0,
			})
			enemy.Cooldown = rand.Intn(250) + 250
		} else if enemy.Cooldown == 0 {
			g.Projectiles = append(g.Projectiles, Projectile{
				X:  enemy.PlayerX - 8, // center of player
				Y:  enemy.PlayerY + 8,
				VX: -1, // pixels per frame
				VY: 0,
			})
			enemy.Cooldown = rand.Intn(50) + 10
		} else {
			enemy.Cooldown--
		}

		//calculate damage to enemies
		for j, p := range g.Projectiles {
			if isColliding(p.X, p.Y, 4, 4, enemy.PlayerX, enemy.PlayerY, 8, 8) && p.VX > 0 {
				enemy.Health -= 5
				enemy.Hurt = true
				g.Projectiles[j] = Projectile{}
				g.Score += 50
			}
		}
		if enemy.Health > 0 && enemy.PlayerX > 0 {
			newEnemies = append(newEnemies, enemy)
		}
	}
	g.Enemies = newEnemies

	// Update projectiles
	for i, p := range g.Projectiles {
		if isColliding(p.X, p.Y, 3, 3, g.Player.PlayerX, g.Player.PlayerY, 8, 8) && p.VX < 0 && g.Wave < 4 {
			g.Player.Health -= 5
			g.Player.Hurt = true
			g.Projectiles[i] = Projectile{}
		} else if isColliding(p.X, p.Y, 4, 4, g.Player.PlayerX, g.Player.PlayerY, 8, 8) && p.VX < 0 && g.Wave >= 4 {
			g.Player.Health -= 10
			g.Player.Hurt = true
			g.Projectiles[i] = Projectile{}
		} else {
			g.Projectiles[i].X += g.Projectiles[i].VX
			g.Projectiles[i].Y += g.Projectiles[i].VY
		}
	}

	// (Optional) Remove off-screen projectiles
	var newProjectiles []Projectile
	for _, p := range g.Projectiles {
		if p.X >= 0 && p.X <= screenWidth && p.Y >= 0 && p.Y <= screenHeight && p.VX != 0 {
			newProjectiles = append(newProjectiles, p)
		}
	}
	g.Projectiles = newProjectiles

	//update wave
	if len(g.Enemies) <= 0 {
		g.Wave += 1
		g.NextWave = true
	}

	if g.NextWave {
		wave(g)
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	// Clear screen
	screen.Fill(color.RGBA{0, 0, 0, 255})

	// Draw player as a white rectangle for now
	if g.Player.Hurt {
		vector.DrawFilledRect(screen, float32(g.Player.PlayerX), float32(g.Player.PlayerY), 8, 8, color.RGBA{255, 0, 0, 255}, false)
	} else {
		vector.DrawFilledRect(screen, float32(g.Player.PlayerX), float32(g.Player.PlayerY), 8, 8, color.White, false)
	}

	//draw health
	for i := range g.Player.Health / 5 {
		vector.DrawFilledRect(screen, float32(((i + 1) * 5)), 10, 3, 5, color.RGBA{255, 0, 0, 255}, false)
	}

	//draw enemies
	if g.Wave == 4 || g.Wave == 8 {
		vector.DrawFilledRect(screen, g.Enemies[0].PlayerX, g.Enemies[0].PlayerY, 16, 16, color.RGBA{0, 0, 255, 255}, false)
		vector.DrawFilledRect(screen, 10, screenHeight-10, float32(g.Enemies[0].Health), 3, color.RGBA{255, 0, 0, 255}, false)
	} else {
		for _, e := range g.Enemies {
			if e.Hurt {
				vector.DrawFilledRect(screen, e.PlayerX, e.PlayerY, 7, 7, color.RGBA{0, 50, 255, 255}, false)
			} else {
				vector.DrawFilledRect(screen, e.PlayerX, e.PlayerY, 7, 7, color.RGBA{0, 0, 255, 255}, false)
			}
		}
	}

	//draw text
	message := "SCORE: " + strconv.Itoa(g.Score)
	text.Draw(screen, message, basicfont.Face7x13, screenWidth-100, 10, color.White)
	message = "WAVE: " + strconv.Itoa(g.Wave)
	text.Draw(screen, message, basicfont.Face7x13, screenWidth-100, screenHeight-10, color.White)
	if g.Lose {
		message = "GAME OVER"
		text.Draw(screen, message, basicfont.Face7x13, screenWidth/2-25, screenHeight/2, color.White)
		return
	}

	//draw projectiles
	for _, p := range g.Projectiles {
		if p.VX > 0 {
			vector.DrawFilledCircle(screen, p.X, p.Y, 3, color.RGBA{0, 255, 0, 255}, false) // green bullet
		} else if p.VX == -1 {
			vector.DrawFilledCircle(screen, p.X, p.Y, 5, color.RGBA{255, 0, 0, 255}, false) // big enemy blast
		} else {
			vector.DrawFilledCircle(screen, p.X, p.Y, 1, color.RGBA{255, 0, 0, 255}, false) // enemy pellet
		}
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func main() {
	game := &Game{
		Player: Player{
			Health:   30,
			Hurt:     false,
			PlayerX:  screenWidth / 2,
			PlayerY:  screenHeight / 2,
			Velocity: 1.0,
			Cooldown: 15,
		},
		Enemies:  []Player{},
		Wave:     8,
		NextWave: false,
		Lose:     false,
	}
	wave(game)

	ebiten.SetWindowSize(screenWidth*2, screenHeight*2) // scale up
	ebiten.SetWindowTitle("2D Pixel Game")

	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}

func isColliding(x1, y1, w1, h1, x2, y2, w2, h2 float32) bool {
	return x1 < x2+w2 &&
		x1+w1 > x2 &&
		y1 < y2+h2 &&
		y1+h1 > y2
}

func wave(g *Game) {
	switch g.Wave {
	case 1:
		for i := range 10 {
			newEnemy := Player{
				Health:   10,
				PlayerX:  screenWidth - 10,
				PlayerY:  float32(20 + i*((screenHeight-40)/9)),
				Speed:    (r1.Float32() / 2),
				Cooldown: rand.Intn(300) + 100,
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
				Velocity: float32(1 - 2*(rand.Intn(2))),
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
