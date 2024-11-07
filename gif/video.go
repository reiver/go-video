package gif

import (
	"image"
	"image/draw"
	gogif "image/gif"
	"io"
	"time"

	"github.com/reiver/go-video"
)

type Video struct {
	internal *gogif.GIF
}

var _ video.Video = &Video{}

func NewVideo(reader io.Reader) (*Video, error) {
	if nil == reader {
		return nil, errNilReader
	}

	anigif, err := gogif.DecodeAll(reader)
	if nil != err {
		return nil, err
	}

	{
		var video = Video{
			internal:anigif,
		}

		return &video, nil
	}
}

func (receiver *Video) Delay(index int) time.Duration {
	var noDuration time.Duration

	if nil == receiver {
		return noDuration
	}

	var internal *gogif.GIF
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

	var internal *gogif.GIF
	if nil == internal {
		return nada
	}

	{
		var disposal byte = internal.Disposal[index-1]

		switch disposal {
		case gogif.DisposalNone:
			return draw.Over
		case gogif.DisposalBackground:
			return draw.Src
		case gogif.DisposalPrevious:
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

	var internal *gogif.GIF
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

	var internal *gogif.GIF
	if nil == internal {
		return 0
	}

	return len(internal.Image)
}

func (receiver *Video) LoopCount() int {
	if nil == receiver {
		return 1
	}

	var internal *gogif.GIF
	if nil == internal {
		return 1
	}

	return internal.LoopCount
}

func (receiver *Video) Size() (width int, height int) {
	if nil == receiver {
		return 0,0
	}

	var internal *gogif.GIF
	if nil == internal {
		return 0,0
	}

	{
		var cfg = internal.Config

		width, height = cfg.Width, cfg.Height

		if 0 != width || 0 != height {
			return width, height
		}
	}

	if receiver.Len() <= 0 {
		return 0,0
	}

	{
		var img image.Image = receiver.Image(0)
		if nil == img {
			return 0,0
		}

		var rect image.Rectangle = img.Bounds()

		width = rect.Max.X - rect.Min.X
		height = rect.Max.Y - rect.Min.Y

		return width, height
	}
}
