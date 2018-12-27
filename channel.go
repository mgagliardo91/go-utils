package utils

// StopChannel is a simple channel used to gracefully stop go-routines, handling
// the handshaking required to ensure that the go-routine has completed
type StopChannel struct {
	OnRequest chan bool
	stop      chan struct{}
}

// NewStopChannel creates an instance of a StopChannel to be used across go-routines
func NewStopChannel() StopChannel {
	return StopChannel{
		OnRequest: make(chan bool),
		stop:      make(chan struct{}),
	}
}

// RequestStop attempts to trigger a stop on the StopChannel and wait for it to complete
func (sC StopChannel) RequestStop() {
	go func() {
		sC.OnRequest <- true
	}()

	<-sC.stop
}

// Stop will close the StopChannel
func (sC StopChannel) Stop() {
	close(sC.stop)
}
