package police

// Arrival limits the count of goroutine processed at same time.
type Arrival struct {
	Ch chan struct{}
}

// Limit creates a channel limited capacity.
func (a *Arrival) Limit(i int) {
	a.Ch = make(chan struct{}, i)
}

// Block adds struct to channel.
func (a *Arrival) Block() {
	a.Ch <- struct{}{}
}

// Release remove struct from channel.
func (a *Arrival) Release() {
	<-a.Ch
}
