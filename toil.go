package watchdog


// toil is used as a message which the Toil method (defined in this file)
// sends to its respective Watcher.
//
// The toil message contains a "done channel" which the Toil methods
// use to make sure the toiling is completed before returning.
type toil struct {
	done chan struct{}
}


// Toil is 1 of 4 methods needed by wdt to implement the Watcher interface.
// Toil is also 1 of 2 methods needed by wdt to implement the Toiler interface.
func (dog *wdt) Toil() {
	done := make(chan struct{})

	dog.message <- toil{
		done:done,
	}

	<-done
}


// watchedToil "runs" a Toiler in such a way that it can detect and
// capture a panic() from the Toiler's Toil method and, through the
// use of the crashedFn func notify the Watcher about the panic().
//
// Each Watcher uses this watchedToil func to detect and capture a
// panic() from a Toiler's Toil method. The Watcher passes a
// special crashedFn (closure) func to the watchedToil func so that
// it can be notified when a Toiler's Toil method has crashes, and
// it (the Watcher) should take the appropriate action.
//
// The crashedFn func the Watcher passes to the watchedToil func is
// a closure that invokes the Watcher's crashed method with the
// Toiler as the parameter.
//
// watchedToil is also able to detect if the Toiler's Toil method
// has returned and also notifies the Watcher about that too.
//
// Each Watcher uses this watchedToil func to detect if the Toiler's
// Toil method has returned. The Watcher passes a secial returnFn
// (closure) func to this watchedToil func so that it can be
// notified when the Toiler's Toil methods has returned, and it
// (the Watcher) should take the appropriate action.
//
// The returnedFn func the Watcher passes to the watchedToil func
// is a closure that invokes the Watcher's returned method with the
// Toiler as the parameter.
func watchedToil(toiler Toiler, crashedFn func(interface{}), returnedFn func()) {

	go func(){
		defer func() {
			if r := recover(); nil != r {
				crashedFn(r)
			}
		}()

		toiler.Toil()

		returnedFn()
	}()
}
