package writer

type Timer struct {
	Time int // sec
}

func NewTimer() *Timer {
	return &Timer{Time: 0}
}

func (t *Timer) Tick() {
	t.Time += 1
}
