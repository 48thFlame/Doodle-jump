package main

import (
	_ "embed"
	"fmt"

	"github.com/avitar64/Doodle-jump/engine"
	pix "github.com/faiface/pixel"
	pixgl "github.com/faiface/pixel/pixelgl"
)

var (
	plrRegularFrame         = pix.R(1, 0, 93, 90)
	plrRegularSmooshedFrame = pix.R(125, 9, 218, 91)
)

//go:embed assets/player.png
var plrPicBytes []byte

func newPlayer() *player {
	plrPic, err := engine.ImageToPictureData(plrPicBytes)
	if err != nil {
		panic(fmt.Errorf("error loading player image: %v", err))
	}

	return &player{
		pos: pix.V(windowWidth/2, windowHeight/2),
		costumes: []*pix.Sprite{
			pix.NewSprite(plrPic, plrRegularFrame),
			pix.NewSprite(plrPic, plrRegularSmooshedFrame),
		},
		costumeIndex: 0,
	}
}

type player struct {
	costumes     []*pix.Sprite
	costumeIndex int

	pos pix.Vec // bottom left corner
}

func (p *player) sprite() *pix.Sprite {
	return p.costumes[p.costumeIndex]
}

func (p *player) Update(g *engine.Game) {
	if g.Win.JustPressed(pixgl.KeySpace) {
		p.costumeIndex = (p.costumeIndex + 1) % len(p.costumes)
	}
	p.sprite().Draw(g.Win, pix.IM.Moved(p.pos))
}

func (p *player) Rect() pix.Rect {
	sFrame := p.sprite().Frame()
	return pix.R(p.pos.X, p.pos.Y, sFrame.W(), p.pos.Y+sFrame.H())
}
