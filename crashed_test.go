package watchdog


import "testing"
import "sync"


type testToilerForCrashed struct{
	death chan struct{}
	terminateWaitgroup *sync.WaitGroup
	toilWaitgroup *sync.WaitGroup
}
func (td *testToilerForCrashed) Terminate() {
	td.terminateWaitgroup.Done()
}
func (td *testToilerForCrashed) Toil() {
	td.toilWaitgroup.Done()

	<-td.death
	panic("CHAOS!!!")
}


func TestOneForOneCrashed(t *testing.T) {

	// Create a watchdog.
		watchdog := NewOneForOne()

	// Create death channel.
	//
	// (Eventually) used to kill just one of the children.
		death := make(chan struct{})

	// Create wait groups.
		var terminateWg sync.WaitGroup
		var toilWg sync.WaitGroup

	// Create 10 toilers and watch them.
		const numChildToilers = 10
		for i:=1; i<=numChildToilers; i++ {

			// Create dummy toiler.
			//
			// This is what we'll send to .Watch()
				toiler := &testToilerForCrashed{
					death:death,
					terminateWaitgroup:&terminateWg,
					toilWaitgroup:&toilWg,
				}

			// Get the watchdog to watch the toiler.
				toilWg.Add(1)
				watchdog.Watch(toiler)
		}

	// Toil.
		go watchdog.Toil()

	// Confirm.
	//
	// The way this confirms is that if this is "hanging", then the Go runtime
	// will panic with the error: "fatal error: all goroutines are asleep - deadlock!".
		toilWg.Wait()


	// Crash one child and confirm.
	//
	// For the 'one for one' watchdog, only the crashed toilers should be restarted.
	//
	// The way this confirms is that if this is "hanging", then the Go runtime
	// will panic with the error: "fatal error: all goroutines are asleep - deadlock!".
	//
	// We do this test 10 times.
		for i:=0; i<10; i++ {
			terminateWg.Add(1)
			toilWg.Add(1)

			death <- struct{}{}

			terminateWg.Wait()
			toilWg.Wait()
		}
}


func TestOneForAllCrashed(t *testing.T) {

	// Create a watchdog.
		watchdog := NewOneForAll()

	// Create death channel.
	//
	// (Eventually) used to kill just one of the children.
		death := make(chan struct{})

	// Create wait groups.
		var terminateWg sync.WaitGroup
		var toilWg sync.WaitGroup

	// Create 10 toilers and watch them.
		const numChildToilers = 10
		for i:=1; i<=numChildToilers; i++ {

			// Create dummy toiler.
			//
			// This is what we'll send to .Watch()
				toiler := &testToilerForCrashed{
					death:death,
					terminateWaitgroup:&terminateWg,
					toilWaitgroup:&toilWg,
				}

			// Get the watchdog to watch the toiler.
				toilWg.Add(1)
				watchdog.Watch(toiler)
		}

	// Toil.
		go watchdog.Toil()

	// Confirm.
	//
	// The way this confirms is that if this is "hanging", then the Go runtime
	// will panic with the error: "fatal error: all goroutines are asleep - deadlock!".
		toilWg.Wait()

	// Crash one child and confirm.
	//
	// For the 'one for all' watchdog, all the toilers should be restarted.
	//
	// The way this confirms is that if this is "hanging", then the Go runtime
	// will panic with the error: "fatal error: all goroutines are asleep - deadlock!".
	//
	// We do this test 10 times.
		for i:=0; i<10; i++ {
			terminateWg.Add(numChildToilers)
			toilWg.Add(numChildToilers)

			death <- struct{}{}

			terminateWg.Wait()
			toilWg.Wait()
		}
}
