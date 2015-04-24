package watchdog


// crashed is used as a message which the crashed methods (defined in this file)
// send to their respective Watchers.
//
// The crashed message contains the Toiler that crashed, due to a panic(), and a
// "done channel" which the crashed methods use to make sure the crash is fully
// handled before returning.
type crashed struct{
	done chan struct{}
	toiler Toiler
}


// crashed is an internal method used to tell a Watcher that a Toiler it is
// watching crashed, due to a panic().
//
// This crashed method is used by its Watcher in a closure func it passes to
// the watchedToil func. The watchedToil method is able to detect and capture
// a panic() from the Toiler's Toil method. And if that happens, this crashed
// method sends a crashed message to the Watcher, which includes the Toiler
// which crashed. The Watcher then takes the appropriate action.
//
// For this Watcher -- the "one for one" watcher -- the appropriate action
// is to try to restart just the Toiler that crashed, due to a panic(), and
// leave all the other Toilers alone.
func (dog *oneForOne) crashed(toiler Toiler) {
	done := make(chan struct{})

	dog.message <- crashed{
		done:done,
		toiler:toiler,
	}

	<-done
}

// crashed is an internal method used to tell a Watcher that a Toiler it is
// watching crashed, due to a panic().
//
// This crashed method is used by its Watcher in a closure func it passes to
// the watchedToil func. The watchedToil method is able to detect and capture
// a panic() from the Toiler's Toil method. And if that happens, this crashed
// method sends a crashed message to the Watcher, which includes the Toiler
// which crashed. The Watcher then takes the appropriate action.
//
// For this Watcher -- the "one for all" watcher -- the appropriate action
// is to try to restart all the Toilers it is watching.
func (dog *oneForAll) crashed(toiler Toiler) {
	done := make(chan struct{})

	dog.message <- crashed{
		done:done,
		toiler:toiler,
	}

	<-done
}
