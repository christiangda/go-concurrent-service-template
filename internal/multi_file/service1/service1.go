package service1

import (
	"log/slog"
	"math/rand"
	"time"
)

type Server1 struct {
	stopCh      chan struct{}
	waitStopCh  chan struct{}
	waitStartCh chan struct{}
}

func NewServer1() *Server1 {
	return &Server1{
		stopCh:      make(chan struct{}),
		waitStopCh:  make(chan struct{}),
		waitStartCh: make(chan struct{}),
	}
}

func (s *Server1) Start() {
	slog.Info("starting service...", "service", "1")

	// simulate the provisioning of the service
	slog.Info("waiting start...", "service", "1")
	time.Sleep(time.Duration(rand.Intn(3)) * time.Second)

	// closing the channel will notify the server that the service is started
	close(s.waitStartCh)
	slog.Info("...service started", "service", "1")

	// do something here for long running tasks
	// like a gRPC server

	// blocked to wait until channel is closed to stop the service
	<-s.stopCh

	// close the channel to notify the server that the service is stopped
	close(s.waitStopCh)
}

func (s *Server1) Stop() {
	slog.Warn("stopping services...", "service", "1")

	slog.Info("waiting stop...", "service", "1")
	// simulate the time spent to stop gracefully shutdown the service
	time.Sleep(time.Duration(rand.Intn(4)) * time.Second)

	close(s.stopCh)
	slog.Warn("...service stopped", "service", "1")
}

func (s *Server1) WaitStart() {
	<-s.waitStartCh
}

func (s *Server1) WaitStop() {
	<-s.waitStopCh
}
