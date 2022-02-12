package ease

import (
	"time"
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
	isFinished    bool
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
	t.isFinished = false
}

func (t *Tween) Update(delta float64) {
	if t.isFinished {
		return
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
		for _, f := range t.finishedFuncs {
			f(t)
		}
	}
}

func (t *Tween) Invert(value bool) {
	// If the inversion operation actually changes anything then we need
	// to recalculate the elapsed time to make sure the value doesn't
	// suddenly jump to a completely different one
	if t.isInverted != value {
		t.elapsed = t.duration - t.elapsed
	}

	t.isInverted = value
	t.isFinished = false

	t.Reverse(false)
}

func (t *Tween) Reverse(value bool) {
	t.isReversed = value
	t.isFinished = false
}

func (t *Tween) InvertOrReverse(value bool) {
	if t.isFinished {
		t.Invert(value)
	} else {
		t.Reverse(!t.isReversed)
	}
}

func (t *Tween) IsForward() bool {
	return !t.isInverted && !t.isReversed || t.isInverted && t.isReversed
}

func (t *Tween) IsBackward() bool {
	return !t.IsForward()
}

func (t *Tween) SetDuration(duration time.Duration) {
	t.duration = duration.Seconds()
}

func (t *Tween) SetEasing(easing Func) {
	if easing == nil {
		easing = Linear
	}

	t.easing = easing
}

func (t *Tween) OnFinished(f TweenHook) {
	t.finishedFuncs = append(t.finishedFuncs, f)
}
