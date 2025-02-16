package main

import (
	"bytes"
	_ "embed"
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/hajimehoshi/ebiten/v2/audio/vorbis"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

//go:embed blip.ogg
var blip_ogg []byte

type Game struct {
	ball Ball

	audioContext *audio.Context
	blipPlayer   *audio.Player
}

func Init() *Game {
	g := &Game{}

	g.ball.Init()
	if g.ball.g == nil {
		g.ball.g = g
	}

	// Set up the audio stuff.
	if g.audioContext == nil {
		g.audioContext = audio.NewContext(44100)
	}
	blipD, err := vorbis.DecodeF32(bytes.NewReader(blip_ogg))
	if err != nil {
		log.Fatal(err)
	}
	g.blipPlayer, err = g.audioContext.NewPlayerF32(blipD)
	if err != nil {
		log.Fatal(err)
	}

	return g
}

func (g *Game) Update() error {
	g.ball.Update()
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{0x36, 0x45, 0x4F, 0xFF})
	g.ball.Draw(screen)
	ebitenutil.DebugPrint(screen, "Go and Ebitengine are pretty cool!")
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 320, 240
}

func main() {
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)
	ebiten.SetWindowTitle("Ball Bounce v0.1 | kfjustis")

	if err := ebiten.RunGame(Init()); err != nil {
		log.Fatal(err)
	}
}
