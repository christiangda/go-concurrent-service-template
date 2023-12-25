package single_file

import (
	"log/slog"
	"os"
	"os/signal"
	"syscall"
)

// this is like a main function
func Run() {
	slog.Info("Starting server...")

	// create a channel to listen for os signals
	osSigCh := make(chan os.Signal, 1)
	signal.Notify(osSigCh, os.Interrupt, syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)

	// used to stop the server and their services
	// once a signal is received from the os
	// the stopCh channel will be closed and this will
	// trigger the stop of the server and their services
	stopCh := make(chan struct{})

	// listen for os signals in a separate goroutine
	go func() {
		for {
			select {
			case sig := <-osSigCh:
				slog.Info("Received signal: %v", "signal", sig)

				switch sig {
				case os.Interrupt, syscall.SIGINT, syscall.SIGTERM:
					slog.Info("Stopping server...")
					close(stopCh)
					return
				case syscall.SIGHUP:
					slog.Info("Reloading configuration...")
					return
				}
			case <-stopCh:
				return
			}
		}
	}()

	// start service 1 in a separate goroutine
	slog.Info("Starting service 1...")
	go func() {
		// do something here for long running tasks
		// like a gRPC server

		<-stopCh
		slog.Info("Stopping service 1...")
	}()

	// start service 2 in a separate goroutine
	slog.Info("Starting service 2...")
	go func() {
		// do something here for long running tasks
		// like a http server

		<-stopCh
		slog.Info("Stopping service 2...")
	}()

	// start service 3 in a separate goroutine
	slog.Info("Starting service 3...")
	go func() {
		// do something here for long running tasks
		// like a TCP server

		<-stopCh
		slog.Info("Stopping service 3...")
	}()

	// wait for stop signal
	slog.Info("Server started.")
	<-stopCh
}
