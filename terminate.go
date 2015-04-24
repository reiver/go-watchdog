package watchdog


// terminate is used as a message which the Terminate methods (defined in this file)
// send to their respective Watchers.
//
// The terminate message contains a "done channel" which the Terminate methods
// use to make sure the terminating is completed before returning.
type terminate struct {
        done chan struct{}
}


// Terminate is 1 of 3 methods needed by oneForOne to implement the Watcher interface.
func (dog *oneForOne) Terminate() {
	done := make(chan struct{})

	dog.terminate <- terminate{
		done:done,
	}

	<-done
}

// Terminate is 1 of 3 methods needed by oneForAll to implement the Watcher interface.
func (dog *oneForAll) Terminate() {
	done := make(chan struct{})

	dog.terminate <- terminate{
		done:done,
	}

	<-done
}
