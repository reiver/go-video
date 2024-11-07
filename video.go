package video

import (
	"image"
	"image/draw"
	"time"
)

// Video is a series of image.Image.
type Video interface {
	Delay(int) time.Duration
	DrawOperation(int) draw.Op
	Image(int) image.Image
	Len() int
	LoopCount() int
}
