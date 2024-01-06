package single_file

import (
	"log/slog"
	"math/rand"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// this simulates the main() function
func Run() {
	slog.Info("starting server...")

	// create a channel to listen for Operating System (OS) signals
	osSigCh := make(chan os.Signal, 1)
	signal.Notify(osSigCh, os.Interrupt, syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)

	// used to stop the server and their services
	// once a signal is received from the os
	// the serverStopCh channel will be closed and this will
	// trigger the stop of the server and their services
	serverStopCh := make(chan struct{})

	// create a channel for each service in order to stop, wait for stop and wait for start
	s1StopCh := make(chan struct{})
	s1WaitStopCh := make(chan struct{})
	s1WaitStartCh := make(chan struct{})

	s2StopCh := make(chan struct{})
	s2WaitStopCh := make(chan struct{})
	s2WaitStartCh := make(chan struct{})

	s3StopCh := make(chan struct{})
	s3WaitStopCh := make(chan struct{})
	s3WaitStartCh := make(chan struct{})

	// this is the main logic of the server
	// it will listen for OS signals and stop the services and the server
	// when a signal is received
	go func() {
		for {
			select {
			case sig := <-osSigCh:
				slog.Info("received OS signal", "type", sig)

				switch sig {
				case os.Interrupt, syscall.SIGINT, syscall.SIGTERM:
					slog.Info("stopping server...")
					close(serverStopCh)

					return
				case syscall.SIGHUP:
					slog.Info("reloading configuration...")
					// do something here to reload the configuration

					return
				}

			// if the stopCh channel is closed, the server is stopped
			// exit from the goroutine. A closed channel always return a zero value
			// and could be read without blocking
			case <-serverStopCh:
				return
			}
		}
	}()

	// start service 1 in a separate goroutine
	go func() {
		slog.Info("starting service...", "service", "1")

		// simulate the provisioning of the service
		slog.Info("waiting start...", "service", "1")
		time.Sleep(time.Duration(rand.Intn(3)) * time.Second)

		// closing the channel will notify the server that the service is started
		close(s1WaitStartCh)
		slog.Info("...service started", "service", "1")

		// do something here for long running tasks
		// like a gRPC server

		// blocked to wait until channel is closed to stop the service
		<-s1StopCh

		slog.Warn("stopping services...", "service", "1")

		// simulate the time spent to stop gracefully shutdown the service
		slog.Info("waiting stop...", "service", "1")
		time.Sleep(time.Duration(rand.Intn(4)) * time.Second)

		// close the channel to notify the server that the service is stopped
		close(s1WaitStopCh)
		slog.Warn("...service stopped", "service", "1")
	}()

	// start service 2 in a separate goroutine
	go func() {
		slog.Info("starting service...", "service", "2")

		// simulate the provisioning of the service
		slog.Info("waiting start...", "service", "2")
		time.Sleep(time.Duration(rand.Intn(3)) * time.Second)

		// closing the channel will notify the server that the service is started
		close(s2WaitStartCh)
		slog.Info("...service started", "service", "2")

		// do something here for long running tasks
		// like a http server

		// blocked to wait until channel is closed to stop the service
		<-s2StopCh

		slog.Warn("stopping services...", "service", "2")

		// simulate the time spent to stop gracefully shutdown the service
		slog.Info("waiting stop...", "service", "2")
		time.Sleep(time.Duration(rand.Intn(4)) * time.Second)

		// close the channel to notify the server that the service is stopped
		close(s2WaitStopCh)
		slog.Warn("...service stopped", "service", "2")
	}()

	// start service 3 in a separate goroutine
	go func() {
		slog.Info("starting service...", "service", "3")

		// simulate the provisioning of the service
		slog.Info("waiting start...", "service", "3")
		time.Sleep(time.Duration(rand.Intn(4)) * time.Second)

		// closing the channel will notify the server that the service is started
		close(s3WaitStartCh)
		slog.Info("...service started", "service", "3")

		// do something here for long running tasks
		// like a TCP server

		// blocked to wait until channel is closed to stop the service
		<-s3StopCh

		slog.Warn("stopping services...", "service", "3")

		// simulate the time spent to stop gracefully shutdown the service
		slog.Info("waiting stop...", "service", "3")
		time.Sleep(time.Duration(rand.Intn(2)) * time.Second)

		// close the channel to notify the server that the service is stopped
		close(s3WaitStopCh)
		slog.Warn("...service stopped", "service", "3")
	}()

	// blocked main to wait until all services are started
	// in the order s1, s2, s3 because these are blocking in that order
	// the main goroutine
	<-s1WaitStartCh
	<-s2WaitStartCh
	<-s3WaitStartCh
	slog.Info("...server started")

	slog.Warn("-> to stop the server press `CTRL+C`")

	// blocked main to wait for stop the server
	// the serverStopCh channel is closed when a signal is received
	// in a different goroutine was started before
	<-serverStopCh

	// notify the services to stop asynchronously
	go close(s1StopCh)
	go close(s2StopCh)
	go close(s3StopCh)

	// blocked main to wait for stop each service
	// the channels are closed when the services are stopped
	// in a different goroutine was started before
	// the wait is in the order s1,s2,s3 because these are blocking in that order
	<-s1WaitStopCh
	<-s2WaitStopCh
	<-s3WaitStopCh

	// all channels are closed (released), the server is stopped
	slog.Warn("...server stopped")
}
