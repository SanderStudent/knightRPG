package main

import (
	"fmt"

	"github.com/veandco/go-sdl2/sdl"
)

const (
	playerSize = 100
	playerHP   = 20
)

type player struct {
	character
	justMoved bool
}

type spriteRenderer struct {
	container     *player
	tex           *sdl.Texture
	width, height int
}

func newPlayer(renderer *sdl.Renderer, position vector) player {
	p := player{
		character: character{
			position:  position,
			size:      playerSize,
			currentHP: playerHP,
			maxHP:     playerHP,
		},
	}
	err := p.newSpriteRenderer(renderer, "sprites/player.bmp")
	if err != nil {
		panic("failed making new player")
	}
	return p
}

func (p *player) newSpriteRenderer(renderer *sdl.Renderer, filename string) error {
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
	p.spriteRenderer = sr
	return nil
}

func (p *player) update() {
	keys := sdl.GetKeyboardState()
	if p.fightMode {
		if keys[sdl.SCANCODE_SPACE] == 1 {
			attack()
		}
		return
	}
	if keys[sdl.SCANCODE_LEFT] == 1 {
		if p.position.x > 0 {
			p.position.x -= blockSize
			p.justMoved = true
		}
	} else if keys[sdl.SCANCODE_RIGHT] == 1 {
		if p.position.x < (screenWidth-1)*blockSize {
			p.position.x += blockSize
			p.justMoved = true
		}
	} else if keys[sdl.SCANCODE_UP] == 1 {
		if p.position.y > 0 {
			p.position.y -= blockSize
			p.justMoved = true
		}
	} else if keys[sdl.SCANCODE_DOWN] == 1 {
		if p.position.y < (screenHeight-1)*blockSize {
			p.position.y += blockSize
			p.justMoved = true
		}
	}
}

func attack() {

}

func (p *player) draw(renderer *sdl.Renderer) error {
	x := p.position.x
	y := p.position.y

	if err := renderer.CopyEx(
		p.spriteRenderer.tex,
		&sdl.Rect{X: 0, Y: 0, W: int32(p.spriteRenderer.width), H: int32(p.spriteRenderer.height)},
		&sdl.Rect{X: int32(x), Y: int32(y), W: int32(p.size), H: int32(p.size)}, 0,
		&sdl.Point{X: int32(p.spriteRenderer.width) / 2, Y: int32(p.spriteRenderer.height) / 2},
		sdl.FLIP_NONE); err != nil {
		return err
	}
	err := showHP(renderer, x, y, p.character)
	return err
}
