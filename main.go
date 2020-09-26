package main

import (
	"time"

	"github.com/veandco/go-sdl2/sdl"

	e "github.com/OriolBallesteros/gamingingo/errhndl"
)

const (
	screenWidht  = 550
	screenHeight = 700

	targetTickPerSecond = 60
)

var delta float64

func main() {

	//SDL init
	err := sdl.Init(sdl.INIT_EVERYTHING)
	e.ErrHndl("Initializing SDL:", err)

	//Window creation
	window, err := sdl.CreateWindow(
		"Gaming in Go - 2",
		sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
		screenWidht, screenHeight,
		sdl.WINDOW_OPENGL)
	e.ErrHndl("Initializing window:", err)
	defer window.Destroy()

	//Renderer creation
	renderer, err := sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	e.ErrHndl("Initializing renderer:", err)
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
		frameStartTime := time.Now()

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
				e.ErrHndl("Updating elment", err)

				err = elem.draw(renderer)
				e.ErrHndl("Drawing element", err)
			}
		}

		for _, bul := range bulletPool {
			bul.draw(renderer)
			bul.update()
		}

		if err := checkCollisions(); err != nil {
			e.ErrHndl("Checking collisions:", err)
		}

		renderer.Present()

		delta = time.Since(frameStartTime).Seconds() * targetTickPerSecond

	}
}
