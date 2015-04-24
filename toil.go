package watchdog


// toil is used as a message which the Toil methods (defined in this file)
// send to their respective Watchers.
//
// The toil message contains a "done channel" which the Toil methods
// use to make sure the toiling is completed before returning.
type toil struct {
	done chan struct{}
}


// Toil is 1 of 3 methods needed by oneForOne to implement the Watcher interface.
// Toil is also 1 of 2 methods needed by oneForOne to implement the Toiler interface.
func (dog *oneForOne) Toil() {
	done := make(chan struct{})

	dog.message <- toil{
		done:done,
	}

	<-done
}

// Toil is 1 of 3 methods needed by oneForAll to implement the Watcher interface.
// Toil is also 1 of 2 methods needed by oneForAll to implement the Toiler interface.
func (dog *oneForAll) Toil() {
	done := make(chan struct{})

	dog.message <- toil{
		done:done,
	}

	<-done
}


// watchedToil "runs" a Toiler in such a way that it can detect and
// capture a panic() from the Toiler's Toil method and, through the
// use of the crashFn func. notify the Watcher about the panic().
//
// Each Watcher uses this watchedToil func to detect and capture a
// panic() from a Toiler's Toil method. The Watcher passes a
// special crashFn (closure) func to the watchedToil func so that
// it can be notified when a Toiler's Toil method has crashes, and
// it (the Watcher) should take the appropriate action.
//
// The crashFn func the Watcher passes to the watchedToil func is
// a closure that invokes the Watcher's crashed method with the
// Toiler as the parameter.
func watchedToil(toiler Toiler, crashFn func()) {

	go func(){
		defer func() {
			if r := recover(); nil != r {
				crashFn()
			}
		}()

		toiler.Toil()
	}()
}
