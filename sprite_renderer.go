package main

import (
	"fmt"

	"github.com/veandco/go-sdl2/sdl"
)

type spriteRenderer struct {
	container *element
	tex       *sdl.Texture

	width, height float64
}

func newSpriteRenderer(container *element, renderer *sdl.Renderer, filename string) *spriteRenderer {
	tex := textureFromBMP(renderer, filename)

	_, _, width, height, err := tex.Query()
	if err != nil {
		panic(fmt.Errorf("Querying texture: %v", err))
	}

	return &spriteRenderer{
		container: container,
		tex:       textureFromBMP(renderer, filename),
		width:     float64(width),
		height:    float64(height),
	}
}

func (sr *spriteRenderer) onDraw(renderer *sdl.Renderer) error {
	return drawTexture(
		sr.tex,
		sr.container.position,
		sr.container.rotation,
		renderer)
}

func (sr *spriteRenderer) onCollision(other *element) error {
	return nil
}

func (sr *spriteRenderer) onUpdate() error {
	return nil
}

//Get and set img
func textureFromBMP(renderer *sdl.Renderer, filename string) *sdl.Texture {
	img, err := sdl.LoadBMP(filename)
	if err != nil {
		panic(fmt.Errorf("loading %v: %v", filename, err))
	}
	defer img.Free()

	tex, err := renderer.CreateTextureFromSurface(img)
	if err != nil {
		panic(fmt.Errorf("Creating texture from %v: %v", filename, err))
	}

	return tex
}
