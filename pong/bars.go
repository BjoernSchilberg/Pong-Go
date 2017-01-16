package pong

import (
	// "fmt"
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten"
)

//STATES OF GAME
// type STATES int

// const (
// 	SETUP STATES = iota
// 	ONPLAY
// 	ONSTOP
// )

type Bar struct {
	width     int
	height    int
	posX      float64
	posY      float64
	velX      float64
	velY      float64
	isPlaying STATES
}

const (
	BAR_WIDTH  = 15
	BAR_HEIGHT = 45

	BAR_POSX = 0
	BAR_POSY = 20

	BAR_VELX = 2
	BAR_VELY = 2
)

var (
	drawing *ebiten.Image
)

func (bar *Bar) initBar(srcPosX float64, srcPosY float64) {

	bar.width = BAR_WIDTH
	bar.height = BAR_HEIGHT
	bar.posX = BAR_POSX
	bar.posY = (srcPosY) - BAR_HEIGHT/2
	bar.velX = BAR_VELX
	bar.velY = BAR_VELY
}

func (bar *Bar) Draw(srcScreen *ebiten.Image /* screen */, scrPosX float64, scrPosY float64) error {

	var err error

	if bar.isPlaying == SETUP {
		bar.initBar(scrPosX, scrPosY)
	}

	drawing, err = ebiten.NewImage(bar.width, bar.height, ebiten.FilterNearest)

	if err != nil {
		log.Fatal(err)
		return err
	}

	if err = drawing.Fill(color.White); err != nil {
		log.Fatal(err)
		return err
	}
	if bar.isPlaying == ONPLAY {
		bar.onPlay(srcScreen)
	}
	if bar.isPlaying == ONSTOP {
		bar.onStop(srcScreen)
	}

	return nil
}

func (bar *Bar) Play() {

	if bar.isPlaying != ONPLAY {
		bar.isPlaying = ONPLAY
	}
}

func (bar *Bar) Stop() {

	if bar.isPlaying != ONSTOP {
		bar.isPlaying = ONSTOP
	}
}

func (bar *Bar) onPlay(srcScreen *ebiten.Image) {
	var err error

	if ebiten.IsKeyPressed(ebiten.KeyW) {
		// fmt.Println("keyPressed W down")
		bar.posY = bar.posY - bar.velY
	}
	if ebiten.IsKeyPressed(ebiten.KeyS) {
		// fmt.Println("keyPressed S down")
		bar.posY = bar.posY + bar.velY
	}

	physicsBallOp := &ebiten.DrawImageOptions{}
	physicsBallOp.GeoM.Translate(bar.posX, bar.posY)

	if err = srcScreen.DrawImage(drawing, physicsBallOp); err != nil {
		log.Fatal(err)
	}
}

func (bar *Bar) onStop(srcScreen *ebiten.Image) {
	// Do something
}
