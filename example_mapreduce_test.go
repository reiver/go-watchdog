package watchdog_test

import "fmt"
import "github.com/reiver/go-watchdog"
import "sync"



type Reducer struct {
	wg *sync.WaitGroup
	inputChannel chan int
	resultChannel chan int
}
func NewReducer() *Reducer {
	var wg sync.WaitGroup

	inputCh := make(chan int)
	resultCh := make(chan int)

	reducer := Reducer{
		wg:&wg,
		inputChannel:inputCh,
		resultChannel:resultCh,
	}

	return &reducer
}
func (r *Reducer) AddMapper(delta int) {
	r.wg.Add(delta)
}
func (r *Reducer) MapperDone() {
	r.wg.Done()
}
func (r *Reducer) Result() int {
	r.wg.Wait()
	close(r.inputChannel)

	return <-r.resultChannel
}
func (r *Reducer) Send(n int) {
	r.inputChannel <- n
}
func (r *Reducer) Toil() {
	count := 0

	for n := range r.inputChannel {
		count += n
	}

	r.resultChannel <- count
	close(r.resultChannel)
}



type Mapper struct{
	reducer *Reducer

	inputCh   chan   int
}
func NewMapper(inputChannel chan int, reducer *Reducer) *Mapper {
	mapper := Mapper{
		reducer:reducer,
		inputCh:inputChannel,
	}
	return &mapper
}
func (m *Mapper) Terminate() {
	// Nothing here.
}
func (m *Mapper) Toil() {

	for n := range m.inputCh {
		if isEven := 0 == (n % 2); isEven {
			m.reducer.Send(n)
		}
	}

	m.reducer.MapperDone()
}
func (m *Mapper) Map(n int) {
	m.inputCh <- n
}
		


// This example shows usage of watchdog for a toy MapReduce.
//
// The mapper phase filters out odd numbers, and only sends even
// numbers to the reducer.
//
// The reducer adds up all the numbers it receives.
//
// So conceptually we are doing something like the following:
//
//	7, 6, 5, 4, 3, 2, 1 -> Map -> 6, 4, 2 -> Reduce -> 12 (= 2+4+6)
//
func Example_mapReduce() {

	// Create the reducer.
		reducer := NewReducer()
		go reducer.Toil() // ← We are not using a watchdog for the reducer, so we call .Toil() manually.

	// Create mappers.
	//
	// The mappers use watchdog!
	//
	// The 'inputChannel' is the input to the mappers.
	//
	// Notice that we are using a 'wait group' to determine
	// when the mappers are done.
	//
		inputChannel := make(chan int)

		mappers := watchdog.NewOneForOne()

		for i:=0; i<5; i++ {
			var toiler watchdog.Toiler = NewMapper(inputChannel, reducer)

			mappers.Watch(toiler)
			reducer.AddMapper(1)
		}

		go mappers.Toil() // ← Don't forget this!

	// Send numbers to mappers.
	//
	// We close 'inputChannel' to tell the
	// mappers we are done sending numbers
	// to them.
	//
		inputChannel <- 1
		inputChannel <- 2
		inputChannel <- 3
		inputChannel <- 4
		inputChannel <- 5
		inputChannel <- 6
		inputChannel <- 7

		close(inputChannel) // ← This is important!

		mappers.Terminate()

	// Get the result of the MapReduce.
		count := reducer.Result()

	// Output.
		fmt.Printf("The count is %d.\n", count)

	// Output:
	// The count is 12.
}
