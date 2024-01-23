package timer

import "time"

type Func func()

type task struct {
	elapsed  float64
	interval float64
	limit    int
	runs     int
	f        Func
}

type Timer struct {
	afters []*task
	everys []*task
}

func New() *Timer {
	return &Timer{}
}

func (t *Timer) After(interval time.Duration, f Func) {
	t.afters = append(t.afters, &task{
		interval: interval.Seconds(),
		f:        f,
	})
}

func (t *Timer) Every(interval time.Duration, limit int, f Func) {
	t.everys = append(t.everys, &task{
		interval: interval.Seconds(),
		limit:    limit,
		f:        f,
	})
}

func (t *Timer) Clear() {
	t.afters = make([]*task, 0)
	t.everys = make([]*task, 0)
}

func (t *Timer) Reset() {
	for _, task := range t.afters {
		task.elapsed = 0
		task.runs = 0
	}

	for _, task := range t.everys {
		task.elapsed = 0
		task.runs = 0
	}
}

func (t *Timer) Update(delta float64) {
	afters := t.afters[:0]
	for _, task := range t.afters {
		task.elapsed += delta

		if task.elapsed > task.interval {
			task.f()
		} else {
			afters = append(afters, task)
		}
	}
	t.afters = afters

	everys := t.everys[:0]
	for _, task := range t.everys {
		task.elapsed += delta

		for task.elapsed >= task.interval {
			if task.limit > 0 && task.runs >= task.limit {
				break
			}

			task.runs++
			task.f()

			task.elapsed -= task.interval
		}

		if task.limit <= 0 || task.runs < task.limit {
			everys = append(everys, task)
		}
	}
	t.everys = everys
}
