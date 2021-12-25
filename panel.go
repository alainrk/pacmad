package main

import (
	"fmt"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"github.com/faiface/pixel/text"
	"golang.org/x/image/font/basicfont"
)

type Panel struct {
	y             float64
	hearthSprites []*pixel.Sprite
	game          *Game
	win           *pixelgl.Window
}

func NewPanel(y float64, game *Game, win *pixelgl.Window) *Panel {
	p := &Panel{
		y:             y,
		game:          game,
		win:           win,
		hearthSprites: []*pixel.Sprite{},
	}
	p.loadHearthSprites()
	return p
}

func (p *Panel) Draw() {
	sprite := p.hearthSprites[0]
	for i := 0; i < p.game.lives; i++ {
		sprite.Draw(
			p.win,
			pixel.IM.Moved(
				pixel.V(
					p.win.Bounds().W()-float64(i*32+i*5)-WindowBoundaryDelta-10,
					WindowBoundaryDelta+10,
				),
			),
		)
	}

	basicAtlas := text.NewAtlas(basicfont.Face7x13, text.ASCII)
	basicTxt := text.New(pixel.V(20, 40), basicAtlas)

	fmt.Fprintf(basicTxt, "SCORE: %d\n", p.game.points)
	fmt.Fprintln(basicTxt, "Click or [Space] to shoot")
	basicTxt.Draw(p.win, pixel.IM.Scaled(basicTxt.Orig, 1.6))
}
