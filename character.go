package main

import (
	"fmt"

	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
)

type character struct {
	position       vector
	size           float64
	spriteRenderer spriteRenderer
	currentHP      int
	maxHP          int
}

func drawText(renderer *sdl.Renderer, x, y float64, c character) error {
	if err := ttf.Init(); err != nil {
		return err
	}
	font, err := ttf.OpenFont("fonts/OpenSans-Regular.ttf", 12)
	if err != nil {
		return err
	}
	defer font.Close()
	hpString := fmt.Sprintf("%d/%d", c.currentHP, c.maxHP)
	surfaceMessage, err := font.RenderUTF8Solid("HP: "+hpString, sdl.Color{R: 255, G: 0, B: 0, A: 255})
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
		&sdl.Rect{X: int32(x), Y: int32(y - 40), W: int32(c.size), H: int32(50)}, 0,
		&sdl.Point{X: int32(c.spriteRenderer.width) / 2, Y: int32(c.spriteRenderer.height) / 2},
		sdl.FLIP_NONE); err != nil {
		return err
	}
	return nil
}