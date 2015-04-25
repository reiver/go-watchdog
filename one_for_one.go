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

	return newWatchDog(oneForOneCrashedStrategy)
}


// oneForOneCrashedStrategy implements the strategy to use to handle crashing of a Toiler
// for the "one for one" Watcher.
//
// Only the Toiler that crashed is restarted. The rest are left alone.
func oneForOneCrashedStrategy(watchedToiler WatchedToiler) {
	watchedToiler.Terminate()
	watchedToiler.Toil()
}
