package util

import "time"

type ProcessDuration struct {
	start time.Time
	end   time.Time
}

func (p *ProcessDuration) SetStart(t time.Time) {
	p.start = t
}
func (p *ProcessDuration) SetEnd(t time.Time) {
	p.end = t
}
func (p *ProcessDuration) GetDuration() time.Duration {
	duration := p.end.Sub(p.start)
	return duration
}
