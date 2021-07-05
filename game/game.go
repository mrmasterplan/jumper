package main

import (
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	gameconfig "github.com/mrmasterplan/jumper/game/config"


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
	return gameconfig.GameWindow.X, gameconfig.GameWindow.Y
}

func main() {
	ebiten.SetWindowSize(gameconfig.GameWindowWorldView.X, gameconfig.GameWindowWorldView.Y)
	ebiten.SetWindowTitle("my title")
	game := &Game{}
	if err := ebiten.RunGame(game); err != nil {
		panic(err)
	}
}
