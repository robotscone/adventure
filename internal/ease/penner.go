// Package ease contains simplified versions of the equations by Robert Penner
// and some basic tweening functionality.
//
// You can find more information about the easing functions at:
//
//	http://robertpenner.com/easing/
//
// The original functions where written using the following parameters:
//
//	t = current time in a tween
//	b = the original beginning value
//	c = the change between b and the destination value (dst - b)
//	d = duration of the tween
//	s = overshoot amount (optional value in "back" functions)
//	a = amplitude (optional value in "elastic" functions)
//	p = period (optional value in "elastic" functions)
//
// Since these functions are used to interpolate a value between two points, if
// we assume that...
//
//	b = 0
//	c = 1
//	d = 1
//
// ...then the equations can be simplified to take a normalised t value
// Because of this the t value in these function is expected to be
// in the range [0, 1] and the returned value will also be normalised
package ease

import "math"

type Func func(t float64) float64

func To(t, src, dst float64, easing Func) float64 {
	return (dst-src)*easing(t) + src
}

func Linear(t float64) float64 {
	return t
}

func QuadIn(t float64) float64 {
	return t * t
}

func QuadOut(t float64) float64 {
	return -1 * t * (t - 2)
}

func QuadInOut(t float64) float64 {
	t *= 2

	if t < 1 {
		return 0.5 * t * t
	}

	t--

	return -0.5 * (t*(t-2) - 1)
}

func CubicIn(t float64) float64 {
	return t * t * t
}

func CubicOut(t float64) float64 {
	t--

	return t*t*t + 1
}

func CubicInOut(t float64) float64 {
	t *= 2

	if t < 1 {
		return 0.5 * t * t * t
	}

	t -= 2

	return 0.5 * (t*t*t + 2)
}

func QuartIn(t float64) float64 {
	return t * t * t * t
}

func QuartOut(t float64) float64 {
	t--

	return -1 * (t*t*t*t - 1)
}

func QuartInOut(t float64) float64 {
	t *= 2

	if t < 1 {
		return 0.5 * t * t * t * t
	}

	t -= 2

	return -0.5 * (t*t*t*t - 2)
}

func QuintIn(t float64) float64 {
	return t * t * t * t * t
}

func QuintOut(t float64) float64 {
	t--

	return t*t*t*t*t + 1
}

func QuintInOut(t float64) float64 {
	t *= 2

	if t < 1 {
		return 0.5 * t * t * t * t * t
	}

	t -= 2

	return 0.5 * (t*t*t*t*t + 2)
}

func SineIn(t float64) float64 {
	return -1*math.Cos(t*(math.Pi/2)) + 1
}

func SineOut(t float64) float64 {
	return math.Sin(t * (math.Pi / 2))
}

func SineInOut(t float64) float64 {
	return -0.5 * (math.Cos(math.Pi*t) - 1)
}

func ExpoIn(t float64) float64 {
	if t == 0 {
		return 0
	}

	return math.Pow(2, 10*(t-1))
}

func ExpoOut(t float64) float64 {
	if t == 1 {
		return 1
	}

	return -math.Pow(2, -10*t) + 1
}

func ExpoInOut(t float64) float64 {
	if t == 0 {
		return 0
	}

	if t == 1 {
		return 1
	}

	t *= 2

	if t < 1 {
		return 0.5 * math.Pow(2, 10*(t-1))
	}

	t--

	return 0.5 * (-math.Pow(2, -10*t) + 2)
}

func CircIn(t float64) float64 {
	return -1 * (math.Sqrt(1-t*t) - 1)
}

func CircOut(t float64) float64 {
	t--

	return math.Sqrt(1 - t*t)
}

func CircInOut(t float64) float64 {
	t *= 2

	if t < 1 {
		return -0.5 * (math.Sqrt(1-t*t) - 1)
	}

	t -= 2

	return 0.5 * (math.Sqrt(1-t*t) + 1)
}

func ElasticIn(t float64) float64 {
	if t == 0 {
		return 0
	}

	var a float64
	var p float64

	if t == 1 {
		return 1
	}

	if p == 0 {
		p = 0.3
	}

	var s float64

	if a < 1 {
		a = 1
		s = p / 4
	} else {
		s = p / (2 * math.Pi) * math.Asin(1/a)
	}

	t--

	return -(a * math.Pow(2, 10*t) * math.Sin((t*1-s)*(2*math.Pi)/p))
}

func ElasticOut(t float64) float64 {
	if t == 0 {
		return 0
	}

	var a float64
	var p float64

	if t == 1 {
		return 1
	}

	if p == 0 {
		p = 0.3
	}

	var s float64

	if a < 1 {
		a = 1
		s = p / 4
	} else {
		s = p / (2 * math.Pi) * math.Asin(1/a)
	}

	return a*math.Pow(2, -10*t)*math.Sin((t*1-s)*(2*math.Pi)/p) + 1
}

func ElasticInOut(t float64) float64 {
	if t == 0 {
		return 0
	}

	var a float64
	var p float64

	t *= 2

	if t == 2 {
		return 1
	}

	if p == 0 {
		p = 0.3 * 1.5
	}

	var s float64

	if a < 1 {
		a = 1
		s = p / 4
	} else {
		s = p / (2 * math.Pi) * math.Asin(1/a)
	}

	if t < 1 {
		t--

		return -0.5 * (a * math.Pow(2, 10*t) * math.Sin((t*1-s)*(2*math.Pi)/p))
	}

	t--

	return a*math.Pow(2, -10*t)*math.Sin((t*1-s)*(2*math.Pi)/p)*0.5 + 1
}

func BackIn(t float64) float64 {
	s := 1.70158

	return t * t * ((s+1)*t - s)
}

func BackOut(t float64) float64 {
	t--
	s := 1.70158

	return t*t*((s+1)*t+s) + 1
}

func BackInOut(t float64) float64 {
	t *= 2
	s := 1.70158
	s *= 1.525

	if t < 1 {
		return 0.5 * (t * t * ((s+1)*t - s))
	}

	t -= 2

	return 0.5 * (t*t*((s+1)*t+s) + 2)
}

func BounceIn(t float64) float64 {
	return 1 - BounceOut(1-t)
}

func BounceOut(t float64) float64 {
	const n1 = 7.5625
	const d1 = 2.75

	if t < 1/d1 {
		return n1 * t * t
	} else if t < 2/d1 {
		t -= 1.5 / d1

		return n1*t*t + 0.75
	} else if t < 2.5/d1 {
		t -= 2.25 / d1

		return n1*t*t + 0.9375
	}

	t -= 2.625 / d1

	return n1*t*t + 0.984375
}

func BounceInOut(t float64) float64 {
	if t < 0.5 {
		return BounceIn(t*2) * 0.5
	}

	return BounceOut(t*2-1)*0.5 + 0.5
}
