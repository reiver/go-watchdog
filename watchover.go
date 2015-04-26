package watchdog


// watchover is what gives a Watcher life. The Watcher's watchover method
// is spun up as a goroutine. It receives messages (from other goroutines)
// via 2 channels: the Watcher's message channel and the Watcher's terminate
// channel. These 2 channels aren't sent on directly, but instead other
// methods of the Watcher send messages on those channels.
func (dog *wdt) watchover() {

	toiling := false
	toilListeners := make([]chan struct{}, 0, defaultLengthOfToilListenersSlice)

	Loop: for {
		select {
			case m := <- dog.message:
				switch msg := m.(type) {
					case crashed:
						dog.crashedStrategyFunc( newWatchedToiler(dog, msg.toiler) )
						close(msg.done)
					case returned:
						dog.doUnwatch(msg.toiler)
						close(msg.done)
						if 0 >= len(dog.toilers) {
							for _,toilDone := range toilListeners {
								close(toilDone)
							}
							break Loop
						}
					case toil:
						if !toiling {
							for _,toiler := range dog.toilers {
								watchedToil(toiler, func(exception interface{}){
									dog.crashed(toiler)
								}, func(){
									dog.returned(toiler)
								})
							}
							toiling = true
						}
						toilListeners = append(toilListeners, msg.done)
					case watch:
						dog.toilers = append(dog.toilers, msg.toiler)
						if toiling {
							watchedToil(msg.toiler, func(exception interface{}){
								dog.crashed(msg.toiler)
							}, func(){
								dog.returned(msg.toiler)
							})
						}
						close(msg.done)
					default:
						//@TODO
				}

			case trmnt := <- dog.terminate:
				for _,toiler := range dog.toilers {
					toiler.Terminate()
				}
				close(trmnt.done)
				break Loop
				//@TODO
		}
	}
}
