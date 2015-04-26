package watchdog


import "testing"
import "sync"


type testToilerForToil struct{
	waitgroup *sync.WaitGroup
}
func (td *testToilerForToil) Terminate() {}
func (td *testToilerForToil) Toil() {
	td.waitgroup.Done()
}


func TestOneForOneToil(t *testing.T) {

	// Create a watchdog.
		watchdog := NewOneForOne()

	// Create a wait group.
		var wg sync.WaitGroup

	// Create 10 toilers.
		const numChildToilers = 10
		for i:=1; i<=numChildToilers; i++ {

			// Create dummy toiler.
			//
			// This is what we'll send to .Watch()
				toiler := &testToilerForToil{waitgroup:&wg}

			// Get the watchdog to watch the toiler.
				wg.Add(1)
				watchdog.Watch(toiler)
		}

	// Toil
		go watchdog.Toil()

	// Confirm.
	//
	// The way this confirms is that if this is "hanging", then the Go runtime
	// will panic with the error: "fatal error: all goroutines are asleep - deadlock!".
		wg.Wait()


	// Add another Toiler.
	//
	// We've already called the Watcher's Toil method.
	// We want to make sure that if we watch another
	// Toiler after the the Watcher's Toil method is
	// called that this Toiler has its Toil method
	// called as part of the Watch.
		toiler := &testToilerForToil{waitgroup:&wg}
		wg.Add(1)
		watchdog.Watch(toiler)

	// Confirm.
	//
	// The way this confirms is that if this is "hanging", then the Go runtime
	// will panic with the error: "fatal error: all goroutines are asleep - deadlock!".
		wg.Wait()
}


func TestOneForAllToil(t *testing.T) {

	// Create a watchdog.
		watchdog := NewOneForAll()

	// Create a wait group.
		var wg sync.WaitGroup

	// Create 10 toilers.
		const numChildToilers = 10
		for i:=1; i<=numChildToilers; i++ {

			// Create dummy toiler.
			//
			// This is what we'll send to .Watch()
				toiler := &testToilerForToil{waitgroup:&wg}

			// Get the watchdog to watch the toiler.
				wg.Add(1)
				watchdog.Watch(toiler)
		}

	// Toil
		go watchdog.Toil()

	// Confirm.
	//
	// The way this confirms is that if this is "hanging", then the Go runtime
	// will panic with the error: "fatal error: all goroutines are asleep - deadlock!".
		wg.Wait()


	// Add another Toiler.
	//
	// We've already called the Watcher's Toil method.
	// We want to make sure that if we watch another
	// Toiler after the the Watcher's Toil method is
	// called that this Toiler has its Toil method
	// called as part of the Watch.
		toiler := &testToilerForToil{waitgroup:&wg}
		wg.Add(1)
		watchdog.Watch(toiler)

	// Confirm.
	//
	// The way this confirms is that if this is "hanging", then the Go runtime
	// will panic with the error: "fatal error: all goroutines are asleep - deadlock!".
		wg.Wait()
}
