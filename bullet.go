package main

import (
	"github.com/veandco/go-sdl2/sdl"
)

const (
	bulletSize  = 32
	bulletSpeed = 6
)

func newBullet(renderer *sdl.Renderer) *element {
	bullet := &element{}

	sr := newSpriteRenderer(bullet, renderer, "sprites/player_bullet.bmp")
	bullet.addComponent(sr)

	mover := newBulletMover(bullet, bulletSpeed)
	bullet.addComponent(mover)

	col := circle{
		center: bullet.position,
		radius: 6,
	}
	bullet.collisions = append(bullet.collisions, col)

	bullet.tag = "bullet"

	return bullet
}

var bulletPool []*element

//ammo
func initBulletPool(renderer *sdl.Renderer) {
	for i := 0; i < 50; i++ {
		bul := newBullet(renderer)
		elements = append(elements, bul)
		bulletPool = append(bulletPool, bul)
	}
}

//ammo going down
func bulletFromPool() (*element, bool) {
	for _, bul := range bulletPool {
		if !bul.active {
			return bul, true
		}
	}

	return nil, false
}
