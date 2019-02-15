package signal

import (
	"os"
	osig "os/signal"
	"syscall"
)

func WaitForInterrupt(cb func()) {
	signalCh := make(chan os.Signal)
	osig.Notify(signalCh, os.Interrupt, syscall.SIGTERM)
	<-signalCh

	if cb != nil {
		cb()
	}
}
