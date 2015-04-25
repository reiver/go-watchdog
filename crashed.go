package watchdog


// crashed is used as a message which the crashed method (defined in this file)
// sends to its respective Watchers.
//
// The crashed message contains the Toiler that crashed, due to a panic(), and a
// "done channel" which the crashed method use to make sure the crash is fully
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
// Depending on the watcher, the "appopriate action" could be just restarting
// the Toiler that crashed, or restarting all the Toilers, or something else.
func (dog *wdt) crashed(toiler Toiler) {
	done := make(chan struct{})

	dog.message <- crashed{
		done:done,
		toiler:toiler,
	}

	<-done
}
