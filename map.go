package watchdog


// Map is 1 of 4 methods needed by wdt to implement the Watcher interface.
func (dog *wdt) Map(fn func(WatchedToiler)) {

	for _,toiler := range dog.toilers {
		fn( newWatchedToiler(dog, toiler) )

	}
}
