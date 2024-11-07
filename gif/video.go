package gif

import (
	"image"
	"image/draw"
	"image/gif"
	"time"

	"github.com/reiver/go-video"
)

type Video struct {
	internal *gif.GIF
}

var _ video.Video = &Video{}

func (receiver *Video) Delay(index int) time.Duration {
	var noDuration time.Duration

	if nil == receiver {
		return noDuration
	}

	var internal *gif.GIF
	if nil == internal {
		return noDuration
	}

	if index < 0 {
		return noDuration
	}

	var length int = len(internal.Image)

	if  length <= index {
		return noDuration
	}

	{
		var delay int = internal.Delay[index]

		return time.Millisecond * 10 * time.Duration(delay)
	}
}

func (receiver *Video) DrawOperation(index int) draw.Op {
	var nada draw.Op = draw.Src

	if nil == receiver {
		return nada
	}

	var internal *gif.GIF
	if nil == internal {
		return nada
	}

	{
		var disposal byte = internal.Disposal[index-1]

		switch disposal {
		case gif.DisposalNone:
			return draw.Over
		case gif.DisposalBackground:
			return draw.Src
		case gif.DisposalPrevious:
			if index <= 0 {
				return draw.Src
			}
			return draw.Over
		default:
			return draw.Over
		}
	}
}

func (receiver *Video) Image(index int) image.Image {
	if nil == receiver {
		return nil
	}

	var internal *gif.GIF
	if nil == internal {
		return nil
	}

	if index < 0 {
		return nil
	}

	var length int = len(internal.Image)

	if  length <= index {
		return nil
	}

	return internal.Image[index]

}

func (receiver *Video) Len() int {
	if nil == receiver {
		return 0
	}

	var internal *gif.GIF
	if nil == internal {
		return 0
	}

	return len(internal.Image)
}

func (receiver *Video) LoopCount() int {
	if nil == receiver {
		return 1
	}

	var internal *gif.GIF
	if nil == internal {
		return 1
	}

	return internal.LoopCount
}
