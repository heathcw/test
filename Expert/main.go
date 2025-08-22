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
	Win         bool
	Timer       Timer
	Stars       []Projectile
	Powerups    map[string]PowerUp
}

func (g *Game) Update() error {
	if g.Player.Health <= 0 || g.Timer.Time == 0 {
		g.Lose = true
		return nil
	}
	if g.Win {
		return nil
	}

	if g.Player.Cooldown > 0 {
		g.Player.Cooldown--
	}

	if g.Player.Hurt {
		g.Player.HurtCooldown--
	}
	if g.Player.HurtCooldown == 0 {
		g.Player.Hurt = false
		g.Player.HurtCooldown = 30
	}

	if g.Timer.Cooldown > 0 {
		g.Timer.Cooldown--
	} else {
		g.Timer.Time--
		g.Timer.Cooldown = 60
	}

	handleInputs(g)

	//update enemies
	var newEnemies []Player
	for _, enemy := range g.Enemies {

		//calculate movement
		enemyMovement(g, &enemy)

		//shoot projectile
		shootProjectiles(g, &enemy)

		//calculate damage to enemies
		newEnemies = calculateEnemyDamage(g, &enemy, newEnemies)
	}
	g.Enemies = newEnemies

	// Update projectiles
	updateProjectiles(g)

	//update powerups
	for i, u := range g.Powerups {
		newPowerUp := u
		if g.Player.Hurt {
			newPowerUp.Got = false
		}
		if isColliding(u.X, u.Y, 4, 4, g.Player.PlayerX, g.Player.PlayerY, 8, 8) && !u.Got && u.Draw {
			newPowerUp.Got = true
			newPowerUp.Draw = false
		} else {
			newPowerUp.X += u.VX
			newPowerUp.Y += u.VY
		}
		g.Powerups[i] = newPowerUp
	}

	//update wave
	if len(g.Enemies) <= 0 {
		g.Wave += 1
		g.NextWave = true
	}
	if g.NextWave {
		wave(g)
	}

	//update stars
	var newStars []Projectile
	for i := range g.Stars {
		g.Stars[i].X += g.Stars[i].VX
		if g.Stars[i].X > 0 {
			newStars = append(newStars, g.Stars[i])
		} else {
			newStars = append(newStars, Projectile{
				X:  screenWidth,
				Y:  g.Stars[i].Y,
				VX: g.Stars[i].VX,
			})
		}
	}
	g.Stars = newStars

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	// Clear screen
	screen.Fill(color.RGBA{0, 0, 0, 255})
	for _, star := range g.Stars {
		alpha := uint8(32 + r1.Intn(64))
		vector.DrawFilledRect(screen, star.X, star.Y, 1, 1, color.RGBA{255, 255, 255, alpha}, false)

	}

	// Draw player as a white rectangle for now
	if g.Player.Hurt {
		vector.DrawFilledRect(screen, float32(g.Player.PlayerX), float32(g.Player.PlayerY), 8, 8, color.RGBA{231, 193, 193, 255}, false)
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
	} else if g.Wave == 12 || g.Wave == 16 {
		vector.DrawFilledRect(screen, g.Enemies[0].PlayerX, g.Enemies[0].PlayerY, 7, 7, color.RGBA{0, 0, 255, 255}, false)
		vector.DrawFilledRect(screen, 10, screenHeight-10, float32(g.Enemies[0].Health/2), 3, color.RGBA{255, 0, 0, 255}, false)
	} else if g.Wave == 20 {
		vector.DrawFilledRect(screen, g.Enemies[0].PlayerX, g.Enemies[0].PlayerY, 32, 32, color.RGBA{0, 0, 255, 255}, false)
		vector.DrawFilledRect(screen, 10, screenHeight-10, float32(g.Enemies[0].Health/16), 3, color.RGBA{255, 0, 0, 255}, false)
	} else {
		for _, e := range g.Enemies {
			if g.Wave == 17 || g.Wave == 18 {
				if e.Hurt {
					vector.DrawFilledRect(screen, e.PlayerX, e.PlayerY, 16, 16, color.RGBA{0, 0, 113, 255}, false)
				} else {
					vector.DrawFilledRect(screen, e.PlayerX, e.PlayerY, 16, 16, color.RGBA{0, 0, 255, 255}, false)
				}
			} else {
				if e.Hurt {
					vector.DrawFilledRect(screen, e.PlayerX, e.PlayerY, 7, 7, color.RGBA{0, 0, 113, 255}, false)
				} else {
					vector.DrawFilledRect(screen, e.PlayerX, e.PlayerY, 7, 7, color.RGBA{0, 0, 255, 255}, false)
				}
			}
		}
	}

	//draw text
	message := "SCORE: " + strconv.Itoa(g.Score)
	text.Draw(screen, message, basicfont.Face7x13, screenWidth-100, 10, color.White)
	message = "WAVE: " + strconv.Itoa(g.Wave)
	text.Draw(screen, message, basicfont.Face7x13, screenWidth-75, screenHeight-10, color.White)
	message = "TIME: " + strconv.Itoa(g.Timer.Time)
	text.Draw(screen, message, basicfont.Face7x13, screenWidth-175, 10, color.White)
	if g.Lose {
		message = "GAME OVER"
		text.Draw(screen, message, basicfont.Face7x13, screenWidth/2-25, screenHeight/2, color.White)
		return
	}
	if g.Win {
		message = "YOU WIN"
		text.Draw(screen, message, basicfont.Face7x13, screenWidth/2-25, screenHeight/2, color.White)
		return
	}

	//draw projectiles
	for _, p := range g.Projectiles {
		if p.VX > 0 {
			if g.Powerups["Big"].Got || g.Powerups["Super"].Got {
				vector.DrawFilledCircle(screen, p.X, p.Y, 6, color.RGBA{0, 255, 0, 255}, false) // big green blast
			} else {
				vector.DrawFilledCircle(screen, p.X, p.Y, 3, color.RGBA{0, 255, 0, 255}, false) // green bullet
			}
		} else if g.Wave == 20 && p.VX == -1 {
			vector.DrawFilledCircle(screen, p.X, p.Y, 10, color.RGBA{255, 0, 0, 255}, false) // ultra enemy blast
		} else if p.VX == -1 || (g.Wave == 16 && p.VX < 0) || (g.Wave == 20 && p.VX < 0) {
			vector.DrawFilledCircle(screen, p.X, p.Y, 5, color.RGBA{255, 0, 0, 255}, false) // big enemy blast
		} else {
			vector.DrawFilledCircle(screen, p.X, p.Y, 2, color.RGBA{255, 0, 0, 255}, false) // enemy pellet
		}
	}

	//draw powerups
	for _, u := range g.Powerups {
		if u.Draw {
			vector.DrawFilledCircle(screen, u.X, u.Y, 3, color.RGBA{245, 40, 145, 255}, false)
		}
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func main() {
	game := &Game{
		Player: Player{
			Health:       30,
			Hurt:         false,
			PlayerX:      screenWidth / 2,
			PlayerY:      screenHeight / 2,
			Velocity:     1.0,
			Cooldown:     15,
			HurtCooldown: 30,
		},
		Enemies:  []Player{},
		Wave:     1,
		NextWave: false,
		Lose:     false,
		Win:      false,
		Timer:    Timer{Time: 600, Cooldown: 60},
		Powerups: make(map[string]PowerUp),
	}
	wave(game)

	for range 99 {
		newStar := Projectile{
			X:  float32(r1.Intn(screenWidth)),
			Y:  float32(r1.Intn(screenHeight)),
			VX: -.2,
		}
		game.Stars = append(game.Stars, newStar)
	}

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

func handleInputs(g *Game) {
	var move float32 = playerSpeed
	if g.Powerups["Speed"].Got || g.Powerups["Super"].Got {
		move = playerSpeed * 2.0
	}
	// Input handling
	if ebiten.IsKeyPressed(ebiten.KeyD) && g.Player.PlayerX < screenWidth-playerSpeed {
		g.Player.updateX(move)
	}
	if ebiten.IsKeyPressed(ebiten.KeyA) && g.Player.PlayerX > playerSpeed {
		g.Player.updateX(-move)
	}
	if ebiten.IsKeyPressed(ebiten.KeyS) && g.Player.PlayerY < screenHeight-playerSpeed {
		g.Player.updateY(move)
		g.Player.Velocity = 1.0
	}
	if ebiten.IsKeyPressed(ebiten.KeyW) && g.Player.PlayerY > playerSpeed {
		g.Player.updateY(-move)
		g.Player.Velocity = -1.0
	}
	// Shoot (spacebar)
	if ebiten.IsKeyPressed(ebiten.KeySpace) && g.Player.Cooldown == 0 && !g.Player.Hurt {
		// Add a projectile moving to the right
		g.Projectiles = append(g.Projectiles, Projectile{
			X:  g.Player.PlayerX + 8, // center of player
			Y:  g.Player.PlayerY + 5,
			VX: weaponSpeed, // pixels per frame
			VY: 0,
		})

		if g.Powerups["Spread"].Got || g.Powerups["Super"].Got {
			g.Projectiles = append(g.Projectiles, Projectile{
				X:  g.Player.PlayerX + 8, // center of player
				Y:  g.Player.PlayerY - 10,
				VX: weaponSpeed, // pixels per frame
				VY: 0,
			})
			g.Projectiles = append(g.Projectiles, Projectile{
				X:  g.Player.PlayerX + 8, // center of player
				Y:  g.Player.PlayerY + 20,
				VX: weaponSpeed, // pixels per frame
				VY: 0,
			})
		}

		if g.Powerups["Blast"].Got || g.Powerups["Super"].Got {
			g.Player.Cooldown = 5
		} else {
			g.Player.Cooldown = 15
		}
	}
}

func enemyMovement(g *Game, enemy *Player) {
	if g.Wave == 1 || g.Wave == 5 {
		enemy.updateX(-enemy.Speed)
	} else if g.Wave == 2 || g.Wave == 6 || g.Wave == 13 || g.Wave == 14 || g.Wave == 15 || g.Wave == 17 || g.Wave == 19 {
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
	} else if g.Wave == 8 || g.Wave == 9 || g.Wave == 10 || g.Wave == 11 || g.Wave == 12 || g.Wave == 18 {
		if enemy.PlayerY < g.Player.PlayerY {
			enemy.Velocity = 1.0
			enemy.updateY(enemy.Velocity * enemy.Speed)
		} else if enemy.PlayerY > g.Player.PlayerY {
			enemy.Velocity = -1.0
			enemy.updateY(enemy.Velocity * enemy.Speed)
		}
	} else if g.Wave == 16 {
		if enemy.PlayerY < g.Player.PlayerY {
			enemy.Velocity = 1.0
			enemy.updateY(enemy.Velocity * enemy.Speed)
		} else if enemy.PlayerY > g.Player.PlayerY {
			enemy.Velocity = -1.0
			enemy.updateY(enemy.Velocity * enemy.Speed)
		}
		if enemy.PlayerX < g.Player.PlayerX+50 {
			enemy.Velocity = 1.0
			enemy.updateX(enemy.Velocity * enemy.Speed)
		} else if enemy.PlayerX > g.Player.PlayerX+50 {
			enemy.Velocity = -1.0
			enemy.updateX(enemy.Velocity * enemy.Speed)
		}
	} else if g.Wave == 20 {
		if enemy.PlayerY <= 100 {
			enemy.Velocity = 1.0
		}
		if enemy.PlayerY >= screenHeight-100 {
			enemy.Velocity = -1.0
		}
		enemy.updateY(enemy.Velocity * enemy.Speed)
	}
}

func shootProjectiles(g *Game, enemy *Player) {
	if enemy.Cooldown == 0 && g.Wave < 4 {
		g.Projectiles = append(g.Projectiles, Projectile{
			X:  enemy.PlayerX - 8, // center of player
			Y:  enemy.PlayerY + 8,
			VX: -weaponSpeed, // pixels per frame
			VY: 0,
		})
		enemy.Cooldown = r1.Intn(300) + 100
	} else if enemy.Cooldown == 0 && g.Wave < 8 {
		g.Projectiles = append(g.Projectiles, Projectile{
			X:  enemy.PlayerX - 8, // center of player
			Y:  enemy.PlayerY + 8,
			VX: -1, // pixels per frame
			VY: 0,
		})
		enemy.Cooldown = r1.Intn(250) + 250
	} else if enemy.Cooldown == 0 && g.Wave == 12 {
		g.Projectiles = append(g.Projectiles, Projectile{
			X:  enemy.PlayerX - 8, // center of player
			Y:  enemy.PlayerY + 4,
			VX: -weaponSpeed, // pixels per frame
			VY: 0,
		})
		enemy.Cooldown = r1.Intn(50) + 10
	} else if enemy.Cooldown == 0 && g.Wave > 12 && g.Wave < 16 {
		if enemy.Speed == 2 {
			g.Projectiles = append(g.Projectiles, Projectile{
				X:  enemy.PlayerX - 8, // center of player
				Y:  enemy.PlayerY + 4,
				VX: -weaponSpeed, // pixels per frame
				VY: 0,
			})
		} else {
			g.Projectiles = append(g.Projectiles, Projectile{
				X:  enemy.PlayerX + 4, // center of player
				Y:  enemy.PlayerY,
				VX: 0,
				VY: enemy.Velocity * weaponSpeed * 1.5,
			})
		}
		enemy.Cooldown = r1.Intn(50) + 10
	} else if enemy.Cooldown == 0 && g.Wave == 16 {
		g.Projectiles = append(g.Projectiles, Projectile{
			X:  enemy.PlayerX - 8, // center of player
			Y:  enemy.PlayerY + 8,
			VX: -1.5, // pixels per frame
			VY: 0,
		})
		enemy.Cooldown = 30
	} else if enemy.Cooldown == 0 && g.Wave == 19 {
		if enemy.Speed == 2 {
			g.Projectiles = append(g.Projectiles, Projectile{
				X:  enemy.PlayerX - 8, // center of player
				Y:  enemy.PlayerY + 4,
				VX: -1, // pixels per frame
				VY: .25 * g.Player.Velocity,
			})
		} else {
			g.Projectiles = append(g.Projectiles, Projectile{
				X:  enemy.PlayerX + 4, // center of player
				Y:  enemy.PlayerY,
				VX: 0,
				VY: enemy.Velocity * weaponSpeed * 1.5,
			})
		}
		enemy.Cooldown = r1.Intn(50) + 10
	} else if enemy.Cooldown == 0 && g.Wave == 20 {
		if enemy.Speed == 0 {
			g.Projectiles = append(g.Projectiles, Projectile{
				X:  enemy.PlayerX + 1, // center of player
				Y:  enemy.PlayerY + 1,
				VX: -2 * r1.Float32(), // pixels per frame
				VY: .25 * g.Player.Velocity,
			})
			enemy.Cooldown = r1.Intn(50) + 25
		} else {
			g.Projectiles = append(g.Projectiles, Projectile{
				X:  enemy.PlayerX - 8, // center of player
				Y:  enemy.PlayerY + 15,
				VX: -1, // pixels per frame
				VY: 0,
			})
			enemy.Cooldown = r1.Intn(50) + 100
		}
	} else if enemy.Cooldown == 0 {
		g.Projectiles = append(g.Projectiles, Projectile{
			X:  enemy.PlayerX - 8, // center of player
			Y:  enemy.PlayerY + 8,
			VX: -1, // pixels per frame
			VY: 0,
		})
		enemy.Cooldown = r1.Intn(50) + 10
	} else {
		enemy.Cooldown--
	}
}

func calculateEnemyDamage(g *Game, enemy *Player, newEnemies []Player) []Player {
	//enemy hurt cooldown
	if enemy.Hurt {
		enemy.HurtCooldown--
	}
	if enemy.HurtCooldown == 0 {
		enemy.Hurt = false
		enemy.HurtCooldown = 30
	}

	//calculate damage to enemies
	for j, p := range g.Projectiles {
		if g.Wave == 20 {
			if g.Powerups["Big"].Got || g.Powerups["Super"].Got {
				if isColliding(p.X, p.Y, 7, 7, enemy.PlayerX, enemy.PlayerY, 30, 30) && p.VX > 0 && enemy.Speed != 0 {
					enemy.Health -= 10
					enemy.Hurt = true
					enemy.HurtCooldown = 30
					g.Projectiles[j] = Projectile{}
					g.Score += 100
				}
			} else if isColliding(p.X, p.Y, 4, 4, enemy.PlayerX, enemy.PlayerY, 30, 30) && p.VX > 0 && enemy.Speed != 0 {
				enemy.Health -= 5
				enemy.Hurt = true
				enemy.HurtCooldown = 30
				g.Projectiles[j] = Projectile{}
				g.Score += 50
			}
			if enemy.Health <= 0 {
				g.Win = true
			}
		} else if g.Wave == 4 || g.Wave == 8 || g.Wave == 17 || g.Wave == 18 {
			if g.Powerups["Big"].Got || g.Powerups["Super"].Got {
				if isColliding(p.X, p.Y, 7, 7, enemy.PlayerX, enemy.PlayerY, 10, 10) && p.VX > 0 && enemy.Speed != 0 {
					enemy.Health -= 10
					enemy.Hurt = true
					enemy.HurtCooldown = 30
					g.Projectiles[j] = Projectile{}
					g.Score += 100
				}
			} else if isColliding(p.X, p.Y, 4, 4, enemy.PlayerX, enemy.PlayerY, 10, 10) && p.VX > 0 && enemy.Speed != 0 {
				enemy.Health -= 5
				enemy.Hurt = true
				enemy.HurtCooldown = 30
				g.Projectiles[j] = Projectile{}
				g.Score += 50
			}
		} else {
			if g.Powerups["Big"].Got || g.Powerups["Super"].Got {
				if isColliding(p.X, p.Y, 7, 7, enemy.PlayerX, enemy.PlayerY, 8, 8) && p.VX > 0 && enemy.Speed != 0 {
					enemy.Health -= 10
					enemy.Hurt = true
					enemy.HurtCooldown = 30
					g.Projectiles[j] = Projectile{}
					g.Score += 100
				}
			} else if isColliding(p.X, p.Y, 4, 4, enemy.PlayerX, enemy.PlayerY, 8, 8) && p.VX > 0 && enemy.Speed != 0 {
				enemy.Health -= 5
				enemy.Hurt = true
				enemy.HurtCooldown = 30
				g.Projectiles[j] = Projectile{}
				g.Score += 50
			}
		}

	}
	if enemy.Health > 0 && enemy.PlayerX > 0 {
		newEnemies = append(newEnemies, *enemy)
	}
	if enemy.PlayerX < 0 {
		g.Player.Health -= 2
	}

	return newEnemies
}

func updateProjectiles(g *Game) {
	for i, p := range g.Projectiles {
		if isColliding(p.X, p.Y, 3, 3, g.Player.PlayerX, g.Player.PlayerY, 8, 8) && p.VX < 0 && g.Wave < 4 && !g.Player.Hurt {
			g.Player.Health -= 5
			g.Player.Hurt = true
			g.Projectiles[i] = Projectile{}
		} else if isColliding(p.X, p.Y, 4, 4, g.Player.PlayerX, g.Player.PlayerY, 8, 8) && p.VX < 0 && g.Wave >= 4 && !g.Player.Hurt {
			if g.Powerups["Super"].Got {
				g.Player.Health -= 3
			} else {
				g.Player.Health -= 7
			}
			g.Player.Hurt = true
			g.Projectiles[i] = Projectile{}
		} else if isColliding(p.X, p.Y, 3, 3, g.Player.PlayerX, g.Player.PlayerY, 8, 8) && p.VX == 0 && g.Wave >= 4 && !g.Player.Hurt {
			g.Player.Health -= 5
			g.Player.Hurt = true
			g.Projectiles[i] = Projectile{}
		} else if isColliding(p.X-5, p.Y-7, 14, 14, g.Player.PlayerX, g.Player.PlayerY, 8, 8) && p.VX == -1 && g.Wave == 20 && !g.Player.Hurt {
			if g.Powerups["Super"].Got {
				g.Player.Health -= 8
			} else {
				g.Player.Health -= 15
			}
			g.Player.Hurt = true
			g.Projectiles[i] = Projectile{}
		} else {
			g.Projectiles[i].X += g.Projectiles[i].VX
			g.Projectiles[i].Y += g.Projectiles[i].VY
		}
	}

	// Remove off-screen projectiles
	var newProjectiles []Projectile
	for _, p := range g.Projectiles {
		if p.X >= 0 && p.X <= screenWidth && p.Y >= 0 && p.Y <= screenHeight /*&& p.VX != 0*/ {
			newProjectiles = append(newProjectiles, p)
		}
	}
	g.Projectiles = newProjectiles
}
