package main

import (
	"fmt"

	"github.com/veandco/go-sdl2/sdl"
)

const (
	screenWidth  = 12  //in blocks
	screenHeight = 8   //in blocks
	blockSize    = 100 //in pixels
)

type vector struct {
	x, y float64
}

func main() {
	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		fmt.Println(err)
	}
	window, err := sdl.CreateWindow("KnightRPG", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED, screenWidth*blockSize, screenHeight*blockSize, sdl.WINDOW_OPENGL)
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

	p := newPlayer(renderer, vector{blockSize * 1, blockSize * 1})
	e := newEnemy(renderer, vector{blockSize * (screenWidth - 2), blockSize * (screenHeight - 2)})

	for {
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch event.(type) {
			case *sdl.QuitEvent:
				return
			}
		}
		if err = renderer.SetDrawColor(0, 0, 100, 255); err != nil {
			return
		}
		if err = renderer.Clear(); err != nil {
			return
		}
		p.update(renderer, e)
		e.update(renderer, p)
		checkCollision(&p, &e)
		if p.fightMode {
			printText("FIGHT!", p.position.x, p.position.y-blockSize, renderer)
		}
		if p.currentHP <= 0 {
			p.fightMode = false
			e.fightMode = false
			printText("You died!", p.position.x, p.position.y-blockSize, renderer)
		}
		if e.currentHP <= 0 {
			p.fightMode = false
			e.fightMode = false
			printText("Enemy died!", e.position.x, e.position.y-blockSize, renderer)
		}
		if err = p.draw(renderer); err != nil {
			fmt.Println("drawing player:", err)
			return
		}
		if err = e.draw(renderer); err != nil {
			fmt.Println("drawing enemy:", err)
			return
		}
		renderer.Present()
		if p.justMoved {
			sdl.Delay(300)
			p.justMoved = false
		}
		if p.justFought {
			sdl.Delay(1500)
			p.justFought = false
			e.justFought = false
		}
	}
}
