package watchdog


// Toiler is the interface that groups Toil and Terminate methods.
//
// A Toiler is the basic abstraction of something that toils (i.e., does work)
// and can be terminated, to stop the toiling.
//
// When the Toil method is invoked, the Toiler should start toiling.
//
// When the Terminate method is invoked, the Toiler should stop toiling.
//
// Any initialization that needs to be done should be done when the
// Toil method is first invoked.
//
// Any clean up that needs to be done should be done when the
// Terminate method invoked.
//
// The Toiler is allowed to panic() from the Toil method.
// In fact, to take advantage of what 'watchdog' provides, the Toiler
// should be panic()ing for certain kinds of errors.
//
// The Toiler should be written in a way where Terminate and Toil can
// be invoked mutliple times.
type Toiler interface {
	Terminate()
	Toil()
}
