package service2

import (
	"log/slog"
	"math/rand"
	"time"
)

type Server2 struct {
	stopCh      chan struct{}
	waitStopCh  chan struct{}
	waitStartCh chan struct{}
}

func NewServer2() *Server2 {
	return &Server2{
		stopCh:      make(chan struct{}),
		waitStopCh:  make(chan struct{}),
		waitStartCh: make(chan struct{}),
	}
}

func (s *Server2) Start() {
	slog.Info("starting service...", "service", "2")
	// simulate the provisioning of the service
	time.Sleep(time.Duration(rand.Intn(3)) * time.Second)
	slog.Info("...service started", "service", "2")

	// closing the channel will notify the server that the service is started
	close(s.waitStartCh)

	// do something here for long running tasks
	// like a gRPC server

	// blocked to wait until channel is closed to stop the service
	<-s.stopCh

	// close the channel to notify the server that the service is stopped
	close(s.waitStopCh)
}

func (s *Server2) Stop() {
	slog.Warn("stopping services...", "service", "2")

	// simulate the time spent to stop gracefully shutdown the service
	time.Sleep(time.Duration(rand.Intn(4)) * time.Second)
	slog.Warn("...service stopped", "service", "2")

	close(s.stopCh)
}

func (s *Server2) WaitStart() {
	<-s.waitStartCh
}

func (s *Server2) WaitStop() {
	<-s.waitStopCh
}
