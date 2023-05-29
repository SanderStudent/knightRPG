package main

import (
	"fmt"
	"math/rand"
	"strconv"

	"github.com/veandco/go-sdl2/sdl"
)

const (
	enemySize    = 100
	enemyHP      = 10
	enemyAttack  = 5
	enemyDefence = 5
)

type enemy struct {
	character
	justFought bool
}

func newEnemy(renderer *sdl.Renderer, position vector) enemy {
	e := enemy{
		character: character{
			position:  position,
			size:      enemySize,
			currentHP: enemyHP,
			maxHP:     enemyHP,
			attack:    enemyAttack,
			defence:   enemyDefence,
		},
	}
	err := e.newSpriteRenderer(renderer, "sprites/enemy.bmp")
	if err != nil {
		panic("failed making new enemy")
	}
	return e
}

func (e *enemy) newSpriteRenderer(renderer *sdl.Renderer, filename string) error {
	sr := spriteRenderer{}
	img, err := sdl.LoadBMP(filename)
	if err != nil {
		return fmt.Errorf("loading %v: %v", filename, err)
	}
	defer img.Free()
	sr.tex, err = renderer.CreateTextureFromSurface(img)
	if err != nil {
		return fmt.Errorf("creating texture from %v: %v", filename, err)
	}
	_, _, width, height, err := sr.tex.Query()
	if err != nil {
		return fmt.Errorf("querying texture: %v", err)
	}
	sr.width = int(width)
	sr.height = int(height)
	e.spriteRenderer = sr
	return nil
}

func (e *enemy) update(renderer *sdl.Renderer, p player) int {
	if e.fightMode {
		damage := e.defendFromPlayer(renderer, p)
		return damage
	}
	return 0
}

func (e *enemy) defendFromPlayer(renderer *sdl.Renderer, p player) int {
	damageRoll := rand.Intn(p.attack*p.attack/e.defence + 1)
	e.justFought = true
	e.currentHP -= damageRoll
	printText(strconv.Itoa(damageRoll), e.position.x, e.position.y, renderer)
	return damageRoll
}

func (e *enemy) draw(renderer *sdl.Renderer) error {
	x := e.position.x
	y := e.position.y
	if err := renderer.CopyEx(
		e.spriteRenderer.tex,
		&sdl.Rect{X: 0, Y: 0, W: int32(e.spriteRenderer.width), H: int32(e.spriteRenderer.height)},
		&sdl.Rect{X: int32(x), Y: int32(y), W: int32(e.size), H: int32(e.size)}, 0,
		&sdl.Point{X: int32(e.spriteRenderer.width) / 2, Y: int32(e.spriteRenderer.height) / 2},
		sdl.FLIP_NONE); err != nil {
		return err
	}
	err := showHP(renderer, x, y, e.character)
	return err
}
