package main

import (
	"fmt"
	"time"

	"github.com/veandco/go-sdl2/sdl"
)

const (
	screenWidth          = 1200
	screenHeight         = 800
	targetTicksPerSecond = 60
)

type vector struct {
	x, y float64
}

func main() {
	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		fmt.Println(err)
	}
	window, err := sdl.CreateWindow("KnightRPG", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED, screenWidth, screenHeight, sdl.WINDOW_OPENGL)
	if err != nil {
		fmt.Println("Initializing window", err)
		return
	}
	defer func(window *sdl.Window) {
		err := window.Destroy()
		if err != nil {
			panic("error destroying window")
		}
	}(window)

	renderer, err := sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	if err != nil {
		fmt.Println("initializing renderer:", err)
		return
	}
	defer func(renderer *sdl.Renderer) {
		err := renderer.Destroy()
		if err != nil {
			panic("error destroying renderer")
		}
	}(renderer)

	elements = append(elements, newPlayer(renderer, vector{screenWidth * 1.0 / 3.0, screenHeight * 1.0 / 3.0}))
	elements = append(elements, newEnemy(renderer, vector{screenWidth * 2.0 / 3.0, screenHeight * 2.0 / 3.0}))

	for {
		frameStartTime := time.Now()
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch event.(type) {
			case *sdl.QuitEvent:
				return
			}
		}
		//add comment
		err := renderer.SetDrawColor(0, 0, 100, 255)
		if err != nil {
			return
		}
		err = renderer.Clear()
		if err != nil {
			return
		}
		for _, elem := range elements {
			if elem.active {
				err = elem.update()
				if err != nil {
					fmt.Println("updating element:", err)
					return
				}
				err = elem.draw(renderer)
				if err != nil {
					fmt.Println("drawing element:", elem)
					return
				}
			}
		}
		if err := checkCollisions(); err != nil {
			fmt.Println("checking collisions:", err)
			return
		}
		renderer.Present()
		delta = time.Since(frameStartTime).Seconds() * targetTicksPerSecond
	}
}
