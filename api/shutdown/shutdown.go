package shutdown

import (
	"os"
	"os/signal"
	"syscall"
)

type Callback func() error

/*
*
Wait is a function that waits for signal from OS to stop the program.
It accepts callback function to gracefully stop the program.
Put shutdown of the server, worker-pool stop function execution, db connection close function etc.
Returns an error in case callback function is failed to execute.
*/
func Wait(callback Callback) error {
	shutdownChannel := make(chan os.Signal, 1)
	signal.Notify(shutdownChannel, os.Interrupt, os.Kill, syscall.SIGHUP, syscall.SIGTERM)
	<-shutdownChannel

	return callback()
}
