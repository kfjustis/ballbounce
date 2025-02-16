package main

import (
	"fmt"
	"image/color"
	"math/rand/v2"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type Ball struct {
	x, y, w, h, speed, xdir, ydir float32
	g                             *Game
}

func (b *Ball) Init() {
	var minx float32 = 0
	var miny float32 = 0
	var maxx float32 = 320
	var maxy float32 = 240
	b.x = minx + rand.Float32()*(maxx-minx)
	b.y = miny + rand.Float32()*(maxy-miny)
	b.w = 16.0
	b.h = 16.0
	b.speed = 1.0
	b.xdir = 1.0
	b.ydir = 1.0

	fmt.Printf("rand x: %f\n", b.x)
	fmt.Printf("rand y: %f\n", b.y)
}

func (b *Ball) Update() error {

	// Resolve directions
	if b.x+b.w > 320 || b.x < 0 {
		b.xdir *= -1.0
		b.g.blipPlayer.Rewind()
		b.g.blipPlayer.Play()
	}

	if b.y+b.h > 240 || b.y < 0 {
		b.ydir *= -1.0
		b.g.blipPlayer.Rewind()
		b.g.blipPlayer.Play()
	}

	b.x += b.speed * b.xdir
	b.y += b.speed * b.ydir

	return nil
}

func (b *Ball) Draw(screen *ebiten.Image) {
	vector.DrawFilledRect(screen, b.x, b.y, b.w, b.h, color.RGBA{0xFF, 0xFF, 0xFF, 0xFF}, false)
}
