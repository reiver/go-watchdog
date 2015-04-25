package watchdog


// watch is used as a message which the Watch method (defined in this file)
// send to its respective Watcher.
//
// The watch message contains the Toiler to be watched and a "done channel"
// which the Watch method uses to make sure the watching is completed before
// returning.
type watch struct {
	toiler Toiler
	done chan struct{}
}


// Watch is 1 of 4 methods needed by wdt to implement the Watcher interface.
func (dog *wdt) Watch(toiler Toiler) {
	done := make(chan struct{})

	dog.message <- watch{
		toiler:toiler,
		done:done,
	}

	<-done
}
