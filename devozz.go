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

type Background struct {
	x int
	y int
}

func (p *Background) Position() (int, int) {
	return p.x, p.y
}

func (p *Background) MoveLeft() {
	p.x -= 2
}

func (p *Background) MoveRight() {
	p.x += 2
}

func (p *Background) MoveDown() {
	p.y += 2
}

func (p *Background) MoveUp() {
	p.y -= 2
}

type Game struct {
	count      int
	character  Character
	background Background
}

func (g *Game) init() error {
	g.character = Character{}
	g.background = Background{}
	return nil
}

func yBorderUp(character *Character) bool {
	distanceFromBorder := screenHeight/2 + character.y
	return distanceFromBorder <= 0
}

func yBorderDown(character *Character) bool {
	distanceFromBorder := screenHeight/2 + character.y
	return distanceFromBorder > 210
}

func xBorderRight(character *Character) bool {
	distanceFromBorder := screenWidth/2 - character.x
	return distanceFromBorder < 26
}

func xBorderLeft(character *Character) bool {
	distanceFromBorder := screenWidth/2 - character.x
	return distanceFromBorder > 328
}

func (g *Game) Update() error {
	g.count++

	if !xBorderLeft(&g.character) && ebiten.IsKeyPressed(ebiten.KeyArrowLeft) {
		g.character.x -= 2
		// g.background.MoveLeft()
	}

	if !xBorderRight(&g.character) && ebiten.IsKeyPressed(ebiten.KeyArrowRight) {
		g.character.x += 2
		// g.background.MoveRight()
	}

	if !yBorderUp(&g.character) && ebiten.IsKeyPressed(ebiten.KeyArrowUp) {
		g.character.y -= 2
		// g.background.MoveUp()
	}

	if !yBorderDown(&g.character) && ebiten.IsKeyPressed(ebiten.KeyArrowDown) {
		g.character.y += 2
		// g.background.MoveDown()
	}

	return nil
}

func (g *Game) DrawBg(screen *ebiten.Image) {

	x16, y16 := g.background.Position()
	offsetX, offsetY := float64(-x16)/16, float64(-y16)/16

	// Draw bgImage on the screen repeatedly.
	const repeat = 3
	w, h := bgImage.Bounds().Dx(), bgImage.Bounds().Dy()
	for j := 0; j < repeat; j++ {
		for i := 0; i < repeat; i++ {
			op := &ebiten.DrawImageOptions{}
			op.GeoM.Translate(float64(w*i), float64(h*j))
			op.GeoM.Translate(offsetX, offsetY)
			screen.DrawImage(bgImage, op)
		}
	}
}

func (g *Game) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	i := (g.count / 5) % frameCount
	sx, sy := frameOX+i*frameWidth, frameOY
	op.GeoM.Translate(float64(g.character.x)-float64(frameWidth)/40, float64(g.character.y)-float64(frameHeight)/40)
	op.GeoM.Translate(screenWidth/2, screenHeight/2)
	subrunner := runnerImage.SubImage(image.Rect(sx, sy, sx+frameWidth, sy+frameHeight))
	g.DrawBg(screen)
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
