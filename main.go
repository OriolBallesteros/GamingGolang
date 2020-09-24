package main

import (
	"github.com/veandco/go-sdl2/sdl"
)

const (
	screenWidht  = 550
	screenHeight = 700
)

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
	elements = append(elements, player)

	for i := 0; i < 5; i++ {
		for j := 0; j < 3; j++ {
			x := (float64(i)/5)*screenWidht + (basicEnemySize / 2)
			y := float64(j)*basicEnemySize + (basicEnemySize / 2)

			enemy := newBasicEnemy(renderer, vector{x, y})

			elements = append(elements, enemy)
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

		for _, elem := range elements {
			if elem.active {
				err = elem.update()
				errHndl("Updating elment", err)

				err = elem.draw(renderer)
				errHndl("Drawing element", err)
			}
		}

		for _, bul := range bulletPool {
			bul.draw(renderer)
			bul.update()
		}

		renderer.Present()
	}
}
