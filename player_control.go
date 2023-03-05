package main

import (
	"github.com/veandco/go-sdl2/sdl"
)

type keyboardMover struct {
	container *element
	speed     float64
	sr        *spriteRenderer
}

func newKeyboardMover(container *element, speed float64) *keyboardMover {
	return &keyboardMover{
		container: container,
		speed:     speed,
		sr:        container.getComponent(&spriteRenderer{}).(*spriteRenderer),
	}
}

func (mover *keyboardMover) onDraw(elem *element, renderer *sdl.Renderer) error {
	return nil
}

func (mover *keyboardMover) onUpdate() error {
	keys := sdl.GetKeyboardState()

	cont := mover.container
	if keys[sdl.SCANCODE_LEFT] == 1 {
		if cont.position.x-cont.size/2.0 > 0 {
			cont.position.x -= mover.speed * delta
		}
	} else if keys[sdl.SCANCODE_RIGHT] == 1 {
		if cont.position.x+cont.size/2.0 < screenWidth {
			cont.position.x += mover.speed * delta
		}
	}
	if keys[sdl.SCANCODE_UP] == 1 {
		if cont.position.y-cont.size/2.0 > 0 {
			cont.position.y -= mover.speed * delta
		}
	} else if keys[sdl.SCANCODE_DOWN] == 1 {
		if cont.position.y+cont.size/2.0 < screenHeight {
			cont.position.y += mover.speed * delta
		}
	}
	return nil
}

func (mover *keyboardMover) onCollision(other *element) error {
	//TODO implement me
	panic("implement me")
}

type keyboardFighter struct {
	//container *element
	//cooldown  time.Duration
	//lastShot  time.Time
}

func (fighter *keyboardFighter) onUpdate() error {
	keys := sdl.GetKeyboardState()

	if keys[sdl.SCANCODE_SPACE] == 1 {
		//battle for example
	}

	return nil
}

func (fighter *keyboardFighter) onDraw(renderer *sdl.Renderer) error {
	return nil
}
