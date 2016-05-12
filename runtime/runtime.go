package runtime

import (
	"os/signal"
	"syscall"
	"os"
)

//-------------------------------------
//
//
//
//-------------------------------------
func WaitExitSignal() os.Signal {
	signalChan := make(chan os.Signal)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)
	return <-signalChan
}
