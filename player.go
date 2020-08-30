package main

import (
	"fmt"

	"github.com/veandco/go-sdl2/sdl"
)

type player struct {
	tex *sdl.Texture
}

func newPlayer(renderer *sdl.Renderer) (p player, err error) {
	//Load img...
	img, err := sdl.LoadBMP("sprites/player.bmp")
	if err != nil {
		return player{}, fmt.Errorf("Loading player sprite: %v", err)
	}
	defer img.Free()

	//...and set it to renderer
	p.tex, err = renderer.CreateTextureFromSurface(img)
	if err != nil {
		return player{}, fmt.Errorf("Creating player texture: %v", err)
	}

	return p, nil
}

func (p *player) draw(renderer *sdl.Renderer) {
	//Copy - what where (which image and where in the screen)
	renderer.Copy(p.tex,
		&sdl.Rect{X: 0, Y: 0, W: 105, H: 105},   //zoom img itself
		&sdl.Rect{X: 50, Y: 50, W: 105, H: 105}) //place img in background
}
