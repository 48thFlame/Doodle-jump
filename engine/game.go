package engine

import (
	"fmt"
	"image/color"
	"time"

	pix "github.com/faiface/pixel"
	pixgl "github.com/faiface/pixel/pixelgl"
)

func Initialize(winConf pixgl.WindowConfig, fps int, bgk *pix.Sprite) *Game {
	g := &Game{}

	g.WinConf = winConf

	win, err := pixgl.NewWindow(g.WinConf)
	if err != nil {
		panic(fmt.Errorf("error creating window: %v", err))
	}

	g.Win = win
	g.Bgk = bgk
	g.millsPerFrame = 1000 / float64(fps)
	g.states = make(statesType)
	g.fields = make(fieldsType)

	return g
}

type statesType map[State][]Component
type fieldsType map[string]interface{}

type Game struct {
	WinConf       pixgl.WindowConfig
	Win           *pixgl.Window
	Bgk           *pix.Sprite
	millsPerFrame float64
	state         State
	states        statesType
	fields        fieldsType
}

func (g *Game) update() {
	for _, c := range g.states[g.state] {
		c.Update(g)
	}
}

func (g *Game) Run() {
	for !g.Win.Closed() {
		g.Win.Clear(color.Opaque)
		g.Bgk.Draw(g.Win, pix.IM.Moved(g.Bgk.Frame().Center()))

		g.update()
		g.Win.Update()

		time.Sleep(time.Millisecond * time.Duration(g.millsPerFrame))
	}

	g.Win.Destroy()
}

func (g *Game) ChangeState(state State) {
	g.state = state
}

func (g *Game) CreateState(state State) {
	g.states[state] = make([]Component, 0)
}

func (g *Game) AddComponentsToState(state State, c ...Component) {
	g.states[state] = append(g.states[state], c...)
}

func (g *Game) SetField(field string, value interface{}) {
	g.fields[field] = value
}

func (g *Game) GetField(field string) interface{} {
	return g.fields[field]
}
