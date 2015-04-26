package watchdog


import "testing"
import "math/rand"
import "sync"
import "time"


type testToilerForReturned struct{
	end chan struct{}
	terminateWaitgroup *sync.WaitGroup
	toilWaitgroup *sync.WaitGroup
}
func newTestToilerForReturned(terminateWg *sync.WaitGroup, toilWg *sync.WaitGroup) Toiler {
	end := make(chan struct{})

	toiler := testToilerForReturned{
		end:end,
		terminateWaitgroup:terminateWg,
		toilWaitgroup:toilWg,
	}

	return &toiler
}
func (td *testToilerForReturned) Terminate() {
	td.terminateWaitgroup.Done()
}
func (td *testToilerForReturned) Toil() {
	td.toilWaitgroup.Done()

	<-td.end
}


func TestOneForOneReturned(t *testing.T) {

	// Seed the random number generator.
	//
	// Random number generator used later.
		rand.Seed( time.Now().UTC().UnixNano())

	// Create a watcher.
		watcher := NewOneForOne()

	// Create wait groups.
		var terminateWg sync.WaitGroup
		var toilWg sync.WaitGroup

	// Create 10 toilers and watch them.
		savedToilers := make([]Toiler, 0, 8)

		const numChildToilers = 10
		for i:=1; i<=numChildToilers; i++ {

			// Create dummy toiler.
			//
			// This is what we'll send to .Watch()
				toiler := newTestToilerForReturned(&terminateWg, &toilWg)

			// Get the watcher to watch the toiler.
				toilWg.Add(1)
				watcher.Watch(toiler)
				savedToilers = append(savedToilers, toiler)
		}

	// Toil.
		watcher.Toil()

	// Confirm.
	//
	// The way this confirms is that if this is "hanging", then the Go runtime
	// will panic with the error: "fatal error: all goroutines are asleep - deadlock!".
		toilWg.Wait()


	// End one child at a time and confirm.
		for i:=0; i<numChildToilers; i++ {

			index := rand.Intn(len(savedToilers))
			targetToiler := savedToilers[index]
			savedToilers = append(savedToilers[:index], savedToilers[1+index:]...)

			watcher.(*wdt).returned(targetToiler)

			if expected, actual := numChildToilers - (1+i), len(watcher.(*wdt).toilers); expected != actual {
				t.Errorf("Ended %d toilers, expected %d toilers left but actually got %d.", i+1, expected, actual)
			}
		}
}

func TestOneForAllReturned(t *testing.T) {

	// Seed the random number generator.
	//
	// Random number generator used later.
		rand.Seed( time.Now().UTC().UnixNano())

	// Create a watcher.
		watcher := NewOneForAll()

	// Create wait groups.
		var terminateWg sync.WaitGroup
		var toilWg sync.WaitGroup

	// Create 10 toilers and watch them.
		savedToilers := make([]Toiler, 0, 8)

		const numChildToilers = 10
		for i:=1; i<=numChildToilers; i++ {

			// Create dummy toiler.
			//
			// This is what we'll send to .Watch()
				toiler := newTestToilerForReturned(&terminateWg, &toilWg)

			// Get the watcher to watch the toiler.
				toilWg.Add(1)
				watcher.Watch(toiler)
				savedToilers = append(savedToilers, toiler)
		}

	// Toil.
		watcher.Toil()

	// Confirm.
	//
	// The way this confirms is that if this is "hanging", then the Go runtime
	// will panic with the error: "fatal error: all goroutines are asleep - deadlock!".
		toilWg.Wait()


	// End one child at a time and confirm.
		for i:=0; i<numChildToilers; i++ {

			index := rand.Intn(len(savedToilers))
			targetToiler := savedToilers[index]
			savedToilers = append(savedToilers[:index], savedToilers[1+index:]...)

			watcher.(*wdt).returned(targetToiler)

			if expected, actual := numChildToilers - (1+i), len(watcher.(*wdt).toilers); expected != actual {
				t.Errorf("Ended %d toilers, expected %d toilers left but actually got %d.", i+1, expected, actual)
			}
		}
}

