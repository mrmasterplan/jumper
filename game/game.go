package main

import (
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/mrmasterplan/jumper/src/config"


)

type Game struct{}

func (g *Game) Update() error {
	// Update the logical state
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	// Render the screen

	x, y := ebiten.CursorPosition()
	ebitenutil.DebugPrint(screen, fmt.Sprintf("%[2]d %[1]d\n", x, y))
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	// Return the game screen size
	return config.GameWindow.X, config.GameWindow.Y
}

func main() {
	ebiten.SetWindowSize(config.GameWindowWorldView.X, config.GameWindowWorldView.Y)
	ebiten.SetWindowTitle("my title")
	game := &Game{}
	if err := ebiten.RunGame(game); err != nil {
		panic(err)
	}
}
