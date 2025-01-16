package statusbar

import (
	"github.com/vbauerster/mpb"
)

// Updater is a convenience wrapper around a mpb.Bar which updates its total per
// a passed function.
//
// Not threadsafe.
type Updater struct {
	bar      *mpb.Bar
	getTotal func() int

	previousTotal int
}

func NewUpdater(bar *mpb.Bar, getTotal func() int) *Updater {
	return &Updater{
		bar:      bar,
		getTotal: getTotal,
	}
}

// Update immediately updates the underlying mpb.Bar.
func (u *Updater) Update() {
	newTotal := u.getTotal()
	increment := newTotal - u.previousTotal
	u.previousTotal = newTotal

	u.bar.IncrBy(increment)
}
