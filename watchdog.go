package watchdog


type wdt struct {
	crashedStrategyFunc func(WatchedToiler)
	terminate chan terminate
	message   chan interface{}
	toilers []Toiler
}


func newWatchDog(crashedStrategyFn func(WatchedToiler)) Watcher {
	terminate := make(chan terminate)
	message   := make(chan interface{})
	toilers   := make([]Toiler, 0, defaultLengthOfToilersSlice)

	watcher := wdt{
		crashedStrategyFunc:crashedStrategyFn,
		terminate:terminate,
		message:message,
		toilers:toilers,
	}

	go watcher.watchover()

	return &watcher
}
