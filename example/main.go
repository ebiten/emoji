// SPDX-License-Identifier: Apache-2.0

package main

import (
	"github.com/ebiten/emoji"
	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct{}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	const zwj = "\u200d"
	lines := [][]string{
		{"🍣", "🤔", "🐣", "⚽", "🍔"},
		{"💵", "👨" + zwj + "👩" + zwj + "👧" + zwj + "👦", "🥺", "💯", "🈲"},
		{"✋" + "🏻", "✋" + "🏼", "✋" + "🏽", "✋" + "🏾", "✋" + "🏿"},
	}

	for j, line := range lines {
		for i, str := range line {
			op := &ebiten.DrawImageOptions{}
			op.GeoM.Translate(float64(i)*128, float64(j)*128)
			screen.DrawImage(emoji.Image(str), op)
		}
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return outsideWidth, outsideHeight
}

func main() {
	ebiten.SetWindowTitle("Ebiten Emoji Test")
	ebiten.SetWindowSize(5*128, 3*128)
	if err := ebiten.RunGame(&Game{}); err != nil {
		panic(err)
	}
}
