package main

import (
	"math"

	"github.com/veandco/go-sdl2/sdl"
)

type bullet struct {
	tex   *sdl.Texture
	x, y  float64
	angle float64

	active bool
}

const (
	bulletSize  = 32
	bulletSpeed = 0.9
)

func newBullet(renderer *sdl.Renderer) (bul bullet) {
	bul.tex = textureFromBMP(renderer, "sprites/player_bullet.bmp")

	return bul
}

func (bul *bullet) draw(renderer *sdl.Renderer) {
	if !bul.active {
		return
	}
	x := bul.x - bulletSize/2.0
	y := bul.y - bulletSize/2.0

	renderer.Copy(bul.tex,
		&sdl.Rect{X: 0, Y: 0, W: bulletSize, H: bulletSize},
		&sdl.Rect{X: int32(x), Y: int32(y), W: bulletSize, H: bulletSize})
}

var bulletPool []*bullet

//ammo
func initBulletPool(renderer *sdl.Renderer) {
	for i := 0; i < 50; i++ {
		bul := newBullet(renderer)
		bulletPool = append(bulletPool, &bul)
	}
}

//ammo going down
func bulletFromPool() (*bullet, bool) {
	for _, bul := range bulletPool {
		if !bul.active {
			return bul, true
		}
	}

	return nil, false
}

//bullet movement
func (bul *bullet) update() {
	bul.x += bulletSpeed * math.Cos(bul.angle) //pre-stablished movement math
	bul.y += bulletSpeed * math.Sin(bul.angle)

	if bul.x > screenWidht || bul.x < 0 || bul.y > screenHeight || bul.y < 0 { //reload
		bul.active = false
	}
}
