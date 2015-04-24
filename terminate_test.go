package watchdog


import "testing"
import "sync"


type testToilerForTerminate struct{
	waitgroup *sync.WaitGroup
}
func (td *testToilerForTerminate) Terminate() {
	td.waitgroup.Done()
}
func (td *testToilerForTerminate) Toil() {}


func TestOneForOneTerminate(t *testing.T) {

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
				toiler := &testToilerForTerminate{waitgroup:&wg}

			// Get the watchdog to watch the toiler.
				watchdog.Watch(toiler)
				wg.Add(1)
		}

	// Terminate.
		watchdog.Terminate()

	// Confirm.
	//
	// The way this confirms is that if this is "hanging", then the Go runtime
	// will panic with the error: "fatal error: all goroutines are asleep - deadlock!".
		wg.Wait()
}


func TestOneForAllTerminate(t *testing.T) {

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
				toiler := &testToilerForTerminate{waitgroup:&wg}

			// Get the watchdog to watch the toiler.
				watchdog.Watch(toiler)
				wg.Add(1)
		}

	// Terminate.
		watchdog.Terminate()

	// Confirm.
	//
	// The way this confirms is that if this is "hanging", then the Go runtime
	// will panic with the error: "fatal error: all goroutines are asleep - deadlock!".
		wg.Wait()
}
