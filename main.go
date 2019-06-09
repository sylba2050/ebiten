package main

import (
	"bytes"
	"image"
	"image/color"
	_ "image/png"
	"log"

	"github.com/hajimehoshi/ebiten"
	"github.com/sylba2050/ebiten/images"
)

const WIDTH = 640
const HEIGHT = 640

var field *ebiten.Image
var line_h *ebiten.Image
var line_w *ebiten.Image
var black_piece *ebiten.Image
var white_piece *ebiten.Image

func init() {
	var err error

	field, err = ebiten.NewImage(WIDTH, HEIGHT, ebiten.FilterDefault)
	if err != nil {
		log.Fatal(err)
	}

	line_h, err = ebiten.NewImage(2, HEIGHT, ebiten.FilterDefault)
	if err != nil {
		log.Fatal(err)
	}

	line_w, err = ebiten.NewImage(WIDTH, 2, ebiten.FilterDefault)
	if err != nil {
		log.Fatal(err)
	}

	field.Fill(color.RGBA{0x1A, 0xC9, 0x2E, 255})
	line_h.Fill(color.RGBA{0x00, 0x00, 0x00, 255})
	line_w.Fill(color.RGBA{0x00, 0x00, 0x00, 255})

	b_img, _, err := image.Decode(bytes.NewReader(images.Black_img))
	if err != nil {
		log.Fatal(err)
	}
	black_piece, _ = ebiten.NewImageFromImage(b_img, ebiten.FilterDefault)

	w_img, _, err := image.Decode(bytes.NewReader(images.White_img))
	if err != nil {
		log.Fatal(err)
	}
	white_piece, _ = ebiten.NewImageFromImage(w_img, ebiten.FilterDefault)
}

func update(screen *ebiten.Image) error {
	if ebiten.IsDrawingSkipped() {
		return nil
	}

	screen.DrawImage(field, &ebiten.DrawImageOptions{})

	op := &ebiten.DrawImageOptions{}
	for i := 0; i < 7; i++ {
		op.GeoM.Translate(float64(WIDTH/8*(i+1)), 0)
		screen.DrawImage(line_h, op)
		op.GeoM.Reset()

		op.GeoM.Translate(0, float64(WIDTH/8*(i+1)))
		screen.DrawImage(line_w, op)
		op.GeoM.Reset()
	}

	op.GeoM.Translate(0, 0)
	screen.DrawImage(white_piece, op)
	op.GeoM.Reset()

	op.GeoM.Translate(80, 0)
	screen.DrawImage(black_piece, op)
	op.GeoM.Reset()

	return nil
}

func main() {
	if err := ebiten.Run(update, WIDTH, HEIGHT, 1, "Fill"); err != nil {
		log.Fatal(err)
	}
}
