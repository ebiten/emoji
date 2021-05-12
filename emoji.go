// SPDX-License-Identifier: Apache-2.0

// Package emoji provides Emoji images for Ebiten.
package emoji

import (
	"embed"
	"fmt"
	"image/png"
	"sync"

	"github.com/hajimehoshi/ebiten/v2"
)

//go:embed image/emoji_*.png
var pngImages embed.FS

var (
	ebitenImagesM sync.Mutex
	ebitenImages  = map[string]*ebiten.Image{}
)

// Image returns an ebiten.Image with the specified string.
// If there is no matched image, Image returns nil.
//
// Ebiten images are cached internally. Then, the same object is returned for the same string.
func Image(str string) *ebiten.Image {
	ebitenImagesM.Lock()
	defer ebitenImagesM.Unlock()

	path := "image/emoji_u"
	for i, r := range str {
		if i > 0 {
			path += "_"
		}
		path += fmt.Sprintf("%x", r)
	}
	path += ".png"

	if img, ok := ebitenImages[path]; ok {
		return img
	}

	f, err := pngImages.Open(path)
	if err != nil {
		// Not found.
		return nil
	}
	img, err := png.Decode(f)
	if err != nil {
		// This should never be reached.
		panic(fmt.Sprintf("emoji: png.Decode failed: %v", err))
		return nil
	}
	eimg := ebiten.NewImageFromImage(img)
	ebitenImages[str] = eimg
	return eimg
}
