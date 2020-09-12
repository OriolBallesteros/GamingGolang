package main

import (
	"fmt"

	"github.com/veandco/go-sdl2/sdl"
)

const (
	screenWidht  = 600
	screenHeight = 800
)

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

func main() {

	//SDL init
	err := sdl.Init(sdl.INIT_EVERYTHING)
	errHndl("Initializing SDL:", err)

	//Window creation
	window, err := sdl.CreateWindow(
		"Gaming in Go - 2",
		sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
		screenWidht, screenHeight,
		sdl.WINDOW_OPENGL)
	errHndl("Initializing window:", err)
	defer window.Destroy()

	//Renderer creation
	renderer, err := sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	errHndl("Initializing renderer:", err)
	defer renderer.Destroy()

	//Load img and set it to renderer
	player := newPlayer(renderer)

	var enemies []basicEnemy
	for i := 0; i < 5; i++ {
		for j := 0; j < 3; j++ {
			x := (float64(i)/5)*screenWidht + (basicEnemySize / 2)
			y := float64(j)*basicEnemySize + (basicEnemySize / 2)

			enemy := newBasicEnemy(renderer, x, y)

			enemies = append(enemies, enemy)
		}
	}

	initBulletPool(renderer)

	for {
		//Set quit command
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch event.(type) {
			case *sdl.QuitEvent:
				return
			}
		}

		//Show, present
		renderer.SetDrawColor(255, 255, 255, 255)
		renderer.Clear()

		//Copy - what where (which image and where in the screen)
		player.draw(renderer)
		player.update()

		for _, enemy := range enemies {
			enemy.draw(renderer)
		}

		for _, bul := range bulletPool {
			bul.draw(renderer)
			bul.update()
		}

		renderer.Present()
	}
}
