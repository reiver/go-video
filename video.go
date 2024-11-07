package video

import (
	"image"
	"time"
)

// Video is a series of image.Image.
type Video interface {
	Delay(int) time.Duration
	Image(int) image.Image
	Len() int
	LoopCount() int
}
