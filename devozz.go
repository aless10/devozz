package main

import (
	"bytes"
	"image"
	_ "image/png"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/examples/resources/images"
)

const (
	screenWidth  = 320
	screenHeight = 240

	frameOX     = 0
	frameOY     = 32
	frameWidth  = 32
	frameHeight = 32
	frameCount  = 4
)

var (
	runnerImage *ebiten.Image
	bgImage     *ebiten.Image
)

type Character struct {
	x int
	y int
}

type Game struct {
	count     int
	character Character
}

func (g *Game) init() error {
	g.character = Character{}
	return nil
}

func (g *Game) Update() error {
	g.count++
	if ebiten.IsKeyPressed(ebiten.KeyArrowLeft) {
		g.character.x -= 2
	}

	if ebiten.IsKeyPressed(ebiten.KeyArrowRight) {
		g.character.x += 2
	}

	if ebiten.IsKeyPressed(ebiten.KeyArrowUp) {
		g.character.y -= 2
	}

	if ebiten.IsKeyPressed(ebiten.KeyArrowDown) {
		g.character.y += 2
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	i := (g.count / 5) % frameCount
	sx, sy := frameOX+i*frameWidth, frameOY
	op.GeoM.Translate(float64(g.character.x)-float64(frameWidth)/40, float64(g.character.y)-float64(frameHeight)/40)
	op.GeoM.Translate(screenWidth/2, screenHeight/2)

	subrunner := runnerImage.SubImage(image.Rect(sx, sy, sx+frameWidth, sy+frameHeight))
	screen.DrawImage(bgImage, op)
	screen.DrawImage(subrunner.(*ebiten.Image), op)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func init() {
	// Decode an image from the image file's byte slice.
	img, _, err := image.Decode(bytes.NewReader(images.Runner_png))
	if err != nil {
		log.Fatal(err)
	}
	runnerImage = ebiten.NewImageFromImage(img)

	bgImg, _, err := image.Decode(bytes.NewReader(images.Tile_png))
	if err != nil {
		log.Fatal(err)
	}
	bgImage = ebiten.NewImageFromImage(bgImg)

}

func main() {

	ebiten.SetWindowSize(screenWidth*2, screenHeight*2)
	ebiten.SetWindowTitle("Animation (Ebitengine Demo)")
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
