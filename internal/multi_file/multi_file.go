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
				slog.Info("Received signal", "signal", sig)

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
	service1 := service1.NewServer1()
	go service1.Start()

	// start service 2 in a separate goroutine
	service2 := service2.NewServer2()
	go service2.Start()

	// start service 3 in a separate goroutine
	service3 := service3.NewServer3()
	go service3.Start()

	// wait for stop signal
	slog.Info("Server started.")
	<-stopCh
	service1.Stop()
	service2.Stop()
	service3.Stop()
}
