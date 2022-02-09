package gfx

import (
	"time"

	"github.com/veandco/go-sdl2/sdl"
)

type Flip byte

const (
	FlipNone       Flip = 0
	FlipHorizontal Flip = 1
	FlipVertical   Flip = 2
)

type Frame struct {
	src  sdl.Rect
	flip Flip
}

type Animation struct {
	elapsed  float64
	period   float64
	duration time.Duration
	idx      int
	frame    *Frame
	frames   []*Frame
}

func (a *Animation) AddFrame(x, y, w, h int, flip Flip) {
	a.frames = append(a.frames, &Frame{
		src: sdl.Rect{
			X: int32(x),
			Y: int32(y),
			W: int32(w),
			H: int32(h),
		},
		flip: flip,
	})

	if a.frame == nil {
		a.frame = a.frames[0]
	}

	a.SetDuration(a.duration)
}

func (a *Animation) SetDuration(duration time.Duration) {
	a.duration = duration
	a.period = a.duration.Seconds() / float64(len(a.frames))
}

func (a *Animation) Reset() {
	a.elapsed = 0.0
	a.idx = 0

	if len(a.frames) > 0 {
		a.frame = a.frames[0]
	} else {
		a.frame = nil
	}
}

func (a *Animation) Step(delta float64) {
	a.elapsed += delta

	if a.period <= 0 {
		return
	}

	for a.elapsed >= a.period {
		a.idx++
		a.elapsed -= a.period

		a.frame = a.frames[a.idx%len(a.frames)]
	}
}