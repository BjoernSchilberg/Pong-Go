package main

import (
	"Pong-Go/pong"
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten"
)

const (
	screenWidth  = 640
	screenHeight = 480
	xSpd         = 5
	ySpd         = 5
)

var (
	imgScreen *ebiten.Image

	xPos float64
	yPos float64
	ball pong.Ball
	bar  pong.Bar
)

func update(screen *ebiten.Image) error {

	var err error

	// screen.Fill(color.RGBA{ 100, 100, 100, 255 }) //gray color
	screen.Fill(color.Black)

	if err = ball.Draw(screen, screenWidth/2, screenHeight/2); err != nil {
		log.Fatal(err)
		return err
	}
	ball.Play()

	if err = bar.Draw(screen, screenWidth/2, screenHeight/2); err != nil {
		log.Fatal(err)
		return err
	}
	bar.Play()

	// collision
	// if ball.PosX == 0 || ball.PosX == float64(screenWidth-2*(ball.Radius)) {
	// 	log.Println("Ground")
	// }
	// if ball.PosX

	return nil
}

func main() {

	var err error

	imgScreen, err = ebiten.NewImage(screenWidth, screenHeight, ebiten.FilterNearest)
	if err != nil {
		log.Fatal(err)
	}

	// posPoint image.Point = { 10, 10  }
	//	my_ball = pong.ball{ posPoint , 5  }

	hasError := ebiten.Run(update, screenWidth, screenHeight, 1, "Pong")

	if hasError != nil {
		log.Fatal(hasError)
	}
}
