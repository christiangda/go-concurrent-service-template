package service1

import "log/slog"

type Server1 struct {
	stopCh chan struct{}
}

func NewServer1() *Server1 {
	return &Server1{
		stopCh: make(chan struct{}),
	}
}

func (s *Server1) Start() {
	// do something here for long running tasks
	// like a gRPC server
	slog.Info("Starting service 1...")

	<-s.stopCh
}

func (s *Server1) Stop() {
	// do something here to stop the service
	slog.Info("Stopping service 1...")

	close(s.stopCh)
}
