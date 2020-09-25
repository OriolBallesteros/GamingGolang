package main

import (
	"github.com/veandco/go-sdl2/sdl"
)

type vulnerableToBullet struct {
	container *element
}

func newVulnerableToBullets(container *element) *vulnerableToBullet {
	return &vulnerableToBullet{container: container}
}

func (vtb *vulnerableToBullet) onDraw(renderer *sdl.Renderer) error {
	return nil
}

func (vtb *vulnerableToBullet) onUpdate() error {
	return nil
}

func (vtb *vulnerableToBullet) onCollision(other *element) error {
	if other.tag == "bullet" {
		vtb.container.active = false
	}
	return nil
}
