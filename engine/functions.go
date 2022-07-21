package engine

import (
	"bytes"
	"image"
	_ "image/png"

	pix "github.com/faiface/pixel"
)

func Touching(a, b Component) bool {
	return a.Rect().Intersects(b.Rect())
}

func TouchingEdge(a Component, g *Game) bool {
	rect := a.Rect()
	winRect := g.WinConf.Bounds
	windowWidth, windowHeight := winRect.W(), winRect.H()

	return rect.Min.X <= 0 || rect.Max.X >= windowWidth || rect.Min.Y <= 0 || rect.Max.Y >= windowHeight
}

func ImageToPictureData(pic []byte) (pix.Picture, error) {
	img, _, err := image.Decode(bytes.NewReader(pic))
	if err != nil {
		return nil, err
	}

	return pix.PictureDataFromImage(img), nil
}
