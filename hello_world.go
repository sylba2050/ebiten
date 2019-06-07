package main

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

func update(screen *ebiten.Image) error {
	if err := ebitenutil.DebugPrint(screen, "Hello world!"); err != nil {
		return err
	}
	return nil
}

func main() {
	ebiten.Run(update, 320, 240, 2, "Hello world!")
}
