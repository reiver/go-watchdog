package watchdog


// Water is an interface that groups together Terminate, Toil and Watcher methods together.
//
// A WatchedToiler is a Toiler (as it implements both the Terminate method and Toil method)
// that is "binded" to a Watcher.
//
// This "binding" is expressed in 2 ways. #1: The WatchedToiler has a Watcher method that
// will return the Watcher. And #2: The Toil method is run the underlying Toiler's Toil
// method is a special way so that the Watcher is watching it (for panic()s).
type WatchedToiler interface {
	Terminate()
	Toil()
	Watcher() Watcher
}



type watchedToiler struct {
	toiler Toiler
	watcher Watcher
}

func newWatchedToiler(watcher Watcher, toiler Toiler) WatchedToiler {

	wt := watchedToiler{
		toiler:toiler,
		watcher:watcher,
	}

	return &wt
}

func (wt *watchedToiler) Terminate() {
	wt.toiler.Terminate()
}

func (wt *watchedToiler) Toil() {
	watchedToil(wt.toiler, func(exception interface{}){
		wt.Watcher().(*wdt).crashed(wt.toiler)
	})
}

func (wt *watchedToiler) Watcher() Watcher {
	return wt.watcher
}
