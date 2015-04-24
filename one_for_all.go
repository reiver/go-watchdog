package watchdog


type oneForAll struct {
	terminate chan terminate
	message   chan interface{}
	toilers []Toiler
}


// NewOneForAll creates a new "one for all" Watcher (i.e., 'watchdog') that (potentially) watches a number of Toilers
// and if one of them fails then all the Toilers will be restarted, 
// by calling the Toiler's Terminate method and then its Toil method.
func NewOneForAll() Watcher {
	terminate := make(chan terminate)
	message   := make(chan interface{})
	toilers   := make([]Toiler, 0, defaultLengthOfToilersSlice)

	watcher := oneForAll{
		terminate:terminate,
		message:message,
		toilers:toilers,
	}

	go watcher.watchover()

	return &watcher
}
