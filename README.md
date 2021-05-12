# Ebiten Emoji

Package emoji provides Emoji images for Ebiten.

![Example](example/screenshot.png)

## Usage

```go
func (*YourGame) Draw(screen *ebiten.Image) {
    screen.DrawImage(emoji.Get("🍣", nil))
}
```
