package domain

type Settings struct {
	Difficulty int
	Mode       int
	Duration   int
}

type Quote struct {
	Text   string
	Author string
}

func (s Settings) GetDuration() int {
	return (s.Duration + 1) * 15
}
