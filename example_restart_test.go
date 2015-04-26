package watchdog_test

import "fmt"
import "github.com/reiver/go-watchdog"



type DummyToiler struct{
	name string
}
func NewDummyToiler(name string) *DummyToiler {
	toiler := DummyToiler{
		name:name,
	}
	return &toiler
}
func (t *DummyToiler) Name() string {
	return t.name
}
func (t *DummyToiler) Terminate() {
	// Nothing here.
}
func (t *DummyToiler) Toil() {
	// Nothing here.
}
		


// This example shows how you would restart all the Toilers
// in a Watcher.
func Example_restart() {

	// Create a watcher.
		watcher := watchdog.NewOneForOne()

	// Add some Toilers to the Watcher.
		const numToilers = 5
		for i:=1; i<=numToilers; i++ {
			toilerName := fmt.Sprintf("Toiler #%d", i)

			toiler := NewDummyToiler(toilerName)

			watcher.Watch(toiler)
		}

		go watcher.Toil() // ← Don't forget this!

	// Restart all of the Toilers being watched by the Watcher.
		count := 0

		watcher.Map(func(watchedToiler watchdog.WatchedToiler){

			watchedToiler.Terminate()
			watchedToiler.Toil()

			count += 1
		})

	// Output.
		fmt.Printf("Restarted %d toilers.\n", count)

	// Output:
	// Restarted 5 toilers.
}
