package watchdog


// Water is an interface that groups together Terminate, Toil and Watch methods together.
//
// A Watcher is the basic abstraction of something that watches. (I.e., the 'watchdog'.)
//
// Notice that every Watcher is also a Toiler, since every Watcher has a Terminate and Toil method.
// This allows one Watcher to watch another Watcher.
type Watcher interface {
	Terminate()
	Toil()
	Watch(Toiler)
}
