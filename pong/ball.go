package pong

import (
	// "image"
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten"
)

//STATES OF GAME
type STATES int

const (
	SETUP STATES = iota
	ONPLAY
	ONSTOP
)

type Ball struct {
	Radius int

	PosX float64
	PosY float64

	velX float64
	velY float64

	accX int
	accY int

	isPlaying STATES
}

const (
	RADIUS = 15

	POSX = 20
	POSY = 20

	VELX = 2
	VELY = 2

	ACCX = 2
	ACCY = 2
)

//private method
func (b *Ball) initBall(srcPosX float64, srcPosY float64) {

	b.Radius = RADIUS
	b.PosX = srcPosX - float64(b.Radius*2)
	b.PosY = srcPosY - float64(b.Radius*2)

	b.velX = VELX
	b.velY = VELY

	b.accX = ACCX
	b.accY = ACCY

	b.isPlaying = ONPLAY
}

var (
	squareDrawing *ebiten.Image
)

//public method
func (b *Ball) Draw(srcScreen *ebiten.Image /*screen*/, scrPosX float64, scrPosY float64) error {

	var err error

	if b.isPlaying == SETUP {

		b.initBall(scrPosX, scrPosY)
		// b.onStop(srcScreen)
	}

	squareDrawing, err = ebiten.NewImage(2*b.Radius, 2*b.Radius, ebiten.FilterNearest)

	if err != nil {
		log.Fatal(err)
		return err
	}

	if err = squareDrawing.Fill(color.White); err != nil {
		log.Fatal(err)
		return err
	}

	if b.isPlaying == ONPLAY {
		b.onPlay(srcScreen)

		// log.Printf("Supposed to be executed" )
	}

	if b.isPlaying == ONSTOP {
		b.onPlay(srcScreen)

		// log.Printf("Supposed to be executed" )
	}

	return nil
}

func (b *Ball) Play() {

	if b.isPlaying != ONPLAY {
		b.isPlaying = ONPLAY
	}

	//log.Printf("Play(): %v", b.isPlaying )
}

func (b *Ball) Stop() {

	if b.isPlaying != ONSTOP {
		b.isPlaying = ONSTOP
	}
}

func (b *Ball) onPlay(srcScreen *ebiten.Image) {

	var err error
	b.PosX = b.PosX + b.velX
	b.PosY = b.PosY + b.velY

	var scrWidth, scrHeight int
	scrWidth, scrHeight = srcScreen.Size()

	if b.PosX == 0 {
		b.velX = b.velX * -1
	}
	if b.PosY == 0 {
		b.velY = b.velY * -1
	}

	if b.PosX == float64(scrWidth-(b.Radius*2)) {
		b.velX = b.velX * -1
	}
	if b.PosY == float64(scrHeight-(b.Radius*2)) {
		b.velY = b.velY * -1
	}

	physicsBallOp := &ebiten.DrawImageOptions{}
	physicsBallOp.GeoM.Translate(b.PosX, b.PosY)

	if err = srcScreen.DrawImage(squareDrawing, physicsBallOp); err != nil {
		log.Fatal(err)
	}

}

func (b *Ball) onStop(srcScreen *ebiten.Image) {

	var err error

	b.PosX = 50
	b.PosY = 50

	physicsBallOp := &ebiten.DrawImageOptions{}
	physicsBallOp.GeoM.Translate(b.PosX, b.PosY)

	if err = srcScreen.DrawImage(squareDrawing, physicsBallOp); err != nil {
		log.Fatal(err)
	}

}
