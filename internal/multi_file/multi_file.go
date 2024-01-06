package multi_file

import (
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"github.com/christiangda/go-concurrent-service-template/internal/multi_file/service1"
	"github.com/christiangda/go-concurrent-service-template/internal/multi_file/service2"
	"github.com/christiangda/go-concurrent-service-template/internal/multi_file/service3"
)

// this simulates the main() function
func Run() {
	slog.Info("starting server...")

	// create a channel to listen for os signals
	osSigCh := make(chan os.Signal, 1)
	signal.Notify(osSigCh, os.Interrupt, syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)

	// used to stop the server and their services
	// once a signal is received from the os
	// the serverStopCh channel will be closed and this will
	// trigger the stop of the server and their services
	serverStopCh := make(chan struct{})

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
	service1 := service1.NewServer1()
	go service1.Start()

	// start service 2 in a separate goroutine
	service2 := service2.NewServer2()
	go service2.Start()

	// start service 3 in a separate goroutine
	service3 := service3.NewServer3()
	go service3.Start()

	// blocked main to wait until all services are started
	service1.WaitStart()
	service2.WaitStart()
	service3.WaitStart()
	slog.Info("...server started")

	slog.Warn("-> to stop the server press `CTRL+C`")

	// blocked main to wait for stop the server
	<-serverStopCh

	// notify the services to stop asynchronously
	go service1.Stop()
	go service2.Stop()
	go service3.Stop()

	// blocked main to wait for stop each service
	service1.WaitStop()
	service2.WaitStop()
	service3.WaitStop()

	// all channels are closed (released), the server is stopped
	slog.Info("...server stopped")
}
