package engine

import pix "github.com/faiface/pixel"

// Anything that gets updated every frame such as UI elements and Entities etc.
type Component interface {
	Update(*Game) // ? error
	Rect() pix.Rect
}

type State int
