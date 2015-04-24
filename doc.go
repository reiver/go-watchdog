/*
Package watchdog provides supervisor tree capabilities, similar to those one might find in Erlang or Scala/Akka.

There are 2 main interfaces in this package: Toiler and Watcher

Toiler

A Toiler is an abstraction (an interface) that represents something that toils (i.e., does work).

Anything that you would want to be watched by a supervisor tree you would turn into a Toiler by implementing
Toiler's 2 methods: Terminate and Toil.

For example, here is a simple (and rather useless) toiler.

	type Outputter struct{
		inputCh  chan int
	}
	
	func NewOutputter() *Outputter {
	
		inputCh  := make(chan int)
	
		outputter := Outputter{
			inputCh:inputCh,
		}
	
		return &outputter
	}
	
	func (o *Outputter) Send(n int) {
		o.inputCh <- n
	}
	
	func (o *Outputter) Terminate() {
		// Nothing here.
	}
	
	func (o *Outputter) Toil() {
		for n := range o.inputCh {
			fmt.Printf("Received a %d.\n", n)
		}
	}

Notice that the example Toiler has both Terminate and Toil methods implemented.

Watcher

There are 2 types of built-in watchers: "one for all" and "one for one".

One For All

A "one for one" watcher is created with the NewOneForAll() func. As in:

	watcher := watchdog.NewOneForAll()

A "one for all" watcher can (potentially) watch many toilers.

	watcher.Watch(toiler1)
	watcher.Watch(toiler2)
	watcher.Watch(toiler3)
	// Etc

If one of the watched toilers dies, by throwing a panic(), then the "one for all" watcher will
restart all the toilers.

Conceptually, it does this by first calling the Terminate methods on all the toilers
and then calling the Toil method on all the toilers.

One For One

A "one for one" watcher is created with the NewOneForOne() func. As in:

	watcher := watchdog.NewOneForOne()

A "one for one" watcher can (potentially) watch many toilers.

	watcher.Watch(toiler1)
	watcher.Watch(toiler2)
	watcher.Watch(toiler3)
	// Etc

If one of the watched toilers dies, by throwing a panic(), then the "one for one" watcher will
restart only that toiler. (The rest of the toilers will be left alone.)

Nesting

A Watcher is also a Toiler. So you can nest watchers. For example:

	root := watchdog.NewOneForOne()
	
	
	left := watchdog.NewOneForMany()
	root.Watch(left) // ← One watcher just watched another watcher!
	
	left.Watch(toiler1)
	left.Watch(toiler2)
	
	
	right := watchdog.NewOneForOne()
	root.Watch(right) // ← We made one watcher watch another watcher again!
	
	right.Watch(toiler3)
	right.Watch(toiler4)
	right.Watch(toiler5)

*/
package watchdog
