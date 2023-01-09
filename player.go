package main

import (
	"github.com/veandco/go-sdl2/sdl"
)

const (
	playerSpeed = 5
	playerSize  = 200
)

var delta = 0.5

func newPlayer(renderer *sdl.Renderer, position vector) *element {
	player := &element{
		position: position,
		rotation: 0,
		active:   true,
		size:     playerSize,
	}

	sr := newSpriteRenderer(player, renderer, "sprites/player.bmp")
	player.addComponent(sr)

	mover := newKeyboardMover(player, playerSpeed)
	player.addComponent(mover)

	return player
}
