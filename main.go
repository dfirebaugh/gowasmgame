package main

import (
	"fmt"
	"gowasmgame/atlas"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

const (
	screenWidth  = 240
	screenHeight = screenWidth
)

// Game is the main game struct
type Game struct {
	Atlas   *atlas.Atlas
	pressed []ebiten.Key
}

const movementInterval = .8

// Update runs before Draw
func (g *Game) Update() error {
	g.pressed = nil
	for k := ebiten.Key(0); k <= ebiten.KeyMax; k++ {
		if ebiten.IsKeyPressed(k) {
			g.pressed = append(g.pressed, k)
			if k == ebiten.KeyRight {
				horizontal += movementInterval
			}
			if k == ebiten.KeyLeft {
				horizontal -= movementInterval
			}
			if k == ebiten.KeyUp {
				vertical -= movementInterval
			}
			if k == ebiten.KeyDown {
				vertical += movementInterval
			}
		}
	}

	return nil
}

var horizontal float64 = 0
var vertical float64 = 0

// Draw runs after update and draws the tiles.
func (g *Game) Draw(screen *ebiten.Image) {
	g.Atlas.Draw(screen, screenWidth)
	g.Atlas.Spr(screen, screenWidth, 1, horizontal, vertical)

	ebitenutil.DebugPrint(screen, fmt.Sprintf("TPS: %0.2f", ebiten.CurrentTPS()))
}

// Layout sets the screen width and screen Height.
func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func main() {
	atlas := atlas.New("./resources/images/tiles1.png")

	g := &Game{
		Atlas: atlas,
	}

	ebiten.SetWindowSize(screenWidth*4, screenHeight*4)
	ebiten.SetWindowResizable(true)
	ebiten.SetWindowTitle("Tiles (Ebiten Demo)")
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
