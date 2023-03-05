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
	col := circle{
		center: enemy.position,
		radius: 30,
	}
	enemy.collisions = append(enemy.collisions, col)
	return enemy
}
