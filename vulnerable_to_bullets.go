package main

import (
	"github.com/veandco/go-sdl2/sdl"
)

type vulnerableToBullet struct {
	container *element
	animator  *animator
}

func newVulnerableToBullets(container *element) *vulnerableToBullet {
	return &vulnerableToBullet{
		container: container,
		animator:  container.getComponent(&animator{}).(*animator),
	}
}

func (vtb *vulnerableToBullet) onDraw(renderer *sdl.Renderer) error {
	return nil
}

func (vtb *vulnerableToBullet) onUpdate() error {
	if vtb.animator.finished && vtb.animator.current == "destroy" {
		vtb.container.active = false
	}

	return nil
}

func (vtb *vulnerableToBullet) onCollision(other *element) error {
	if other.tag == "bullet" {
		vtb.animator.setSequence("destroy")
	}
	return nil
}
