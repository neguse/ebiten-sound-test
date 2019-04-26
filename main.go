package main

import (
	"fmt"
	// "image"
	"io/ioutil"
	"log"
	// "math"
	// "math/rand"

	_ "github.com/neguse/ebi/statik"
	"github.com/rakyll/statik/fs"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/audio"
	"github.com/hajimehoshi/ebiten/audio/wav"
	"github.com/hajimehoshi/ebiten/ebitenutil" // This is required to draw debug texts.
)

const (
	width  = 240
	height = 240
	scale  = 2
)

// type Bullet struct {
// 	x, y   float64
// 	vx, vy float64
// }

// func (b *Bullet) Update(dt float64) {
// 	b.x += b.vx * dt
// 	b.y += b.vy * dt
// 	if b.x < 0 || width < b.x {
// 		b.vx = -b.vx
// 	}
// 	if b.y < 0 || height < b.y {
// 		b.vy = -b.vy
// 	}
// }

var (
	audioCtx    *audio.Context
	audioPlayer *audio.Player
	// ballImage   *ebiten.Image
	// t           float64
	// num         int
	// bullets     []*Bullet
)

func init2() {
	// num = 1000
	statikFs, err := fs.New()
	if err != nil {
		log.Panic(err)
	}
	audioCtx, err := audio.NewContext(44100)
	if err != nil {
		log.Panic(err)
	}
	f, err := statikFs.Open("/damage.wav")
	if err != nil {
		log.Panic(err)
	}
	defer f.Close()
	wavData, err := ioutil.ReadAll(f)
	if err != nil {
		log.Panic(err)
	}
	// f2, err := statikFs.Open("/1.png")
	// if err != nil {
	// 	log.Panic(err)
	// }
	// defer f2.Close()
	// img, _, err := image.Decode(f2)
	// if err != nil {
	// 	log.Panic(err)
	// }
	// ballImage, err = ebiten.NewImageFromImage(img, ebiten.FilterNearest)
	// if err != nil {
	// 	log.Panic(err)
	// }
	d, err := wav.Decode(audioCtx, audio.BytesReadSeekCloser(wavData))
	if err != nil {
		log.Panic(err)
	}
	audioPlayer, err = audio.NewPlayer(audioCtx, d)
	if err != nil {
		log.Panic(err)
	}
}

func update(screen *ebiten.Image) error {
	// if ebiten.IsKeyPressed(ebiten.KeyLeft) {
	// 	num--
	// }
	// if ebiten.IsKeyPressed(ebiten.KeyRight) {
	// 	num++
	// }
	if !audioPlayer.IsPlaying() {
		audioPlayer.Rewind()
		audioPlayer.Play()
	}
	// for i := 0; i < num; i++ {
	// 	if len(bullets) <= i {
	// 		a := rand.Float64() * math.Pi * 2
	// 		v := 300.0 + rand.Float64()*100.0
	// 		bullets = append(bullets, &Bullet{x: width / 2, y: height / 2, vx: math.Cos(a) * v, vy: math.Sin(a) * v})
	// 	}
	// 	bullets[i].Update(0.016)
	// }
	if ebiten.IsDrawingSkipped() {
		return nil
	}
	// t += 0.016

	// for i := 0; i < num; i++ {
	// 	opt := &ebiten.DrawImageOptions{}
	// 	opt.GeoM.Translate(bullets[i].x, bullets[i].y)
	// 	screen.DrawImage(ballImage, opt)
	// }

	ebitenutil.DebugPrint(screen, fmt.Sprintf("TPS:%2.2f FPS:%2.2f", ebiten.CurrentTPS(), ebiten.CurrentFPS()))
	// ebitenutil.DebugPrint(screen, fmt.Sprintf("TPS:%2.2f FPS:%2.2f\nnum:%d", ebiten.CurrentTPS(), ebiten.CurrentFPS(), num))
	return nil
}

func main() {
	init2()
	if err := ebiten.Run(update, width, height, scale, "Hello world!"); err != nil {
		panic(err)
	}
}
