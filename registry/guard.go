package registry

import "sync"

const (
	_percentThreshold float64 = 0.85
)

// Guard count the renew of all operations for self protection
type Guard struct {
	expPerMin    int64
	expThreshold int64
	facInMin     int64
	facLastMin   int64
	lock         sync.RWMutex
}

func (g *Guard) incrExp() {
	g.lock.Lock()
	g.expPerMin = g.expPerMin + 2
	g.expThreshold = int64(float64(g.expPerMin) * _percentThreshold)
	g.lock.Unlock()
}
