package main

import (
	"fmt"

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

	// sr := newSpriteRenderer(basicEnemy, renderer, "sprites/basic_enemy.bmp")
	// basicEnemy.addComponent(sr)
	idleSequence, err := newSequence("sprites/idle", 10, true, renderer)
	if err != nil {
		panic(fmt.Errorf("Creating idle sequence: %v", err))
	}
	destroySequence, err := newSequence("sprites/animation", 10, false, renderer)
	if err != nil {
		panic(fmt.Errorf("Creating destroy sequence: %v", err))
	}

	sequences := map[string]*sequence{
		"idle":    idleSequence,
		"destroy": destroySequence,
	}

	animator := newAnimator(basicEnemy, sequences, "idle")
	basicEnemy.addComponent(animator)

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
