package main

import (
	"github.com/veandco/go-sdl2/sdl"
)

const (
	enemySize = 100
)

func newEnemy(renderer *sdl.Renderer, position vector) *element {
	enemy := &element{
		position: position,
		rotation: 0,
		active:   true,
		size:     enemySize,
	}

	sr := newSpriteRenderer(enemy, renderer, "sprites/enemy.bmp")
	enemy.addComponent(sr)

	return enemy
}

//func (e *enemy) draw(renderer *sdl.Renderer) {
//	x := e.x - enemySize/2.0
//	y := e.y - enemySize/2.0
//
//	if err := renderer.Copy(e.tex,
//		&sdl.Rect{W: 580, H: 580},
//		&sdl.Rect{X: int32(x), Y: int32(y), W: enemySize, H: enemySize},
//	); err != nil {
//		panic("render copy failed")
//	}
//}
