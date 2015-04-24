package watchdog


type oneForOne struct {
	terminate chan terminate
	message   chan interface{}
	toilers []Toiler
}


// NewOneForOne creates a new "one for one" Watcher (i.e., 'watchdog') that (potentially) watches a number of Toilers
// and if one of them fails then only that one Toiler will be restarted,
// by calling the Toiler's Terminate method and then its Toil method.
// (The other watched Toilers will be left alone.)
func NewOneForOne() Watcher {
	terminate := make(chan terminate)
	message   := make(chan interface{})
	toilers   := make([]Toiler, 0, defaultLengthOfToilersSlice)

	watcher := oneForOne{
		terminate:terminate,
		message:message,
		toilers:toilers,
	}

	go watcher.watchover()

	return &watcher
}
