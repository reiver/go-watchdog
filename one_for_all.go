package watchdog


// NewOneForAll creates a new "one for all" Watcher (i.e., 'watchdog') that (potentially) watches a number of Toilers
// and if one of them fails then all the Toilers will be restarted, 
// by calling the Toiler's Terminate method and then its Toil method.
func NewOneForAll() Watcher {

	return newWatchDog(oneForAllCrashedStrategy)
}


// oneForAllCrashedStrategy implements the strategy to use to handle crashing of a Toiler
// for the "one for all" Watcher.
//
// All the Toilers are restarted.
func oneForAllCrashedStrategy(watchedToiler WatchedToiler) {
	watcher := watchedToiler.Watcher()

	watcher.Map(func(watchedToiler WatchedToiler){
		watchedToiler.Terminate()
		watchedToiler.Toil()
	})
}

