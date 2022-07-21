package main

import (
	"fmt"

	_ "embed"

	"github.com/avitar64/Doodle-jump/engine"
	pix "github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
)

const (
	windowWidth  = 576
	windowHeight = 812
)

func main() {
	pixelgl.Run(run)
}

//go:embed assets/background.png
var backgroundPicBytes []byte

func newBackground() *pix.Sprite {
	pic, err := engine.ImageToPictureData(backgroundPicBytes)
	if err != nil {
		panic(fmt.Errorf("error loading entity image: %v", err))
	}

	return pix.NewSprite(pic, pic.Bounds())
}

//go:embed assets/icon.png
var iconPicBytes []byte

func run() {
	icon, err := engine.ImageToPictureData(iconPicBytes)
	if err != nil {
		panic(fmt.Errorf("error loading icon image: %v", err))
	}

	game := engine.Initialize(
		pixelgl.WindowConfig{
			Title:  "Doodle jump",
			Icon:   []pix.Picture{icon},
			Bounds: pix.R(0, 0, windowWidth, windowHeight),
		},
		60,
		newBackground(),
	)

	game.CreateState(1)
	game.AddComponentsToState(1, newPlayer())

	game.ChangeState(1)
	game.Run()
}
