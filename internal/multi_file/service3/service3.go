package service3

import (
	"log/slog"
	"math/rand"
	"time"
)

type Server3 struct {
	stopCh      chan struct{}
	waitStopCh  chan struct{}
	waitStartCh chan struct{}
}

func NewServer3() *Server3 {
	return &Server3{
		stopCh:      make(chan struct{}),
		waitStopCh:  make(chan struct{}),
		waitStartCh: make(chan struct{}),
	}
}

func (s *Server3) Start() {
	slog.Info("starting service...", "service", "3")
	// simulate the provisioning of the service
	time.Sleep(time.Duration(rand.Intn(3)) * time.Second)
	slog.Info("...service started", "service", "3")

	// closing the channel will notify the server that the service is started
	close(s.waitStartCh)

	// do something here for long running tasks
	// like a gRPC server

	// blocked to wait until channel is closed to stop the service
	<-s.stopCh

	// close the channel to notify the server that the service is stopped
	close(s.waitStopCh)
}

func (s *Server3) Stop() {
	slog.Warn("stopping services...", "service", "3")

	// simulate the time spent to stop gracefully shutdown the service
	time.Sleep(time.Duration(rand.Intn(4)) * time.Second)
	slog.Warn("...service stopped", "service", "3")

	close(s.stopCh)
}

func (s *Server3) WaitStart() {
	<-s.waitStartCh
}

func (s *Server3) WaitStop() {
	<-s.waitStopCh
}
