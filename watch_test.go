package watchdog


import "testing"


type testToilerForWatch struct{}
func (td *testToilerForWatch) Terminate() {}
func (td *testToilerForWatch) Toil() {}


func TestOneForOneWatch(t *testing.T) {

	// Create a watchdog.
		watchdog := NewOneForOne()

	// Confirm.
		if expected, actual := 0, len(watchdog.(*oneForOne).toilers); expected != actual {
			t.Errorf("For initial watchdog, expected %d toilers but got %d.", expected, actual)
		}


	// Do this test 10 times.
		const numChildToilers = 10
		for i:=1; i<=numChildToilers; i++ {

			// Create dummy toiler.
			//
			// This is what we'll send to .Watch()
				toiler := &testToilerForWatch{}

			// Get the watchdog to watch the toiler.
				watchdog.Watch(toiler)

			// Confirm.
				if expected, actual := i, len(watchdog.(*oneForOne).toilers); expected != actual {
					t.Errorf("After adding toiler to watchdog, expected %d toilers but got %d.", expected, actual)
				}
		}
}


func TestOneForAllWatch(t *testing.T) {

	// Create a watchdog.
		watchdog := NewOneForAll()

	// Confirm.
		if expected, actual := 0, len(watchdog.(*oneForAll).toilers); expected != actual {
			t.Errorf("For initial watchdog, expected %d toilers but got %d.", expected, actual)
		}


	// Do this test 10 times.
		const numChildToilers = 10
		for i:=1; i<=numChildToilers; i++ {

			// Create dummy toiler.
			//
			// This is what we'll send to .Watch()
				toiler := &testToilerForWatch{}

			// Get the watchdog to watch the toiler.
				watchdog.Watch(toiler)

			// Confirm.
				if expected, actual := i, len(watchdog.(*oneForAll).toilers); expected != actual {
					t.Errorf("After adding toiler to watchdog, expected %d toilers but got %d.", expected, actual)
				}
		}
}
