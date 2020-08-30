package main

import (
	"github.com/veandco/go-sdl2/sdl"
)

const (
	screenWidht  = 600
	screenHeight = 800
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
	player, err := newPlayer(renderer)
	errHndl("Creating player:", err)

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

		renderer.Present()
	}
}
