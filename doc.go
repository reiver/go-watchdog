/*
Package watchdog provides supervisor tree capabilities.

A supervisor tree can help you create software that follows the 'crash-only software' design pattern,
by allowing you to build a system made up of 'microrebootable' and 'microrecoverable' components.

(More is said about the 'crash-only software' design pattern later in this documentation.)

Some might also recognize this library as being similar to facilities one might find in Erlang/OTP or Scala/Akka.

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

Motivation

The motivation for this library is to make it easier to create software that follows the 'crash-only software' design pattern.

Crash-Only Software

Crash-only software can be described as:

"Crash-only programs crash safely and recover quickly.
There is only one way to stop such software—by crashing it—and only one way to bring it up—by initiating recovery.
Crash-only systems are built from crash-only components, and the use of transparent component-level retries hides intra-system component crashes from end users." [1]

One way of thinking of at least part of this is that, many kinds of software failures can be
recovered from  by rebooting the entire software system. However, it is not always necessary to
reboot the entire software system to recover. Sometimes rebooting (small) components within the
software system (instead of rebooting the entire software system) can equally recover from the
failure. The advantage of rebooting some (small) components is that it is orders of magnitude
faster, than rebooting the entire software system. This (increased) speed of recovery is a useful
property for those interested in 'highly available' software system. Also 'high availability' is
a concern for those leaning on the "A"-side of 'CAP theorem'.

There are properties of 'crash-only software' that [1] lists:

Property (a): All important non-volatile state is managed by dedicated state stores.

Property (b): Components have externally enforced boundaries.

Property (c): All interactions between components have a timeout.

Property (d): All resources are leased.

Property (e): Requests are entirely self-describing.

This library does not (and really cannot) enforce any of these properies.
It is up to the person making use of this library to design and architect these properties into their software system.

You can find more information about 'crash-only software' at:

[1] http://www.usenix.org/events/hotos03/tech/full_papers/candea/candea.pdf

[2] http://brooker.co.za/blog/2012/01/22/crash-only.html

[3] http://web.archive.org/web/20060426230247/http://crash.stanford.edu/

Toiler as Crash-Only Software

The Toiler interface tries to embody some of these ideas.
It only contains 2 methods: Terminate and Toil.

The Toiler's Terminate method can be thought of (manually) crashing the Toiler.
(Although the Toiler can crash itself by calling panic().)

The Toiler's Toil method can be thought of 'bringing up' the Toiler.

*/
package watchdog
