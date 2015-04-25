package watchdog


// terminate is used as a message which the Terminate method (defined in this file)
// sends to its respective Watcher.
//
// The terminate message contains a "done channel" which the Terminate method
// uses to make sure the terminating is completed before returning.
type terminate struct {
        done chan struct{}
}


// Terminate is 1 of 4 methods needed by wdt to implement the Watcher interface.
// Terminate is also 1 of 2 methods needed by wdt to implement the Toiler interface.
func (dog *wdt) Terminate() {
	done := make(chan struct{})

	dog.terminate <- terminate{
		done:done,
	}

	<-done
}
