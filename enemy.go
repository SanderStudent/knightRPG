package main

import (
	"fmt"

	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
)

const (
	enemySize = 100
	enemyHP   = 10
)

type enemy struct {
	position       vector
	size           float64
	spriteRenderer spriteRenderer
	healthPoints   int
}

func newEnemy(renderer *sdl.Renderer, position vector) enemy {
	e := enemy{
		position:     position,
		size:         enemySize,
		healthPoints: enemyHP,
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

func (e *enemy) update() error {

	return nil
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
	font, err := ttf.OpenFont("Sans.ttf", 24)
	if err != nil {
		return err
	}
	defer font.Close()
	surfaceMessage, err := font.RenderUTF8Solid("hello world", sdl.Color{R: 100, G: 100, B: 100, A: 255})
	if err != nil {
		return err
	}
	defer surfaceMessage.Free()
	if err = surfaceMessage.Blit(nil, surfaceMessage, &sdl.Rect{X: 400, Y: 300, W: 0, H: 0}); err != nil {
		return err
	}
	messageTexture, err := renderer.CreateTextureFromSurface(surfaceMessage)
	if err != nil {
		return err
	}
	if err = renderer.CopyEx(
		messageTexture,
		nil,
		&sdl.Rect{X: int32(x), Y: int32(y), W: int32(e.size), H: int32(e.size)}, 0,
		&sdl.Point{X: int32(e.spriteRenderer.width) / 2, Y: int32(e.spriteRenderer.height) / 2},
		sdl.FLIP_NONE); err != nil {
		return err
	}

	return nil
}
