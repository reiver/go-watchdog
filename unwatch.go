package watchdog


// doUnwatch does the actual work of unwatching a Toiler for the Watcher.
//
// THIS METHOD SHOULD ONLY BE CALLED FROM THE watchover METHOD!
func (dog *wdt) doUnwatch(toiler Toiler) {

	toilers := dog.toilers
	length := len(toilers)

	for i:=0; i<length; i++ {
		if toiler == toilers[i] {
			toilers = append(toilers[:i], toilers[1+i:]...)

			length = len(toilers)
		}
	}

	dog.toilers = toilers
}
