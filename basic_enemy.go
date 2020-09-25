package main

import (
	"github.com/veandco/go-sdl2/sdl"
)

type basicEnemy struct {
	tex  *sdl.Texture
	x, y float64
}

const (
	basicEnemySize = 105
)

func newBasicEnemy(renderer *sdl.Renderer, position vector) *element {
	basicEnemy := &element{}

	basicEnemy.position = position
	basicEnemy.rotation = 180

	sr := newSpriteRenderer(basicEnemy, renderer, "sprites/basic_enemy.bmp")
	basicEnemy.addComponent(sr)

	vtb := newVulnerableToBullets(basicEnemy)
	basicEnemy.addComponent(vtb)

	col := circle{
		center: basicEnemy.position,
		radius: 32,
	}
	basicEnemy.collisions = append(basicEnemy.collisions, col)

	basicEnemy.active = true

	return basicEnemy
}
