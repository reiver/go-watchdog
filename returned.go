package watchdog


// returned is used as a message which the returned method (defined in this file)
// sends to its respective Watchers.
//
// The returned message contains the Toiler that returned, due to a panic(), and a
// "done channel" which the returned method use to make sure the return is fully
// handled before returning.
type returned struct{
	done chan struct{}
	toiler Toiler
}


// returned is an internal method used to tell a Watcher that a Toiler it is
// watching returned, due to Toiler's Toil method either calling return or
// simply running to the end of the func.
//
// This returned method is used by its Watcher in a closure func it passes to
// the watchedToil func. The watchedToil method is able to detect and capture
// a panic() from the Toiler's Toil method. And if that happens, this returned
// method sends a returned message to the Watcher, which includes the Toiler
// which returned. The Watcher then takes the appropriate action.
func (dog *wdt) returned(toiler Toiler) {
	done := make(chan struct{})

	dog.message <- returned{
		done:done,
		toiler:toiler,
	}

	<-done
}
