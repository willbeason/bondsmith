package statusbar

import (
	"time"
)

// AutoUpdater is a convenience wrapper for Updater, automatically updating the
// underlying mpb.Bar per a passed tick channel.
//
// Not threadsafe.
type AutoUpdater struct {
	*Updater
}

// NewAutoUpdater instantiates an AutoUpdater.
// Does not start updating the bar on instantiation; use Tick() to do so.
func NewAutoUpdater(updater *Updater) *AutoUpdater {
	return &AutoUpdater{
		Updater: updater,
	}
}

// Tick signals the status bar to begin updating per the passed ticker.
// Blocks until the passed ticker is closed.
func (sb *AutoUpdater) Tick(ticker <-chan time.Time) {
	for range ticker {
		sb.Update()
	}
}
