package watchdog


// watch is used as a message which the Watch methods (defined in this file)
// send to their respective Watchers.
//
// The watch message contains the Toiler to be watched and a "done channel"
// which the Watch methods use to make sure the watching is completed before
// returning.
type watch struct {
	toiler Toiler
	done chan struct{}
}


// Watch is 1 of 3 methods needed by oneForOne to implement the Watcher interface.
func (dog *oneForOne) Watch(toiler Toiler) {
	done := make(chan struct{})

	dog.message <- watch{
		toiler:toiler,
		done:done,
	}

	<-done
}

// Watch is 1 of 3 methods needed by oneForAll to implement the Watcher interface.
func (dog *oneForAll) Watch(toiler Toiler) {
	done := make(chan struct{})

	dog.message <- watch{
		toiler:toiler,
		done:done,
	}

	<-done
}
