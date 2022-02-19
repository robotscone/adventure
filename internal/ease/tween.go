package ease

import (
	"time"
)

type Direction byte

const (
	Forward Direction = iota
	Backward
	Opposite
)

type TweenHook func(t *Tween)

type Tween struct {
	elapsed       float64
	duration      float64
	easing        Func
	From          float64
	To            float64
	Value         float64
	isReversed    bool
	isInverted    bool
	isStarted     bool
	isFinished    bool
	isPaused      bool
	startedFuncs  []TweenHook
	finishedFuncs []TweenHook
}

func NewTween(from, to float64, duration time.Duration, easing Func) *Tween {
	var t Tween

	t.SetDuration(duration)
	t.SetEasing(easing)

	t.From = from
	t.To = to
	t.Value = t.From

	return &t
}

func (t *Tween) Reset() {
	t.elapsed = 0
	t.Value = t.From
	t.isReversed = false
	t.isInverted = false
	t.isStarted = false
	t.isFinished = false
}

func (t *Tween) Play() {
	t.isPaused = false
}

func (t *Tween) Pause() {
	t.isPaused = true
}

func (t *Tween) Update(delta float64) {
	if t.isPaused || t.isFinished {
		return
	}

	if !t.isStarted {
		t.isStarted = true

		for _, f := range t.startedFuncs {
			f(t)
		}
	}

	if t.isReversed {
		t.elapsed -= delta

		if t.elapsed <= 0 {
			t.elapsed = 0
			t.isFinished = true
		}
	} else {
		t.elapsed += delta

		if t.elapsed >= t.duration {
			t.elapsed = t.duration
			t.isFinished = true
		}
	}

	if t.isInverted {
		t.Value = To(t.elapsed/t.duration, t.To, t.From, t.easing)
	} else {
		t.Value = To(t.elapsed/t.duration, t.From, t.To, t.easing)
	}

	if t.isFinished {
		t.isStarted = false

		for _, f := range t.finishedFuncs {
			f(t)
		}
	}
}

func (t *Tween) Direction() Direction {
	if !t.isInverted && !t.isReversed || t.isInverted && t.isReversed {
		return Forward
	}

	return Backward
}

func (t *Tween) SetDirection(dir Direction) {
	if dir == Opposite {
		if t.Direction() == Forward {
			dir = Backward
		} else {
			dir = Forward
		}
	}

	switch dir {
	case Forward:
		if t.Direction() == Backward {
			if t.isFinished {
				t.setInverted(false)
			} else {
				t.setReversed(!t.isReversed)
			}
		}
	case Backward:
		if t.Direction() == Forward {
			if t.isFinished {
				t.setInverted(true)
			} else {
				t.setReversed(!t.isReversed)
			}
		}
	}
}

func (t *Tween) setInverted(value bool) {
	// If the inversion operation actually changes anything then we need
	// to recalculate the elapsed time to make sure the value doesn't
	// suddenly jump to a completely different one
	if t.isInverted != value {
		t.elapsed = t.duration - t.elapsed
	}

	t.isInverted = value
	t.isReversed = false

	if t.isFinished {
		t.isFinished = false
	}
}

func (t *Tween) setReversed(value bool) {
	t.isReversed = value

	if t.isFinished {
		t.isFinished = false
	}
}

func (t *Tween) Duration() time.Duration {
	return time.Duration(t.duration * float64(time.Second))
}

func (t *Tween) SetDuration(duration time.Duration) {
	t.duration = duration.Seconds()
}

func (t *Tween) Easing() Func {
	return t.easing
}

func (t *Tween) SetEasing(easing Func) {
	if easing == nil {
		easing = Linear
	}

	t.easing = easing
}

func (t *Tween) OnStarted(f TweenHook) {
	t.startedFuncs = append(t.startedFuncs, f)
}

func (t *Tween) OnFinished(f TweenHook) {
	t.finishedFuncs = append(t.finishedFuncs, f)
}
