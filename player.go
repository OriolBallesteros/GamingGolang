package main

import (
	"time"

	"github.com/veandco/go-sdl2/sdl"
)

const (
	playerSpeed = 6
	playerSize  = 105

	playerShotCooldDown = time.Millisecond * 250 //4 shoots/second
)

func newPlayer(renderer *sdl.Renderer) *element {
	player := &element{}

	player.position = vector{
		x: screenWidht / 2.0,
		y: screenHeight - playerSize,
	}
	player.active = true

	sr := newSpriteRenderer(player, renderer, "sprites/player.bmp")
	player.addComponent(sr)

	mover := newKeyboardMover(player, playerSpeed)
	player.addComponent(mover)

	shooter := newKeyboardShooter(player, playerShotCooldDown)
	player.addComponent(shooter)

	return player
}
