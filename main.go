package main

import (
	"bytes"
	"image"
	"image/color"
	_ "image/png"
	"log"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/inpututil"
	"github.com/sylba2050/ebiten/images"
)

const WIDTH = 640
const HEIGHT = 640

var back_ground *ebiten.Image
var line_h *ebiten.Image
var line_w *ebiten.Image
var black_piece *ebiten.Image
var white_piece *ebiten.Image
var turn int

const (
	BG    = 0
	BLACK = 1
	WHITE = 2
)

type Direction struct {
	x int
	y int
}

var up = Direction{x: 0, y: -1}
var down = Direction{x: 0, y: 1}
var left = Direction{x: -1, y: 0}
var right = Direction{x: 1, y: 0}
var upright = Direction{x: 1, y: -1}
var upleft = Direction{x: -1, y: -1}
var downright = Direction{x: 1, y: 1}
var downleft = Direction{x: -1, y: 1}

// 0: none, 1: black, 2: white
var status [8][8]int = [8][8]int{
	{0, 0, 0, 0, 0, 0, 0, 0},
	{0, 0, 0, 0, 0, 0, 0, 0},
	{0, 0, 0, 0, 0, 0, 0, 0},
	{0, 0, 0, BLACK, WHITE, 0, 0, 0},
	{0, 0, 0, WHITE, BLACK, 0, 0, 0},
	{0, 0, 0, 0, 0, 0, 0, 0},
	{0, 0, 0, 0, 0, 0, 0, 0},
	{0, 0, 0, 0, 0, 0, 0, 0},
}

func IsCanPutWithDirection(s [8][8]int, x, y, turn int, d Direction) bool {
	x += d.x
	y += d.y

	if x < 0 || y < 0 || x > 7 || y > 7 {
		return false
	}

	if s[y][x] == BG {
		return false
	}

	if s[y][x] == turn {
		return false
	}

	for {
		x += d.x
		y += d.y

		if x < 0 || y < 0 || x > 7 || y > 7 {
			break
		}

		if s[y][x] == BG {
			return false
		}

		if s[y][x] == turn {
			return true
		}
	}

	return false
}

func IsCanPut(s [8][8]int, x, y, turn int) bool {
	if IsCanPutWithDirection(s, x, y, turn, up) {
		return true
	}
	if IsCanPutWithDirection(s, x, y, turn, down) {
		return true
	}
	if IsCanPutWithDirection(s, x, y, turn, left) {
		return true
	}
	if IsCanPutWithDirection(s, x, y, turn, right) {
		return true
	}
	if IsCanPutWithDirection(s, x, y, turn, upleft) {
		return true
	}
	if IsCanPutWithDirection(s, x, y, turn, upright) {
		return true
	}
	if IsCanPutWithDirection(s, x, y, turn, downleft) {
		return true
	}
	if IsCanPutWithDirection(s, x, y, turn, downright) {
		return true
	}

	return false
}

func reverseWithDirection(s *[8][8]int, x, y, turn int, d Direction) {
	for {
		x += d.x
		y += d.y

		if x < 0 || y < 0 || x > 7 || y > 7 {
			break
		}

		if s[y][x] == turn {
			break
		}

		s[y][x] = turn
	}

	return
}

func reverse(s *[8][8]int, x, y, turn int) {
	if IsCanPutWithDirection(*s, x, y, turn, up) {
		reverseWithDirection(s, x, y, turn, up)
	}
	if IsCanPutWithDirection(*s, x, y, turn, down) {
		reverseWithDirection(s, x, y, turn, down)
	}
	if IsCanPutWithDirection(*s, x, y, turn, left) {
		reverseWithDirection(s, x, y, turn, left)
	}
	if IsCanPutWithDirection(*s, x, y, turn, right) {
		reverseWithDirection(s, x, y, turn, right)
	}
	if IsCanPutWithDirection(*s, x, y, turn, upleft) {
		reverseWithDirection(s, x, y, turn, upleft)
	}
	if IsCanPutWithDirection(*s, x, y, turn, upright) {
		reverseWithDirection(s, x, y, turn, upright)
	}
	if IsCanPutWithDirection(*s, x, y, turn, downleft) {
		reverseWithDirection(s, x, y, turn, downleft)
	}
	if IsCanPutWithDirection(*s, x, y, turn, downright) {
		reverseWithDirection(s, x, y, turn, downright)
	}
}

func init() {
	var err error

	back_ground, err = ebiten.NewImage(WIDTH, HEIGHT, ebiten.FilterDefault)
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

	back_ground.Fill(color.RGBA{0x1A, 0xC9, 0x2E, 255})
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

	turn = WHITE
}

func update(screen *ebiten.Image) error {
	is_click := inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft)
	mouse_x, mouse_y := ebiten.CursorPosition()

	if ebiten.IsDrawingSkipped() {
		return nil
	}

	screen.DrawImage(back_ground, &ebiten.DrawImageOptions{})

	op := &ebiten.DrawImageOptions{}
	for i := 0; i < 7; i++ {
		op.GeoM.Translate(float64(WIDTH/8*(i+1)), 0)
		screen.DrawImage(line_h, op)
		op.GeoM.Reset()

		op.GeoM.Translate(0, float64(WIDTH/8*(i+1)))
		screen.DrawImage(line_w, op)
		op.GeoM.Reset()
	}

	for j := 0; j < 8; j++ {
		for k := 0; k < 8; k++ {
			if status[j][k] == 0 {
				continue
			}
			if status[j][k] == 1 {
				op.GeoM.Translate(float64(k*80), float64(j*80))
				screen.DrawImage(black_piece, op)
				op.GeoM.Reset()
			}
			if status[j][k] == 2 {
				op.GeoM.Translate(float64(k*80), float64(j*80))
				screen.DrawImage(white_piece, op)
				op.GeoM.Reset()
			}
		}
	}

	if is_click && mouse_x >= 0 && mouse_x <= WIDTH && mouse_y >= 0 && mouse_y <= HEIGHT {
		if IsCanPut(status, mouse_x/80, mouse_y/80, turn) {

			status[mouse_y/80][mouse_x/80] = turn
			if turn == WHITE {
				turn = BLACK
			} else {
				turn = WHITE
			}
			reverse(&status, mouse_x/80, mouse_y/80, turn)
		}

	}

	return nil
}

func main() {
	if err := ebiten.Run(update, WIDTH, HEIGHT, 1, "Fill"); err != nil {
		log.Fatal(err)
	}
}
