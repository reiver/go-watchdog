package watchdog


// watchover is what gives a Watcher life. The Watcher's watchover method
// is spun up as a goroutine. It receives messages (from other goroutines)
// via 2 channels: the Watcher's message channel or the Watcher's terminate
// channel. These 2 channels aren't sent on directly, but instead other
// methods of the Watcher send messages on those channels.
func (dog *oneForOne) watchover() {

	Loop: for {
		select {
			case m := <- dog.message:
				switch msg := m.(type) {
					case crashed:
						msg.toiler.Terminate()
						watchedToil(msg.toiler, func(){
							dog.crashed(msg.toiler)
						})
						close(msg.done)
					case toil:
						for _,toiler := range dog.toilers {
							watchedToil(toiler, func(){
								dog.crashed(toiler)
							})
						}
						close(msg.done)
					case watch:
						dog.toilers = append(dog.toilers, msg.toiler)
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

// watchover is what gives a Watcher life. The Watcher's watchover method
// is spun up as a goroutine. It receives messages (from other goroutines)
// via 2 channels: the Watcher's message channel or the Watcher's terminate
// channel. These 2 channels aren't sent on directly, but instead other
// methods of the Watcher send messages on those channels.
func (dog *oneForAll) watchover() {

	Loop: for {
		select {
			case m := <- dog.message:
				switch msg := m.(type) {
					case crashed:
						for _,toiler := range dog.toilers {
							toiler.Terminate()
						}
						for _,toiler := range dog.toilers {
							watchedToil(toiler, func(){
								dog.crashed(toiler)
							})
						}
						close(msg.done)

					case toil:
						for _,toiler := range dog.toilers {
							watchedToil(toiler, func(){
								dog.crashed(toiler)
							})
						}
						close(msg.done)
					case watch:
						dog.toilers = append(dog.toilers, msg.toiler)
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
